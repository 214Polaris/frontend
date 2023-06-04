<script setup>
import NavBar from "./components/NavBar.vue";
import Footer from "./components/Footer.vue";
import Head from "@/views/Layout/components/head.vue";
</script>

<template>
  <Head />
  <NavBar />
  <RouterView />
  <Footer />
</template>
