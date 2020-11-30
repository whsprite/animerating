import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [{
    path: '/',
    name: 'Login',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import( /* webpackChunkName: "login" */ '../views/Login.vue')
  },
  {
    path: '/home',
    name: 'Home',
    redirect: '/all',
    component: () => import( /* webpackChunkName: "Home" */ '../views/Home.vue'),
    children: [{
        path: '/all',
        name: 'All',
        component: () => import( /* webpackChunkName: "All" */ '../views/All.vue')
      },
      {
        path: '/ranking',
        name: 'Ranking',
        component: () => import( /* webpackChunkName: "All" */ '../views/Ranking.vue')
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

export default router