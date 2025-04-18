<script setup>
import {RouterView} from "vue-router";
import {useThemeStore} from "@/stores/themeStore";
import {theme} from "ant-design-vue";
const themeStore = useThemeStore()

themeStore.initTheme()
import zhCN from 'ant-design-vue/es/locale/zh_CN';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
dayjs.locale('zh-cn');
</script>
<template>
  <a-config-provider
      :theme="{
      algorithm: themeStore.isDark ?theme.darkAlgorithm:theme.defaultAlgorithm,
    }"
      :locale="zhCN"
  >
  </a-config-provider>
  <div class="overflow-y-auto bkg bkg-theme-switch">
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