<template>
  <div class="w-full min-h-screen bg-white px-4 pt-4 pb-24 overflow-x-hidden">
    <Banner class="mt-6" text="MAX — приложение дня" />

    <div v-for="cat in filteredCategories" :key="cat.category_name" class="mt-10">
      <SectionBlock :title="cat.category_name">

        <div v-if="cat.category_name.toLowerCase().includes('банк')">
          <div class="flex flex-col divide-y divide-gray-100">
            <AppRowWhite v-for="app in cat.apps" :key="app.app_id" :app="app" />
          </div>
        </div>

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
  components: { Banner, SectionBlock, AppRowWhite, SwipeCarousel },

  data() {
    return {
      query: "",
      categories: json.categories
    }
  },

  computed: {
    filteredCategories() {
      const q = this.query.toLowerCase()
      return this.categories
          .map(cat => ({
            ...cat,
            apps: cat.apps.filter(app =>
                app.app_name.toLowerCase().includes(q) ||
                app.app_category.toLowerCase().includes(q)
            )
          }))
          .filter(cat => cat.apps.length)
    }
  }
}
</script>