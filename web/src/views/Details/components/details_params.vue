<template>
  <main class="bg_gray">
    <div class="bg_white" style="text-align: center">
      <div class="container">
        <div class="col-lg-12" style="width: 30vw">
          <!--拿去商品信息后放在这里-->
          <div class="prod_options version_2">
            <h1 class="g-name" >{{ GoodName }}</h1>
            <p class="g-desc">{{ GoodIntroduce }}</p>
            <el-divider />
            <div class="row">
              <label
                class="col-xl-7 col-lg-5 col-md-6 col-6"
                style="font-size: large"
                ><strong>选项</strong> - 参数选择
              </label>
              <div
                class="col-xl-5 col-lg-5 col-md-6 col-6"
                style="padding-bottom: 10px"
              >
                <a href="#0" data-bs-toggle="modal" data-bs-target="#size-modal"
                  ><i class="iconfont icon-arrow-double-right"></i
                ></a>
              </div>
            </div>
            <el-divider />
            <div class="row">
              <label
                class="col-xl-7 col-lg-5 col-md-6 col-6"
                style="font-size: large"
                ><strong>选项</strong> - 大小选择
              </label>
              <div
                class="col-xl-5 col-lg-5 col-md-6 col-6"
                style="padding-bottom: 10px"
              >
                <a href="#0" data-bs-toggle="modal" data-bs-target="#size-modal"
                  ><i class="iconfont icon-arrow-double-right"></i
                ></a>
              </div>
            </div>
            <el-divider />
            <div class="row">
              <label
                class="col-xl-7 col-lg-5 col-md-6 col-6"
                style="font-size: large"
                ><strong>数量</strong></label
              >
              <el-input-number v-model="num" :min="1" :max="99" @change="handleChange" />
            </div>
            <el-divider />
            <div class="row mt-3">
              <div class="col-lg-7 col-md-6">
                <div class="price_main">
                  <span class="new_price">${{price}}</span
                  ><span class="percentage">-20%</span>
                  <span class="old_price">${{props.GoodPrice}}</span>
                </div>
              </div>
              <div class="col-lg-5 col-md-6">
                <div class="btn_add_to_cart">
                  <a href="#0" class="btn_1">添加到购物车</a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <!-- /row -->
    </div>
  </main>
  <!--选项表-->
  <div
    class="modal fade"
    tabindex="-1"
    role="dialog"
    aria-labelledby="size-modal"
    id="size-modal"
    aria-hidden="true"
  >
    <div class="modal-dialog modal-dialog-centered modal-lg">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title">参数选择</h5>
          <button
            type="button"
            class="btn-close"
            data-bs-dismiss="modal"
            aria-label="Close"
          ></button>
        </div>
        <div class="modal-body" style="font-size: x-small;">
          <el-row :gutter="12" >
            <el-col
              v-for="(item, index) in goods"
              :key="index"
              :span="8"
            >
              <el-card :body-style="{ padding: '0px' }">
                <img
                style="object-fit: contain;width:auto;height: 210px;"
                  :src="item.image"
                  class="image"
                />
                <div style="padding: 12px">
                  <span>{{item.introduce}}</span>
                  <div class="bottom">
                    <time class="time">{{ GoodId }}</time>
                    <el-button text class="button">确认</el-button>
                  </div>
                </div>
              </el-card>
            </el-col>
          </el-row>
          <!-- /table -->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import  axios  from 'axios'
import { defineProps, onMounted, watch } from 'vue';

const props = defineProps({
  GoodName: String,
  GoodIntroduce: String,
  GoodPrice: String,
  GoodId: String
});

// 处理点击事件
const handleChange = (num) => {
  console.log(num);
};

onMounted(() => {
  console.log(props.GoodId);
});

watch(() => props.GoodId, (newId, oldId) => {
  console.log(newId);

  axios({
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded',
    },
    method: 'post',
    url: '/api/details/info',
    data: {
      productID: newId,
    },
  })
    .then((response) => {
      const goods = response.data;
      console.log(goods);
      // 执行其他逻辑
    })
    .catch((error) => {
      console.error(error);
    });
});
</script>

<style lang="scss" scoped>
@import "@/assets/variables.scss";

.g-name {
  font-size: 35px;
}

.g-desc {
  color: #999;
  margin-top: 10px;
}

.iconfont::before {
  font-family: iconfont;
}

.iconfont {
  font-size: xx-large;
}
</style>
