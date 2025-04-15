<script setup>
import {onBeforeMount, reactive} from "vue";
import {useProfileStore} from "@/stores/profileStore.js";
import template1 from '@/assets/imgs/templates/template1.webp'
import template2 from '@/assets/imgs/templates/template2.webp'
import template3 from '@/assets/imgs/templates/template3.webp'
import template4 from '@/assets/imgs/templates/template4.webp'

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
  isEditing: false,
  addFormActiveKey: 1,
  currentEditingProfile: {
    name: '马越越',
    jobName: 'Java后端工程师',
    phoneNumber: '15695533649',
    emailAddress: 'w2084151024@gmail.com',
    position: '芜湖',
    birthDate: dayjs('2004/08/13', 'YYYY/MM/DD'),
    webSite: 'https://flyinsky.wiki',
    github: 'https://github.com/Flyinsky2004',
    linkIn: 'https://www.linkedin.com/in/%E5%90%89%E9%A2%96-%E7%8E%8B-16a669295/',
    workExperience: [],
    skills: [],
    education: [],
    races: [],
    projects: [],
    avatar: ''
  },
  currentEditingExperience: {
    workFor: 'FastCV',
    jobName: '后端工程师',
    position: '芜湖',
    description: 'API 设计与实现：负责设计和实现服务端 API，确保数据传输高效、稳定，同时提供简洁易用的接口文档。\n' +
        '\n' +
        '数据库设计与管理：负责数据库模型的设计和优化，确保数据存储的结构合理性，并使用 MyBatis 或 JPA 等持久层框架来进行数据库交互操作。\n' +
        '\n' +
        '业务逻辑实现：根据产品需求实现复杂的业务逻辑，包括用户管理、简历解析、职位匹配、面试推荐等功能模块。\n' +
        '\n' +
        '权限控制：通过过滤器和拦截器等方式，实现细粒度的权限控制，以保证数据安全。\n' +
        '\n' +
        '系统性能优化：分析性能瓶颈并进行优化，例如数据库索引优化、缓存（如 Redis）引入等，以提升系统响应速度和并发能力。\n' +
        '\n' +
        '单元测试与集成测试：编写测试用例，确保核心功能模块的可靠性和稳定性，使用 JUnit、Mockito 等工具进行单元测试和集成测试。\n' +
        '\n' +
        '错误处理与日志记录：为系统设置全面的错误处理机制，并实现日志系统，用于错误追踪和性能监控，保障系统稳定性。\n' +
        '\n' +
        '接口安全与数据保护：实现 JWT 认证等安全机制，确保用户数据隐私和接口安全性。\n' +
        '\n' +
        '与前端及移动端协作：与前端和移动端开发团队紧密合作，确保数据接口和业务逻辑的准确性，实现更好的用户体验。\n' +
        '\n' +
        '持续集成与部署：参与 CI/CD 流程，确保代码的高效发布与上线，以及系统的稳定运行。',
    startDate: dayjs(new Date(2024, 11, 5)),
    endDate: dayjs(new Date(2025, 4, 15))
  },
  currentEditingSkill: {
    title: '熟悉Spring,SpringMVC,SpringBoot等流行框架',
    description: '能够对业务流程和逻辑有自己的理解和抽象能力；熟悉Spring,SpringMVC,SpringBoot框架开发，对关系型数据库Mysql以及非关系型数据库Redis有一定的理解，熟悉高并发,MVCC，锁机制...了解SpringCloud分布式框架，熟悉SpringCloud Alibaba开发流程'
  },
  currentEditingEducation: {
    name: '安徽信息工程学院',
    level: '本科',
    profess: '计算机科学与技术',
    startDate: dayjs(new Date(2022, 10, 1)),
    endDate: dayjs(new Date(2026, 7, 1))
  },
  currentEditingRace: {
    name: '第十四届全国软件和信息技术专业人才大赛',
    level: '省级三等奖',
    date: dayjs(new Date(2023, 5, 19))
  },
  currentEditingProjects: {
    title: '智能之翼打卡平台',
    description: 'AI职探是我开发的一款用于实验室成员打卡，记录成长的平台，具有交流社区，在线编写，调试和运行算法题代码，验证题解，值日表功能。'
  }
})
const initData = () => {
  options.profiles = JSON.parse(localStorage.getItem("profiles"))
  if (options.profiles === null) {
    options.profiles = []
    options.profiles.push(defaultProfile)
  }

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
  localStorage.removeItem("profiles")
  options.profiles = []
  options.isDeleteModalOpen = false
  messageApi.success("所有资料已清空")
}
const saveDataToLocalStorage = () => {
  try {
    localStorage.setItem("profiles", JSON.stringify(options.profiles));
  } catch (e) {
    messageApi.error("发生错误：" + e)
  } finally {
    messageApi.success("保存成功！")
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
</script>
<template>
  <contextHolder/>
  <a-modal v-model:open="options.isEditing" title="添加新资料"
           @ok="options.profiles.push(options.currentEditingProfile);options.isEditing = false;" cancel-text="取消"
           ok-text="添加">
    <a-collapse v-model:activeKey="options.addFormActiveKey">
      <a-collapse-panel key="1" header="基本资料">
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>头像</a>
          <input
              class="basic-blue-input"
              type="file"
              @change="handleFileSelect"
              accept="image/*"
          >
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>姓名</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.name"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>应聘岗位</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.jobName"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>电话</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.phoneNumber"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>电子邮箱</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.emailAddress"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>城市</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.position"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>生日</a>
          <a-date-picker class="basic-blue-input" v-model:value="options.currentEditingProfile.birthDate"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>个人网站</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.webSite"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>Github主页</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.github"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>领英主页</a>
          <input class="basic-blue-input" v-model="options.currentEditingProfile.linkIn"/>
        </div>
      </a-collapse-panel>
      <a-collapse-panel key="2" header="工作经历">
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>工作单位</a>
          <input class="basic-blue-input" v-model="options.currentEditingExperience.workFor"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>岗位</a>
          <input class="basic-blue-input" v-model="options.currentEditingExperience.jobName"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>城市</a>
          <input class="basic-blue-input" v-model="options.currentEditingExperience.position"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>描述</a>
          <textarea class="basic-blue-input" v-model="options.currentEditingExperience.description"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>起始时间</a>
          <a-date-picker class="basic-blue-input" v-model:value="options.currentEditingExperience.startDate"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>离职时间</a>
          <a-date-picker class="basic-blue-input" v-model:value="options.currentEditingExperience.endDate"/>
        </div>
        <div>
          <button @click="options.currentEditingProfile.workExperience.push(options.currentEditingExperience);
          messageApi.success('保存成功！您可以继续编写下一个工作经历！');"
                  class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700">保存并继续
          </button>
        </div>
      </a-collapse-panel>
      <a-collapse-panel key="3" header="竞赛经历">
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>竞赛名称</a>
          <input class="basic-blue-input" v-model="options.currentEditingRace.name"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>奖项</a>
          <input class="basic-blue-input" v-model="options.currentEditingRace.level"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>获得时间</a>
          <a-date-picker class="basic-blue-input" v-model:value="options.currentEditingRace.date"/>
        </div>
        <div>
          <button @click="options.currentEditingProfile.races.push(options.currentEditingRace);
          messageApi.success('保存成功！您可以继续编写下一个工作经历！');"
                  class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700">保存并继续
          </button>
        </div>
      </a-collapse-panel>
      <a-collapse-panel key="4" header="教育经历">
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>学校名称</a>
          <input class="basic-blue-input" v-model="options.currentEditingEducation.name"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>学历</a>
          <input class="basic-blue-input" v-model="options.currentEditingEducation.level"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>专业方向</a>
          <input class="basic-blue-input" v-model="options.currentEditingEducation.profess"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>入学时间</a>
          <a-date-picker class="basic-blue-input" v-model:value="options.currentEditingEducation.startDate"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>毕业时间</a>
          <a-date-picker class="basic-blue-input" v-model:value="options.currentEditingEducation.endDate"/>
        </div>
        <div>
          <button @click="options.currentEditingProfile.education.push(options.currentEditingEducation);
          messageApi.success('保存成功！您可以继续编写下一个工作经历！');"
                  class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700">保存并继续
          </button>
        </div>
      </a-collapse-panel>
      <a-collapse-panel key="5" header="个人项目">
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>项目名称</a>
          <input class="basic-blue-input" v-model="options.currentEditingProjects.title"/>
        </div>
        <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
          <a>简介</a>
          <textarea class="basic-blue-input" v-model="options.currentEditingProjects.description"/>
        </div>
        <div>
          <button @click="options.currentEditingProfile.projects.push(options.currentEditingProjects);
          messageApi.success('保存成功！您可以继续编写下一个工作经历！');"
                  class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700">保存并继续
          </button>
        </div>
      </a-collapse-panel>
    </a-collapse>
  </a-modal>
  <a-modal v-model:open="options.isDeleteModalOpen" title="确认删除" @ok="clearData()" cancel-text="取消"
           ok-text="我心已决" ok-type="danger">
    <a>你确定要</a><a class="text-red-600">删除所有</a><a>已有资料吗</a>
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
          <button @click="options.isEditing = true"
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
             v-for="profile in options.profiles"
             @click="profileStore.setProfile(profile);messageApi.success('成功选择资料:'+profile.jobName+profile.name)">
          <img class="w-20 h-20 my-auto mt-4" :src="profile.avatar" alt="">
          <div class="font-bold mt-1">{{profile.name}}</div>
          <div class="mt-1">{{profile.jobName}}</div>
          <div class="mt-1 mb-4">{{profile.position}}</div>
        </div>
      </div>
      <div class="text-center text-gray-500 h-32 flex" v-else>
        <span class="mx-auto my-auto">阿哦，你好像还没有任何录入的个人资料...<a
            class="text-blue-500 hover:underline cursor-pointer font-bold">录入一个</a></span>
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