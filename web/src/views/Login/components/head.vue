<template>
  <div class="container">
    <div class="left">
      <router-link to="/"
        ><img src="@/assets/logo.png" alt="Logo"
      /></router-link>
    </div>
    <div class="right">
      <router-link to="/">
        暂不登陆，返回主页面
        <i class="iconfont icon-arrow-double-right"></i>
      </router-link>
    </div>
  </div>
</template>

<style scoped lang="scss">
.container {
  height: 60px;
  width: auto;
  display: flex;
  background-color: #fff;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;

  .left {
    margin-left: 60px;
    img {
      height: 60px;
      width: auto;
      object-fit: contain;
    }
  }

  .right {
    margin-right: 30px;
    a {
      font-size: 16px;
      color: #000;
    }
  }
}
</style>
