<script setup>
import axios from "axios";
import { computed, onBeforeMount, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import { RegionColumns } from "v-region";
import { ElMessage } from "element-plus";
import { renderPopper } from "element-plus/lib/el-popper";
const userID = localStorage.getItem("userID");
const dataGet = ref(false);
const editform = ref();
const totalPrice = ref();
const route = useRoute();
const form = ref({
  account: "",
  userName: "",
  phone: "",
  email: "",
  birthday: "",
  address: "",
  realName: "",
});
const orderInfo = ref([]);
const orderID = route.params.orderID;
const region = ref({
  province: "",
  city: "",
  area: "",
  town: "",
});
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
const editFormRules = ref({
  phone: [
    { required: true, validator: requiredFunction, trigger: "blur" },
    {
      pattern: /(^1(3[0-9]|5[0-3,5-9]|7[1-3,5-8]|8[0-9])\d{8}$)/,
      message: "请输入正确的号码",
      trigger: "blur",
    },
  ],
  realName: [{ required: true, validator: requiredFunction, trigger: "blur" }],
  address: [
    { required: true, validator: requiredFunction, trigger: "blur" },
    { validator: regionSelect, trigger: "blur" },
  ],
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
      console.log(response);
      form.value.account = userID;
      form.value.phone = response.data.mobile;
      const regions = response.data.address.split("-");
      region.value.province = regions[0];
      region.value.city = regions[1];
      region.value.area = regions[2];
      region.value.town = regions[3];
      form.value.address = regions[4];
      dataGet.value = true;
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
const getOrderInfo = () => {
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/GetOrders",
    data: {
      orderId: orderID,
    },
  })
    .then((response) => {
      orderInfo.value = response.data.entitylist;
      totalPrice.value = orderInfo.value[0].totalPrice;
      console.log(response);
    })
    .catch(() => {
      console.log("查询失败！");
    });
};
// function change(data) {
//   console.log(data.province);
//   console.log(data);
// }
const goPay = async () => {
  await editform.value.validate((valid) => {
    if (valid) {
      //
      //支付发起的函数写在这里
      //
      //下面的删掉
      ElMessage({
        message: "支付成功",
        type: "success",
      });
    }
  });
};
onMounted(() => {
  getUserInfo();
  getOrderInfo();
});
</script>

<template>
  <div class="container">
    <h2 style="margin-top: 40px; margin-bottom: 40px">
      <i class="iconfont icon-zhifu"></i>支付订单
    </h2>
    <h4 style="margin-bottom: 40px">以下是订单号 {{ orderID }} 包含的商品</h4>
    <el-table
      :data="orderInfo"
      style="width: 80%; margin: auto"
      max-height="500px"
      stripe
      fix
      size="large"
    >
      <el-table-column
        label="下单商品"
        width="452px"
        cell-style="min-height: 50px"
      >
        <template #default="scope">
          <router-link
            :to="{ name: 'Details', params: { goodId: scope.row.productID } }"
            ><a class="good_detial">
              <img :src="scope.row.url" class="good_img" />
              <span class="good_decr" :title="scope.row.productName">{{
                scope.row.productName
              }}</span>
            </a></router-link
          >
        </template>
      </el-table-column>
      <el-table-column label="选择类型" width="150px">
        <template #default="scope">
          <span style="font-size: 15px"
            >{{ scope.row.value1 }} + {{ scope.row.value2 }}</span
          >
        </template>
      </el-table-column>
      <el-table-column
        prop="productCount"
        label="商品数量"
        width="350px"
        align="right"
      ></el-table-column>
      <!-- <el-table-column label="总价格/元" width="100px" align="right">
        <template #default="scope">
          <p
            style="position: relative; top: 10px; color: brown; font-size: 16px"
          >
            {{ scope.row.totalPrice }}
          </p>
        </template>
      </el-table-column> -->
    </el-table>
    <div class="form">
      <h4><i class="iconfont icon-address"></i>配送地址</h4>
      <el-form
        :model="form"
        label-width="100px"
        :rules="editFormRules"
        ref="editform"
        :inline="true"
      >
        <el-form-item label="收货人姓名" prop="realName">
          <el-input
            v-model="form.realName"
            autocomplete="off"
            autosize
          ></el-input>
        </el-form-item>
        <el-form-item label="电话号码" prop="phone">
          <el-input v-model="form.phone" autocomplete="off" autosize></el-input>
        </el-form-item>
        <el-form-item label="详细地址：" prop="address">
          <RegionColumns v-model="region" :town="true" v-if="dataGet" />
          <el-input
            v-model="form.address"
            autocomplete="off"
            type="textarea"
            autosize
          ></el-input>
        </el-form-item>
        <el-form-item style="margin-left: 1000px">
          <p style="font-size: 20px">总价格</p>
          <sapn v-if="dataGet" class="price">￥{{ totalPrice }}</sapn>
        </el-form-item>
        <el-form-item style="margin-left: 1000px">
          <el-button
            type="primary"
            style="height: 70px; width: auto; font-size: 20px"
            @click="goPay()"
            >使用<i class="iconfont icon-zhifubaozhifu"></i
            >支付宝支付</el-button
          >
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<style scoped lang="scss">
h2 .iconfont {
  font-size: 40px;
}
h4 .iconfont {
  font-size: 40px;
}
.good_detial {
  max-height: 80px;
  color: #6b778c;
  display: flex;
  align-items: center;
  .good_img {
    max-height: 70px;
    max-width: 70px;
    margin-right: 10px;
  }
  .good_decr {
    overflow: hidden; /*内容会被修剪，并且其余内容是不可见的。*/
    text-overflow: ellipsis; /*显示省略符号来代表被修剪的文本*/
    white-space: nowrap; /*wrap :换行*/
  }
}
.form {
  max-width: 1200px;
  margin-top: 50px;
}

.iconfont {
  font-size: 20x;
}
.iconfont::before {
  font-family: iconfont;
}
.price {
  font-size: 30px;
  color: brown;
}
</style>
