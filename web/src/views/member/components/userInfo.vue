<script setup>
import { onMounted, ref, h } from "vue";
import list from "@/static/goodsList.json";
import { ElMessage, ElMessageBox } from "element-plus";
import axios from "axios";

const addressDialogVisible = ref(false);
const infoDialogVisible = ref(false);
const passwordDialogVisible = ref(false);
const userID = localStorage.getItem("userID");
const form = ref({
  account: "",
  userName: "",
  phone: "",
  email: "",
  birthday: "",
  address: "",
  realName: "",
});
const keyOfCheck = [
  "account",
  "userName",
  "phone",
  "email",
  "birthday",
  "address",
  "realName",
];
const passForm = ref({
  old_password: "",
  new_password: "",
  confirm_password: "",
});
const goodList = ref([]);

const getUserInfo = () => {
  console.log(userID);
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/showUser",
    data: {
      id: userID,
    },
  })
    .then((response) => {
      if (response.data.code === 200) {
        form.value[mobile] = response.data.mobile;
        form.value[address] = response.data.address;
        form.value[userName] = response.data.userName;
        form.value[email] = response.data.userEmail;
        form.value[birthday] = response.data.userBornDate;
        form.value[realName] = response.data.userRealName;
      }
    })
    .catch(() => {
      console.log("查询失败！");
      ElMessage.error("Oops, this is a error message.");
    });
};

const getgoodList = () => {
  goodList.value = list;
};

//确认提交信息函数
const infoConfirm = () => {
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/Edit",
    data: {
      id: parseInt(userID),
      mobile: form.value.phone,
      address: form.value.address,
      name: form.value.userName,
      email: form.value.email,
      bornDate: form.value.birthday,
      realname: form.value.realName,
    },
  })
    .then(() => {
      ElMessage({
        message: "修改成功！",
        type: "success",
      });
    })
    .catch(function () {
      ElMessage.error("修改失败！");
    });
};

//确认提交密码修改函数
const passConfirm = () => {};

const is_new = () => {
  var _new = false;
  for (const key of keyOfCheck) {
    if (
      form.value[key] == null ||
      form.value[key] == undefined ||
      !form.value[key]
    ) {
      _new = true;
    }
  }
  if (_new) {
    ElMessage({
      message: "请尽快完善您的个人信息，帮助我们更好为您服务",
      type: "warning",
    });
  }
};

onMounted(() => {
  getUserInfo();
  getgoodList();
  is_new();
});
</script>

with
<template>
  <div class="home-overview">
    <!-- 用户信息 -->
    <div class="user-meta">
      <div class="avatar">
        <img src="" />
      </div>
      <h4>{{ form.userName }}</h4>
    </div>
    <div class="item">
      <a href="javascript:;">
        <span
          class="iconfont icon-gerenxinxi1"
          @click="
            () => {
              infoDialogVisible = true;
              getUserInfo();
            }
          "
        ></span>
        <p>个人信息</p>
      </a>
      <a href="javascript:;">
        <span
          class="iconfont icon-setting"
          @click="
            () => {
              passwordDialogVisible = true;
            }
          "
        ></span>
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
        <span
          class="iconfont icon-address"
          @click="
            () => {
              addressDialogVisible = true;
              getUserInfo();
            }
          "
        ></span>
        <p>地址管理</p>
      </a>
    </div>
  </div>
  <!-- 修改个人信息窗口 -->
  <el-dialog v-model="infoDialogVisible" title="修改个人信息">
    <el-form v-model="form" label-width="100px">
      <el-form-item label="用户名：">
        <el-input v-model="form.userName"></el-input>
      </el-form-item>
      <el-form-item label="电话号码：">
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
      <el-form-item style="margin-left: 310px">
        <el-button type="primary" @click="infoConfirm()">提交修改</el-button>
        <el-button
          @click="
            () => {
              infoDialogVisible = false;
            }
          "
          >取消</el-button
        >
      </el-form-item>
    </el-form>
    <!-- 修改密码窗口 -->
  </el-dialog>
  <el-dialog v-model="passwordDialogVisible" title="修改密码">
    <el-form v-model="form" label-width="100px">
      <el-form-item label="原密码：">
        <el-input v-model="passForm.old_password" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="新密码：">
        <el-input
          v-model="passForm.new_password"
          autocomplete="off"
          type="password"
        ></el-input>
      </el-form-item>
      <el-form-item label="确认密码：">
        <el-input
          v-model="passForm.confirm_password"
          autocomplete="off"
          type="password"
          autosize
        ></el-input>
      </el-form-item>
      <el-form-item style="margin-left: 310px">
        <el-button type="primary" @click="passConfirm()">提交修改</el-button>
        <el-button
          @click="
            () => {
              passwordDialogVisible = false;
            }
          "
          >取消</el-button
        >
      </el-form-item>
    </el-form>
    <!-- 修改地址窗口 -->
  </el-dialog>
  <el-dialog v-model="addressDialogVisible" title="修改地址">
    <el-form v-model="form" label-width="100px">
      <el-form-item label="收货姓名：">
        <el-input v-model="form.realName" autocomplete="off"></el-input>
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
      <el-form-item style="margin-left: 310px">
        <el-button type="primary" @click="infoConfirm()">提交修改</el-button>
        <el-button
          @click="
            () => {
              addressDialogVisible = false;
            }
          "
          >取消</el-button
        >
      </el-form-item>
    </el-form>
  </el-dialog>
  <div class="like-container">
    <div class="home-panel">
      <div class="header">
        <h4>猜你喜欢</h4>
      </div>
      <!-- <template> -->
      <div class="total">
        <div class="layer_good">
          <div
            class="recommend"
            v-for="(item, index) in goodList.slice(0, 4)"
            :key="index"
          >
            <router-link :to="{ name: 'Details', params: { goodId: item.id } }">
              <ul>
                <li class="img_li">
                  <img :src="item.image" class="item_img" />
                </li>
                <li class="introduce">{{ item.introduce }}</li>
                <li class="price">￥{{ item.price }}</li>
              </ul>
            </router-link>
          </div>
        </div>
      </div>
      <!-- </template> -->
    </div>
  </div>
</template>

<style scoped lang="scss">
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

.home-overview .item a .iconfont::before {
  font-family: iconfont;
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

.total {
  display: flex;
  flex-direction: column;
  flex-wrap: wrap;
}

.layer_good {
  display: flex;
  flex-grow: 1;
  margin-bottom: 64px;
  align-content: stretch;
  padding-left: 12px;
}

.recommend {
  margin: 8px;
  display: inline-block;
}

li {
  list-style: none;
}

.recommend img {
  height: 300px;
  text-align: center;
}

.introduce {
  color: black;
  width: 170px;
  text-align: center;
  font-size: 12px;
  margin-top: 6px;
  margin-bottom: 6px;
}

.price {
  color: black;
  width: 150px;
  text-align: center;
  font-size: 18px;
  color: #f00;
}
ul {
  width: 220px;
  height: auto;
  overflow: hidden;
  padding: 0;
  padding-right: 15px;
}
.img_li {
  width: 220px;
  height: 220px;
  overflow: hidden;
  .item_img {
    position: relative;
    margin: auto;
    max-width: 200px;
    max-height: 200px;
    object-fit: cover;
    box-shadow: 5px 5px 5px rgba(57, 56, 56, 0.5);
  }
}
</style>
