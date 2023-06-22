<script setup>
import { onMounted, ref } from "vue";
import axios from "axios";
import ol from "@/static/orderList.json";
import { ElMessage } from "element-plus";
const orderList = ref([]);
const one_orderList = ref([]);
const userID = localStorage.getItem("userID");
const tabTypes = [
  { name: "all", label: "全部订单" },
  { name: "unpay", label: "待付款" },
  { name: "deliver", label: "待发货" },
  { name: "receive", label: "待收货" },
  { name: "comment", label: "待评价" },
  { name: "complete", label: "已完成" },
  { name: "cancel", label: "已取消" },
];

const getOrderList = () => {
  orderList.value = ol;
  // axios({
  //   headers: {
  //     "Content-Type": "application/x-www-form-urlencoded",
  //   },
  //   method: "post",
  //   url: "/api/getuserorders",
  //   data: {
  //     id: userID,
  //   },
  // })
  //   .then((response) => {
  //     console.log(response);
  //   })
  //   .catch(() => {
  //     ElMessage.error("无订单");
  //   });
};

const getOneOrder = () => {
  axios({
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    method: "post",
    url: "/api/getuserorders",
    data: {
      id: userID,
    },
  })
    .then((response) => {
      console.log(response);
    })
    .catch(() => {
      ElMessage.error("无订单");
    });
};

//根据订单类型返回状态描述
const displayState = (state) => {
  return tabTypes[state].label;
};

const orderPay = () => {};

//根据不同类型选择，显示不同订单
const selectList = (state) => {
  var t_list = [];
  if (state == 0) {
    t_list = orderList.value;
    return t_list;
  }
  for (var i = 0; i < orderList.value.length; i++) {
    if (parseInt(orderList.value[i].state) == state) {
      t_list.push(orderList.value[i]);
    }
  }
  return t_list;
};

onMounted(() => {
  getOrderList();
});
</script>

<template>
  <el-tabs class="demo-tabs" stretch>
    <el-tab-pane
      v-for="(tab, index) in tabTypes"
      :key="tab.name"
      :label="tab.label"
      class="el-tabs__content"
    >
      <el-table
        :data="selectList(index)"
        style="width: 100%"
        max-height="500px"
        stripe
        lazy
        v-loading="loading"
      >
        <!-- <el-table-column
          fixed
          prop="order_id"
          label="订单编号"
          width="150px"
        ></el-table-column> -->
        <template #empty>
          <el-empty></el-empty>
        </template>
        <el-table-column
          label="下单商品"
          width="300px"
          cell-style="min-height: 50px"
        >
          <template #default="scope">
            <router-link
              :to="{ name: 'Details', params: { goodId: scope.row.good_id } }"
              ><a class="good_detial">
                <img :src="scope.row.good_img" class="good_img" />
                <span class="good_decr" :title="scope.row.good_name">{{
                  scope.row.good_name
                }}</span>
              </a></router-link
            >
          </template>
        </el-table-column>
        ><el-table-column
          prop="good_type"
          label="选择类型"
          width="80px"
        ></el-table-column>
        <el-table-column
          prop="good_num"
          label="商品数量"
          width="120px"
        ></el-table-column>
        <el-table-column
          prop="total_price"
          label="总价格"
          width="100px"
        ></el-table-column
        ><el-table-column
          prop="create_time"
          label="下单时间"
          width="150px"
        ></el-table-column>
        <el-table-column
          label="订单状态"
          width="150px"
          v-if="tab.label == '全部订单'"
        >
          <template #default="scope">
            <p style="position: relative; top: 10px; color: brown">
              {{ displayState(scope.row.state) }}
            </p>
          </template>
        </el-table-column>
        <el-table-column
          label="操作"
          width="150px"
          v-if="tab.label != '全部订单'"
        >
          <template #default>
            <div class="btnGroup">
              <a
                v-if="tab.label == '待付款'"
                href="javascript:"
                class="go_to_pay"
                @click="orderPay()"
                >去付款</a
              >
              <a
                v-if="tab.label == '待发货'"
                href="javascript:"
                class="go_to_pay"
                >提醒发货</a
              >
              <a
                v-if="tab.label == '待收货'"
                href="javascript:"
                class="go_to_pay"
                >查看物流</a
              >
              <a
                v-if="tab.label == '待评价'"
                href="javascript:"
                class="go_to_pay"
                >去评价</a
              >
              <a
                v-if="tab.label == '已完成'"
                href="javascript:"
                class="go_to_pay"
                >售后服务</a
              >
              <a
                v-if="tab.label == '已取消'"
                href="javascript:"
                class="go_to_pay"
                >重新购买</a
              >
            </div>
          </template>
        </el-table-column>
      </el-table>
    </el-tab-pane>
  </el-tabs>
</template>

<style scoped lang="scss">
.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.good_detial {
  max-height: 44px;
  color: #6b778c;
  display: flex;
  align-items: center;
  .good_img {
    max-height: 50px;
    max-width: 50px;
  }
  .good_decr {
    overflow: hidden; /*内容会被修剪，并且其余内容是不可见的。*/
    text-overflow: ellipsis; /*显示省略符号来代表被修剪的文本*/
    white-space: nowrap; /*wrap :换行*/
  }
}

.go_to_pay {
  color: #66b1ff;
}
</style>
