<script setup>
import { onBeforeMount, reactive, ref } from "vue";
import { useProfileStore } from "@/stores/profileStore.js";
import template1 from "@/assets/imgs/templates/template1.webp";
import template2 from "@/assets/imgs/templates/template2.webp";
import template3 from "@/assets/imgs/templates/template3.webp";
import template4 from "@/assets/imgs/templates/template4.webp";
import { get, post, del } from "@/util/request.js";
import { useMotion } from '@vueuse/motion'

import { message } from "ant-design-vue";
import router from "@/router/index.js";
import { dayjs } from "element-plus";
import { defaultProfile } from "@/hooks/data";

const [messageApi, contextHolder] = message.useMessage();
const profileStore = useProfileStore();
const templates = [
  {
    id: 1,
    name: "蓝调映像",
    description: "极简线条，精简干练",
    cover: template1,
  },
  {
    id: 2,
    name: "垂直橙线",
    description: "橘橙线条，向前冲冲",
    cover: template2,
  },
  {
    id: 3,
    name: "灰色商务",
    description: "商务风格，专业团队",
    cover: template3,
  },
  {
    id: 4,
    name: "清新职道",
    description: "简洁明快，优雅大方",
    cover: template4,
  },
];

const options = reactive({
  profiles: [],
  isDeleteModalOpen: false,
  isDeleteSingleProfileModalOpen: false,
  deletingProfileIndex: -1,
});
const initData = () => {
  get(
    "/api/profile/get",
    null,
    (message, data) => {
      options.profiles = data || [];
      if (options.profiles.length === 0) {
        options.profiles.push(defaultProfile);
      } else {
        options.profiles.forEach((profile) => {
          try {
            profile.content = JSON.parse(profile.content);
          } catch (error) {
            console.error(
              `Error parsing content for profile: ${profile}`,
              error
            );
          }
        });
      }
    },
    (error) => {
      messageApi.error("获取资料失败：" + error);
      options.profiles = [defaultProfile];
    }
  );
};
onBeforeMount(() => {
  initData();
});
const templateClickHandler = (id) => {
  if (!profileStore.getCurrentProfile.avatar) {
    messageApi.warning("需要先选择资料哦！");
  } else {
    router.push("/template/" + id);
  }
};
const clearData = () => {
  del(
    "/api/profile/delete",
    null,
    () => {
      options.profiles = [];
      options.isDeleteModalOpen = false;
      messageApi.success("所有资料已清空");
    },
    (error) => {
      messageApi.error("清空资料失败：" + error);
    }
  );
};

const fileToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result);
    reader.onerror = (error) => reject(error);
    reader.readAsDataURL(file);
  });
};

// 处理文件选择事件
const handleFileSelect = async (event) => {
  const file = event.target.files[0];
  if (!file) {
    console.error("No file selected");
    return;
  }

  try {
    const base64String = await fileToBase64(file);
    options.currentEditingProfile.avatar = base64String; // 使用.value来设置ref的值
  } catch (error) {
    console.error("Error converting image:", error);
  }
};

// 删除单个资料
const deleteSingleProfile = () => {
  if (
    options.deletingProfileIndex >= 0 &&
    options.deletingProfileIndex < options.profiles.length
  ) {
    const profileToDelete = options.profiles[options.deletingProfileIndex];
    del(
      "/api/profile/delete",
      { id: profileToDelete.id },
      () => {
        options.profiles.splice(options.deletingProfileIndex, 1);
        messageApi.success("删除成功");
      },
      (error) => {
        messageApi.error("删除失败：" + error);
      }
    );
  }
  options.isDeleteSingleProfileModalOpen = false;
  options.deletingProfileIndex = -1;
};

// 添加跳转到编辑页面的函数
const navigateToEdit = (index) => {
  router.push(`/edit/${options.profiles[index].ID}`);
};

// 跳转到添加新资料页面
const navigateToAddNew = () => {
  router.push("/edit/-1");
};

// 添加动画引用
const profilesSection = ref(null)
const templatesSection = ref(null)

// 设置动画
const profilesMotion = useMotion(profilesSection, {
  initial: {
    opacity: 0,
    y: 50,
  },
  enter: {
    opacity: 1,
    y: 0,
    transition: {
      duration: 800,
    },
  },
})

const templatesMotion = useMotion(templatesSection, {
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
  <contextHolder />
  <a-modal
    v-model:open="options.isDeleteModalOpen"
    title="确认删除"
    @ok="clearData()"
    cancel-text="取消"
    ok-text="我心已决"
    ok-type="danger"
  >
    <a>你确定要</a><a class="text-red-600">删除所有</a><a>已有资料吗</a>
  </a-modal>
  <a-modal
    v-model:open="options.isDeleteSingleProfileModalOpen"
    title="确认删除"
    @ok="deleteSingleProfile()"
    cancel-text="取消"
    ok-text="删除"
    ok-type="danger"
  >
    <p>
      你确定要删除
      <span class="text-red-600 font-bold">{{
        options.deletingProfileIndex >= 0
          ? options.profiles[options.deletingProfileIndex].name
          : ""
      }}</span>
      的资料吗？
    </p>
    <p class="text-gray-500">此操作不可撤销。</p>
  </a-modal>
  <div class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-7xl mx-auto">
      <!-- 个人资料部分 -->
      <div ref="profilesSection" class="bg-white rounded-2xl shadow-xl p-6 mb-8 animate__animated animate__fadeInUp">
        <div class="flex items-center justify-between mb-6">
          <h1 class="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-blue-600 to-indigo-600 animate__animated animate__fadeInLeft">个人资料</h1>
          <div class="flex space-x-4 animate__animated animate__fadeInRight">
            <button
              @click="navigateToAddNew()"
              class="px-4 py-2 bg-gradient-to-r from-blue-500 to-blue-600 text-white rounded-lg hover:shadow-lg transition-all duration-300 transform hover:-translate-y-1 animate__animated animate__pulse animate__infinite animate__slow"
            >
              新建数据
            </button>
            <button
              class="px-4 py-2 bg-gradient-to-r from-red-500 to-red-600 text-white rounded-lg hover:shadow-lg transition-all duration-300 transform hover:-translate-y-1"
              @click="options.isDeleteModalOpen = true"
            >
              清空数据
            </button>
          </div>
        </div>

        <div v-if="options.profiles.length > 0" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div
            v-for="(profile, index) in options.profiles"
            :key="index"
            class="bg-white rounded-xl shadow-md hover:shadow-xl transition-all duration-300 overflow-hidden group animate__animated animate__fadeIn"
            :style="{ 'animation-delay': index * 0.1 + 's' }"
          >
            <div 
              class="p-6 cursor-pointer"
              @click="
                profileStore.setProfile(profile.content);
                messageApi.success('成功选择资料:' + profile.content.jobName + profile.content.name);
              "
            >
              <div class="flex flex-col items-center">
                <img
                  class="w-24 h-24 rounded-full object-cover mb-4 group-hover:scale-105 transition-transform duration-300"
                  :src="profile.content.avatar"
                  alt=""
                />
                <h3 class="text-xl font-semibold text-gray-800">{{ profile.content.name }}</h3>
                <p class="text-gray-600 mt-1">{{ profile.content.jobName }}</p>
                <p class="text-gray-500 mt-1">{{ profile.content.position }}</p>
              </div>
            </div>
            <div class="flex justify-center space-x-3 p-4 bg-gray-50">
              <button
                class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600 transition-colors duration-300"
                @click.stop="navigateToEdit(index)"
              >
                编辑
              </button>
              <button
                class="px-4 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 transition-colors duration-300"
                @click.stop="
                  options.deletingProfileIndex = index;
                  options.isDeleteSingleProfileModalOpen = true;
                "
              >
                删除
              </button>
            </div>
          </div>
        </div>
        <div v-else class="flex items-center justify-center h-32 bg-gray-50 rounded-lg animate__animated animate__fadeIn">
          <p class="text-gray-500">
            阿哦，你好像还没有任何录入的个人资料...
            <button
              class="text-blue-500 hover:text-blue-600 font-semibold hover:underline focus:outline-none animate__animated animate__pulse animate__infinite animate__slow"
              @click="navigateToAddNew()"
            >
              录入一个
            </button>
          </p>
        </div>
      </div>

      <!-- 简历模版部分 -->
      <div ref="templatesSection" class="bg-white rounded-2xl shadow-xl p-6 animate__animated animate__fadeInUp animate__delay-1s">
        <div class="flex items-center justify-between mb-6">
          <h1 class="text-3xl font-bold bg-clip-text text-transparent bg-gradient-to-r from-blue-600 to-indigo-600 animate__animated animate__fadeInLeft animate__delay-1s">简历模版</h1>
          <div>
            <h1 v-if="!profileStore.isSelected" class="text-red-500 font-medium animate__animated animate__headShake animate__infinite">
              警告！您还没有选择模版渲染的个人资料!
            </h1>
            <h1 v-else class="text-blue-600 font-medium animate__animated animate__fadeInRight animate__delay-1s">
              简历使用资料：{{ profileStore.getCurrentProfile.name }}-{{
                profileStore.getCurrentProfile.jobName
              }}
            </h1>
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <div
            v-for="template in templates"
            :key="template.id"
            :class="[
              'group bg-white rounded-xl shadow-md overflow-hidden transition-all duration-500 ease-in-out transform hover:-translate-y-2 hover:shadow-xl animate__animated animate__zoomIn',
              profileStore.getCurrentProfile.name === null ? 'cursor-not-allowed opacity-50' : 'cursor-pointer'
            ]"
            :style="{ 'animation-delay': template.id * 0.1 + 1 + 's' }"
            @click="templateClickHandler(template.id)"
          >
            <div class="relative overflow-hidden">
              <img 
                class="h-48 mx-auto object-cover transition-all duration-500 ease-in-out transform group-hover:scale-110" 
                :src="template.cover" 
                alt="" 
              />
            </div>
            <div class="p-4">
              <h3 class="text-lg font-semibold text-gray-800">{{ template.name }}</h3>
              <p class="text-sm text-gray-600 mt-1">{{ template.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.basic-button {
  @apply px-4 py-2 text-white rounded-lg transition-all duration-300 hover:shadow-lg;
}
</style>
