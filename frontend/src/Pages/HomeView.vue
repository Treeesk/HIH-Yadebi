<template>
  <div class="w-full min-h-screen bg-white pb-10">

    <!-- 🔥 НОВЫЙ ВЕРХНИЙ BAR -->
    <div class="w-full flex items-center justify-between py-3 px-4 bg-white">

      <!-- Стрелка назад НЕ показывается на главной -->
      <button class="opacity-0 pointer-events-none text-xl">←</button>

      <!-- Поле поиска (НЕ input, а кнопка) -->
      <div
          class="flex items-center bg-gray-100 rounded-full px-4 py-2 w-full mx-3 active:scale-95 transition cursor-pointer"
          @click="$router.push('/search')"
      >
        <span class="text-gray-400 mr-3 text-lg">🔍</span>
        <span class="text-gray-500 text-base flex-1">Поиск приложений</span>
      </div>

      <!-- Профиль -->
      <button
          @click="$router.push('/profile')"
          class="text-gray-700 text-2xl active:scale-90 transition"
      >
        👤
      </button>
    </div>

    <!-- БАННЕР -->
    <Banner text="MAX — приложение дня" class="px-4" />

    <!-- КАТЕГОРИИ -->
    <div
        v-for="cat in filteredCategories"
        :key="cat.category_name"
        class="mt-8 px-4"
    >
      <SectionBlock :title="cat.category_name">

        <!-- Банки: обычный список -->
        <template v-if="cat.category_name === 'Банки'">
          <div class="flex flex-col gap-3">
            <AppRowWhite
                v-for="app in cat.apps"
                :key="app.app_id"
                :app="app"
            />
          </div>
        </template>

        <!-- Остальные: свайп-карусель -->
        <template v-else>
          <SwipeCarousel :apps="cat.apps" />
        </template>

      </SectionBlock>
    </div>

  </div>
</template>

<script>
import Banner from "@/components/UI/Banner.vue"
import SectionBlock from "@/components/UI/SectionBlock.vue"
import SwipeCarousel from "@/components/SwipeCarousel.vue"
import AppRowWhite from "@/components/AppRowWhite.vue"

import json from "@/data/apps.json"

export default {
  name: "HomeView",

  components: {
    Banner,
    SectionBlock,
    SwipeCarousel,
    AppRowWhite
  },

  data() {
    return {
      categories: json.categories
    }
  },

  computed: {
    filteredCategories() {
      return this.categories
    }
  }
}
</script>

<style scoped>
</style>