<template>
  <el-menu
    :default-active="activeMenu"
    :unique-opened="true"
    class="el-menu-demo"
    mode="horizontal"
    active-text-color="#409EFF"
    :ellipsis="false"
    router
    style="overflow: hidden; display: flex"
  >
    <el-menu-item index="MainPage" @click="goback()">首页</el-menu-item>
    <el-submenu index="2">
      <template #title>我的用户</template>
      <!--到时这里记得改index，实现高亮-->
      <el-menu-item index="2.1" @click="MyOrders()">我的订单</el-menu-item>
      <el-menu-item index="2.2" @click="MyInfo()">个人信息</el-menu-item>
      <el-menu-item index="2.3" v-if="token" @click="logout()"
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
    <div
      style="
        position: absolute;
        margin-bottom: 10px;
        margin-left: 600px;
        justify-content: center;
      "
    >
      <el-autocomplete
        v-model="state"
        :fetch-suggestions="querySearchAsync"
        placeholder="请输入想要搜索的商品"
        @select="handleSelect"
        style="width: 400px; margin-top: 10px"
      ></el-autocomplete>
      <el-button
        type="primary"
        @click="search"
        style="margin-top: 10px; margin-left: 5px"
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
        >登录
      </el-button>
      <!--用户图标，之后在这里添加点击到用户主页面-->
      <div v-if="token">
        <el-image
          :src="require('../../../assets/user.png')"
          style="width: 55px; height: 55px"
          :fit="fit"
        />
      </div>
    </div>
  </el-menu>
</template>

<script>
import { ElMessageBox } from "element-plus";
export default {
  data() {
    return {
      token: localStorage.token,
    };
  },
  computed: {
    // 我们使用计算属性来获取到当前点击的菜单的路由路径，然后设置default-active中的值
    // 使得菜单在载入时就能对应高亮
    activeMenu() {
      const route = this.$route;
      const { path, name } = route;
      // if set path, the sidebar will highlight the path you set
      // 可以在路由配置文件中设置自定义的路由路径到meta.activeMenu属性中，来控制菜单自定义高亮显示
      if (name) {
        return name;
      }
      return path;
    },
  },
  methods: {
    //跳转到登陆页面
    login() {
      this.$router.push({ name: "login" });
    },
    //跳转到主页
    goback() {
      this.$router.push({ path: "/" });
    },
    //退出登录
    logout() {
      ElMessageBox.alert("您已退出登录").then(() => {
        this.$store.commit("LOGOUT");
        this.$router.push({ path: "/" });
      });
    },
    //搜索功能（暂不可用）
    search() {
      ElMessageBox.alert("当前没有连接到后端").then(() => {
        this.$router.push({ name: "login" });
      });
    },
    //进入订单页面
    MyOrders() {
      if (!this.token) {
        ElMessageBox.alert("请先登录").then(() => {
          this.$router.push({ name: "login" });
        });
      }
      //补充跳转到订单页面
      else {
      }
    },
    //进入个人信息页面
    MyInfo() {
      if (!this.token) {
        ElMessageBox.alert("请先登录").then(() => {
          this.$router.push({ name: "login" });
        });
      }
      //补充跳转到个人信息页面
      else {
      }
    },
  },
};
</script>

<style>
@import "../../../assets/base.css";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.user {
  display: flex;
  position: absolute;
  right: 10px;
  top: 2px;
}
</style>
