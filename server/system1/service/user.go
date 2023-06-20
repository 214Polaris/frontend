package service

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"system/data"
	"system/model"
)

type UserSrv interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	HandleRequest(c *gin.Context)
	Edit(ctx *gin.Context)
}

type UserService struct {
	Repo data.UserRepoInterface
}

func (srv *UserService) Login(ctx *gin.Context) {
	var u model.User
	u.Username = ctx.PostForm("user")
	u.Password = ctx.PostForm("pass")
	//判断用户是否存在
	if srv.Repo.Query(u) == false {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//判断密码是否正确
	if srv.Repo.CheckPassword(u) == false {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不正确"})
		return
	}

	//发放token给前端
	token, err := GenerateToken(u.Username)
	if err != nil {
		log.Fatal(err)
	}
	UserID := srv.Repo.GetID(u)
	// c.JSON：返回JSON格式的数据
	ctx.JSON(200, gin.H{
		"code":   200,
		"token":  token,
		"userID": UserID,
		"msg":    "登录成功",
	})
}

func (srv *UserService) Register(ctx *gin.Context) {
	var u model.User
	u.Username = ctx.PostForm("user")
	u.Password = ctx.PostForm("pass")
	//密码加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	//将加密后的密码赋给u.Password
	u.Password = string(hasePassword)
	//判断用户是否存在
	if srv.Repo.Query(u) == true {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	//添加用户
	srv.Repo.Insert(u)

	// c.JSON：返回JSON格式的数据
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

func (srv *UserService) HandleRequest(ctx *gin.Context) {
	//如果是GET函数，不受中间件管理
	m := ctx.Request.Method
	if m == "GET" {
		return
	}
	// 从请求标头中获取JWT令牌
	tokenString := ctx.GetHeader("token")
	print(tokenString)
	//验证JWT
	token, claims, err := VerifyToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(401, gin.H{"code": 401, "msg": "Token过期"})
		ctx.Abort()
		return
	}
	//验证token，并存储在请求中
	ctx.Set("user", claims)
}

// 修改个人信息
func (srv *UserService) Edit(ctx *gin.Context) {
	//获取各种信息
	var u model.User
	UserIDStr := ctx.PostForm("id")
	UserID, _ := strconv.Atoi(UserIDStr)
	u.UserID = UserID
	//fmt.Print(u.UserID)
	u.Mobile = ctx.PostForm("mobile")
	u.Password = ctx.PostForm("password")
	u.Address = ctx.PostForm("address")
	u.Username = ctx.PostForm("name")
	u.UserEmail = ctx.PostForm("email")
	u.UserBornDate = ctx.PostForm("bornDate")
	u.UserRealName = ctx.PostForm("realName")
	if srv.Repo.Edit(u) == false {
		ctx.JSON(401, gin.H{"msg": "修改失败"})
		return
	}
	ctx.JSON(200, gin.H{"msg": "修改成功"})
}
