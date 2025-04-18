<script setup>
import {useProfileStore} from "@/stores/profileStore.js";
import github from '@/assets/imgs/github-mark.svg'
import dateToString from "../../util/dateFormatter.js";
import {onBeforeMount, reactive, ref} from "vue";

const profileStore = useProfileStore()
const options = reactive({
  profile: {}
})
onBeforeMount(() => {

  if (!profileStore.isSelected) options.profile = JSON.parse(localStorage.getItem("lastProfile"))
  else options.profile = profileStore.getCurrentProfile
})


</script>
<template>
  <div class="max-w-5xl mx-auto bg-white shadow-lg rounded-lg print:shadow-none">
    <div class="h-2 bg-gradient-to-r from-orange-500 to-orange-600 rounded-t-lg"></div>
    
    <div class="p-6 space-y-8">
      <!-- 主要个人资料 -->
      <div class="flex flex-col md:flex-row items-start gap-6">
        <img class="w-24 h-24 rounded-lg object-cover shadow-md" :src="options.profile.avatar" alt="Profile Photo">
        <div class="space-y-2">
          <h1 class="text-4xl font-bold text-gray-800">{{ options.profile.name }}</h1>
          <h2 class="text-xl text-orange-600 font-medium">{{ options.profile.jobName }}</h2>
        </div>
      </div>

      <!-- 联系方式和社交链接 -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div class="flex items-center space-x-3 text-gray-600">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="w-5 h-5 text-orange-500">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"/>
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z"/>
          </svg>
          <span>{{ options.profile.position }}</span>
        </div>
        
        <div class="flex items-center space-x-3">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="w-5 h-5 text-orange-500">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75"/>
          </svg>
          <a :href="'mailto:'+options.profile.emailAddress" class="text-orange-600 hover:text-orange-700 transition-colors">
            {{ options.profile.emailAddress }}
          </a>
        </div>

        <div class="flex items-center space-x-3 text-gray-600">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="w-5 h-5 text-orange-500">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5v2.25Z"/>
          </svg>
          <span>{{ options.profile.phoneNumber }}</span>
        </div>

        <div class="flex items-center space-x-3 text-gray-600">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="w-5 h-5 text-orange-500">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M12 8.25v-1.5m0 1.5c-1.355 0-2.697.056-4.024.166C6.845 8.51 6 9.473 6 10.608v2.513m6-4.871c1.355 0 2.697.056 4.024.166C17.155 8.51 18 9.473 18 10.608v2.513M15 8.25v-1.5m-6 1.5v-1.5m12 9.75-1.5.75a3.354 3.354 0 0 1-3 0 3.354 3.354 0 0 0-3 0 3.354 3.354 0 0 1-3 0 3.354 3.354 0 0 0-3 0 3.354 3.354 0 0 1-3 0L3 16.5m15-3.379a48.474 48.474 0 0 0-6-.371c-2.032 0-4.034.126-6 .371m12 0c.39.049.777.102 1.163.16 1.07.16 1.837 1.094 1.837 2.175v5.169c0 .621-.504 1.125-1.125 1.125H4.125A1.125 1.125 0 0 1 3 20.625v-5.17c0-1.08.768-2.014 1.837-2.174A47.78 47.78 0 0 1 6 13.12M12.265 3.11a.375.375 0 1 1-.53 0L12 2.845l.265.265Zm-3 0a.375.375 0 1 1-.53 0L9 2.845l.265.265Zm6 0a.375.375 0 1 1-.53 0L15 2.845l.265.265Z"/>
          </svg>
          <span>{{ dateToString(options.profile.birthDate) }}</span>
        </div>

        <!-- 社交链接 -->
        <div class="flex space-x-6 col-span-3">
          <a v-if="options.profile.webSite" :href="options.profile.webSite" 
             class="flex items-center space-x-2 text-orange-600 hover:text-orange-700 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                 stroke="currentColor" class="w-5 h-5">
              <path stroke-linecap="round" stroke-linejoin="round"
                    d="M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622 1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244"/>
            </svg>
            <span>个人网站</span>
          </a>

          <a v-if="options.profile.github" :href="options.profile.github"
             class="flex items-center space-x-2 text-orange-600 hover:text-orange-700 transition-colors">
            <img class="w-5 h-5" :src="github" alt="GitHub">
            <span>Github</span>
          </a>

          <a v-if="options.profile.linkIn" :href="options.profile.linkIn"
             class="flex items-center space-x-2 text-orange-600 hover:text-orange-700 transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
                 stroke="currentColor" class="w-5 h-5">
              <path stroke-linecap="round" stroke-linejoin="round"
                    d="M3.75 21h16.5M4.5 3h15M5.25 3v18m13.5-18v18M9 6.75h1.5m-1.5 3h1.5m-1.5 3h1.5m3-6H15m-1.5 3H15m-1.5 3H15M9 21v-3.375c0-.621.504-1.125 1.125-1.125h3.75c.621 0 1.125.504 1.125 1.125V21"/>
            </svg>
            <span>LinkedIn</span>
          </a>
        </div>
      </div>

      <!-- 主要内容区域 -->
      <div class="grid grid-cols-1 lg:grid-cols-[3fr,2fr] gap-8">
        <!-- 左侧内容 -->
        <div class="space-y-8">
          <!-- 工作经历 -->
          <section>
            <div class="border-b border-gray-200 pb-2 mb-6">
              <h2 class="text-2xl font-bold text-gray-800">工作经历</h2>
            </div>
            <div class="space-y-6">
              <div v-for="exp in options.profile.workExperience" class="grid grid-cols-[auto,1fr] gap-4">
                <div class="relative">
                  <div class="h-full w-0.5 bg-orange-200 absolute left-0"></div>
                  <div class="w-4 h-4 rounded-full bg-orange-500"></div>
                  <div class="text-sm text-gray-500 mt-2">
                    {{ dateToString(exp.startDate) }}<br>至<br>{{ dateToString(exp.endDate) }}
                  </div>
                </div>
                <div class="space-y-2">
                  <h3 class="text-lg font-semibold text-orange-600">{{ exp.workFor }}</h3>
                  <div class="text-gray-700">{{ exp.jobName }}</div>
                  <p class="text-gray-600">{{ exp.description }}</p>
                </div>
              </div>
            </div>
          </section>

          <!-- 个人项目 -->
          <section>
            <div class="border-b border-gray-200 pb-2 mb-6">
              <h2 class="text-2xl font-bold text-gray-800">个人项目</h2>
            </div>
            <div class="space-y-6">
              <div v-for="pro in options.profile.projects" class="space-y-2">
                <h3 class="text-lg font-semibold text-orange-600">{{ pro.title }}</h3>
                <p class="text-gray-600">{{ pro.description }}</p>
              </div>
            </div>
          </section>
        </div>

        <!-- 右侧内容 -->
        <div class="space-y-8">
          <!-- 教育经历 -->
          <section>
            <div class="border-b border-gray-200 pb-2 mb-6">
              <h2 class="text-2xl font-bold text-gray-800">教育经历</h2>
            </div>
            <div class="space-y-6">
              <div v-for="edu in options.profile.education" class="space-y-2">
                <h3 class="text-lg font-semibold text-orange-600">{{ edu.level }}</h3>
                <div class="text-sm text-gray-500">
                  {{ dateToString(edu.startDate) }} 至 {{ dateToString(edu.endDate) }}
                </div>
                <div class="flex justify-between text-gray-600">
                  <span>{{ edu.name }}</span>
                  <span>{{ edu.profess }}</span>
                </div>
              </div>
            </div>
          </section>

          <!-- 专业技能 -->
          <section>
            <div class="border-b border-gray-200 pb-2 mb-6">
              <h2 class="text-2xl font-bold text-gray-800">专业技能</h2>
            </div>
            <div class="space-y-4">
              <div v-for="skill in options.profile.skills" class="space-y-2">
                <h3 class="text-lg font-semibold text-orange-600">{{ skill.title }}</h3>
                <p class="text-gray-600">{{ skill.description }}</p>
              </div>
            </div>
          </section>

          <!-- 竞赛经历 -->
          <section>
            <div class="border-b border-gray-200 pb-2 mb-6">
              <h2 class="text-2xl font-bold text-gray-800">竞赛经历</h2>
            </div>
            <div class="space-y-4">
              <div v-for="race in options.profile.races" class="space-y-2">
                <h3 class="text-lg font-semibold text-orange-600">{{ race.name }}</h3>
                <div class="flex justify-between text-gray-600">
                  <span>{{ race.level }}</span>
                  <span class="text-sm">{{ dateToString(race.date) }}</span>
                </div>
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
  </div>
</template>