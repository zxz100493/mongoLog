import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/log',
    name: 'log',
    component: () => import(/* webpackChunkName: "log" */ '../views/LogIndexView.vue'),
    children: [
      {
        path: '/about',
        name: 'about',
        component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
      },
      {
        path: '/systemInfo',
        name: 'systemInfo',
        component: () => import(/* webpackChunkName: "systemInfo" */ '../views/systemInfo/systemInfo.vue')
      },
      {
        path: '/dbManager',
        name: 'dbmanager',
        component: () => import(/* webpackChunkName: "dbmanager" */ '../views/dbManager/dbManager.vue')
      },
      {
        path: '/queryLog',
        name: 'querylog',
        component: () => import(/* webpackChunkName: "querylog" */ '../views/queryLog/queryLog.vue')
      }
    ]
  }
  //   {
  //     path: '/about',
  //     name: 'about',
  //     // route level code-splitting
  //     // this generates a separate chunk (about.[hash].js) for this route
  //     // which is lazy-loaded when the route is visited.
  //     component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  //   }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
