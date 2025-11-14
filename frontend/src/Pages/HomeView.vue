<template>
  <div class="mx-auto px-4 py-4 w-full max-w-[480px] bg-white">

    <!-- SEARCH BAR -->
    <div class="bg-gray-100 w-full rounded-2xl px-4 py-3 flex items-center gap-3 shadow-sm">
      <input
          type="text"
          v-model="query"
          placeholder="Поиск"
          class="flex-1 bg-transparent outline-none text-gray-900"
      />
      <span class="text-gray-500 text-xl">🔍</span>
    </div>

    <!-- BIG PROMO BANNER -->
    <div class="w-full h-40 mt-6 rounded-3xl bg-gradient-to-r from-indigo-500 to-blue-500 flex items-center justify-center shadow-md">
      <p class="text-white text-lg font-semibold">Игра которую мы предлагаем</p>
    </div>

    <!-- POPULAR -->
    <SectionBlock title="Категория 1 (популярное)">
      <div class="flex flex-col divide-y divide-gray-200">
        <AppRowWhite
            v-for="app in popularFiltered"
            :key="app.app_id"
            :app="app"
        />
      </div>
    </SectionBlock>

    <!-- BANKS -->
    <SectionBlock title="Банки">
      <div class="w-full h-40 bg-gray-100 rounded-3xl flex items-center justify-center shadow-inner text-gray-700 font-medium">
        Банковская подборка
      </div>
    </SectionBlock>

    <!-- EDITORS CHOICE -->
    <SectionBlock title="Выбор редакции">
      <div class="flex flex-col divide-y divide-gray-200">
        <AppRowWhite
            v-for="app in editorsChoice"
            :key="app.app_id"
            :app="app"
        />
      </div>
    </SectionBlock>

    <!-- RECENT -->
    <SectionBlock title="Вы недавно смотрели">
      <div class="flex flex-col divide-y divide-gray-200">
        <AppRowWhite
            v-for="app in recent"
            :key="app.app_id"
            :app="app"
        />
      </div>
    </SectionBlock>

  </div>
</template>

<script>
import appsJson from "@/data/apps.json";
import SectionBlock from "@/components/UI/SectionBlock.vue";
import AppRowWhite from "@/components/AppRowWhite.vue";

export default {
  components: { SectionBlock, AppRowWhite },

  data() {
    return {
      query: "",
      apps: appsJson
    };
  },

  computed: {
    popularFiltered() {
      return this.apps.filter(a => a.popular);
    },
    editorsChoice() {
      return this.apps.filter(a => a.editorsChoice);
    },
    recent() {
      return this.apps.filter(a => a.recent);
    }
  }
};
</script>
