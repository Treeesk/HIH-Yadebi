<template>
  <div class="mx-auto px-4 py-4 w-full max-w-[420px]">

    <SearchBar v-model="query" />

    <Banner text="Игра которую мы предлагаем" />

    <SectionBlock title="Категория 1 (популярное)">
      <div class="flex flex-col gap-4">
        <Card
            v-for="app in popularFiltered"
            :key="app.app_id"
            :app="app"
            @click="openApp(app.app_id)"
        />
      </div>
    </SectionBlock>

    <SectionBlock title="Банки">
      <div class="w-full h-40 bg-gray-200 rounded-xl flex items-center justify-center shadow-inner">
        <span class="text-gray-600 font-medium">Банковская подборка</span>
      </div>
    </SectionBlock>

    <SectionBlock title="Выбор редакции">
      <div class="flex flex-col gap-4">
        <Card
            v-for="app in editorsChoice"
            :key="app.app_id"
            :app="app"
            @click="openApp(app.app_id)"
        />
      </div>
    </SectionBlock>

    <SectionBlock title="Вы недавно смотрели">
      <div class="flex flex-col gap-4">
        <Card
            v-for="app in recent"
            :key="app.app_id"
            :app="app"
            @click="openApp(app.app_id)"
        />
      </div>
    </SectionBlock>

  </div>
</template>

<script>
import appData from "@/data/apps.json";
import SearchBar from "@/components/UI/SearchBar.vue";
import Banner from "@/components/UI/Banner.vue";
import SectionBlock from "@/components/UI/SectionBlock.vue";
import Card from "@/components/Card/Card.vue";

export default {
  components: { SearchBar, Banner, SectionBlock, Card },

  data() {
    return {
      query: "",
      apps: appData
    };
  },

  computed: {
    popularFiltered() {
      return this.apps
          .filter(a => a.popular)
          .filter(a => a.app_name.toLowerCase().includes(this.query.toLowerCase()));
    },
    editorsChoice() {
      return this.apps.filter(a => a.editorsChoice);
    },
    recent() {
      return this.apps.filter(a => a.recent);
    },
  },

  methods: {
    openApp(id) {
      this.$router.push(`/app/${id}`);
    }
  }
};
</script>
