<template>
  <div class="carousel no-scrollbar">

    <!-- Каждая страница = 3 приложения -->
    <div
        v-for="(page, index) in pages"
        :key="index"
        class="carousel-page"
    >
      <div class="flex flex-col gap-3">
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
import AppRowWhite from "@/components/AppRowWhite.vue"

export default {
  name: "SwipeCarousel",
  components: { AppRowWhite },
  props: {
    apps: Array
  },

  computed: {
    pages() {
      const chunkSize = 3
      const result = []
      for (let i = 0; i < this.apps.length; i += chunkSize) {
        result.push(this.apps.slice(i, i + chunkSize))
      }
      return result
    }
  }
}
</script>

<style scoped>
.carousel {
  display: flex;
  overflow-x: auto;
  scroll-snap-type: x mandatory;
  -webkit-overflow-scrolling: touch;
  gap: 16px;
  padding-bottom: 4px;
}

.carousel-page {
  flex-shrink: 0;
  width: 100%;
  scroll-snap-align: start;
  padding-right: 4px;
}

/* убираем скроллбар */
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>