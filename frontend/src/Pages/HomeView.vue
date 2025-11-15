<template>
  <!-- ФОН НА ВЕСЬ ЭКРАН -->
  <div class="w-full min-h-screen bg-white">

    <!-- ЦЕНТРАЛЬНЫЙ КОНТЕЙНЕР -->
    <div class="mx-auto w-full max-w-[600px] px-3 py-4">

      <!-- SEARCH -->
      <div class="bg-gray-100 w-full rounded-2xl px-4 py-3 flex items-center gap-3 shadow-sm mb-6">
        <input
            type="text"
            v-model="query"
            placeholder="Поиск"
            class="flex-1 bg-transparent outline-none text-gray-900"
        />
        <span class="text-gray-500 text-xl">🔍</span>
      </div>

      <!-- БАННЕР: фиксированно показываем MAX -->
      <Banner text="MAX — приложение дня" />

      <!-- ВСЕ КАТЕГОРИИ ИЗ JSON -->
      <div
          v-for="cat in filteredCategories"
          :key="cat.category_name"
      >
        <SectionBlock :title="cat.category_name">

          <!-- ЕСЛИ ЭТО КАТЕГОРИЯ 'Банки' — ПОКАЗЫВАЕМ ПО ОДНОМУ -->
          <div v-if="cat.category_name === 'Банки'">
            <div class="flex flex-col divide-y divide-gray-200">
              <AppRowWhite
                  v-for="app in cat.apps"
                  :key="app.app_id"
                  :app="app"
              />
            </div>
          </div>

          <!-- ДЛЯ ВСЕХ ОСТАЛЬНЫХ КАТЕГОРИЙ — КАРУСЕЛЬ -->
          <div v-else>
            <SwipeCarousel :apps="cat.apps" />
          </div>

        </SectionBlock>
      </div>

    </div>
  </div>
</template>

<script>
import Banner from "@/components/UI/Banner.vue"
import SectionBlock from "@/components/UI/SectionBlock.vue"
import SwipeCarousel from "@/components/SwipeCarousel.vue"
import AppRowWhite from "@/components/AppRowWhite.vue"

// JSON от бэка, формата:
// { "categories": [ { "category_name": "...", "apps": [ { app_id, app_name, app_category } ] } ] }
import json from "@/data/apps.json"

export default {
  components: { Banner, SectionBlock, SwipeCarousel, AppRowWhite },

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

      // фильтруем вложенные приложения по имени/жанру
      return this.categories
          .map(category => ({
            ...category,
            apps: category.apps.filter(app =>
                app.app_name.toLowerCase().includes(q) ||
                app.app_category.toLowerCase().includes(q)
            )
          }))
          // убираем категории, где после фильтрации нет приложений
          .filter(category => category.apps.length > 0)
    }
  }
}
</script>