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
          path: '/login',
          component: () => import('@/views/auth/Login.vue')
        },{
          name: 'register',
          path: '/register',
          component: () => import('@/views/auth/Register.vue')
        }
      ]
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
  ],
})

export default router
