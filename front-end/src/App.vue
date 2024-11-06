<script setup>
import {useThemeStore} from "@/stores/themeStore.js";
import ThemeSwitch from "@/components/commmon/themeSwitch.vue";
const themeStore = useThemeStore()
themeStore.initTheme()
</script>
<template>
  <div class="h-screen overflow-y-hidden bkg">
<!--    <ThemeSwitch class="fixed top-0 right-0"/>-->
    <RouterView/>
  </div>

</template>
<style scoped>
.bkg{
  background-image:  linear-gradient(#bfbfcf 2px, transparent 2px), linear-gradient(90deg, #bfbfcf 2px, transparent 2px), linear-gradient(#bfbfcf 1px, transparent 1px), linear-gradient(90deg, #bfbfcf 1px, #ffffff 1px);
  background-size: 50px 50px, 50px 50px, 10px 10px, 10px 10px;
  background-position: -2px -2px, -2px -2px, -1px -1px, -1px -1px;
}
</style>