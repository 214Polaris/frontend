<template>
  <div
    class="tb-booth"
    @mouseover="onMouseOver"
    @mouseout="onMouseOut"
    @mousemove="onMouseMove"
    ref="boothRef"
  >
    <!--此处导入图片-->
    <img :src="image" alt="" v-if="image" class="smallImg"/>
    <!--大图容器-->
    <div class="mask" ref="mask" v-show="boxShow" />
    <div class="big-img_box" ref="bigImgBox" v-show="boxShow">
      <!--此处导入图片-->
      <img class="big-img" ref="bigImg" :src="image" />
    </div>
  </div>
</template>
<script>
import { reactive, toRefs, ref } from "vue";
export default {
  //使用props定义要从上一个页面传入的参数
  props: {
    image: String,
  },
  setup() {
    const boothRef = ref(null);
    const mask = ref(null);
    const bigImg = ref(null);
    const bigImgBox = ref(null);
    const state = reactive({
      boxShow: false,
    });
    //监听鼠标移动
    const onMouseOver = () => {
      state.boxShow = true;
    };
    const onMouseOut = () => {
      state.boxShow = false;
    };

    const onMouseMove = (e) => {
      let x = e.pageX - boothRef.value.offsetLeft;
      let y = e.pageY - boothRef.value.offsetTop;
      let maskX = x - mask.value.offsetWidth / 2;
      let maskY = y - mask.value.offsetHeight / 2;
      // mask的x最大移动距离
      let maskXMaxMove = boothRef.value.offsetWidth - mask.value.offsetWidth;
      let maskYMaxMove = boothRef.value.offsetHeight - mask.value.offsetHeight;
      let bigImgXMaxMove =
        bigImgBox.value.offsetWidth - bigImg.value.offsetWidth;
      let bigImgYMaxMove =
        bigImgBox.value.offsetHeight - bigImg.value.offsetHeight;
      if (maskX <= 0) {
        maskX = 0;
      } else if (maskX >= maskXMaxMove) {
        maskX = maskXMaxMove;
      }
      if (maskY <= 0) {
        maskY = 0;
      } else if (maskY >= maskYMaxMove) {
        maskY = maskYMaxMove;
      }
      mask.value.style.left = maskX + "px";
      mask.value.style.top = maskY + "px";
      // 大图片移动距离 = mask的移动距离*大盒子最大移动距离 / mask的x最大移动距离
      let bixImgXMove = (maskX * bigImgXMaxMove) / maskXMaxMove;
      let bixImgYMove = (maskY * bigImgYMaxMove) / maskYMaxMove;
      bigImg.value.style.left = bixImgXMove + "px";
      bigImg.value.style.top = bixImgYMove + "px";
    };
    return {
      ...toRefs(state),
      boothRef,
      mask,
      bigImg,
      bigImgBox,
      onMouseOver,
      onMouseOut,
      onMouseMove,
    };
  },
};
</script>
<style scoped lang="scss">
.tb-booth {
  top:5vw;
  width: 420px;
  height: 420px;
  position: relative;
  border: 1px solid #cccccc;
  .smallImg{
    object-fit: contain;
    height: 420px;
    width: auto;
  }
}
.mask {
  position: absolute;
  top: 0;
  left: 0;
  width: 210px;
  height: 210px;
  background-color: rgb(61, 110, 206);
  opacity: 0.5;
  cursor: move;
}

.big-img_box {
  transition-delay: 1s;
  position: absolute;
  top: 0;
  left: 500px;
  width: 400px;
  height: 400px;
  background-color: #fff;
  border: 1px solid #cccccc;
  overflow: hidden;
  z-index: 99999999;
}
.big-img {

  position: absolute;
  width:200%;
  height: 200%;
  left: 0;
  top: 0;
}
</style>
