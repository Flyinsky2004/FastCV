<script setup>
import {ref, onMounted} from 'vue'
import Layout from '@/components/Layout.vue'
import router from '@/router'
import { useMotion } from '@vueuse/motion'

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

// 添加动画引用
const heroSection = ref(null)
const buttonsSection = ref(null)

// 设置动画
const heroMotion = useMotion(heroSection, {
  initial: {
    opacity: 0,
    y: 100,
  },
  enter: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 800,
    },
  },
})

const buttonsMotion = useMotion(buttonsSection, {
  initial: {
    opacity: 0,
    y: 50,
  },
  enter: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 800,
      delay: 200,
    },
  },
})
</script>

<template>
  <div class="min-h-screen bg-gray-100/80 py-12 px-4">
    <div ref="heroSection" class="flex flex-col h-1/2 select-none mb-16">
      <div class="mx-auto my-auto text-center">
        <div class="typewriter mb-8">
          <a class="dark:text-white">你好，</a><span class="text-blue-500 font-bold">{{ currentText }}</span><span class="cursor">|</span>
        </div>
        <div class="w-full flex mt-8 flex-col space-y-3">
          <h1 class="mx-auto text-gray-700 dark:text-gray-300 text-xl font-medium">一个纯前端，由Vue.js构建的简历模版渲染工具。</h1>
          <h1 class="mx-auto text-gray-600 dark:text-gray-400 text-lg"><a class="line-through">本工具无任何联网功能，</a></h1>
          <h1 class="mx-auto text-red-500 dark:text-red-400 text-lg font-semibold">2025.4.15Update: 新增AI润色和模拟答辩功能</h1>
          <h1 class="mx-auto text-gray-600 dark:text-gray-400 text-lg">依赖浏览器localStorage持久缓存数据。</h1>
        </div>
      </div>
    </div>
    
    <div ref="buttonsSection" class="max-w-4xl mx-auto px-4">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
        <button 
          class="group relative overflow-hidden rounded-2xl bg-gradient-to-r from-blue-500 to-blue-600 p-[2px] transition-all duration-300 hover:shadow-lg hover:shadow-blue-500/25"
          @click="router.push('/main')"
        >
          <div class="relative rounded-2xl bg-gray-900 p-6 transition-all duration-300 group-hover:bg-opacity-90">
            <div class="flex flex-col items-center text-white">
              <span class="text-2xl font-bold mb-2">开始使用</span>
              <span class="text-sm opacity-90">创建或编辑你的简历</span>
            </div>
          </div>
        </button>
        
        <button 
          class="group relative overflow-hidden rounded-2xl bg-gradient-to-r from-purple-500 to-purple-600 p-[2px] transition-all duration-300 hover:shadow-lg hover:shadow-purple-500/25"
          @click="router.push('/virtual')"
        >
          <div class="relative rounded-2xl bg-gray-900 p-6 transition-all duration-300 group-hover:bg-opacity-90">
            <div class="flex flex-col items-center text-white">
              <span class="text-2xl font-bold mb-2">模拟面试</span>
              <span class="text-sm opacity-90">AI驱动的面试训练</span>
            </div>
          </div>
        </button>
      </div>

      <div class="flex flex-col sm:flex-row justify-center space-y-4 sm:space-y-0 sm:space-x-6">
        <button 
          class="group relative overflow-hidden rounded-xl bg-gradient-to-r from-green-500 to-green-600 p-[2px] transition-all duration-300 hover:shadow-lg hover:shadow-green-500/25"
          @click="openLink('https://1024110.xyz/')"
        >
          <div class="relative rounded-xl bg-gray-900 px-6 py-3 transition-all duration-300 group-hover:bg-opacity-90">
            <span class="text-lg font-medium text-white">作者博客</span>
          </div>
        </button>

        <button 
          class="group relative overflow-hidden rounded-xl bg-gradient-to-r from-gray-600 to-gray-700 p-[2px] transition-all duration-300 hover:shadow-lg hover:shadow-gray-500/25"
          @click="openLink('https://github.com/Flyinsky2004/FastCV')"
        >
          <div class="relative rounded-xl bg-gray-900 px-6 py-3 transition-all duration-300 group-hover:bg-opacity-90">
            <span class="text-lg font-medium text-white">开源仓库</span>
          </div>
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.typewriter {
  font-family: ui-monospace, SFMono-Regular, Menlo, Monaco, Consolas, "Liberation Mono", "Courier New", monospace;
  font-size: 2.5rem;
  @apply sm:text-5xl font-bold;
  display: inline-block;
  position: relative;
}

.cursor {
  @apply text-blue-500;
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