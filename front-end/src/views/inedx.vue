<script setup>
import {ref, onMounted} from 'vue'
import Layout from '@/components/Layout.vue'
import router from '@/router'

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
    <div class="flex flex-col h-1/2 select-none">
      <div class="mx-auto my-auto">
        <div class="typewriter">
          <a class="dark:text-white">你好，</a><span class="text-blue-500">{{ currentText }}</span><span class="cursor">|</span>
        </div>
        <div class="w-full flex mt-4 flex-col space-y-2">
          <h1 class="mx-auto text-gray-600 text-lg">一个纯前端，由Vue.js构建的简历模版渲染工具。</h1>
          <h1 class="mx-auto text-gray-600 text-lg"><a class="line-through">本工具无任何联网功能，</a></h1>
          <h1 class="mx-auto text-red-500 text-lg font-medium">2025.4.15Update: 新增AI润色和模拟答辩功能</h1>
          <h1 class="mx-auto text-gray-600 text-lg">依赖浏览器localStorage持久缓存数据。</h1>
        </div>
      </div>
    </div>
    <div class="mt-12 px-4">
      <div class="max-w-4xl mx-auto">
        <div class="grid grid-cols-2 gap-8 mb-8">
          <button class="primary-button bg-gradient-to-r from-blue-500 to-blue-600" @click="router.push('/main')">
            <div class="flex flex-col items-center">
              <span class="text-2xl font-semibold mb-2">开始使用</span>
              <span class="text-sm opacity-90">创建或编辑你的简历</span>
            </div>
          </button>
          <button class="primary-button bg-gradient-to-r from-purple-500 to-purple-600" @click="router.push('/virtual')">
            <div class="flex flex-col items-center">
              <span class="text-2xl font-semibold mb-2">模拟面试</span>
              <span class="text-sm opacity-90">AI驱动的面试训练</span>
            </div>
          </button>
        </div>
        <div class="flex justify-center space-x-8">
          <button class="secondary-button bg-gradient-to-r from-green-500 to-green-600" @click="openLink('https://flyinsky.wiki/')">
            <span class="text-lg">作者博客</span>
          </button>
          <button class="secondary-button bg-gradient-to-r from-gray-500 to-gray-600" @click="openLink('https://github.com/Flyinsky2004/FastCV')">
            <span class="text-lg">开源仓库</span>
          </button>
        </div>
      </div>
    </div>
</template>

<style scoped>
.typewriter {
  font-family: monospace;
  font-size: 3.5rem;
  display: inline-block;
  position: relative;
}

.cursor {
  color: rgb(78, 128, 238);
  display: inline-block;
  animation: blink 0.8s infinite;
}

.primary-button {
  @apply flex items-center justify-center px-8 py-8 rounded-2xl text-white shadow-lg 
         transition-all duration-300 transform hover:scale-105 hover:shadow-xl
         border border-white/10 backdrop-blur-sm;
}

.secondary-button {
  @apply flex items-center justify-center px-8 py-4 rounded-xl text-white shadow-md 
         transition-all duration-300 transform hover:scale-105 hover:shadow-lg
         border border-white/10;
}

.primary-button:active, .secondary-button:active {
  @apply transform scale-95;
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