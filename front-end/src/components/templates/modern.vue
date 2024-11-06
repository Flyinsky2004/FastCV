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

  <div class="bg-white">
    <div class="h-8"></div>
    <div class="h-[2px] bg-orange-600"></div>
    <div class="grid grid-cols-1 mt-2">
      <!--       主要个人资料-->
      <div class="grid grid-cols-1 place-items-start">
        <!--          姓名，岗位-->
        <div class="grid grid-cols-2 place-items-center">
          <img class="h-20 w-20" :src="options.profile.avatar" alt="">
          <div class="grid grid-cols-1">
            <h1 class="text-4xl font-bold">{{ options.profile.name }}</h1>
            <h1 class="text-orange-600">{{ options.profile.jobName }}</h1>
          </div>
        </div>
      </div>
      <div class="grid grid-cols-2 gap-4 place-items-start">
        <div class="flex">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="size-5">
            <path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z"/>
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z"/>
          </svg>
          {{ options.profile.position }}
        </div>
        <div class="flex place-items-center text-orange-500 ">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="size-5">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M21.75 6.75v10.5a2.25 2.25 0 0 1-2.25 2.25h-15a2.25 2.25 0 0 1-2.25-2.25V6.75m19.5 0A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25m19.5 0v.243a2.25 2.25 0 0 1-1.07 1.916l-7.5 4.615a2.25 2.25 0 0 1-2.36 0L3.32 8.91a2.25 2.25 0 0 1-1.07-1.916V6.75"/>
          </svg>
          <a :href="'mailto:'+options.profile.emailAddress">{{ options.profile.emailAddress }}</a>
        </div>
        <div class="flex place-items-center">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="size-5">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M2.25 6.75c0 8.284 6.716 15 15 15h2.25a2.25 2.25 0 0 0 2.25-2.25v-1.372c0-.516-.351-.966-.852-1.091l-4.423-1.106c-.44-.11-.902.055-1.173.417l-.97 1.293c-.282.376-.769.542-1.21.38a12.035 12.035 0 0 1-7.143-7.143c-.162-.441.004-.928.38-1.21l1.293-.97c.363-.271.527-.734.417-1.173L6.963 3.102a1.125 1.125 0 0 0-1.091-.852H4.5A2.25 2.25 0 0 0 2.25 4.5v2.25Z"/>
          </svg>
          <a>{{ options.profile.phoneNumber }}</a>
        </div>
        <div class="flex place-items-center">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="size-5">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M12 8.25v-1.5m0 1.5c-1.355 0-2.697.056-4.024.166C6.845 8.51 6 9.473 6 10.608v2.513m6-4.871c1.355 0 2.697.056 4.024.166C17.155 8.51 18 9.473 18 10.608v2.513M15 8.25v-1.5m-6 1.5v-1.5m12 9.75-1.5.75a3.354 3.354 0 0 1-3 0 3.354 3.354 0 0 0-3 0 3.354 3.354 0 0 1-3 0 3.354 3.354 0 0 0-3 0 3.354 3.354 0 0 1-3 0L3 16.5m15-3.379a48.474 48.474 0 0 0-6-.371c-2.032 0-4.034.126-6 .371m12 0c.39.049.777.102 1.163.16 1.07.16 1.837 1.094 1.837 2.175v5.169c0 .621-.504 1.125-1.125 1.125H4.125A1.125 1.125 0 0 1 3 20.625v-5.17c0-1.08.768-2.014 1.837-2.174A47.78 47.78 0 0 1 6 13.12M12.265 3.11a.375.375 0 1 1-.53 0L12 2.845l.265.265Zm-3 0a.375.375 0 1 1-.53 0L9 2.845l.265.265Zm6 0a.375.375 0 1 1-.53 0L15 2.845l.265.265Z"/>
          </svg>
          <a>{{ dateToString(options.profile.birthDate) }}</a>
        </div>
        <div v-if="options.profile.webSite" class="flex flex-nowrap text-orange-500  place-items-center">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="size-5">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M13.19 8.688a4.5 4.5 0 0 1 1.242 7.244l-4.5 4.5a4.5 4.5 0 0 1-6.364-6.364l1.757-1.757m13.35-.622 1.757-1.757a4.5 4.5 0 0 0-6.364-6.364l-4.5 4.5a4.5 4.5 0 0 0 1.242 7.244"/>
          </svg>
          <a :href="options.profile.webSite">个人网站</a>
        </div>
        <div v-if="options.profile.github" class="flex flex-nowrap text-orange-500  place-items-center">
          <img class="bg-white rounded-xl" width="20" :src="github">
          <a :href="options.profile.github">Github</a>
        </div>
        <div v-if="options.profile.webSite" class="flex flex-nowrap text-orange-500  place-items-center">
          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5"
               stroke="currentColor" class="size-5">
            <path stroke-linecap="round" stroke-linejoin="round"
                  d="M3.75 21h16.5M4.5 3h15M5.25 3v18m13.5-18v18M9 6.75h1.5m-1.5 3h1.5m-1.5 3h1.5m3-6H15m-1.5 3H15m-1.5 3H15M9 21v-3.375c0-.621.504-1.125 1.125-1.125h3.75c.621 0 1.125.504 1.125 1.125V21"/>
          </svg>
          <a :href="options.profile.linkIn">LinkIn</a>
        </div>
      </div>
    </div>

    <div class="grid grid-cols-[6fr,4fr] place-items-start">
      <div class="grid grid-cols-1 gap-4">
        <div class="p-1 font-bold w-full">
          <div class="h-[1px] bg-gray-400 mt-2 mb-2"></div>
          <h1 class="text-2xl">工作经历</h1>
          <div class="grid grid-cols-[1fr,5fr]"  v-for="exp in options.profile.workExperience">
            <div>
              <div class="h-[2px] bg-orange-600"></div>
              <span>{{ dateToString(exp.startDate) }}<br>至{{ dateToString(exp.endDate) }}</span>
            </div>
            <div class="grid grid-cols-1 gap-2">
              <div class="text-orange-500">
                {{ exp.workFor }}
              </div>
              <div class="text-gray-500">
                {{ exp.jobName }}
              </div>
              <div class="font-light">
                {{ exp.description }}
              </div>
            </div>
          </div>
        </div>
        <div class="p-1 font-bold w-full">
          <div class="h-[1px] bg-gray-400 mt-2 mb-2"></div>
          <h1 class="text-2xl">个人项目</h1>
          <div class="grid grid-cols-[1fr,5fr]"  v-for="pro in options.profile.projects">
            <div class="place-items-end p-2">
              <div class="h-2 w-2 rounded"></div>
            </div>
            <div class="grid grid-cols-1 mb-2 gap-1">
              <div class="text-orange-500">
                {{ pro.title }}
              </div>
              <div class="text-black font-light">
                {{ pro.description }}
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="grid grid-cols-1 gap-4">
        <div class="p-1 font-bold w-full">
          <div class="h-[1px] bg-gray-400 mt-2 mb-2"></div>
          <h1 class="text-2xl">教育经历</h1>
          <div v-for="edu in options.profile.education" class="grid grid-cols-1 gap-2">
            <div class="text-orange-500">
              {{ edu.level }}
            </div>
            <div class="text-gray-500 font-light">
              {{ dateToString(edu.startDate) }}至{{ dateToString(edu.endDate) }}
            </div>
            <div class="text-gray-500 font-light flex">
              <div class="flex-grow">
                {{ edu.name }}
              </div>
              <div>
                {{edu.profess}}
              </div>
            </div>
          </div>
        </div>
        <div class="p-1 font-bold w-full">
          <div class="h-[1px] bg-gray-400 mt-2 mb-2"></div>
          <h1 class="text-2xl">专业技能</h1>
          <div v-for="skill in options.profile.skills" class="mt-2">
            <div class="font-bold mt-2 text-orange-700">
              {{ skill.title }}
            </div>
            <div class="mt-2 font-light">
              {{ skill.description }}
            </div>
          </div>
        </div>
        <div class="p-1 font-bold w-full">
          <div class="h-[1px] bg-gray-400 mt-2 mb-2"></div>
          <h1 class="text-2xl">竞赛经历</h1>
          <div v-for="race in options.profile.races" class="mt-2 grid grid-cols-1">
            <div class="font-bold mt-2 text-orange-700">
              {{ race.name }}
            </div>
            <div class="mt-2 font-light flex">
              <div class="flex-grow">
                {{ race.level }}
              </div>
              <div>
                {{ dateToString(race.date) }}
              </div>
            </div>
          </div>
        </div>
      </div>
      </div>


    <div class="h-4"></div>
  </div>
</template>