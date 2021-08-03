import { createRouter, createWebHistory } from 'vue-router';
import PawnByCategory from '../views/PawnByCategory.vue';
import ImportFile from '../views/ImportFile.vue';
import ExportFile from '../views/ExportFile.vue';
import NotFound from '../views/NotFound.vue';
import store from '../store';

const routes = [
  {
    path: '/',
    name: 'home',
    redirect: { name: 'import-file' },
  },
  {
    path: '/import-file',
    name: 'import-file',
    component: ImportFile,
  },
  {
    path: '/export-file',
    name: 'export-file',
    component: ExportFile,
  },
  {
    path: '/pawns/:category(main|first|second)',
    name: 'pawns',
    component: PawnByCategory,
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: NotFound,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
  linkActiveClass: 'active',
  linkExactActiveClass: 'exact-active',
});

// eslint-disable-next-line no-unused-vars
router.beforeEach((to, from, next) => {
  const requiresUserFile = [
    'export-file',
    'pawns',
  ];

  if (requiresUserFile.includes(to.name) && store.state.hasUserFile === false) {
    next({ name: 'home' });
  } else {
    next();
  }
});

export default router;
