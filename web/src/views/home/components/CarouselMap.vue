<!-- 这里实现轮播图 -->

<script setup>
import axios from "axios";
import { onMounted, ref } from "vue";

const banner_list = ref([]);

const get_banner = () => {
  axios
    .get("/api/CarouselMap")
    .then((response) => {
      console.log(response);
      banner_list.value = response.data;
    })
    .catch((error) => {
      console.error(error);
    });
};

const goToDetial = () => {};

onMounted(() => {
  get_banner();
});
</script>

<template>
  <div class="container">
    <el-carousel
      height="500px"
      autoplay
      indicator-position="outside"
      initial-index="0"
      v-if="banner_list[0]"
    >
      <el-carousel-item v-for="(item, index) in banner_list" :key="index">
        <router-link
          :to="{ name: 'Details', params: { goodId: item.productId } }"
          ><img :src="item.productLink" @click="goToDetial(item.productId)" />
        </router-link>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<style scoped>
img {
  max-height: 500px;
  width: auto;
}

.container {
  width: 100vw;
}
</style>
