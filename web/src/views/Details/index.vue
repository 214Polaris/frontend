<template>
  <div class="container">
    <div class="common-layout">
      <div class="goods-info">
        <el-container>
          <el-aside width="auto">
            <!-- 商品图片组件 -->
            <div class="media">
              <GoodsImage :image="good.productLink" />
            </div>
          </el-aside>
          <el-main>
            <!--商品信息组件-->
            <div class="spec">
              <GoodsIntroduce />
            </div>
          </el-main>
        </el-container>
      </div>
    </div>
  </div>
</template>
<script>
import axios from 'axios'
import GoodsImage from './components/details_image.vue'
import GoodsIntroduce from './components/details_introduce.vue'
export default {
  data () {
    return{
      //从上个页面获取点击的商品id
      good_id: this.$route.params.goodId,
      good:null
    }
  },
  //向后端请求返回对应商品
  mounted(){
    axios({
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
          },
          method: 'post',
          url: '/api/details',
          data: {
            good_id: this.good_id
          }
        }).then((response)=>{
          this.good = response.data
        }).catch(error => {
      console.error(error)
    })
},
  components:{
    GoodsIntroduce,
    GoodsImage
  }
}
</script>
<style scoped lang="scss">
@import "@/assets/variables.scss";
.goods-info {
  min-height: 600px;
  background: #fff;
}
.media {
  width: 580px;
  height: 600px;
}
.spec {
  flex: 1;
}
.container {
  display: flex;
  justify-content: center;
  align-items: center;
}
</style>