<script setup>
import {onBeforeMount, reactive, ref} from "vue";
import {message} from 'ant-design-vue';
import router from "@/router/index.js";
import dayjs from "dayjs";
import {defaultProfile} from '@/hooks/data';
import {get, put, postJSON, del} from '@/util/request';

const [messageApi, contextHolder] = message.useMessage();

// 获取路由参数
const profileIndex = ref(parseInt(router.currentRoute.value.params.index));
const isNewProfile = ref(profileIndex.value === -1);

// 编辑表单数据
const editingProfile = reactive({
  profile: {
    name: '',
    jobName: '',
    phoneNumber: '',
    emailAddress: '',
    position: '',
    birthDate: null,
    webSite: '',
    github: '',
    linkIn: '',
    workExperience: [],
    skills: [],
    education: [],
    races: [],
    projects: [],
    avatar: ''
  },
  addFormActiveKey: 1,
  currentEditingExperience: {
    workFor: '',
    jobName: '',
    position: '',
    description: '',
    startDate: dayjs(new Date()),
    endDate: dayjs(new Date())
  },
  currentEditingSkill: {
    title: '',
    description: ''
  },
  currentEditingEducation: {
    name: '',
    level: '',
    profess: '',
    startDate: dayjs(new Date()),
    endDate: dayjs(new Date())
  },
  currentEditingRace: {
    name: '',
    level: '',
    date: dayjs(new Date())
  },
  currentEditingProjects: {
    title: '',
    description: ''
  },
  // 记录当前正在编辑的对象索引，-1表示新增
  currentEditingIndex: {
    workExperience: -1,
    skill: -1,
    education: -1,
    race: -1,
    project: -1
  },
  // 是否处于编辑模式
  isEditing: {
    workExperience: false,
    skill: false,
    education: false,
    race: false,
    project: false
  }
});

// 页面加载时初始化数据
onBeforeMount(() => {
  initData();
});

// 初始化数据
const initData = () => {
  if (!isNewProfile.value && profileIndex.value >= 0) {
    // 从后端获取现有资料
    get(`/api/profile/${profileIndex.value}/get`, null, 
      (message, data) => {
        // 深拷贝以避免直接修改引用
        editingProfile.profile = JSON.parse(data.content);
        // 处理日期格式
        if (editingProfile.profile.birthDate) {
          editingProfile.profile.birthDate = dayjs(editingProfile.profile.birthDate);
        }
        
        // 处理工作经历中的日期
        if (editingProfile.profile.workExperience && editingProfile.profile.workExperience.length > 0) {
          editingProfile.profile.workExperience.forEach(exp => {
            if (exp.startDate) exp.startDate = dayjs(exp.startDate);
            if (exp.endDate) exp.endDate = dayjs(exp.endDate);
          });
        }
        
        // 处理教育经历中的日期
        if (editingProfile.profile.education && editingProfile.profile.education.length > 0) {
          editingProfile.profile.education.forEach(edu => {
            if (edu.startDate) edu.startDate = dayjs(edu.startDate);
            if (edu.endDate) edu.endDate = dayjs(edu.endDate);
          });
        }
        
        // 处理竞赛经历中的日期
        if (editingProfile.profile.races && editingProfile.profile.races.length > 0) {
          editingProfile.profile.races.forEach(race => {
            if (race.date) race.date = dayjs(race.date);
          });
        }
        console.log(editingProfile.profile)
      },
      (error) => {
        messageApi.error("获取资料失败：" + error);
      }
    );
  }
};

// 保存编辑的资料
const saveProfile = () => {
  const profileData = JSON.parse(JSON.stringify(editingProfile.profile));
  const str = JSON.stringify(profileData)
  if (isNewProfile.value) {
    // 创建新资料
    postJSON('/api/profile/create', {
      content: str
    },
      (message) => {
        messageApi.success("已成功添加新资料");
        router.push('/main');
      },
      (error) => {
        messageApi.error("保存失败：" + error);
      }
    );
  } else {
    // 更新现有资料
    put('/api/profile/update', {
      profile_id: profileIndex.value,
      content: str
    },
      (message) => {
        messageApi.success("资料修改成功");
        router.push('/main');
      },
      (error) => {
        messageApi.error("保存失败：" + error);
      }
    );
  }
};

// 取消编辑，返回主页
const cancelEdit = () => {
  router.push('/main');
};

// 处理文件选择，将图片转换为 Base64
const handleFileSelect = async (event) => {
  const file = event.target.files[0];
  if (!file) {
    console.error('未选择文件');
    return;
  }

  try {
    const base64String = await fileToBase64(file);
    editingProfile.profile.avatar = base64String;
  } catch (error) {
    console.error('图片转换错误:', error);
    messageApi.error('图片转换失败');
  }
};

// 文件转 Base64
const fileToBase64 = (file) => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve(reader.result);
    reader.onerror = (error) => reject(error);
    reader.readAsDataURL(file);
  });
};

// 编辑工作经历
const editWorkExperience = (index) => {
  const experience = editingProfile.profile.workExperience[index];
  editingProfile.currentEditingExperience = { ...experience };
  editingProfile.currentEditingIndex.workExperience = index;
  editingProfile.isEditing.workExperience = true;
};

// 更新工作经历
const updateWorkExperience = () => {
  const index = editingProfile.currentEditingIndex.workExperience;
  if (index >= 0) {
    editingProfile.profile.workExperience[index] = { ...editingProfile.currentEditingExperience };
    messageApi.success('工作经历更新成功！');
  } else {
    editingProfile.profile.workExperience.push({ ...editingProfile.currentEditingExperience });
    messageApi.success('保存成功！您可以继续编写下一个工作经历！');
  }
  resetWorkExperienceForm();
};

// 重置工作经历表单
const resetWorkExperienceForm = () => {
  editingProfile.currentEditingExperience = {
    workFor: '',
    jobName: '',
    position: '',
    description: '',
    startDate: dayjs(new Date()),
    endDate: dayjs(new Date())
  };
  editingProfile.currentEditingIndex.workExperience = -1;
  editingProfile.isEditing.workExperience = false;
};

// 编辑技能
const editSkill = (index) => {
  const skill = editingProfile.profile.skills[index];
  editingProfile.currentEditingSkill = { ...skill };
  editingProfile.currentEditingIndex.skill = index;
  editingProfile.isEditing.skill = true;
};

// 更新技能
const updateSkill = () => {
  const index = editingProfile.currentEditingIndex.skill;
  if (index >= 0) {
    editingProfile.profile.skills[index] = { ...editingProfile.currentEditingSkill };
    messageApi.success('技能更新成功！');
  } else {
    editingProfile.profile.skills.push({ ...editingProfile.currentEditingSkill });
    messageApi.success('保存成功！您可以继续编写下一个技能！');
  }
  resetSkillForm();
};

// 重置技能表单
const resetSkillForm = () => {
  editingProfile.currentEditingSkill = {
    title: '',
    description: ''
  };
  editingProfile.currentEditingIndex.skill = -1;
  editingProfile.isEditing.skill = false;
};

// 编辑教育经历
const editEducation = (index) => {
  const education = editingProfile.profile.education[index];
  editingProfile.currentEditingEducation = { ...education };
  editingProfile.currentEditingIndex.education = index;
  editingProfile.isEditing.education = true;
};

// 更新教育经历
const updateEducation = () => {
  const index = editingProfile.currentEditingIndex.education;
  if (index >= 0) {
    editingProfile.profile.education[index] = { ...editingProfile.currentEditingEducation };
    messageApi.success('教育经历更新成功！');
  } else {
    editingProfile.profile.education.push({ ...editingProfile.currentEditingEducation });
    messageApi.success('保存成功！您可以继续编写下一个教育经历！');
  }
  resetEducationForm();
};

// 重置教育经历表单
const resetEducationForm = () => {
  editingProfile.currentEditingEducation = {
    name: '',
    level: '',
    profess: '',
    startDate: dayjs(new Date()),
    endDate: dayjs(new Date())
  };
  editingProfile.currentEditingIndex.education = -1;
  editingProfile.isEditing.education = false;
};

// 编辑竞赛经历
const editRace = (index) => {
  const race = editingProfile.profile.races[index];
  editingProfile.currentEditingRace = { ...race };
  editingProfile.currentEditingIndex.race = index;
  editingProfile.isEditing.race = true;
};

// 更新竞赛经历
const updateRace = () => {
  const index = editingProfile.currentEditingIndex.race;
  if (index >= 0) {
    editingProfile.profile.races[index] = { ...editingProfile.currentEditingRace };
    messageApi.success('竞赛经历更新成功！');
  } else {
    editingProfile.profile.races.push({ ...editingProfile.currentEditingRace });
    messageApi.success('保存成功！您可以继续编写下一个竞赛经历！');
  }
  resetRaceForm();
};

// 重置竞赛经历表单
const resetRaceForm = () => {
  editingProfile.currentEditingRace = {
    name: '',
    level: '',
    date: dayjs(new Date())
  };
  editingProfile.currentEditingIndex.race = -1;
  editingProfile.isEditing.race = false;
};

// 编辑项目
const editProject = (index) => {
  const project = editingProfile.profile.projects[index];
  editingProfile.currentEditingProjects = { ...project };
  editingProfile.currentEditingIndex.project = index;
  editingProfile.isEditing.project = true;
};

// 更新项目
const updateProject = () => {
  const index = editingProfile.currentEditingIndex.project;
  if (index >= 0) {
    editingProfile.profile.projects[index] = { ...editingProfile.currentEditingProjects };
    messageApi.success('项目更新成功！');
  } else {
    editingProfile.profile.projects.push({ ...editingProfile.currentEditingProjects });
    messageApi.success('保存成功！您可以继续编写下一个项目！');
  }
  resetProjectForm();
};

// 重置项目表单
const resetProjectForm = () => {
  editingProfile.currentEditingProjects = {
    title: '',
    description: ''
  };
  editingProfile.currentEditingIndex.project = -1;
  editingProfile.isEditing.project = false;
};
</script>

<template>
  <contextHolder/>
  <div class="h-screen overflow-y-auto bg-white bg-opacity-70 animate__animated animate__fadeIn">  
    <!-- 主体内容 -->
    <div class="h-fit min-h-screen bg-blue-200 bg-opacity-10 backdrop-blur-md shadow-xl w-3/4 mx-auto rounded-xl mt-2 p-4 hover:shadow-lg transition-all transition-duration-300">
      <div class="flex mb-4">
        <h1 class="font-bold text-2xl flex-grow">{{ isNewProfile ? '创建新资料' : '编辑资料' }}</h1>
        <div class="grid grid-cols-[1fr,1fr] gap-2">
          <button @click="saveProfile" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700">保存</button>
          <button @click="cancelEdit" class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700">取消</button>
        </div>
      </div>
      
      <!-- 表单内容 -->
      <a-collapse v-model:activeKey="editingProfile.addFormActiveKey">
        <a-collapse-panel key="1" header="基本资料">
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>头像</a>
            <div class="flex items-center">
              <input class="basic-blue-input" type="file" @change="handleFileSelect" accept="image/*">
              <img v-if="editingProfile.profile.avatar" :src="editingProfile.profile.avatar" class="w-16 h-16 ml-4 rounded-full object-cover border-2 border-gray-300">
            </div>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>姓名</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.name"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>应聘岗位</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.jobName"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>电话</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.phoneNumber"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>电子邮箱</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.emailAddress"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>城市</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.position"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>生日</a>
            <a-date-picker class="basic-blue-input" v-model:value="editingProfile.profile.birthDate"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>个人网站</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.webSite"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>Github主页</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.github"/>
          </div>
          <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
            <a>领英主页</a>
            <input class="basic-blue-input" v-model="editingProfile.profile.linkIn"/>
          </div>
        </a-collapse-panel>
        
        <a-collapse-panel key="2" header="工作经历">
          <!-- 已添加的工作经历列表 -->
          <div v-if="editingProfile.profile.workExperience && editingProfile.profile.workExperience.length > 0" class="mb-4 border-b pb-4">
            <h3 class="font-bold mb-2">已添加的工作经历</h3>
            <div v-for="(exp, index) in editingProfile.profile.workExperience" :key="index" class="p-3 mb-2 bg-gray-100 rounded-md relative">
              <h4 class="font-bold">{{ exp.jobName }} @ {{ exp.workFor }}</h4>
              <p>{{ exp.position }} · {{ exp.startDate?.format('YYYY-MM-DD') || '未设置' }} 至 {{ exp.endDate?.format('YYYY-MM-DD') || '未设置' }}</p>
              <p class="text-sm text-gray-600 mt-1 line-clamp-2">{{ exp.description }}</p>
              <div class="absolute top-2 right-2 flex">
                <button @click="editWorkExperience(index)" class="text-blue-500 hover:text-blue-700 mr-2">
                  <span class="text-sm">✎</span>
                </button>
                <button @click="editingProfile.profile.workExperience.splice(index, 1)" class="text-red-500 hover:text-red-700">
                  <span class="text-xl">×</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- 添加新工作经历表单 -->
          <div>
            <h3 class="font-bold mb-2">{{ editingProfile.isEditing.workExperience ? '编辑' : '添加' }}工作经历</h3>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>工作单位</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingExperience.workFor"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>岗位</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingExperience.jobName"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>城市</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingExperience.position"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>描述</a>
              <textarea class="basic-blue-input" v-model="editingProfile.currentEditingExperience.description" rows="5"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>起始时间</a>
              <a-date-picker class="basic-blue-input" v-model:value="editingProfile.currentEditingExperience.startDate"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>离职时间</a>
              <a-date-picker class="basic-blue-input" v-model:value="editingProfile.currentEditingExperience.endDate"/>
            </div>
            <div class="mt-3 flex">
              <button @click="updateWorkExperience" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700 mr-2">
                {{ editingProfile.isEditing.workExperience ? '更新' : '保存并继续' }}
              </button>
              <button v-if="editingProfile.isEditing.workExperience" 
                      @click="resetWorkExperienceForm" 
                      class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700">
                取消编辑
              </button>
            </div>
          </div>
        </a-collapse-panel>
        
        <a-collapse-panel key="3" header="竞赛经历">
          <!-- 已添加的竞赛经历列表 -->
          <div v-if="editingProfile.profile.races && editingProfile.profile.races.length > 0" class="mb-4 border-b pb-4">
            <h3 class="font-bold mb-2">已添加的竞赛经历</h3>
            <div v-for="(race, index) in editingProfile.profile.races" :key="index" class="p-3 mb-2 bg-gray-100 rounded-md relative">
              <h4 class="font-bold">{{ race.name }}</h4>
              <p>{{ race.level }} · {{ race.date?.format('YYYY-MM-DD') || '未设置' }}</p>
              <div class="absolute top-2 right-2 flex">
                <button @click="editRace(index)" class="text-blue-500 hover:text-blue-700 mr-2">
                  <span class="text-sm">✎</span>
                </button>
                <button @click="editingProfile.profile.races.splice(index, 1)" class="text-red-500 hover:text-red-700">
                  <span class="text-xl">×</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- 添加新竞赛经历表单 -->
          <div>
            <h3 class="font-bold mb-2">{{ editingProfile.isEditing.race ? '编辑' : '添加' }}竞赛经历</h3>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>竞赛名称</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingRace.name"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>奖项</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingRace.level"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>获得时间</a>
              <a-date-picker class="basic-blue-input" v-model:value="editingProfile.currentEditingRace.date"/>
            </div>
            <div class="mt-3 flex">
              <button @click="updateRace" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700 mr-2">
                {{ editingProfile.isEditing.race ? '更新' : '保存并继续' }}
              </button>
              <button v-if="editingProfile.isEditing.race" 
                      @click="resetRaceForm" 
                      class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700">
                取消编辑
              </button>
            </div>
          </div>
        </a-collapse-panel>
        
        <a-collapse-panel key="4" header="教育经历">
          <!-- 已添加的教育经历列表 -->
          <div v-if="editingProfile.profile.education && editingProfile.profile.education.length > 0" class="mb-4 border-b pb-4">
            <h3 class="font-bold mb-2">已添加的教育经历</h3>
            <div v-for="(edu, index) in editingProfile.profile.education" :key="index" class="p-3 mb-2 bg-gray-100 rounded-md relative">
              <h4 class="font-bold">{{ edu.name }}</h4>
              <p>{{ edu.level }} - {{ edu.profess }}</p>
              <p>{{ edu.startDate?.format('YYYY-MM-DD') || '未设置' }} 至 {{ edu.endDate?.format('YYYY-MM-DD') || '未设置' }}</p>
              <div class="absolute top-2 right-2 flex">
                <button @click="editEducation(index)" class="text-blue-500 hover:text-blue-700 mr-2">
                  <span class="text-sm">✎</span>
                </button>
                <button @click="editingProfile.profile.education.splice(index, 1)" class="text-red-500 hover:text-red-700">
                  <span class="text-xl">×</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- 添加新教育经历表单 -->
          <div>
            <h3 class="font-bold mb-2">{{ editingProfile.isEditing.education ? '编辑' : '添加' }}教育经历</h3>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>学校名称</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingEducation.name"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>学历</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingEducation.level"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>专业方向</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingEducation.profess"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>入学时间</a>
              <a-date-picker class="basic-blue-input" v-model:value="editingProfile.currentEditingEducation.startDate"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>毕业时间</a>
              <a-date-picker class="basic-blue-input" v-model:value="editingProfile.currentEditingEducation.endDate"/>
            </div>
            <div class="mt-3 flex">
              <button @click="updateEducation" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700 mr-2">
                {{ editingProfile.isEditing.education ? '更新' : '保存并继续' }}
              </button>
              <button v-if="editingProfile.isEditing.education" 
                      @click="resetEducationForm" 
                      class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700">
                取消编辑
              </button>
            </div>
          </div>
        </a-collapse-panel>
        
        <a-collapse-panel key="5" header="个人项目">
          <!-- 已添加的项目列表 -->
          <div v-if="editingProfile.profile.projects && editingProfile.profile.projects.length > 0" class="mb-4 border-b pb-4">
            <h3 class="font-bold mb-2">已添加的项目</h3>
            <div v-for="(project, index) in editingProfile.profile.projects" :key="index" class="p-3 mb-2 bg-gray-100 rounded-md relative">
              <h4 class="font-bold">{{ project.title }}</h4>
              <p class="text-sm text-gray-600 mt-1 line-clamp-2">{{ project.description }}</p>
              <div class="absolute top-2 right-2 flex">
                <button @click="editProject(index)" class="text-blue-500 hover:text-blue-700 mr-2">
                  <span class="text-sm">✎</span>
                </button>
                <button @click="editingProfile.profile.projects.splice(index, 1)" class="text-red-500 hover:text-red-700">
                  <span class="text-xl">×</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- 添加新项目表单 -->
          <div>
            <h3 class="font-bold mb-2">{{ editingProfile.isEditing.project ? '编辑' : '添加' }}项目</h3>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>项目名称</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingProjects.title"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>简介</a>
              <textarea class="basic-blue-input" v-model="editingProfile.currentEditingProjects.description" rows="5"/>
            </div>
            <div class="mt-3 flex">
              <button @click="updateProject" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700 mr-2">
                {{ editingProfile.isEditing.project ? '更新' : '保存并继续' }}
              </button>
              <button v-if="editingProfile.isEditing.project" 
                      @click="resetProjectForm" 
                      class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700">
                取消编辑
              </button>
            </div>
          </div>
        </a-collapse-panel>
        
        <a-collapse-panel key="6" header="技能">
          <!-- 已添加的技能列表 -->
          <div v-if="editingProfile.profile.skills && editingProfile.profile.skills.length > 0" class="mb-4 border-b pb-4">
            <h3 class="font-bold mb-2">已添加的技能</h3>
            <div v-for="(skill, index) in editingProfile.profile.skills" :key="index" class="p-3 mb-2 bg-gray-100 rounded-md relative">
              <h4 class="font-bold">{{ skill.title }}</h4>
              <p class="text-sm text-gray-600 mt-1 line-clamp-2">{{ skill.description }}</p>
              <div class="absolute top-2 right-2 flex">
                <button @click="editSkill(index)" class="text-blue-500 hover:text-blue-700 mr-2">
                  <span class="text-sm">✎</span>
                </button>
                <button @click="editingProfile.profile.skills.splice(index, 1)" class="text-red-500 hover:text-red-700">
                  <span class="text-xl">×</span>
                </button>
              </div>
            </div>
          </div>
          
          <!-- 添加新技能表单 -->
          <div>
            <h3 class="font-bold mb-2">{{ editingProfile.isEditing.skill ? '编辑' : '添加' }}技能</h3>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>技能名称</a>
              <input class="basic-blue-input" v-model="editingProfile.currentEditingSkill.title"/>
            </div>
            <div class="grid grid-cols-[1fr,3fr] place-items-center mt-1">
              <a>技能描述</a>
              <textarea class="basic-blue-input" v-model="editingProfile.currentEditingSkill.description" rows="5"/>
            </div>
            <div class="mt-3 flex">
              <button @click="updateSkill" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700 mr-2">
                {{ editingProfile.isEditing.skill ? '更新' : '保存并继续' }}
              </button>
              <button v-if="editingProfile.isEditing.skill" 
                      @click="resetSkillForm" 
                      class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700">
                取消编辑
              </button>
            </div>
          </div>
        </a-collapse-panel>
      </a-collapse>
      
      <!-- 底部保存按钮 -->
      <div class="flex justify-center mt-6">
        <button @click="saveProfile" class="basic-button bg-green-500 hover:bg-green-600 active:bg-green-700 mr-4 px-8">保存资料</button>
        <button @click="cancelEdit" class="basic-button bg-gray-500 hover:bg-gray-600 active:bg-gray-700 px-8">取消</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.basic-blue-input {
  @apply border p-2 rounded-md focus:outline-none focus:border-blue-600 w-full;
}
.basic-button {
  @apply rounded-md py-2 px-4 text-white font-medium transition-colors duration-200;
}
.line-clamp-2 {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}
</style> 