<script setup>
import { onBeforeMount, reactive } from "vue";
import { useProfileStore } from "@/stores/profileStore.js";
import template1 from "@/assets/imgs/templates/template1.webp";
import template2 from "@/assets/imgs/templates/template2.webp";
import template3 from "@/assets/imgs/templates/template3.webp";
import template4 from "@/assets/imgs/templates/template4.webp";
import { get, post, del } from "@/util/request.js";

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
  <div
    class="h-screen overflow-y-auto bg-white bg-opacity-70 animate__animated animate__fadeIn"
  >
    
    <div
      class="h-fit min-h-screen bg-blue-200 bg-opacity-10 backdrop-blur-md shadow-xl w-3/4 mx-auto rounded-xl mt-2 p-4 hover:shadow-lg transition-all transition-duration-300"
    >
      <div></div>
      <div class="flex">
        <h1 class="font-bold text-2xl flex-grow">个人资料</h1>
        <div class="grid grid-cols-[1fr,1fr,1fr] gap-2">
          <button
            @click="navigateToAddNew()"
            class="basic-button bg-blue-500 hover:bg-blue-600 active:bg-blue-700"
          >
            新建数据
          </button>
          <button
            class="basic-button bg-red-500 hover:bg-red-600 active:bg-red-700"
            @click="options.isDeleteModalOpen = true"
          >
            清空数据
          </button>
        </div>
      </div>
      <div v-if="options.profiles.length > 0" class="grid grid-cols-4 gap-2">
        <div
          class="bg-white p-4 shadow grid grid-cols-1 place-items-center cursor-pointer hover:shadow-lg transition-all transition-duration-300 active:bg-gray-100 select-none"
          v-for="(profile, index) in options.profiles"
          @click="
            profileStore.setProfile(profile.content);
            messageApi.success(
              '成功选择资料:' + profile.content.jobName + profile.content.name
            );
          "
        >
          <img
            class="w-20 h-20 my-auto mt-4"
            :src="profile.content.avatar"
            alt=""
          />
          <div class="font-bold mt-1">{{ profile.content.name }}</div>
          <div class="mt-1">{{ profile.content.jobName }}</div>
          <div class="mt-1 mb-4">{{ profile.content.position }}</div>
          <div class="flex space-x-2">
            <a-button
              type="primary"
              size="small"
              @click.stop="navigateToEdit(index)"
              >编辑</a-button
            >
            <a-button
              type="danger"
              size="small"
              @click.stop="
                options.deletingProfileIndex = index;
                options.isDeleteSingleProfileModalOpen = true;
              "
              >删除</a-button
            >
          </div>
        </div>
      </div>
      <div class="text-center text-gray-500 h-32 flex" v-else>
        <span class="mx-auto my-auto"
          >阿哦，你好像还没有任何录入的个人资料...<a
            class="text-blue-500 hover:underline cursor-pointer font-bold"
            @click="navigateToAddNew()"
            >录入一个</a
          ></span
        >
      </div>
      <div class="flex mt-4">
        <h1 class="font-bold text-2xl flex-grow">简历模版</h1>
        <h1 v-if="!profileStore.isSelected" class="text-red-600">
          警告！您还没有选择模版渲染的个人资料!
        </h1>
        <h1 v-else class="text-blue-600">
          简历使用资料：{{ profileStore.getCurrentProfile.name }}-{{
            profileStore.getCurrentProfile.jobName
          }}
        </h1>
      </div>
      <div class="grid grid-cols-4 gap-4">
        <div
          :class="
            profileStore.getCurrentProfile.name === null ? 'cursor-no-drop' : ''
          "
          class="bg-white p-2 select-none shadow grid grid-cols-1 place-items-center cursor-pointer hover:shadow-lg transition-all transition-duration-300 active:bg-gray-100"
          v-for="template in templates"
          @click="templateClickHandler(template.id)"
        >
          <img class="my-auto mt-4" :src="template.cover" alt="" />
          <div class="font-bold">{{ template.name }}</div>
          <div class="text-sm mb-4">{{ template.description }}</div>
        </div>
      </div>
    </div>
  </div>
</template>
