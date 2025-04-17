<script setup>
import {onBeforeMount, reactive} from "vue";
import {useProfileStore} from "@/stores/profileStore.js";
import template1 from '@/assets/imgs/templates/template1.webp'
import template2 from '@/assets/imgs/templates/template2.webp'
import template3 from '@/assets/imgs/templates/template3.webp'
import template4 from '@/assets/imgs/templates/template4.webp'
import {get, post, del} from "@/util/request.js";

import {message} from 'ant-design-vue';
import router from "@/router/index.js";
import {dayjs} from "element-plus";
import {defaultProfile} from '@/hooks/data'

const [messageApi, contextHolder] = message.useMessage();
const profileStore = useProfileStore()
const templates = [
  {
    id: 1,
    name: '蓝调映像',
    description: '极简线条，精简干练',
    cover: template1
  },{
    id: 2,
    name: '垂直橙线',
    description: '橘橙线条，向前冲冲',
    cover: template2
  },{
    id: 3,
    name: '灰色商务',
    description: '商务风格，专业团队',
    cover: template3
  },{
    id: 4,
    name: '清新职道',
    description: '简洁明快，优雅大方',
    cover: template4
  }
]

const options = reactive({
  profiles: [],
  isDeleteModalOpen: false,
  isDeleteSingleProfileModalOpen: false,
  deletingProfileIndex: -1
})
const initData = () => {
  get('/api/profile/get', null, 
    (message, data) => {
      options.profiles = data || [];
      if (options.profiles.length === 0) {
        options.profiles.push(defaultProfile);
      }
    },
    (error) => {
      messageApi.error('获取资料失败：' + error);
      options.profiles = [defaultProfile];
    }
  );
}
onBeforeMount(() => {
  initData()
})
const templateClickHandler = (id) => {
  if (!profileStore.getCurrentProfile.avatar) {
    messageApi.warning('需要先选择资料哦！')
  } else {
    router.push('/template/' + id)
  }
}
const clearData = () => {
  del('/api/profile/delete', null,
    () => {
      options.profiles = [];
      options.isDeleteModalOpen = false;
      messageApi.success("所有资料已清空");
    },
    (error) => {
      messageApi.error("清空资料失败：" + error);
    }
  );
}
const saveDataToLocalStorage = () => {
  try {
    post('/api/profile/create', options.profiles,
      () => {
        messageApi.success("保存成功！");
      },
      (error) => {
        messageApi.error("保存失败：" + error);
      }
    );
  } catch (e) {
    messageApi.error("发生错误：" + e);
  }
}
const fileToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result);
    reader.onerror = (error) => reject(error);
    reader.readAsDataURL(file);
  });
}

// 处理文件选择事件
const handleFileSelect = async (event) => {
  const file = event.target.files[0];
  if (!file) {
    console.error('No file selected');
    return;
  }

  try {
    const base64String = await fileToBase64(file);
    options.currentEditingProfile.avatar = base64String; // 使用.value来设置ref的值
  } catch (error) {
    console.error('Error converting image:', error);
  }
}

// 删除单个资料
const deleteSingleProfile = () => {
  if (options.deletingProfileIndex >= 0 && options.deletingProfileIndex < options.profiles.length) {
    const profileToDelete = options.profiles[options.deletingProfileIndex];
    del('/api/profile/delete', { id: profileToDelete.id },
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
}

// 添加跳转到编辑页面的函数
const navigateToEdit = (index) => {
  router.push(`/edit/${index}`);
};

// 跳转到添加新资料页面
const navigateToAddNew = () => {
  router.push('/edit/-1');
};
</script>
<template>
  <contextHolder/>
  <a-modal v-model:open="options.isDeleteModalOpen" title="确认删除" @ok="clearData()" cancel-text="取消"
           ok-text="我心已决" ok-type="danger">
    <a>你确定要</a><a class="text-red-600">删除所有</a><a>已有资料吗</a>
  </a-modal>
  <a-modal v-model:open="options.isDeleteSingleProfileModalOpen" title="确认删除" 
           @ok="deleteSingleProfile()" cancel-text="取消"
           ok-text="删除" ok-type="danger">
    <p>你确定要删除 <span class="text-red-600 font-bold">{{ options.deletingProfileIndex >= 0 ? options.profiles[options.deletingProfileIndex].name : '' }}</span> 的资料吗？</p>
    <p class="text-gray-500">此操作不可撤销。</p>
  </a-modal>
  <div class="h-screen overflow-y-auto bg-white bg-opacity-70 animate__animated animate__fadeIn">
    <div class="h-10 bg-white shadow-md shadow-gray-400 place-items-center flex">
      <div class="font-bold text-xl flex-grow text-center title bg-clip-text bg-gradient-to-r from-purple-500 via-red-500 to-pink-500 text-transparent select-none cursor-pointer" @click="router.push('/')"  >FastCV</div>
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
    <div
        class="h-fit min-h-screen bg-blue-200 bg-opacity-10 backdrop-blur-md shadow-xl w-3/4 mx-auto rounded-xl mt-2 p-4 hover:shadow-lg transition-all transition-duration-300">
      <div>

      </div>
      <div class="flex">
        <h1 class="font-bold text-2xl flex-grow">个人资料</h1>
        <div class="grid grid-cols-[1fr,1fr,1fr] gap-2">
          <button @click="navigateToAddNew()"
                  class="basic-button bg-blue-500 hover:bg-blue-600 active:bg-blue-700">新建数据
          </button>
          <button class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700"
                  @click="saveDataToLocalStorage">保存数据
          </button>
          <button class="basic-button bg-red-500 hover:bg-red-600 active:bg-red-700"
                  @click="options.isDeleteModalOpen = true">清空数据
          </button>
        </div>
      </div>
      <div v-if="options.profiles.length > 0" class="grid grid-cols-4 gap-2">
        <div class="bg-white p-4 shadow grid grid-cols-1 place-items-center cursor-pointer hover:shadow-lg transition-all transition-duration-300 active:bg-gray-100 select-none"
             v-for="(profile, index) in options.profiles"
             @click="profileStore.setProfile(profile);messageApi.success('成功选择资料:'+profile.jobName+profile.name)">
          <img class="w-20 h-20 my-auto mt-4" :src="profile.avatar" alt="">
          <div class="font-bold mt-1">{{profile.name}}</div>
          <div class="mt-1">{{profile.jobName}}</div>
          <div class="mt-1 mb-4">{{profile.position}}</div>
          <div class="flex space-x-2">
            <a-button type="primary" size="small" @click.stop="navigateToEdit(index)">编辑</a-button>
            <a-button type="danger" size="small" @click.stop="options.deletingProfileIndex = index; options.isDeleteSingleProfileModalOpen = true">删除</a-button>
          </div>
        </div>
      </div>
      <div class="text-center text-gray-500 h-32 flex" v-else>
        <span class="mx-auto my-auto">阿哦，你好像还没有任何录入的个人资料...<a
            class="text-blue-500 hover:underline cursor-pointer font-bold" @click="navigateToAddNew()">录入一个</a></span>
      </div>
      <div class="flex mt-4">
        <h1 class="font-bold text-2xl flex-grow">简历模版</h1>
        <h1 v-if="!profileStore.isSelected" class="text-red-600">警告！您还没有选择模版渲染的个人资料!</h1>
        <h1 v-else class="text-blue-600">简历使用资料：{{profileStore.getCurrentProfile.name}}-{{profileStore.getCurrentProfile.jobName}}</h1>
      </div>
      <div class="grid grid-cols-4 gap-4">
        <div :class="profileStore.getCurrentProfile.name === null?'cursor-no-drop':''"
             class="bg-white p-2 select-none shadow grid grid-cols-1 place-items-center cursor-pointer hover:shadow-lg transition-all transition-duration-300 active:bg-gray-100"
             v-for="template in templates"
             @click="templateClickHandler(template.id)">
          <img class="my-auto mt-4" :src="template.cover" alt="">
          <div class="font-bold">{{ template.name }}</div>
          <div class="text-sm mb-4">{{ template.description }}</div>
        </div>
      </div>
    </div>
  </div>
</template>