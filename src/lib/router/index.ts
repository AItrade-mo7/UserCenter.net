import { createRouter, createWebHistory } from 'vue-router';
import { getToken } from '@/utils/tools';

const routes: any = [
  {
    path: '/',
    component: () => import('@/pages/HomePage.vue'),
  },
  {
    path: '/login',
    component: () => import('@/pages/LoginPage.vue'),
  },
  {
    path: '/register',
    component: () => import('@/pages/RegisterPage.vue'),
  },
  {
    path: '/about',
    component: () => import('@/pages/About/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'list',
        component: () => import('@/pages/About/ListPage.vue'),
      },
      {
        path: 'pwa',
        description: 'PWA应用安装指南',
        component: () => import('@/pages/About/PWA.vue'),
      },
      {
        path: 'duty',
        description: '用户协议',
        component: () => import('@/pages/About/DutyPage.vue'),
      },
      {
        path: 'release_notes',
        description: '版本说明',
        component: () => import('@/pages/About/ReleaseNotes.vue'),
      },
      {
        path: 'okxkey',
        description: 'Okx 密钥',
        component: () => import('@/pages/About/OkxKey.vue'),
      },
      {
        path: 'FundServer',
        description: 'FundServer 说明',
        component: () => import('@/pages/About/FundServe.vue'),
      },
    ],
  },
  {
    path: '/secret_key',
    isLogin: true,
    component: () => import('@/pages/SecretKey/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'list',
        component: () => import('@/pages/SecretKey/ListPage.vue'),
      },
      {
        path: 'add',
        description: '新增 密钥',
        component: () => import('@/pages/SecretKey/AddPage.vue'),
      },
    ],
  },
  {
    path: '/CoinServe',
    isLogin: true,
    component: () => import('@/pages/CoinServe/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'list',
        component: () => import('@/pages/CoinServe/ListPage.vue'),
      },
      {
        path: 'add',
        description: '新增 服务',
        component: () => import('@/pages/CoinServe/AddPage.vue'),
      },
      {
        path: 'info',
        description: '查看详情',
        component: () => import('@/pages/CoinServe/InfoPage.vue'),
      },
    ],
  },
  {
    path: '/personal',
    isLogin: true,
    component: () => import('@/pages/PersonalPage.vue'),
  },
  {
    path: '/edit_profile',
    isLogin: true,
    component: () => import('@/pages/EditProfile.vue'),
  },
  {
    path: '/change_password',
    component: () => import('@/pages/ChangePassword.vue'),
  },
  {
    path: '/:pathMatch(.*)',
    component: () => import('@/pages/NotFound.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

router.beforeEach((to) => {
  const Token = getToken();
  if (!Token) {
    for (const route of routes) {
      if (to.path.indexOf(route.path) > -1) {
        if (route.isLogin) {
          return { path: '/login' };
        }
      }
    }
  }
});

export { router, routes };
