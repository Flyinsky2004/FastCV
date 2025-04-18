<script setup>
import {useRoute} from "vue-router";
import {reactive, ref} from "vue";
import Standard from "@/components/templates/standard.vue";
import html2pdf from "html2pdf.js";
import {useThemeStore} from "@/stores/themeStore.js";
import router from "@/router/index.js";
import Loading from "@/components/commmon/Loading.vue";
import Modern from "@/components/templates/modern.vue";
import Business from "@/components/templates/business.vue";
import Forest from "@/components/templates/forest.vue";


const themeStore=  useThemeStore()
const route = useRoute()
const currentId = route.params.id
const options = reactive({
  isEditMode : false,
  outputMode : false
})
const exportResumeAsPDF = async () => {
  try {
    options.outputMode = true;
    const element = document.getElementById('resume');
    const isDark = themeStore.isDark;
    const backgroundColor = isDark ? '#000000' : '#ffffff';
    const opt = {
      margin: [10, 0, 0, 10],
      filename: 'download.pdf',
      image: { type: 'jpeg', quality: 1 },
      html2canvas: {
        scale: 5,
        backgroundColor: '#000000'
      },
      jsPDF: {
        unit: 'pt',
        format: [920, 1300],
        orientation: 'portrait'
      }
    };

    await html2pdf().from(element).set(opt).save();
  } catch (error) {
    console.error('生成失败：', error);
  } finally {
    options.outputMode = false;
  }
}


</script>

<template>
  <div class="h-screen overflow-y-auto bg-white bg-opacity-70">
    <div class="h-20 p-4 bg-white w-3/4 mx-auto rounded shadow-lg grid grid-cols-8 mt-2 animate__animated animate__fadeIn">
      <button @click="router.push('/main')"
              class="basic-button bg-blue-500 hover:bg-blue-600 active:bg-blue-700 flex flex-nowrap place-items-center">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4">
          <path stroke-linecap="round" stroke-linejoin="round" d="M9 15 3 9m0 0 6-6M3 9h12a6 6 0 0 1 0 12h-3" />
        </svg>
        返回主页
      </button>
      <button class="basic-button bg-red-500 hover:bg-red-600 active:bg-red-700 flex flex-nowrap place-items-center"
              @click="exportResumeAsPDF">
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="size-4">
          <path stroke-linecap="round" stroke-linejoin="round" d="M3 16.5v2.25A2.25 2.25 0 0 0 5.25 21h13.5A2.25 2.25 0 0 0 21 18.75V16.5M16.5 12 12 16.5m0 0L7.5 12m4.5 4.5V3" />
        </svg>保存为PDF
      </button>
    </div>
    <div
        :class="options.outputMode ? 'w-4/5' : 'w-3/4'"
        id="resume"
        class="mx-auto bg-white mt-4"
    >
      <Standard v-if="(Number)(currentId) === 1"/>
      <Modern v-if="(Number)(currentId) === 2"/>
      <Business v-if="(Number)(currentId) === 3"/>
      <Forest v-if="(Number)(currentId) === 4"/>

    </div>
  </div>
</template>