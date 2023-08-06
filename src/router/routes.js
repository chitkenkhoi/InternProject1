import LayoutLogin from '../layout/LayoutLogin/LayoutLogin.vue'
import LayoutPage from '../layout/LayoutPage/LayoutPage.vue'
import LoginPage from '../views/LoginPage/LoginPage.vue'
import SignupPage from '../views/SignupPage/SignupPage.vue'
import Company from '../views/PageAfterLogin/Company/Company.vue'
import Memberlist from '../views/PageAfterLogin/MemberList/MemberList.vue'
import Error from '../views/Error.vue'
import Work1 from '../views/works/work1.vue'
import Work2 from '../views/works/work2.vue'
import Work3 from '../views/works/work3.vue'
import Work4 from '../views/works/work4.vue'
import Work5 from '../views/works/work5.vue'
import Regular from '../views/PageAfterLogin/MemberList/Profile.vue'
import Admin from '../views/PageAfterLogin/MemberList/ProfileAdmin.vue'
import TwoFA from '../views/LoginPage/TwoFA.vue'
import ForgetPassword from '../views/PasswordForget/mainPage.vue'
import Email_Confirm from '../views/PasswordForget/Email_Confirm.vue'
import { createRouter, createWebHistory } from "vue-router"
const routes = [
    {
      path: '/',
      component: LayoutPage,
      name: 'LayoutPage',
      children: [
        {
            path: 'dashboard',
            component: Work1,
        },
        {
            path: 'folder',
            component: Work2,
        },
        {
            path: 'template',
            component: Work3,
        },
        {
            path: '/management',
            component: Work4,
            children: [
                {
                    path: 'company',
                    component: Company,
                },
                {
                    path: 'member',
                    component: Memberlist,
                },
            ]
        },
        {
            path: 'customer',
            component: Work5,
        },
        {
          path: 'profile.admin/:type/:id',
          component: Admin,
        }
        ,
        { 
          path: 'profile/:type/:id',
          component: Regular,
        }
      ],
    },
  {
      path: '/auth',
      component: LayoutLogin,
      children: [
        {
          path: 'login',
          component: LoginPage,
          name : 'Login'
        },
        {
          path: 'signup',
          component: SignupPage,
        },
        {
          path:'2FA',
          component: TwoFA,
          name: 'TwoFA'
        },
        {
          path:'password_forget',
          component: ForgetPassword,
          name: 'ForgetPassword'
        },
        {
          path:'emailConfirm/:email',
          component: Email_Confirm,
          name: 'Email_Confirm'
        }
      ],
    },
    {path: '/:pathMatch(.*)*',component: Error},
  ]
const router = createRouter({
    history: createWebHistory(),
    routes
})
export default router