import { createRouter, createWebHashHistory } from 'vue-router'
import { useBackstageStore } from '@/stores/backstage'
const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/',
      name: 'userView',
      component: () => import('@/views/UserView.vue')
    },
    {
      path: '/bs',
      name: 'backstageView',
      component: () => import('@/views/BackstageView.vue'),
      strict: true,
      children: [
        {
          path: '',
          name: 'loginView',
          component: () => import('@/views/backstage/LoginView.vue')
        },
        {
          path: 'pendingTasks/',
          name: 'pedingTasksView',
          component: () => import('@/views/backstage/PendingTasksView.vue'),
          meta: {
            requireLogin: true
          }
        },
      ]
    }
  ],
})

router.beforeEach((to) => {
  const store = useBackstageStore()
  // 鉴权需要登录的页面
  if (to.meta.requireLogin && !store.jwt) {
    alert('登录已过期,请先登录')
    return '/bs/'
  }
  // 自动补全结尾的/,以便使用相对路径
  if (to.path == '/bs') {
    return '/bs/'
  }
})

export default router