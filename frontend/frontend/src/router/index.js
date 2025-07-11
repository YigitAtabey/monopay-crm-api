// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore }              from '@/store/auth'

import HomeView        from '@/views/HomeView.vue'
import Login           from '@/views/user/Login.vue'
import Register        from '@/views/user/Register.vue'
import Logout          from '@/views/user/Logout.vue'
import UserPanel       from '@/views/user/UserPanel.vue'
import TaskList        from '@/views/user/TaskList.vue'
import TaskManagement  from '@/views/user/TaskManagement.vue'
import AdminPanel      from '@/views/admin/AdminPanel.vue'

const routes = [
  { path: '/',           name: 'home',        component: HomeView },
  { path: '/login',      name: 'login',       component: Login },
  { path: '/register',   name: 'register',    component: Register },
  { path: '/panel',      name: 'user-panel',  component: UserPanel,      meta: { requiresAuth: true } },
  { path: '/tasks',      name: 'task-list',   component: TaskList,       meta: { requiresAuth: true } },
  { path: '/tasks/manage', name: 'task-manage', component: TaskManagement, meta: { requiresAuth: true } },
  { path: '/admin',      name: 'admin-panel', component: AdminPanel,     meta: { requiresAuth: true, requiresAdmin: true } },
  { path: '/logout',     name: 'logout',      component: Logout,         meta: { requiresAuth: true } },
  { path: '/:pathMatch(.*)*', redirect: '/' }
]
const router = createRouter({
  history: createWebHistory(),
  routes
})
router.beforeEach((to, from, next) => {
  const auth     = useAuthStore()
  const loggedIn = auth.isLoggedIn
  const isAdmin  = auth.user?.role === 'admin'

  if (to.meta.requiresAuth && !loggedIn) {
    return next({ path: '/login' })
  }
  if (to.meta.requiresAdmin && !isAdmin) {
    return next({ path: '/' })
  }
  if ((to.path === '/login' || to.path === '/register') && loggedIn) {
    return next({ path: '/' })
  }
  next()
})

export default router
