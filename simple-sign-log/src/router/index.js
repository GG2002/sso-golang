// Composables
import { createRouter, createWebHistory } from 'vue-router'
import Login from '../views/Login.vue'
import Signup from '../views/Signup.vue'
import Ok from '../views/Ok.vue'

const routes = [
  {
    path: '/',
    component: Login
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/signup',
    name: 'signup',
    component: Signup
  },
  {
    path: '/ok',
    component: Ok
  }
]

const router = createRouter({
  history: createWebHistory('/ssolog/'),
  routes,
})

export default router
