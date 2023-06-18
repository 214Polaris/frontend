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
                  <div class="row justify-content-between">
                    <div class="col-lg-6">
                      <h3>商品详情</h3>
                      {{ GoodIntroduce }}
                    </div>
                    <div class="col-lg-5">
                      <h3>参数</h3>
                      <div class="table-responsive">
                        <table class="table table-sm table-striped">
                          <tbody>
                            <tr>
                              <td><strong>Color</strong></td>
                              <td>Blue, Purple</td>
                            </tr>
                            <tr>
                              <td><strong>Size</strong></td>
                              <td>150x100x100</td>
                            </tr>
                            <tr>
                              <td><strong>Weight</strong></td>
                              <td>0.6kg</td>
                            </tr>
                            <tr>
                              <td><strong>Manifacturer</strong></td>
                              <td>Manifacturer</td>
                            </tr>
                          </tbody>
                        </table>
                      </div>
                      <!-- /table-responsive -->
                    </div>
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
    GoodName: String,
    GoodIntroduce: String,
  },
};
</script>

<style lang="scss" scoped>
@import "@/assets/variables.scss";

</style>
