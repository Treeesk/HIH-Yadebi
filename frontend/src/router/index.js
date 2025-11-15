import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/pages/HomeView.vue'
import OpenBoard from '@/pages/OpenBoard.vue' // Изменил импорт
import { isAndroidWebView } from '/utils/device.js' // Добавь @

const routes = [
    {
        path: '/',
        component: HomeView
    },
    {
        path: '/login',
        component: OpenBoard
    },
]

const router = createRouter({
    history: createWebHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    console.log('Route change:', to.path);
    console.log('Is Android WebView:', isAndroidWebView());

    // ТОЛЬКО Android WebView может видеть login
    if (to.path === '/login' && !isAndroidWebView()) {
        console.log('Non-Android WebView trying to access login - redirecting to Home');
        return next('/');
    }

    // Android WebView всегда начинает с login
    if (isAndroidWebView() && to.path === '/') {
        console.log('Android WebView accessing root - redirecting to login');
        return next('/login');
    }

    next()
})

export default router