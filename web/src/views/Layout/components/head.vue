<script setup>
const token = localStorage.getItem("token");
</script>

<template>
  <div class="Head-container">
    <ul class="head-ul">
      <li class="head-li" v-if="!token">
        <router-link to="/login"><a href="/login">请先登录</a></router-link>
      </li>
      <li class="head-li" v-else>
        <a href="javascript:"
          ><router-link
            to="/member"
            class="icon-user iconfont"
            title="我的信息"
          ></router-link
        ></a>
      </li>
      <li class="head-li">
        <a href="javascript:">帮助中心</a>
      </li>
      <li class="head-li">
        <a href="javascript:">关于我们</a>
      </li>
    </ul>
  </div>
</template>

<style lang="scss">
.head-ul {
  margin: 0;

  display: flex;
  height: 53px;
  justify-content: flex-end;
  align-items: center;
  list-style: none;
}

.head-li {
  color: #cdcdcd;
  text-align: center;
  padding-right: 18px;
  font-size: 14px;
}

a {
  color: #cdcdcd;
}

.head-li::after {
  padding-left: 18px;
  content: "|";
  float: inline-end;
  color: #94959e;
}

.head-li:last-child::after {
  display: none;
}

.head-li a:hover {
  color: #409eff;
}

.head-li a .iconfont {
  color: white;
  font-size: 16px;
  text-align: center;
}

.head-li a .iconfont {
  font-family: iconfont;
}

.Head-container {
  width: auto;
  margin: 0;
  position: relative;
  top: 0;
  background-color: #333333;
}
</style>
