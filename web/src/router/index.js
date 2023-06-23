import Login from "../views/Login/index.vue";
import Activity from "../views/Cart/components/CartProduct.vue";
import Home from "../views/home/index";
import Details from "../views/Details/index.vue";
import Layout from "../views/Layout/index";
import Member from "@/views/member/index.vue";
import cart from "@/views/Cart/index.vue";
import UserInfo from "@/views/member/components/userInfo.vue";
import userOrder from "@/views/member/components/userOrder.vue";
import Search from "@/views/Search/index.vue";
import empty from "@/views/Layout/empty.vue";
import CheckOut from "@/views/CheckOut/index.vue";
import { createRouter, createWebHashHistory } from "vue-router";

const routes = [
  //测试用路由，当连接不到后端时调用这里的路由并更改Activity来进行调试
  {
    path: "/activity",
    name: "Activity",
    component: Activity,
  },
  //空白页，处理搜索时刷新问题
  {
    path: "/empty/:prom",
    name: "empty",
    component: empty,
  },
  // Home组件，显示中心内容
  {
    path: "/",
    name: "Layout",
    component: Layout,
    children: [
      {
        path: "",
        component: Home,
      },
      {
        path: "member",
        name: "member",
        component: Member,
        children: [
          {
            path: "",
            component: UserInfo,
          },
          {
            path: "order",
            name: "order",
            component: userOrder,
          },
        ],
      },
      {
        path: "Details/:goodId",
        name: "Details",
        //添加路由守卫请求
        meta: {
          requireAuth: true,
        },
        component: Details,
      },
      //购物车
      {
        path: "cart",
        name: "cart",
        meta: {
          requireAuth: true,
        },
        component: cart,
      },
      {
        path: "search/:prom",
        name: "search",
        component: Search,
      },
      {
        path: "checkout/:orderID",
        name: "checkout",
        component: CheckOut,
      },
    ],
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

// 路由拦截，要求用户登录之后才可以访问,需要拦截的路由记得requireAuth
router.beforeEach((to, from, next) => {
  const token = localStorage.token;
  if (to.meta.requireAuth) {
    // 判断该路由是否需要登录权限
    if (token) {
      // 通过vuex state获取当前的token是否存在
      next();
    } else {
      window.console.log("该页面需要登陆");
      next({
        path: "/login",
        // query: {redirect: to.fullPath} // 将跳转的路由path作为参数，登录成功后跳转到该路由
      });
    }
  } else {
    next();
  }
});

export default router;
