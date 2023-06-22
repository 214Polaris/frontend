<template></template>
<script setup>
import { onBeforeMount } from "vue";
import { useRoute, useRouter } from "vue-router";
const route = useRoute();
const router = useRouter();
const key = route.params.prom;
onBeforeMount(() => {
  router.push({ name: "search", params: key });
});
</script>
