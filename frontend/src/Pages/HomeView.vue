<template>
  <div class="w-full min-h-screen bg-white px-4 pt-4 pb-24 overflow-x-hidden">

    <!-- Поиск -->
    <div class="bg-gray-100 w-full rounded-2xl px-4 py-3 flex items-center gap-3 shadow-sm">
      <span class="text-gray-500 text-xl">🔍</span>
      <input
          type="text"
          v-model="query"
          placeholder="Поиск приложений"
          class="flex-1 bg-transparent outline-none text-gray-900"
      />
      <span class="text-gray-600 text-xl">👤</span>
    </div>

    <!-- Баннер -->
    <Banner class="mt-6" text="MAX — приложение дня" />

    <!-- Категории -->
    <div
        v-for="cat in filteredCategories"
        :key="cat.category_name"
        class="mt-10"
    >
      <SectionBlock :title="cat.category_name">

        <!-- КАТЕГОРИЯ: Банки (вертикальный список) -->
        <div v-if="cat.category_name.toLowerCase() === 'банк' || cat.category_name.toLowerCase() === 'банки' || cat.category_name.toLowerCase() === 'финансы'">
          <div class="flex flex-col divide-y divide-gray-100">
            <AppRowWhite
                v-for="app in cat.apps"
                :key="app.app_id"
                :app="app"
            />
          </div>
        </div>

        <!-- ВСЕ ОСТАЛЬНЫЕ КАТЕГОРИИ -->
        <div v-else>
          <SwipeCarousel :apps="cat.apps" />
        </div>

      </SectionBlock>
    </div>

  </div>
</template>

<script>
import Banner from "@/components/UI/Banner.vue"
import SectionBlock from "@/components/UI/SectionBlock.vue"
import AppRowWhite from "@/components/AppRowWhite.vue"
import SwipeCarousel from "@/components/SwipeCarousel.vue"

import json from "@/data/apps.json"

export default {
  components: {
    Banner,
    SectionBlock,
    AppRowWhite,
    SwipeCarousel,
  },

  data() {
    return {
      query: "",
      categories: json.categories
    }
  },

  computed: {
    filteredCategories() {
      const q = this.query.trim().toLowerCase()
      if (!q) return this.categories

      return this.categories
          .map(cat => ({
            ...cat,
            apps: cat.apps.filter(app =>
                app.app_name.toLowerCase().includes(q) ||
                app.app_category.toLowerCase().includes(q)
            )
          }))
          .filter(cat => cat.apps.length > 0)
    }
  }
}
</script>

<style scoped>
/* без ограничений ширины */
</style>