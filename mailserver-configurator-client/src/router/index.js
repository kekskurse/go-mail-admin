import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Domains from "../views/Domains";
import Alias from "../views/Alias";
import AliasEdit from "../views/AliasEdit";
import Accounts from "../views/Accounts";
import AccountEdit from "../views/AccountEdit";
import TLSPolicy from "../views/TLSPolicy";
import TLSPolicyEdit from "../views/TLSPolicyEdit";
import Login from "../views/Login"
import Logout from "../views/Logout";

Vue.use(VueRouter)

const routes = [
  {
    path: '/domains',
    name: 'Domains',
    component: Domains
  },
  {
    path: '/alias',
    name: 'Alias',
    component: Alias
  },
  {
    path: '/alias/:id',
    name: 'AliasEdit',
    component: AliasEdit
  },
  {
    path: '/account',
    name: 'Accounts',
    component: Accounts
  },
  {
    path: '/account/:id',
    name: 'AccountEdit',
    component: AccountEdit
  },
  {
    path: '/tls',
    name: 'TLS',
    component: TLSPolicy
  },
  {
    path: '/tls/new',
    name: 'TLSNew',
    component: TLSPolicyEdit
  },
  {
    path: '/tls/:id',
    name: 'TLSEdit',
    component: TLSPolicyEdit
  },
  {
    path: '/',
    name: 'Home2',
    component: Home
  },
  {
    path: '/home',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/logout',
    name: 'Logout',
    component: Logout
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
  }
]

const router = new VueRouter({
  //mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
