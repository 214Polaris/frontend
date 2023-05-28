import Login from '../views/Login_user'
import Activity from '../components/index_test'
import Index from '../views/home/MainPage'
import store from '../store/index'
import Details from '../views/home/childHome/Details/details_page'
import { createRouter, createWebHashHistory } from 'vue-router'

const routes = [
  {
    path: '/MainPage/Details/:goodId',
    name: 'Details',
    meta: {
      requireAuth: true
    },
    component: Details
  },
  {
    path: '/',
    name: 'MainPage',
    component: Index
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/activity',
    name: 'activity',
    component: Activity,
    meta: {
      requireAuth: true
    }
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})
// 路由拦截，要求用户登录之后才可以访问
router.beforeEach((to, from, next) => {
  const token = localStorage.token
  if (to.meta.requireAuth) { // 判断该路由是否需要登录权限
    if (token) { // 通过vuex state获取当前的token是否存在
      next()
    } else {
      window.console.log('该页面需要登陆')
      next({
        path: '/login'
        // query: {redirect: to.fullPath} // 将跳转的路由path作为参数，登录成功后跳转到该路由
      })
    }
  } else {
    next()
  }
})

export default router
