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
    <div class="h-10 bg-white shadow-md shadow-gray-400 place-items-center flex">
      <div class="font-bold text-xl flex-grow text-center title bg-clip-text bg-gradient-to-r from-purple-500 via-red-500 to-pink-500 text-transparent">FastCV</div>
      <div>
        <a href="https://github.com/Flyinsky2004/FastCV">
          <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24" aria-hidden="true">
            <path fill-rule="evenodd"
                  d="M12 2C6.477 2 2 6.484 2 12.017c0 4.425 2.865 8.18 6.839 9.504.5.092.682-.217.682-.483 0-.237-.008-.868-.013-1.703-2.782.605-3.369-1.343-3.369-1.343-.454-1.158-1.11-1.466-1.11-1.466-.908-.62.069-.608.069-.608 1.003.07 1.531 1.032 1.531 1.032.892 1.53 2.341 1.088 2.91.832.092-.647.35-1.088.636-1.338-2.22-.253-4.555-1.113-4.555-4.951 0-1.093.39-1.988 1.029-2.688-.103-.253-.446-1.272.098-2.65 0 0 .84-.27 2.75 1.026A9.564 9.564 0 0112 6.844c.85.004 1.705.115 2.504.337 1.909-1.296 2.747-1.027 2.747-1.027.546 1.379.202 2.398.1 2.651.64.7 1.028 1.595 1.028 2.688 0 3.848-2.339 4.695-4.566 4.943.359.309.678.92.678 1.855 0 1.338-.012 2.419-.012 2.747 0 .268.18.58.688.482A10.019 10.019 0 0022 12.017C22 6.484 17.522 2 12 2z"
                  clip-rule="evenodd"></path>
          </svg>
        </a>
      </div>
    </div>
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