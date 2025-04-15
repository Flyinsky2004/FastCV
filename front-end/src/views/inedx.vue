<script setup>
import {ref, onMounted} from 'vue'
import router from "@/router/index.js";

const props = defineProps({
  // 要依次显示的文本数组
  texts: {
    type: Array,
    default: () => ['后端工程师', '前端工程师', '运维工程师', 'UI设计师','IOS工程师','全栈开发者']
  },
  // 打字速度(ms)
  typeSpeed: {
    type: Number,
    default: 150
  },
  // 删除速度(ms)
  deleteSpeed: {
    type: Number,
    default: 100
  },
  // 切换文本前的停顿时间(ms)
  delayBeforeDelete: {
    type: Number,
    default: 2000
  }
})

const currentText = ref('')
const currentIndex = ref(0)
const isDeleting = ref(false)

const typeText = () => {
  const currentFullText = props.texts[currentIndex.value]

  if (!isDeleting.value) {
    // 打字过程
    currentText.value = currentFullText.substring(0, currentText.value.length + 1)

    if (currentText.value === currentFullText) {
      // 完成打字，准备删除
      isDeleting.value = true
      setTimeout(typeText, props.delayBeforeDelete)
      return
    }
  } else {
    // 删除过程
    currentText.value = currentFullText.substring(0, currentText.value.length - 1)

    if (currentText.value === '') {
      // 完成删除，准备打下一个文本
      isDeleting.value = false
      currentIndex.value = (currentIndex.value + 1) % props.texts.length
    }
  }

  // 计算下一次打字/删除的延迟时间
  const delay = isDeleting.value ? props.deleteSpeed : props.typeSpeed
  setTimeout(typeText, delay)
}

onMounted(() => {
  typeText()
})
const openLink = (link) => {
  window.open(link, '_blank')
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
    <div class="flex flex-col h-1/2 select-none">
      <div class="mx-auto my-auto">
        <div class="typewriter">
          你好，<span class="text-blue-500">{{ currentText }}</span><span class="cursor">|</span>
        </div>
        <div class="w-full flex mt-2 flex-col">
          <h1 class="mx-auto text-gray-500">一个纯前端，由Vue.js构建的简历模版渲染工具。</h1>
          <h1 class="mx-auto text-gray-500"><a class="line-through">本工具无任何联网功能，</a></h1>
          <h1 class="mx-auto text-gray-500"><a class="text-red-500">2025.4.15Update:新增AI润色和模拟答辩功能</a></h1>
          <h1 class="mx-auto text-gray-500">依赖浏览器localStorage持久缓存数据。</h1>
        </div>
      </div>
    </div>
    <div>
      <div class="grid grid-cols-[1fr,1fr,1fr] place-items-center w-1/3 mx-auto">
        <button class="basic-button bg-blue-500 hover:bg-blue-600 active:bg-blue-700" @click="router.push('/main')">开始使用</button>
        <button class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700" @click="openLink('https://flyinsky.wiki/')">作者博客</button>
        <button class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700" @click="openLink('https://github.com/Flyinsky2004/FastCV')">开源仓库</button>
      </div>
    </div>

  </div>

</template>

<style scoped>
.typewriter {
  font-family: monospace;
  font-size: 50px;
  display: inline-block;
  position: relative;
}

.cursor {
  color: rgb(78, 128, 238);
  display: inline-block;
  animation: blink 0.8s infinite;
}

@keyframes blink {
  0%, 100% {
    opacity: 1;
  }
  50% {
    opacity: 0;
  }
}
</style>