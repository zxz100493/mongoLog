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
        component: () => import(/* webpackChunkName: "about" */ '../views/AboutViewTwo.vue')
      },
      {
        path: '/about2',
        name: 'about2',
        component: () => import(/* webpackChunkName: "about2" */ '../views/AboutView.vue')
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
