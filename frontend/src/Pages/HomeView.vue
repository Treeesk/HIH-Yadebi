<template>
  <div class="w-full min-h-screen bg-white px-4 pt-4 pb-10">

    <!-- SEARCH -->
    <SearchBar v-model="query" />

    <!-- FIXED BANNER -->
    <Banner text="MAX — приложение дня" class="mb-6" />

    <!-- CATEGORIES FROM JSON -->
    <div
        v-for="cat in filteredCategories"
        :key="cat.category_name"
        class="mb-8"
    >
      <SectionBlock :title="cat.category_name">

        <!-- BANKS — LIST -->
        <template v-if="cat.category_name === 'Банки'">
          <div class="flex flex-col gap-3">
            <AppRowWhite
                v-for="app in cat.apps"
                :key="app.app_id"
                :app="app"
            />
          </div>
        </template>

        <!-- OTHER CATEGORIES — VERTICAL SWIPE BY 3 -->
        <template v-else>
          <SwipeCarousel :apps="cat.apps" />
        </template>

      </SectionBlock>
    </div>

  </div>
</template>

<script>
import SearchBar from "@/components/UI/SearchBar.vue"
import Banner from "@/components/UI/Banner.vue"
import SectionBlock from "@/components/UI/SectionBlock.vue"
import SwipeCarousel from "@/components/SwipeCarousel.vue"
import AppRowWhite from "@/components/AppRowWhite.vue"

import json from "@/data/apps.json"

export default {
  name: "HomeView",

  components: {
    SearchBar,
    Banner,
    SectionBlock,
    SwipeCarousel,
    AppRowWhite
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
          .map(category => ({
            ...category,
            apps: category.apps.filter(app =>
                app.app_name.toLowerCase().includes(q) ||
                app.app_category.toLowerCase().includes(q)
            )
          }))
          .filter(category => category.apps.length > 0)
    }
  }
}
</script>

<style scoped>
/* Optional — чтобы чуть сгладить фон при прокрутке */
</style>