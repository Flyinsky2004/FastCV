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