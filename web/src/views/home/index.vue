<script setup>
import CarouselMap from "./components/CarouselMap.vue";
import Recommend from "./components/Recommend.vue";
</script>

<template>
  <CarouselMap style="margin-top: 20px" />
  <h1 style="color: #collapse-A; margin-top: 30px; margin-left: 30px">
    为您推荐：
  </h1>
  <el-divider />
  <Recommend />
</template>
