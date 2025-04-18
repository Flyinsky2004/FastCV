import { useUserStore } from '@/stores/user'
import { get } from '@/util/request'
import { message } from 'ant-design-vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      name: 'framework',
      path: '/',
      component: () => import('@/components/Layout.vue'),
      children:[
        {
          name: 'index',
          path: '/',
          component: () => import('@/views/inedx.vue')
        },{
          name: 'login',
          path: '/auth/login',
          component: () => import('@/views/auth/Login.vue')
        },{
          name: 'register',
          path: '/auth/register',
          component: () => import('@/views/auth/Register.vue')
        },{
          name: 'virtual',
          path:'/virtual',
          component: () => import('@/views/main/Virtual.vue')
        },{
          name: 'main',
          path: '/main',
          component: () => import('@/views/main/main.vue')
        },{
          name:'template',
          path:'/template/:id',
          component: () => import('@/views/main/render.vue')
        },{
          name: 'edit',
          path: '/edit/:index',
          component: () => import('@/views/main/edit.vue')
        }
      ]
    }
  ],
})

router.beforeEach((to, from, next) => {
  const userStore = useUserStore()
  
  // 处理无需登录的路由
  if (to.path === '/' || to.path.startsWith('/auth')) {
      next()
      return
  }
  
  // 处理需要管理员权限的路由
  if (to.matched.some(record => record.meta.requiresAdmin)) {
      if (userStore.isLogin && userStore.user?.permission >= 1) {
          next()
      } else {
          message.info('此页面需要管理员权限，请使用管理员账号登录！');
          setTimeout(() => {
              next('/auth/login')
          }, 2000)
      }
      return
  }
  
  // 处理普通需登录路由
  if (userStore.isLogin) {
      next()
  } else {
      get('/api/user/myInfo', {},
          (message, data) => {
              userStore.login(data)
              next()
          }, (messager, data) => {
              message.info('您还尚未登录，请先登录！');
              setTimeout(() => {
                  next('/auth/login')
              }, 2000)
          }, (messager, data) => {
              message.info('您还尚未登录，请先登录！');
              setTimeout(() => {
                  next('/auth/login')
              }, 2000)
          }
      )
  }
})
export default router
