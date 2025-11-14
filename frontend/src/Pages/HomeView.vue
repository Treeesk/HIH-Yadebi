<template>
  <div class="mx-auto px-4 py-4
              max-w-[480px]
              md:max-w-3xl
              lg:max-w-5xl">

    <!-- Поиск -->
    <SearchBar v-model="query" />

    <!-- Баннер -->
    <Banner text="Игра которую мы предлагаем" />

    <!-- Популярное -->
    <SectionBlock title="Категория 1 (популярное)">
      <div class="flex flex-col gap-3 md:grid md:grid-cols-2 lg:grid-cols-3">
        <Card
            v-for="app in popularFiltered"
            :key="app.id"
            :app="app"
            @click="openApp(app.id)"
        />
      </div>
    </SectionBlock>

    <!-- Банки -->
    <SectionBlock title="Банки">
      <div class="w-full bg-gray-200 rounded-xl h-48 md:h-64 flex items-center justify-center">
        Банковская подборка
      </div>
    </SectionBlock>

    <!-- Выбор редакции -->
    <SectionBlock title="Выбор редакции">
      <div class="grid grid-cols-3 gap-3 md:grid-cols-4 lg:grid-cols-6">
        <Card
            v-for="app in editorsChoice"
            :key="app.id"
            :app="app"
            size="small"
            @click="openApp(app.id)"
        />
      </div>
    </SectionBlock>

    <!-- Недавно смотрели -->
    <SectionBlock title="Вы недавно смотрели">
      <div class="grid grid-cols-3 gap-3 md:grid-cols-4 lg:grid-cols-6">
        <Card
            v-for="app in recent"
            :key="app.id"
            :app="app"
            size="small"
            @click="openApp(app.id)"
        />
      </div>
    </SectionBlock>

  </div>
</template>


<script>
import apps from "@/data/apps.json";
import SearchBar from "@/components/UI/SearchBar.vue";
import Banner from "@/components/UI/Banner.vue";
import Card from "@/components/Card/Card.vue";
import SectionBlock from "@/components/UI/SectionBlock.vue";

export default {
  components: { SearchBar, Banner, Card, SectionBlock },

  data() {
    return {
      query: "",
      apps
    };
  },

  computed: {
    popularFiltered() {
      return this.apps
          .filter(a => a.popular)
          .filter(a => a.title.toLowerCase().includes(this.query.toLowerCase()));
    },

    editorsChoice() {
      return this.apps.filter(a => a.editorsChoice);
    },

    recent() {
      return this.apps.filter(a => a.recent);
    }
  },

  methods: {
    openApp(id) {
      this.$router.push(`/app/${id}`);
    }
  }
};
</script>
