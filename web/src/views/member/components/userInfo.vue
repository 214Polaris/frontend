<script setup>
import { onMounted, ref } from "vue";
import list from "@/static/goodsList.json";
import { ElMessage } from "element-plus";
import { RegionColumns } from "v-region";

import axios from "axios";

const addressDialogVisible = ref(false);
const infoDialogVisible = ref(false);
const passwordDialogVisible = ref(false);
const userID = localStorage.getItem("userID");
const editform = ref(null);
const passf = ref(null);
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

const region = ref({
  province: "350000",
  city: "350100",
  area: "350104",
  town: "350104008",
});

const getUserInfo = () => {
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
      form.value.account = userID;
      form.value.phone = response.data.mobile;
      const regions = response.data.address.split("-");
      region.value.province = regions[0];
      region.value.city = regions[1];
      region.value.area = regions[2];
      region.value.town = regions[3];
      form.value.address = regions[4];
      form.value.userName = response.data.userName;
      form.value.email = response.data.userEmail;
      form.value.birthday = response.data.userBornDate;
      form.value.realName = response.data.userRealName;
      return;
    })
    .catch(() => {
      console.log("查询失败！");
      ElMessage.error("查询个人信息出错，请检查服务器状态");
    });
};

const getRandom = (max) => {
  var arr = [];
  for (var i = 0; i < 4; i++) {
    arr[i] = parseInt(Math.random() * (max + 1), 10);
    for (var j = 0; j < i; j++) {
      //如果重复则 i-- 重新产生一个
      if (arr[i] == arr[j]) {
        i--;
        break;
      }
    }
  }
  return arr;
};

const getgoodList = () => {
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "get",
    url: "/api/goodslist",
  })
    .then((response) => {
      const len = response.data.length - 1;
      const arr = getRandom(len);
      for (var i = 0; i < 4; i++) {
        goodList.value[i] = response.data[arr[i]];
      }
      return;
    })
    .catch(() => {
      ElMessage.error("获取失败！");
    });
};

//验证回调函数
const requiredFunction = (rule, value, callback) => {
  if (value == "") {
    return callback(new Error("此项为必填项,请输入"));
  }
  return callback();
};

const regionSelect = (rule, value, callback) => {
  if (region.value.town == undefined) {
    callback(new Error("请选择地区"));
  }
  return callback();
};

const validatePass2 = (rule, value, callback) => {
  if (value === "") {
    callback(new Error("请再次输入密码"));
  } else if (value !== passForm.value.new_password) {
    callback(new Error("两次输入密码不一致!"));
  } else {
    callback();
  }
};

//表单验证规则
const editFormRules = ref({
  phone: [
    { required: true, validator: requiredFunction, trigger: "blur" },
    {
      pattern: /(^1(3[0-9]|5[0-3,5-9]|7[1-3,5-8]|8[0-9])\d{8}$)/,
      message: "请输入正确的号码",
      trigger: "blur",
    },
  ],
  userName: [
    { required: true, validator: requiredFunction, trigger: "blur" },
    { min: 5, message: "用户名应在5字符以上", trigger: "blur" },
  ],
  email: [
    { required: true, validator: requiredFunction, trigger: "blur" },
    { type: "email", message: "请输入正确的邮箱地址", trigger: "blur" },
  ],
  realName: [{ required: true, validator: requiredFunction, trigger: "blur" }],
  address: [
    { required: true, validator: requiredFunction, trigger: "blur" },
    { validator: regionSelect, trigger: "blur" },
  ],
  old_password: [{ required: true, message: "请输入原密码", trigger: "blur" }],
  new_password: [{ required: true, message: "请输入新密码", trigger: "blur" }],
  confirm_password: [
    { required: true, validator: validatePass2, trigger: "blur" },
  ],
});

//确认提交信息函数
const infoConfirm = async () => {
  await editform.value.validate((valid) => {
    if (valid) {
      axios({
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        method: "post",
        url: "/api/Edit",
        data: {
          id: parseInt(userID),
          mobile: form.value.phone,
          address:
            region.value.province +
            "-" +
            region.value.city +
            "-" +
            region.value.area +
            "-" +
            region.value.town +
            "-" +
            form.value.address,
          name: form.value.userName,
          email: form.value.email,
          bornDate: form.value.birthday,
          realName: form.value.realName,
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
      console.log("submit");
    } else {
      console.log("error submit");
      return false;
    }
  });
};

//确认提交密码修改函数
const passConfirm = async () => {
  if (
    passForm.value.old_password === passForm.value.new_password &&
    passForm.value.old_password != ""
  ) {
    ElMessage.error("新密码请不要和原密码相同！");
    console.log("error");
    return;
  }
  await passf.value.validate((valid) => {
    if (valid) {
      axios({
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
        },
        method: "post",
        url: "/api/EditPass",
        data: {
          id: userID,
          oldPass: passForm.value.old_password,
          newPass: passForm.value.new_password,
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
    } else {
      console.log("error submit");
      return false;
    }
  });
};

// const is_new = () => {
//   var _new = false;
//   for (const key of keyOfCheck) {
//     if (
//       form.value[key] == null ||
//       form.value[key] == undefined ||
//       !form.value[key]
//     ) {
//       console.log(form.key);
//       console.log("数据不全");
//       _new = true;
//     }
//   }
//   if (_new) {
//     ElMessage({
//       message: "请尽快完善您的个人信息，帮助我们更好为您服务",
//       type: "warning",
//     });
//   }
// };

const editDialogClosed = () => {
  passForm.value.old_password = "";
  passForm.value.new_password = "";
  passForm.value.confirm_password = "";
  passf.value.resetFields();
};
const editDialogClosed_info = () => {
  editform.value.resetFields();
};

function change(data) {
  console.log(data.province);
  console.log(data);
}

onMounted(() => {
  getUserInfo();
  getgoodList();
});
</script>

<template>
  <div class="home-overview">
    <!-- 用户信息 -->
    <div class="user-meta">
      <div class="avatar">
        <img src="../../../assets/user.jpg" />
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
  <el-dialog
    v-model="infoDialogVisible"
    title="修改个人信息"
    @close="editDialogClosed_info()"
  >
    <el-form
      :model="form"
      label-width="100px"
      :rules="editFormRules"
      ref="editform"
    >
      <el-form-item label="用户名：" prop="userName">
        <el-input v-model="form.userName"></el-input>
      </el-form-item>
      <el-form-item label="电话号码：" prop="phone">
        <el-input v-model.number="form.phone"></el-input>
      </el-form-item>
      <el-form-item label="邮箱：" prop="email">
        <el-input v-model="form.email"></el-input>
      </el-form-item>
      <el-form-item label="出生日期：">
        <el-date-picker
          v-model="form.birthday"
          type="date"
          placeholder="选择出生日期"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="详细地址：" prop="address">
        <RegionColumns v-model="region" @change="change" :town="true" />
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
  <el-dialog
    v-model="passwordDialogVisible"
    title="修改密码"
    @close="editDialogClosed()"
  >
    <el-form
      :model="passForm"
      label-width="100px"
      :rules="editFormRules"
      ref="passf"
    >
      <el-form-item label="原密码：" prop="old_password">
        <el-input
          v-model="passForm.old_password"
          autocomplete="off"
          type="password"
          show-password
        ></el-input>
      </el-form-item>
      <el-form-item label="新密码：" prop="new_password">
        <el-input
          v-model="passForm.new_password"
          autocomplete="off"
          type="password"
          show-password
        ></el-input>
      </el-form-item>
      <el-form-item label="确认密码：" prop="confirm_password">
        <el-input
          v-model="passForm.confirm_password"
          autocomplete="off"
          type="password"
          autosize
          show-password
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
  <el-dialog
    v-model="addressDialogVisible"
    title="修改地址"
    @close="editDialogClosed_info()"
  >
    <el-form
      :model="form"
      label-width="100px"
      :rules="editFormRules"
      ref="editform"
    >
      <el-form-item label="收货姓名：" prop="realName">
        <el-input v-model="form.realName" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item label="收货电话：" prop="phone">
        <el-input v-model.number="form.phone"></el-input>
      </el-form-item>
      <el-form-item label="详细地址：" prop="address">
        <RegionColumns v-model="region" @change="change" :town="true" />
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
          <div class="recommend" v-for="(item, index) in goodList" :key="index">
            <router-link
              :to="{ name: 'Details', params: { goodId: item.productId } }"
            >
              <ul>
                <li class="img_li">
                  <img :src="item.productLink" class="item_img" />
                </li>
                <li class="introduce">{{ item.productName }}</li>
                <li class="price">￥{{ item.productPrice }}</li>
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
