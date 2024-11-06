import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      name: 'index',
      path: '/',
      component: () => import('@/views/inedx.vue')
    },{
      name: 'main',
      path: '/main',
      component: () => import('@/views/main/main.vue')
    },{
      name:'template',
      path:'/template/:id',
      component: () => import('@/views/main/render.vue')
    }
  ],
})

export default router
