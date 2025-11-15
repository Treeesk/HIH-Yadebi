import { createRouter, createWebHistory } from 'vue-router'
import AppList from '../Pages/HomeView.vue' // Новые Pages страницы сюда подключать здесь
import OpenBoard from '../Pages/OpenBoard.vue'

const routes = [
    { path: '/', component: AppList },
    { path: '/login', component: OpenBoard},
]

export default createRouter({
    history: createWebHistory(),
    routes,
})
