<template>
  <div class="carousel no-scrollbar">

    <!-- Одна страница = 3 карточки -->
    <div
        v-for="(page, pageIndex) in pages"
        :key="pageIndex"
        class="carousel-page"
    >
      <div class="flex flex-col gap-4">
        <AppRowWhite
            v-for="app in page"
            :key="app.app_id"
            :app="app"
        />
      </div>
    </div>

  </div>
</template>

<script>
import AppRowWhite from "@/components/AppRowWhite.vue";

export default {
  props: {
    apps: { type: Array, required: true }
  },
  components: { AppRowWhite },

  computed: {
    pages() {
      const result = [];
      const size = 3; // 3 карточки на страницу

      for (let i = 0; i < this.apps.length; i += size) {
        result.push(this.apps.slice(i, i + size));
      }

      return result;
    }
  }
};
</script>

<style scoped>
.carousel {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  -webkit-overflow-scrolling: touch;
  gap: 16px;
  padding-bottom: 8px;
  width: 100%;
}

.carousel-page {
  flex: 0 0 100%;
  scroll-snap-align: start;
  padding-right: 6px;
}

/* скрываем scrollbar */
.no-scrollbar::-webkit-scrollbar {
  display: none;
}

.no-scrollbar {
  scrollbar-width: none;
  -ms-overflow-style: none;
}
</style>