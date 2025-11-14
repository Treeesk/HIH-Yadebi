<template>
  <div class="mx-auto px-4 py-4 w-full max-w-[480px] bg-white">

    <!-- SEARCH -->
    <div class="bg-gray-100 w-full rounded-2xl px-4 py-3 flex items-center gap-3 shadow-sm">
      <input
          type="text"
          v-model="query"
          placeholder="Поиск"
          class="flex-1 bg-transparent outline-none text-gray-900"
      />
      <span class="text-gray-500 text-xl">🔍</span>
    </div>

    <Banner text="Игра которую мы предлагаем" />

    <!-- POPULAR -->
    <SectionBlock title="Категория 1 (популярное)">
      <SwipeCarousel :apps="popularFiltered" />
    </SectionBlock>

    <!-- BANKS -->
    <SectionBlock title="Банки">
      <SwipeCarousel :apps="banksFiltered" />
    </SectionBlock>

    <!-- EDITORS -->
    <SectionBlock title="Выбор редакции">
      <SwipeCarousel :apps="editorsChoice" />
    </SectionBlock>

    <!-- RECENT -->
    <SectionBlock title="Вы недавно смотрели">
      <SwipeCarousel :apps="recent" />
    </SectionBlock>

  </div>
</template>

<script>
import appsJson from "@/data/apps.json"
import Banner from "@/components/UI/Banner.vue"
import SectionBlock from "@/components/UI/SectionBlock.vue"
import SwipeCarousel from "@/components/SwipeCarousel.vue"

export default {
  components: { Banner, SectionBlock, SwipeCarousel },

  data() {
    return {
      query: "",
      apps: appsJson
    }
  },

  computed: {
    filteredApps() {
      const q = this.query.trim().toLowerCase()
      if (!q) return this.apps
      return this.apps.filter(a =>
          a.app_name.toLowerCase().includes(q) ||
          a.app_categorie.toLowerCase().includes(q)
      )
    },

    popularFiltered() { return this.filteredApps.filter(a => a.popular) },
    editorsChoice() { return this.filteredApps.filter(a => a.editorsChoice) },
    recent() { return this.filteredApps.filter(a => a.recent) },

    // категория "Банки"
    banksFiltered() { return this.filteredApps.filter(a => a.bank) }
  }
}
</script>
