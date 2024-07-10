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
      children: [
        {
          path: '',
          name: 'loginView',
          component: () => import('@/views/backstage/LoginView.vue')
        },
        {
          path: 'pendingTasks',
          name: 'pedingTasksView',
          component: () => import('@/views/backstage/PendingTasksView.vue'),
          meta: {
            requireLogin: true
          }
        },
      ]
    },
    {
      path: '/:pathMatch(.*)*',
      name: '404',
      component: () => import('@/views/UserView.vue')
    }
  ],
})

router.beforeEach((to, from, next) => {
  const store = useBackstageStore()
  // 鉴权需要登录的页面
  if (to.meta.requireLogin && !store.jwt) {
    alert('登录已过期,请先登录')
    next({ path: '/bs/' })
  }
  // 自动补全结尾的/,以便使用相对路径
  if (to.path == '/bs') {
    next({ path: '/bs/', replace: true })
  }
  // 防止登录后再次登录
  if (to.path == '/bs/' && store.jwt) {
    next({ path: '/bs/pendingTasks', replace: true })
  }
  next()
})

export default router