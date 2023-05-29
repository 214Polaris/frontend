<!-- eslint-disable vue/no-multiple-template-root -->
<!--连接到主页-->
<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    :router="true"
    style="position: relative; overflow: hidden; display: flex"
  >
    <el-menu-item @click.native="goback()">首页</el-menu-item>
    <el-submenu index="2">
      <template #title>我的用户</template>
      <el-menu-item @click.native="MyOrders()">我的订单</el-menu-item>
      <el-menu-item @click.native="MyInfo()">个人信息</el-menu-item>
      <el-menu-item v-if="token" @click.native="logout()"
        >退出登录</el-menu-item
      >
      <el-submenu index="2-4">
        <template #title>选项4</template>
        <el-menu-item index="2-4-1">选项1</el-menu-item>
        <el-menu-item index="2-4-2">选项2</el-menu-item>
        <el-menu-item index="2-4-3">选项3</el-menu-item>
      </el-submenu>
    </el-submenu>
    <el-menu-item index="3" disabled>消息中心</el-menu-item>
    <div style="display: flex; margin-bottom: 10px">
      <el-autocomplete
        v-model="state"
        :fetch-suggestions="querySearchAsync"
        placeholder="请输入想要搜索的商品"
        @select="handleSelect"
        style="width: 400px; margin-left: 300px; margin-top: 10px"
      ></el-autocomplete>
      <el-button
        type="primary"
        @click="search"
        style="margin-left: 20px; margin-top: 10px"
        >搜索</el-button
      >
    </div>
    <!--登录按钮-->
    <div class="user">
      <el-button
        type="primary"
        style="margin-top: 10px; margin-bottom: 10px"
        @click="login"
        v-if="!token"
      >
        <router-link to="/login">登录</router-link>
      </el-button>
      <!--用户图标，之后在这里添加点击到用户主页面-->
      <div v-if="token">
        <el-image
          :src="require('./assets/user.png')"
          style="width: 55px; height: 55px"
          :fit="fit"
        />
      </div>
    </div>
  </el-menu>
  <router-link to="/"></router-link>
  <router-view />
</template>

<script>
import { ElMessageBox } from 'element-plus'
export default {
  data () {
    return {
      token: localStorage.token,
    }
  },
  methods: {
    //跳转到登陆页面
    login () {
      this.$router.replace({ name: 'login' })
    },
    //跳转到主页
    goback () {
      this.$router.replace({ name: 'MainPage' })
    },
    //退出登录
    logout () {
      ElMessageBox.alert("您已退出登录").then(() => {
        this.$store.commit('LOGOUT')
      }
      )
    },
    //搜索功能（暂不可用）
    search () {
      ElMessageBox.alert("当前没有连接到后端").then(() => {
        this.$router.replace({ name: 'login' })
      }
      )
    },
    //进入订单页面
    MyOrders () {
      if (!this.token) {
        ElMessageBox.alert("请先登录").then(() => {
          this.$router.replace({ name: 'login' })
        }
        )
      }
      //补充跳转到订单页面
      else{}
    },
    //进入个人信息页面
    MyInfo(){
      if (!this.token) {
        ElMessageBox.alert("请先登录").then(() => {
          this.$router.replace({ name: 'login' })
        }
        )
      }
      //补充跳转到个人信息页面
      else{}
    }
  }
}
</script>

<style>
@import "./assets/base.css";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.user {
  margin-left: 600px;
  display: flex;
}
</style>
