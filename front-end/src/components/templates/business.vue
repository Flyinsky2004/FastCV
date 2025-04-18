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
  <div class="max-w-5xl mx-auto bg-white shadow-lg rounded-lg overflow-hidden">
    <!-- 顶部个人信息区域 -->
    <div class="grid grid-cols-1 md:grid-cols-[1fr,1fr,1fr] bg-gradient-to-r from-[#7e90a5] to-[#5a6d83] text-white">
      <!-- 联系方式区域 -->
      <div class="p-6 space-y-4">
        <div class="space-y-3">
          <div class="flex items-center space-x-2 hover:text-gray-200 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z" />
            </svg>
            <span>{{options.profile.position}}</span>
          </div>
          <div class="flex items-center space-x-2 hover:text-gray-200 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75" />
            </svg>
            <a :href="'mailto:'+options.profile.emailAddress" class="hover:underline">{{options.profile.emailAddress}}</a>
          </div>
          <div class="flex items-center space-x-2 hover:text-gray-200 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5v2.25Z" />
            </svg>
            <span>{{options.profile.phoneNumber}}</span>
          </div>
          <div class="flex items-center space-x-2 hover:text-gray-200 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M12 8.25v-1.5m0 1.5c-1.355 0-2.697.056-4.024.166C6.845 8.51 6 9.473 6 10.608v2.513m6-4.871c1.355 0 2.697.056 4.024.166C17.155 8.51 18 9.473 18 10.608v2.513M15 8.25v-1.5m-6 1.5v-1.5m12 9.75-1.5.75a3.354 3.354 0 0 1-3 0 3.354 3.354 0 0 0-3 0 3.354 3.354 0 0 1-3 0 3.354 3.354 0 0 0-3 0 3.354 3.354 0 0 1-3 0L3 16.5m15-3.379a48.474 48.474 0 0 0-6-.371c-2.032 0-4.034.126-6 .371m12 0c.39.049.777.102 1.163.16 1.07.16 1.837 1.094 1.837 2.175v5.169c0 .621-.504 1.125-1.125 1.125H4.125A1.125 1.125 0 0 1 3 20.625v-5.17c0-1.08.768-2.014 1.837-2.174A47.78 47.78 0 0 1 6 13.12M12.265 3.11a.375.375 0 1 1-.53 0L12 2.845l.265.265Zm-3 0a.375.375 0 1 1-.53 0L9 2.845l.265.265Zm6 0a.375.375 0 1 1-.53 0L15 2.845l.265.265Z" />
            </svg>
            <span>{{dateToString(options.profile.birthDate)}}</span>
          </div>
        </div>
        <!-- 社交链接 -->
        <div class="grid grid-cols-3 gap-4">
          <a v-if="options.profile.webSite" :href="options.profile.webSite" class="flex items-center space-x-1 hover:text-gray-200 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622 1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244" />
            </svg>
            <span>个人网站</span>
          </a>
          <a v-if="options.profile.github" :href="options.profile.github" class="flex items-center space-x-1 hover:text-gray-200 transition">
            <img class="h-5 w-5 bg-white rounded-full p-0.5" :src="github" alt="GitHub">
            <span>Github</span>
          </a>
          <a v-if="options.profile.linkIn" :href="options.profile.linkIn" class="flex items-center space-x-1 hover:text-gray-200 transition">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 21h16.5M4.5 3h15M5.25 3v18m13.5-18v18M9 6.75h1.5m-1.5 3h1.5m-1.5 3h1.5m3-6H15m-1.5 3H15m-1.5 3H15M9 21v-3.375c0-.621.504-1.125 1.125-1.125h3.75c.621 0 1.125.504 1.125 1.125V21" />
            </svg>
            <span>LinkIn</span>
          </a>
        </div>
      </div>
      
      <!-- 姓名和职位 -->
      <div class="flex flex-col items-center justify-center p-6 text-center">
        <h1 class="text-4xl font-bold mb-2 tracking-wider">{{options.profile.name}}</h1>
        <h2 class="text-xl font-medium text-gray-100">{{options.profile.jobName}}</h2>
      </div>
      
      <!-- 头像 -->
      <div class="flex justify-center items-center p-6 bg-[#5a6d83]">
        <div class="relative w-40 h-40">
          <div class="absolute inset-0 bg-white rounded-full shadow-inner"></div>
          <img class="relative w-36 h-36 rounded-full object-cover border-4 border-white m-2" :src="options.profile.avatar" alt="Profile">
        </div>
      </div>
    </div>

    <!-- 内容区域 -->
    <div class="px-6 py-4">
      <!-- 工作经历 -->
      <div class="mb-8">
        <div class="flex items-center mb-4">
          <div class="flex justify-center items-center w-10 h-10 bg-[#7e90a5] rounded-lg mr-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="#f6de74" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-[#7e90a5]">工作经历</h3>
        </div>
        
        <div v-for="exp in options.profile.workExperience" class="mb-6 last:mb-0">
          <div class="grid grid-cols-1 md:grid-cols-[3fr,7fr,3fr] gap-2 mb-2">
            <div class="font-bold flex items-center space-x-2">
              <span>{{exp.workFor}}-{{exp.jobName}}</span>
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-600" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z" />
              </svg>
              <span class="text-gray-600">{{exp.position}}</span>
            </div>
            <div></div>
            <div class="text-right text-gray-600">
              {{dateToString(exp.startDate)}}至{{dateToString(exp.endDate)}}
            </div>
          </div>
          <p class="text-gray-700 leading-relaxed">{{exp.description}}</p>
        </div>
      </div>

      <!-- 专业技能 -->
      <div class="mb-8">
        <div class="flex items-center mb-4">
          <div class="flex justify-center items-center w-10 h-10 bg-[#7e90a5] rounded-lg mr-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="#f6de74" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-[#7e90a5]">专业技能</h3>
        </div>
        
        <div v-for="skill in options.profile.skills" class="mb-4 last:mb-0">
          <h4 class="font-bold text-gray-800 mb-2">{{skill.title}}</h4>
          <p class="text-gray-700 leading-relaxed">{{skill.description}}</p>
        </div>
      </div>

      <!-- 竞赛经历 -->
      <div class="mb-8">
        <div class="flex items-center mb-4">
          <div class="flex justify-center items-center w-10 h-10 bg-[#7e90a5] rounded-lg mr-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="#f6de74" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-[#7e90a5]">竞赛经历</h3>
        </div>
        
        <div v-for="race in options.profile.races" class="mb-4 last:mb-0">
          <div class="grid grid-cols-1 md:grid-cols-[3fr,3fr,1fr] gap-2">
            <div class="font-bold text-gray-800">{{race.name}}-{{race.level}}</div>
            <div class="flex items-center space-x-2 text-gray-600">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4.26 10.147a60.438 60.438 0 0 0-.491 6.347A48.62 48.62 0 0 1 12 20.904a48.62 48.62 0 0 1 8.232-4.41 60.46 60.46 0 0 0-.491-6.347m-15.482 0a50.636 50.636 0 0 0-2.658-.813A59.906 59.906 0 0 1 12 3.493a59.903 59.903 0 0 1 10.399 5.84c-.896.248-1.783.52-2.658.814m-15.482 0A50.717 50.717 0 0 1 12 13.489a50.702 50.702 0 0 1 7.74-3.342M6.75 15a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Zm0 0v-3.675A55.378 55.378 0 0 1 12 8.443m-7.007 11.55A5.981 5.981 0 0 0 6.75 15.75v-1.5" />
              </svg>
              <span>{{race.level}}</span>
            </div>
            <div class="text-right text-gray-600">
              {{dateToString(race.date)}}
            </div>
          </div>
        </div>
      </div>

      <!-- 教育经历 -->
      <div class="mb-8">
        <div class="flex items-center mb-4">
          <div class="flex justify-center items-center w-10 h-10 bg-[#7e90a5] rounded-lg mr-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="#f6de74" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-[#7e90a5]">教育经历</h3>
        </div>
        
        <div v-for="edu in options.profile.education" class="mb-4 last:mb-0">
          <div class="grid grid-cols-1 md:grid-cols-[1fr,3fr,1fr] gap-2">
            <div class="font-bold text-gray-800">{{edu.name}}-{{edu.level}}</div>
            <div class="flex items-center space-x-2 text-gray-600">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" d="M4.26 10.147a60.438 60.438 0 0 0-.491 6.347A48.62 48.62 0 0 1 12 20.904a48.62 48.62 0 0 1 8.232-4.41 60.46 60.46 0 0 0-.491-6.347m-15.482 0a50.636 50.636 0 0 0-2.658-.813A59.906 59.906 0 0 1 12 3.493a59.903 59.903 0 0 1 10.399 5.84c-.896.248-1.783.52-2.658.814m-15.482 0A50.717 50.717 0 0 1 12 13.489a50.702 50.702 0 0 1 7.74-3.342M6.75 15a.75.75 0 1 0 0-1.5.75.75 0 0 0 0 1.5Zm0 0v-3.675A55.378 55.378 0 0 1 12 8.443m-7.007 11.55A5.981 5.981 0 0 0 6.75 15.75v-1.5" />
              </svg>
              <span>{{edu.level}}</span>
            </div>
            <div class="text-right text-gray-600">
              {{dateToString(edu.startDate)}}至{{dateToString(edu.endDate)}}
            </div>
          </div>
        </div>
      </div>

      <!-- 个人项目 -->
      <div class="mb-8">
        <div class="flex items-center mb-4">
          <div class="flex justify-center items-center w-10 h-10 bg-[#7e90a5] rounded-lg mr-3">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="#f6de74" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" d="M11.48 3.499a.562.562 0 0 1 1.04 0l2.125 5.111a.563.563 0 0 0 .475.345l5.518.442c.499.04.701.663.321.988l-4.204 3.602a.563.563 0 0 0-.182.557l1.285 5.385a.562.562 0 0 1-.84.61l-4.725-2.885a.562.562 0 0 0-.586 0L6.982 20.54a.562.562 0 0 1-.84-.61l1.285-5.386a.562.562 0 0 0-.182-.557l-4.204-3.602a.562.562 0 0 1 .321-.988l5.518-.442a.563.563 0 0 0 .475-.345L11.48 3.5Z" />
            </svg>
          </div>
          <h3 class="text-xl font-bold text-[#7e90a5]">个人项目</h3>
        </div>
        
        <div v-for="pro in options.profile.projects" class="mb-4 last:mb-0">
          <h4 class="font-bold text-gray-800 mb-2">{{pro.title}}</h4>
          <p class="text-gray-700 leading-relaxed">{{pro.description}}</p>
        </div>
      </div>
    </div>
  </div>
</template>