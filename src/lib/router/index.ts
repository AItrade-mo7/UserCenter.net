import { createRouter, createWebHistory } from 'vue-router';
import { getToken } from '@/utils/tools';

const routes: any = [
  {
    path: '/',
    component: () => import('@/pages/HomePage.vue'),
  },
  {
    path: '/Login',
    component: () => import('@/pages/LoginPage.vue'),
  },
  {
    path: '/Register',
    component: () => import('@/pages/RegisterPage.vue'),
  },
  {
    path: '/About',
    component: () => import('@/pages/About/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'List',
        component: () => import('@/pages/About/ListPage.vue'),
      },

      {
        path: 'PWA',
        description: 'PWA应用安装指南',
        component: () => import('@/pages/About/PWA.vue'),
      },
      {
        path: 'Duty',
        description: '用户协议',
        component: () => import('@/pages/About/DutyPage.vue'),
      },
      {
        path: 'ReleaseNotes',
        description: '版本说明',
        component: () => import('@/pages/About/ReleaseNotes.vue'),
      },
      {
        path: 'OkxKey',
        description: 'Okx 密钥',
        component: () => import('@/pages/About/OkxKey.vue'),
      },
      {
        path: 'AIServe',
        description: 'AIServe 说明',
        component: () => import('@/pages/About/AIServe.vue'),
      },
      {
        path: 'CoinServe',
        description: 'CoinServe 使用说明',
        component: () => import('@/pages/About/AIServe.vue'),
      },
    ],
  },
  {
    path: '/OkxKey',
    isLogin: true,
    component: () => import('@/pages/OkxKey/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'list',
        component: () => import('@/pages/OkxKey/ListPage.vue'),
      },
      {
        path: 'Add',
        description: '新增 密钥',
        component: () => import('@/pages/OkxKey/AddPage.vue'),
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
        description: 'List',
        component: () => import('@/pages/CoinServe/ListPage.vue'),
      },
      {
        path: 'Add',
        description: '新增 服务',
        component: () => import('@/pages/CoinServe/AddPage.vue'),
      },
      {
        path: 'Info',
        description: '查看详情',
        component: () => import('@/pages/CoinServe/InfoPage.vue'),
      },
    ],
  },
  {
    path: '/StockServe',
    isLogin: true,
    component: () => import('@/pages/StockServe/IndexPage.vue'),
    children: [
      {
        path: '',
        description: 'List',
        component: () => import('@/pages/StockServe/ListPage.vue'),
      },
    ],
  },
  {
    path: '/Personal',
    isLogin: true,
    component: () => import('@/pages/PersonalPage.vue'),
  },
  {
    path: '/EditProfile',
    isLogin: true,
    component: () => import('@/pages/EditProfile.vue'),
  },
  {
    path: '/ChangePassword',
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
          return { path: '/Login' };
        }
      }
    }
  }
});

export { router, routes };
