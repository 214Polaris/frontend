<template>
  <head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1, shrink-to-fit=no"
    />
    <meta name="description" content="" />
    <meta name="author" content="Ansonika" />
    <title>我的购物车</title>
  </head>

  <body>
    <div id="page">
      <main class="bg_gray">
        <div class="container margin_30">
          <div class="page_header">
            <h1>我的购物车</h1>
          </div>
          <!-- /page_header -->
          <table class="table table-striped cart-list">
            <thead>
              <tr>
                <th style="width: 50px; vertical-align: center">
                  <el-checkbox v-model="selectAll" @change="selectAllItems"
                    >全选</el-checkbox
                  >
                </th>
                <th style="width: 50%; text-align: center">商品</th>
                <th>单价</th>
                <th style="text-align: center; width: 10%">数量</th>
                <th style="padding-left: 5vw; width: 20%">总价</th>
              </tr>
            </thead>
            <tbody>
              <tr v-if="cartlist.length === 0">
                <td colspan="6" class="text-center">购物车为空</td>
              </tr>
              <tr
                v-else
                v-bind="item"
                v-for="(item, index) in cartlist"
                :key="index"
              >
                <td align="center">
                  <!-- 添加选择框 -->
                  <el-checkbox v-model="item.selected"></el-checkbox>
                </td>
                <td>
                  <router-link
                    :to="{
                      name: 'Details',
                      params: { goodId: item.productID },
                    }"
                  >
                    <div class="thumb_cart">
                      <img
                        v-lazy="item.productLink"
                        :data-src="item.productLink"
                        class="lazy"
                        alt="Image"
                      />
                    </div>
                    <span
                      class="item_cart"
                      style="
                        color: black;
                        margin-top: 5px;
                        width: 70%;
                        white-space: nowrap;
                        overflow: hidden;
                        text-overflow: ellipsis;
                      "
                      :title="item.productName"
                      >{{ item.productName }} <br /><br />
                      <div style="text-align: left">
                        {{ item.value1 }} + {{ item.value2 }}
                      </div>
                    </span>
                  </router-link>
                </td>
                <td>
                  <strong>{{ item.productPrice }}</strong>
                </td>
                <td>
                  <el-input-number
                    size="small"
                    v-model="item.productCount"
                    :min="1"
                    :id="'quantity_' + item.productCount"
                    :name="'quantity_' + item.productCount"
                  ></el-input-number>
                </td>
                <td style="padding-left: 5.5vw">
                  <strong style="color: red">{{
                    item.productPrice * item.productCount
                  }}</strong>
                </td>
                <td class="options">
                  <button @click="deleteItem(index)">
                    <i class="ti-trash"></i>
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
          <div class="row add_top_30 flex-sm-row-reverse cart_actions"></div>
        </div>
        <!-- /container -->
        <div class="box_cart fixed-bottom">
          <div class="box_cart">
            <div class="container">
              <div class="row justify-content-end">
                <div class="col-xl-4 col-lg-3 col-md-3">
                  <ul>
                    <li><span>已选商品数：</span>{{ selectedItemCount }}</li>
                    <li><span>总计</span> ${{ calculateTotal() }}</li>
                  </ul>
                  <!-- 添加提交支付按钮 -->
                  <button
                    type="button"
                    class="btn_1 full-width cart"
                    @click="submitPayment"
                  >
                    确认支付
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- /box_cart -->
      </main>
      <!--/main-->
    </div>
  </body>
</template>

<script setup>
import axios from "axios";
import { onMounted, ref, computed } from "vue";
import { useRouter } from "vue-router";
let cartlist = ref([]);
const router = useRouter();
//获取购物车内容
onMounted(() => {
  let id = localStorage.getItem("userID");
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/cart",
    data: {
      userID: id,
    },
  })
    .then((response) => {
      cartlist.value = response.data;
      console.log(cartlist);
    })
    .catch((error) => {
      console.error(error);
    });
});

// 计算总费用
function calculateTotal() {
  let total = 0;
  for (const item of cartlist.value) {
    if (item.selected) {
      total += item.productPrice * item.productCount;
    }
  }
  return total.toFixed(2); // 保留两位小数
}

// 提交支付
function submitPayment() {
  const selectedItems = cartlist.value.filter((item) => item.selected);
  // 计算总价
  let totalprice = calculateTotal();
  //已选择商品的 cartID
  let cartIds = selectedItems.map((item) => (item.cartID));
  // const formattedItems = selectedItems.map((item) => {
  //     return {
  //         cartId: item.cartID,
  //     };
  // });
  //加项目
  for (const item of selectedItems) {
    item.TotalPrice = totalprice;
    item.userAddress = "";
    item.mobile = "";
  }
  // 提交创建订单
  axios({
    headers: {
      "Content-Type": "application/json",
    },
    method: "post",
    url: "/api/createorder",
    data: selectedItems,
  })
    .then((response) => {
      console.log("Payment successful:", response.data);
      let orderId = response.data.orderId;
      console.log(orderId);
      router.push({ name: "checkout", params: { orderID: orderId } });
    })
    .catch((error) => {
      console.error("Payment error:", error);
      // Handle error response
    });
    console.log(cartIds)
    //清除购物车的 api
    axios({
    headers: {
      "Content-Type": "application/json",
    },
    method: "post",
    url: "/api/remove",
    data: {cartId:cartIds}
  })
    .then((response) => {
      console.log("Payment successful:", response.data);
    })
    .catch((error) => {
      console.error("Payment error:", error);
      // Handle error response
    });
}

// 计算已选商品的数量
const selectedItemCount = computed(() => {
  let count = 0;
  for (const item of cartlist.value) {
    if (item.selected) {
      count += item.productCount;
    }
  }
  return count;
});
//单个删除
function deleteItem(index) {
  const itemToDelete = cartlist.value[index];
  let cartid = itemToDelete.cartID;
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/delete",
    data: {
      cartID: cartid,
    },
  })
    .then(() => {
      console.log("Deleting item:", itemToDelete);
      cartlist.value.splice(index, 1);
    })
    .catch((error) => {
      console.error(error);
    });
  location.reload();
  this.$router.go(0);
}

//全选按钮
let selectAll = ref(false);
function selectAllItems() {
  for (const item of cartlist.value) {
    item.selected = selectAll.value;
  }
}
</script>

<style>
.box_cart {
  position: fixed;
  left: 0;
  right: 0;
  bottom: 0;
  padding: 10px 0;
  background-color: #fff;
  z-index: 999;
  /* 确保在其他元素上方显示 */
}
</style>
