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
                  <a href="#0" data-bs-toggle="modal" data-bs-target="#para-modal"
                    ><i class="iconfont icon-arrow-double-right"></i
                  ></a>
                </div>
                <div
                  class="col-xl-5 col-lg-5 col-md-6 col-6"
                  style="padding-bottom: 10px"
                >
                <a href="#0" data-bs-toggle="modal" data-bs-target="#para-modal">{{ store.dataWithPic.option.Value }}</a>
                </div>
              </div>
              <el-divider />
              <div class="row">
                <label
                  class="col-xl-7 col-lg-5 col-md-6 col-6"
                  style="font-size: large"
                  ><strong>选项</strong> - 类型选择
                </label>
                <div
                  class="col-xl-5 col-lg-5 col-md-6 col-6"
                  style="padding-bottom: 10px"
                >
                  <a href="#0" data-bs-toggle="modal" data-bs-target="#size-modal"
                    ><i class="iconfont icon-arrow-double-right"></i
                  ></a>
                </div>
                <div
                  class="col-xl-5 col-lg-5 col-md-6 col-6"
                  style="padding-bottom: 10px"
                >
                <a href="#0" data-bs-toggle="modal" data-bs-target="#para-modal">{{ store.dataWithoutPic.option.Value }}</a>
                </div>
              </div>
              <el-divider />
              <div class="row">
                <label
                  class="col-xl-7 col-lg-5 col-md-6 col-6"
                  style="width:40%; font-size: medium;"
                  ><strong>数量</strong></label
                >
                <el-input-number style="margin-left: 5vw;" v-model="num" :min="1" :max="99" @change="handleChange" />
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
                    <a class="btn_1" @click="addCart()">添加到购物车</a>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <!-- /row -->
      </div>
    </main>
    <!--大小选项表-->
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
            <el-row :gutter="12">
              <el-col v-bind="item" v-for="item in goods.dataWithoutPic" :key="item.option.ID" :span="8">
                  <div class="bottom">
                    <el-button text class="button" @click="selectedsize = item, store_size()" data-bs-dismiss="modal"
              aria-label="Close">{{ item.option.Value }}</el-button>
                  </div>
              </el-col>
            </el-row>
            <!-- /table -->
          </div>
        </div>
      </div>
    </div>
    <!--参数选项表-->
    <div
      class="modal fade"
      tabindex="-1"
      role="dialog"
      aria-labelledby="size-modal"
      id="para-modal"
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
            <el-row :gutter="12">
              <el-col v-bind="item" v-for="item in goods.dataWithPic" :key="item.option.ID" :span="8">
                <el-card :body-style="{ padding: '0px' }">
                  <img
                    style="object-fit: contain; width:auto; height: 210px;"
                    :src="item.image.URL"
                    class="image"
                  />
                  <div style="padding: 12px">
                    <span>{{ item.option.Value }}</span>
                    <div class="bottom">
                      <el-button text class="button" @click="selectedparams = item, store_param()" data-bs-dismiss="modal"
              aria-label="Close">确认</el-button>
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
import { ElMessageBox } from 'element-plus';
  import { defineProps, ref, watch, reactive } from 'vue';
  
  //选择的参数
  let selectedparams = null;
  let selectedsize = null;
  //父组件传来的参数
  const props = defineProps({
    GoodName: String,
    GoodIntroduce: String,
    GoodPrice: String,
    GoodId: String
  });
  let price;
  //后端发来的商品参数
  let goods = reactive({
  dataWithPic: [],
  dataWithoutPic: []
})
  //加入购物车的数量
  let num = ref(0);

  //向后端发送的已选商品参数
  let store = reactive({
  dataWithPic: reactive({image:[],option:[]}),
  dataWithoutPic: reactive({option:[]}),
})

  // 处理点击事件
  const handleChange = (num) => {
    console.log(num);
  };

  //存带图片参数
  function store_param(){
    if (selectedparams) {
      store.dataWithPic = { image: selectedparams.image, option: selectedparams.option };
      console.log(store);
    }
  }
  //存不带图片参数
  function store_size(){
    if (selectedsize) {
      store.dataWithoutPic = { option: selectedsize.option };
      console.log(store);
    }
  }
  //监听props变化并向后端post参数
  watch(() => props.GoodId, (newId, oldId) => {
    console.log(newId);
    //打折之后的价格
    price = parseInt(props.GoodPrice) * 0.8;
    price = Math.floor(price);
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
        goods.dataWithPic = response.data.dataWithPic
        goods.dataWithoutPic = response.data.dataWithoutPic
        console.log(goods.dataWithPic);
        // 执行其他逻辑
      })
      .catch((error) => {
        console.error(error);
      });
  });

  //加入购物车
  function addCart(){
    let opt_img = store.dataWithPic.option;
    let opt =  store.dataWithoutPic.option;
    let img = store.dataWithPic.image;
    if(store.dataWithPic.image.URL!=null&&store.dataWithoutPic.option.Value!=null){
      axios({
      headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
      },
      method: 'post',
      url: '/api/add',
      data: {
        image:img,
        options:opt,
        count:num.value
      },
    })
      .then((response) => {
        ElMessageBox.alert("已加入购物车！",{
        showClose: false
      })
        // 执行其他逻辑
      })
      .catch((error) => {
        console.log(num)
        console.error(error);
      });
    }
    else{
      ElMessageBox.alert("请选择参数！",{
        showClose: false
      })
    }
  }
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
