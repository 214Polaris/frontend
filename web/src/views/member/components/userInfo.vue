<script setup>
import { onMounted, ref } from "vue";

const addressDialogVisible = ref(false);
const infoDialogVisible = ref(false);
const form = ref({
  name: "",
  phone: "",
  email: "",
  birthday: "",
  address: "",
});

onMounted(() => {});
</script>

<template>
  <div class="home-overview">
    <!-- 用户信息 -->
    <div class="user-meta">
      <div class="avatar">
        <img src="" />
      </div>
      <h4>用户名</h4>
    </div>
    <div class="item">
      <a href="javascript:;">
        <span
          class="iconfont icon-gerenxinxi1"
          @click="
            () => {
              infoDialogVisible = true;
            }
          "
        ></span>
        <p>个人信息</p>
      </a>
      <a href="javascript:;">
        <span class="iconfont icon-setting"></span>
        <p>安全设置</p>
      </a>
      <a
        href="javascript:;"
        @click="
          () => {
            addressDialogVisible = true;
          }
        "
      >
        <span class="iconfont icon-address"></span>
        <p>地址管理</p>
      </a>
    </div>
  </div>
  <div class="like-container">
    <div class="home-panel">
      <div class="header">
        <h4 data-v-bcb266e0="">猜你喜欢</h4>
      </div>
    </div>
  </div>
  <el-dialog v-model="infoDialogVisible" title="修改个人信息">
    <el-form v-model="form" label-width="100px">
      <el-form-item label="账号：">
        <el-input v-model="form.name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="用户名：">
        <el-input v-model.number="form.phone"></el-input>
      </el-form-item>
      <el-form-item label="邮箱：">
        <el-input v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="出生日期：">
        <el-date-picker
          v-model="form.birthday"
          type="date"
          placeholder="Select date and time"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="详细地址：">
        <el-input
          v-model="form.address"
          autocomplete="off"
          type="textarea"
          autosize
        ></el-input>
      </el-form-item>
    </el-form>
  </el-dialog>
  <el-dialog v-model="addressDialogVisible" title="修改地址">
    <el-form v-model="form" label-width="100px">
      <el-form-item label="姓名：">
        <el-input v-model="form.name" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="电话：">
        <el-input v-model.number="form.phone"></el-input>
      </el-form-item>
      <el-form-item label="详细地址：">
        <el-input
          v-model="form.address"
          autocomplete="off"
          type="textarea"
          autosize
        ></el-input>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<style scoped>
.home-overview {
  height: 132px;
  background: url("../../../assets/memberBackground.jpg") no-repeat center /
    cover;
  display: flex;
}
.home-overview .user-meta {
  flex: 1;
  display: flex;
  align-items: center;
}
.home-overview .user-meta .avatar {
  width: 85px;
  height: 85px;
  border-radius: 50%;
  overflow: hidden;
  margin-left: 60px;
}
.home-overview .user-meta img {
  width: 100%;
  height: 100%;
}

.home-overview .user-meta h4 {
  padding-left: 26px;
  font-size: 18px;
  font-weight: normal;
  color: white;
}

.home-overview .item {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-around;
}
.home-overview .item:first-child {
  border-right: 1px solid #f4f4f4;
}

.home-overview .item a {
  color: white;
  font-size: 16px;
  text-align: center;
}
.home-overview .item a .iconfont {
  font-size: 36px;
}

.home-overview .item a p {
  line-height: 32px;
}

.like-container {
  margin-top: 20px;
  border-radius: 4px;
  background-color: #fff;
}

.home-panel {
  background-color: #fff;
  padding: 0 20px;
  margin-top: 20px;
  height: 400px;
}
.home-panel .header {
  height: 66px;
  border-bottom: 1px solid #f5f5f5;
  padding: 18px 0;
  display: flex;
  justify-content: space-between;
  align-items: baseline;
}
.home-panel h4 {
  font-size: 22px;
  font-weight: 400;
}

.home-panel .goods-list {
  display: flex;
  justify-content: space-around;
}

/* .el-input {
  width: 300px;
} */
</style>
