package main

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// 数据库配置
const (
	userName = "root"
	password = "chen8574jun"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "shopping_system"
)

var db *sql.DB = nil

// User 相关信息
type User struct {
	Username string
	Password string
}

func InitDB() *sql.DB {
	// 如果已经建立数据库，直接返回
	if db != nil {
		return db
	}
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	// Db数据库连接池
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, err := sql.Open("mysql", path)
	//验证连接
	if err != nil {
		fmt.Println("open database fail")
		return DB
	}
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	fmt.Println("connect success")
	db = DB
	return db
}

// MyQuery 查询
func MyQuery(u User) bool {
	query := "SELECT * FROM user WHERE username = ?"
	rows, _ := db.Query(query, u.Username)
	for rows.Next() {
		return true
	}
	return false
}

// MyInsert 添加
func MyInsert(u User) {
	if MyQuery(u) == true {
		fmt.Println("添加失败，数据库中已有该账号")
		return
	}
	query := "INSERT INTO user (`username`, `password`) VALUES (?, ?)"
	r, err := db.Exec(query, u.Username, u.Password)
	if err != nil {
		fmt.Println("添加失败", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed", err)
	}
	fmt.Println("添加成功", id)
}

// MyDelete 删除
func MyDelete(u User) {
	if MyQuery(u) == false {
		fmt.Println("删除失败，数据库中无该账号")
		return
	}
	query := "DELETE FROM user WHERE username = ?"
	_, err := db.Exec(query, u.Username)
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}
}

func CheckPassword(u User) bool {
	query := "SELECT * FROM user WHERE username = ?"
	rows, _ := db.Query(query, u.Username)
	for rows.Next() {
		var Username, HashPassword string
		rows.Scan(&Username, &HashPassword)
		//如果两个密码值不相同，就密码出错，返回false
		if err := bcrypt.CompareHashAndPassword([]byte(HashPassword), []byte(u.Password)); err != nil {
			return false
		}
		//正确就返回true
		return true
	}
	//程序不会到这里
	return false
}
