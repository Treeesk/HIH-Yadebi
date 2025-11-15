<template>
  <div class="carousel no-scrollbar" ref="container">

    <!-- Каждая страница = 3 карточки В СТОЛБИК -->
    <div
        v-for="(page, i) in pages"
        :key="i"
        class="carousel-page snap-start"
    >
      <div class="flex flex-col gap-2">
        <SmallCard
            v-for="app in page"
            :key="app.app_id"
            :app="app"
        />
      </div>
    </div>

  </div>
</template>

<script>
import SmallCard from "./SmallCard.vue"

export default {
  props: {
    apps: Array
  },

  components: { SmallCard },

  computed: {
    pages() {
      const size = 3 // ⬅ РОВНО 3 КАРТОЧКИ НА СТРАНИЦУ
      const result = []

      for (let i = 0; i < this.apps.length; i += size) {
        result.push(this.apps.slice(i, i + size))
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
  padding-bottom: 6px;
}

/* ширина страницы = ширина экрана */
.carousel-page {
  flex-shrink: 0;
  width: 100%;
  scroll-snap-align: start;
}

/* скрыть полосу прокрутки */
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>