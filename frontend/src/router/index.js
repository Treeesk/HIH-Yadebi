import { createRouter, createWebHistory } from 'vue-router'
import AppList from '../Pages/HomeView.vue' // Новые Pages страницы сюда подключать здесь

const routes = [
    { path: '/', component: AppList },
]

export default createRouter({
    history: createWebHistory(),
    routes,
})