<script setup>
import axios from "axios";
import { ref } from "vue";
import { useRoute } from "vue-router";
const userID = localStorage.getItem("userID");
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
</script>

<template>
  <h1>支付订单</h1>
</template>
