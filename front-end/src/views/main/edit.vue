<script setup>
import {onBeforeMount, reactive, ref} from "vue";
import {message, Tabs} from 'ant-design-vue';
import router from "@/router/index.js";
import dayjs from "dayjs";
import {defaultProfile} from '@/hooks/data';
import {get, put, postJSON, del} from '@/util/request';
import TextEditModal from '@/components/TextEditModal.vue';
import FormField from '@/components/FormField.vue';
import CardSection from '@/components/CardSection.vue';
import ItemList from '@/components/ItemList.vue';

const { TabPane } = Tabs;
const [messageApi, contextHolder] = message.useMessage();

// 获取路由参数
const profileIndex = ref(parseInt(router.currentRoute.value.params.index));
const isNewProfile = ref(profileIndex.value === -1);

// 当前激活的标签页
const activeTab = ref('basic');

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
  <div class="min-h-screen bg-gray-50 py-6 animate__animated animate__fadeIn">
    <div class="max-w-5xl mx-auto px-4">
      <div class="flex items-center justify-between mb-8">
        <h1 class="text-3xl font-bold text-gray-900">{{ isNewProfile ? '创建新资料' : '编辑资料' }}</h1>
        <div class="flex space-x-4">
          <button @click="saveProfile" class="btn-primary">保存</button>
          <button @click="cancelEdit" class="btn-secondary">取消</button>
        </div>
      </div>

      <div class="bg-white rounded-lg shadow-sm">
        <a-tabs v-model:activeKey="activeTab" 
                class="custom-tabs"
                :animated="{ inkBar: true, tabPane: false }">
          <a-tab-pane key="basic" tab="基本资料">
            <div class="p-6">
              <FormField label="头像">
                <div class="flex items-center">
                  <input type="file" @change="handleFileSelect" accept="image/*" class="file-input">
                  <img v-if="editingProfile.profile.avatar" :src="editingProfile.profile.avatar" class="w-16 h-16 ml-4 rounded-full object-cover border-2 border-gray-200">
                </div>
              </FormField>
              
              <FormField label="姓名">
                <input v-model="editingProfile.profile.name" class="form-input"/>
              </FormField>
              
              <FormField label="应聘岗位">
                <input v-model="editingProfile.profile.jobName" class="form-input"/>
              </FormField>
              
              <FormField label="电话">
                <input v-model="editingProfile.profile.phoneNumber" class="form-input"/>
              </FormField>
              
              <FormField label="电子邮箱">
                <input v-model="editingProfile.profile.emailAddress" class="form-input"/>
              </FormField>
              
              <FormField label="城市">
                <input v-model="editingProfile.profile.position" class="form-input"/>
              </FormField>
              
              <FormField label="生日">
                <a-date-picker class="form-input" v-model:value="editingProfile.profile.birthDate"/>
              </FormField>
              
              <FormField label="个人网站">
                <input v-model="editingProfile.profile.webSite" class="form-input"/>
              </FormField>
              
              <FormField label="Github主页">
                <input v-model="editingProfile.profile.github" class="form-input"/>
              </FormField>
              
              <FormField label="领英主页">
                <input v-model="editingProfile.profile.linkIn" class="form-input"/>
              </FormField>
            </div>
          </a-tab-pane>

          <a-tab-pane key="work" tab="工作经历">
            <div class="p-6 grid grid-cols-12 gap-6">
              <!-- 左侧列表 -->
              <div class="col-span-4 border-r pr-4">
                <div class="flex justify-between items-center mb-4">
                  <h3 class="text-lg font-bold">工作经历列表</h3>
                  <button 
                    @click="resetWorkExperienceForm" 
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + 新增
                  </button>
                </div>
                <div class="space-y-3">
                  <div 
                    v-for="(item, index) in editingProfile.profile.workExperience" 
                    :key="index"
                    @click="editWorkExperience(index)"
                    class="p-3 rounded-lg cursor-pointer transition-all duration-200"
                    :class="[
                      editingProfile.currentEditingIndex.workExperience === index 
                        ? 'bg-blue-50 border-blue-200' 
                        : 'hover:bg-gray-50 border-gray-100'
                    ]"
                  >
                    <h4 class="font-medium">{{ item.jobName }}</h4>
                    <p class="text-sm text-gray-600">{{ item.workFor }}</p>
                    <p class="text-xs text-gray-500">
                      {{ item.startDate?.format('YYYY-MM') || '未设置' }} 至 
                      {{ item.endDate?.format('YYYY-MM') || '未设置' }}
                    </p>
                  </div>
                </div>
              </div>

              <!-- 右侧编辑区 -->
              <div class="col-span-8">
                <div class="bg-gray-50 p-4 rounded-lg">
                  <h4 class="font-bold mb-4">{{ editingProfile.currentEditingIndex.workExperience >= 0 ? '编辑' : '添加' }}工作经历</h4>
                  <FormField label="工作单位">
                    <input v-model="editingProfile.currentEditingExperience.workFor" class="form-input"/>
                  </FormField>
                  <FormField label="岗位">
                    <input v-model="editingProfile.currentEditingExperience.jobName" class="form-input"/>
                  </FormField>
                  <FormField label="城市">
                    <input v-model="editingProfile.currentEditingExperience.position" class="form-input"/>
                  </FormField>
                  <FormField label="描述">
                    <TextEditModal
                      v-model="editingProfile.currentEditingExperience.description"
                      title="编辑工作经历描述"
                      :rows="5"
                    />
                  </FormField>
                  <FormField label="起始时间">
                    <a-date-picker class="form-input" v-model:value="editingProfile.currentEditingExperience.startDate"/>
                  </FormField>
                  <FormField label="离职时间">
                    <a-date-picker class="form-input" v-model:value="editingProfile.currentEditingExperience.endDate"/>
                  </FormField>
                  <div class="flex justify-end space-x-3 mt-4">
                    <button @click="resetWorkExperienceForm" class="btn-secondary">取消</button>
                    <button @click="updateWorkExperience" class="btn-primary">保存</button>
                  </div>
                </div>
              </div>
            </div>
          </a-tab-pane>

          <a-tab-pane key="education" tab="教育经历">
            <div class="p-6 grid grid-cols-12 gap-6">
              <!-- 左侧列表 -->
              <div class="col-span-4 border-r pr-4">
                <div class="flex justify-between items-center mb-4">
                  <h3 class="text-lg font-bold">教育经历列表</h3>
                  <button 
                    @click="resetEducationForm" 
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + 新增
                  </button>
                </div>
                <div class="space-y-3">
                  <div 
                    v-for="(item, index) in editingProfile.profile.education" 
                    :key="index"
                    @click="editEducation(index)"
                    class="p-3 rounded-lg cursor-pointer transition-all duration-200"
                    :class="[
                      editingProfile.currentEditingIndex.education === index 
                        ? 'bg-blue-50 border-blue-200' 
                        : 'hover:bg-gray-50 border-gray-100'
                    ]"
                  >
                    <h4 class="font-medium">{{ item.name }}</h4>
                    <p class="text-sm text-gray-600">{{ item.level }} - {{ item.profess }}</p>
                    <p class="text-xs text-gray-500">
                      {{ item.startDate?.format('YYYY') || '未设置' }} - 
                      {{ item.endDate?.format('YYYY') || '未设置' }}
                    </p>
                  </div>
                </div>
              </div>

              <!-- 右侧编辑区 -->
              <div class="col-span-8">
                <div class="bg-gray-50 p-4 rounded-lg">
                  <h4 class="font-bold mb-4">{{ editingProfile.currentEditingIndex.education >= 0 ? '编辑' : '添加' }}教育经历</h4>
                  <FormField label="学校名称">
                    <input v-model="editingProfile.currentEditingEducation.name" class="form-input"/>
                  </FormField>
                  <FormField label="学历">
                    <input v-model="editingProfile.currentEditingEducation.level" class="form-input"/>
                  </FormField>
                  <FormField label="专业方向">
                    <input v-model="editingProfile.currentEditingEducation.profess" class="form-input"/>
                  </FormField>
                  <FormField label="入学时间">
                    <a-date-picker class="form-input" v-model:value="editingProfile.currentEditingEducation.startDate"/>
                  </FormField>
                  <FormField label="毕业时间">
                    <a-date-picker class="form-input" v-model:value="editingProfile.currentEditingEducation.endDate"/>
                  </FormField>
                  <div class="flex justify-end space-x-3 mt-4">
                    <button @click="resetEducationForm" class="btn-secondary">取消</button>
                    <button @click="updateEducation" class="btn-primary">保存</button>
                  </div>
                </div>
              </div>
            </div>
          </a-tab-pane>

          <a-tab-pane key="race" tab="竞赛经历">
            <div class="p-6 grid grid-cols-12 gap-6">
              <!-- 左侧列表 -->
              <div class="col-span-4 border-r pr-4">
                <div class="flex justify-between items-center mb-4">
                  <h3 class="text-lg font-bold">竞赛经历列表</h3>
                  <button 
                    @click="resetRaceForm" 
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + 新增
                  </button>
                </div>
                <div class="space-y-3">
                  <div 
                    v-for="(item, index) in editingProfile.profile.races" 
                    :key="index"
                    @click="editRace(index)"
                    class="p-3 rounded-lg cursor-pointer transition-all duration-200"
                    :class="[
                      editingProfile.currentEditingIndex.race === index 
                        ? 'bg-blue-50 border-blue-200' 
                        : 'hover:bg-gray-50 border-gray-100'
                    ]"
                  >
                    <h4 class="font-medium">{{ item.name }}</h4>
                    <p class="text-sm text-gray-600">{{ item.level }}</p>
                    <p class="text-xs text-gray-500">{{ item.date?.format('YYYY-MM-DD') || '未设置' }}</p>
                  </div>
                </div>
              </div>

              <!-- 右侧编辑区 -->
              <div class="col-span-8">
                <div class="bg-gray-50 p-4 rounded-lg">
                  <h4 class="font-bold mb-4">{{ editingProfile.currentEditingIndex.race >= 0 ? '编辑' : '添加' }}竞赛经历</h4>
                  <FormField label="竞赛名称">
                    <input v-model="editingProfile.currentEditingRace.name" class="form-input"/>
                  </FormField>
                  <FormField label="奖项">
                    <input v-model="editingProfile.currentEditingRace.level" class="form-input"/>
                  </FormField>
                  <FormField label="获得时间">
                    <a-date-picker class="form-input" v-model:value="editingProfile.currentEditingRace.date"/>
                  </FormField>
                  <div class="flex justify-end space-x-3 mt-4">
                    <button @click="resetRaceForm" class="btn-secondary">取消</button>
                    <button @click="updateRace" class="btn-primary">保存</button>
                  </div>
                </div>
              </div>
            </div>
          </a-tab-pane>

          <a-tab-pane key="project" tab="个人项目">
            <div class="p-6 grid grid-cols-12 gap-6">
              <!-- 左侧列表 -->
              <div class="col-span-4 border-r pr-4">
                <div class="flex justify-between items-center mb-4">
                  <h3 class="text-lg font-bold">项目列表</h3>
                  <button 
                    @click="resetProjectForm" 
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + 新增
                  </button>
                </div>
                <div class="space-y-3">
                  <div 
                    v-for="(item, index) in editingProfile.profile.projects" 
                    :key="index"
                    @click="editProject(index)"
                    class="p-3 rounded-lg cursor-pointer transition-all duration-200"
                    :class="[
                      editingProfile.currentEditingIndex.project === index 
                        ? 'bg-blue-50 border-blue-200' 
                        : 'hover:bg-gray-50 border-gray-100'
                    ]"
                  >
                    <h4 class="font-medium">{{ item.title }}</h4>
                    <p class="text-sm text-gray-600 line-clamp-2">{{ item.description }}</p>
                  </div>
                </div>
              </div>

              <!-- 右侧编辑区 -->
              <div class="col-span-8">
                <div class="bg-gray-50 p-4 rounded-lg">
                  <h4 class="font-bold mb-4">{{ editingProfile.currentEditingIndex.project >= 0 ? '编辑' : '添加' }}项目</h4>
                  <FormField label="项目名称">
                    <input v-model="editingProfile.currentEditingProjects.title" class="form-input"/>
                  </FormField>
                  <FormField label="简介">
                    <TextEditModal
                      v-model="editingProfile.currentEditingProjects.description"
                      title="编辑项目简介"
                      :rows="5"
                    />
                  </FormField>
                  <div class="flex justify-end space-x-3 mt-4">
                    <button @click="resetProjectForm" class="btn-secondary">取消</button>
                    <button @click="updateProject" class="btn-primary">保存</button>
                  </div>
                </div>
              </div>
            </div>
          </a-tab-pane>

          <a-tab-pane key="skill" tab="技能">
            <div class="p-6 grid grid-cols-12 gap-6">
              <!-- 左侧列表 -->
              <div class="col-span-4 border-r pr-4">
                <div class="flex justify-between items-center mb-4">
                  <h3 class="text-lg font-bold">技能列表</h3>
                  <button 
                    @click="resetSkillForm" 
                    class="text-blue-600 hover:text-blue-700"
                  >
                    + 新增
                  </button>
                </div>
                <div class="space-y-3">
                  <div 
                    v-for="(item, index) in editingProfile.profile.skills" 
                    :key="index"
                    @click="editSkill(index)"
                    class="p-3 rounded-lg cursor-pointer transition-all duration-200"
                    :class="[
                      editingProfile.currentEditingIndex.skill === index 
                        ? 'bg-blue-50 border-blue-200' 
                        : 'hover:bg-gray-50 border-gray-100'
                    ]"
                  >
                    <h4 class="font-medium">{{ item.title }}</h4>
                    <p class="text-sm text-gray-600 line-clamp-2">{{ item.description }}</p>
                  </div>
                </div>
              </div>

              <!-- 右侧编辑区 -->
              <div class="col-span-8">
                <div class="bg-gray-50 p-4 rounded-lg">
                  <h4 class="font-bold mb-4">{{ editingProfile.currentEditingIndex.skill >= 0 ? '编辑' : '添加' }}技能</h4>
                  <FormField label="技能名称">
                    <input v-model="editingProfile.currentEditingSkill.title" class="form-input"/>
                  </FormField>
                  <FormField label="技能描述">
                    <TextEditModal
                      v-model="editingProfile.currentEditingSkill.description"
                      title="编辑技能描述"
                      :rows="5"
                    />
                  </FormField>
                  <div class="flex justify-end space-x-3 mt-4">
                    <button @click="resetSkillForm" class="btn-secondary">取消</button>
                    <button @click="updateSkill" class="btn-primary">保存</button>
                  </div>
                </div>
              </div>
            </div>
          </a-tab-pane>
        </a-tabs>
      </div>

      <!-- 底部保存按钮 -->
      <div class="flex justify-center mt-8 space-x-4">
        <button @click="saveProfile" class="btn-primary px-8">保存资料</button>
        <button @click="cancelEdit" class="btn-secondary px-8">取消</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.btn-primary {
  @apply bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-medium transition-colors duration-200;
}

.btn-secondary {
  @apply bg-gray-200 hover:bg-gray-300 text-gray-700 px-6 py-2 rounded-lg font-medium transition-colors duration-200;
}

.form-input {
  @apply w-full px-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-blue-500 transition-all duration-200;
}

.file-input {
  @apply block w-full text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:text-sm file:font-medium
  file:bg-blue-50 file:text-blue-700 hover:file:bg-blue-100;
}

:deep(.custom-tabs) {
  @apply bg-white;
}

:deep(.custom-tabs .ant-tabs-nav) {
  @apply mb-0 px-6 pt-4;
}

:deep(.custom-tabs .ant-tabs-tab) {
  @apply text-gray-600 text-base transition-colors duration-200;
}

:deep(.custom-tabs .ant-tabs-tab-active) {
  @apply text-blue-600;
}

:deep(.custom-tabs .ant-tabs-ink-bar) {
  @apply bg-blue-600;
}
</style> 