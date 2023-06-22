<!--商品信息介绍-->
<template>
  <main class="bg_gray">
    <div class="tabs_product bg_white version_2">
      <div class="container">
        <ul class="nav nav-tabs" role="tablist">
          <li class="nav-item">
            <a
              id="tab-A"
              href="#pane-A"
              class="nav-link active"
              data-bs-toggle="tab"
              role="tab"
              >Description</a
            >
          </li>
        </ul>
      </div>
    </div>
    <!-- 简述与评论模块 -->
    <div class="tab_content_wrapper">
      <div class="container">
        <div class="tab-content" role="tablist">
          <div
            id="pane-A"
            class="card tab-pane fade active show"
            role="tabpanel"
            aria-labelledby="tab-A"
          >
            <div class="card-header" role="tab" id="heading-A">
              <h5 class="mb-0">
                <a
                  class="collapsed"
                  data-bs-toggle="collapse"
                  href="#collapse-A"
                  aria-expanded="false"
                  aria-controls="collapse-A"
                >
                  Description
                </a>
              </h5>
            </div>
            <div
              id="collapse-A"
              class="collapse"
              role="tabpanel"
              aria-labelledby="heading-A"
            >
              <div class="card-body">
                <div
                  class="row justify-content-center"
                  style="text-align: center"
                >
                  <h1>商品详情</h1>
                  <el-divider style="left: 0" />
                  <div v-for="(item, index) in dataWithPic" :key="index">
                    <img :src="item.image.URL" style="width: 65%" />
                  </div>
                  <el-divider style="left: 0" />
                  <p style="width: 700px">
                    商家详情页（含主图）以图片或文字形式标注的一口价、促销价、优惠价等价格可能是在使用优惠券、满减或特定优惠活动和时段等情形下的价格，具体请以结算页面的标价、优惠条件或活动规则为准。
                  </p>
                  <!-- /table-responsive -->
                </div>
              </div>
            </div>
          </div>
          <!-- /TAB A -->
          <!-- /tab B -->
        </div>
        <!-- /tab-content -->
      </div>
      <!-- /container -->
    </div>
    <!-- /tab_content_wrapper -->
  </main>
  <!-- /main -->
</template>

<script>
export default {
  name: "GoodsIntroduce",
  //props处获取参数
  props: {
    dataWithPic: {
      type: String,
      required: true,
    },
  },
};
</script>

<style lang="scss" scoped>
@import "@/assets/variables.scss";
</style>
