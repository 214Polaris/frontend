<template>
  <div class="container">
    <div class="common-layout">
      <el-container>
        <div class="goods-info">
          <!-- 商品图片组件 -->
            <div class="media" style="display: flex;">
              <!--<GoodsImage :image="good.productLink" />-->
              <GoodsImage
                image="https://g-search1.alicdn.com/img/bao/uploaded/i4/i4/38522192/O1CN014WSqOl1S3yk3s1IxK_!!38522192.jpg_460x460Q90.jpg_.webp"
              />
              <div style="display: flex; padding-left: 10vw; font-size: large;">
              <GoodsParams 
              GoodName="good.productName"
              GoodIntroduce="good.productIntro"
              GoodPrice = "182"
              GoodId="good_id"
              />
              </div>
            </div>
          <!--商品信息组件-->    
        </div>
      </el-container>
    </div>
    <div class="spec">
      <!-- <GoodsIntroduce
                :GoodName="good.productName"
                :GoodIntroduce="good.productIntro"
                :GoodPrice="good.productPrice"
              /> -->
      <GoodsIntroduce
        GoodName="good.productName"
        GoodIntroduce="good.productIntro"
        GoodPrice="good.productPrice"
      />
    </div>
  </div>
</template>
<script>
import axios from "axios";
import GoodsImage from "./components/details_image.vue";
import GoodsIntroduce from "./components/details_introduce.vue";
import GoodsParams from "./components/details_params.vue";
export default {
  data() {
    return {
      //从上个页面获取点击的商品id
      good_id: this.$route.params.goodId,
      good: null,
    };
  },
  //向后端请求返回对应商品
  mounted() {
    axios({
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
      method: "post",
      url: "/api/details",
      data: {
        good_id: this.good_id,
      },
    })
      .then((response) => {
        this.good = response.data;
      })
      .catch((error) => {
        console.error(error);
      });
  },
  components: {
    GoodsIntroduce,
    GoodsImage,
    GoodsParams,
  },
};
</script>
<style scoped lang="scss">
@import "@/assets/variables.scss";

.goods-info {
  width: 100vw;
  background: #fff;
}

.media {

}

.spec {
  flex: 1;
}

// .container {
//   display: flex;
//   justify-content: center;
//   align-items: center;
// }
</style>
