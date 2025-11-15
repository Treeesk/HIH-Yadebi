<template>
  <div
      class="relative overflow-hidden select-none"
      @touchstart="startSwipe"
      @touchmove="moveSwipe"
      @touchend="endSwipe"
  >
    <div
        class="flex transition-transform duration-500 ease-[cubic-bezier(.16,.84,.44,1)]"
        :style="{ transform: `translateX(-${currentPage * 100}%)` }"
    >
    <!-- каждая страница по 3 приложения -->
      <div
          v-for="(group, index) in pages"
          :key="index"
          class="w-full shrink-0"
      >
        <div class="flex flex-col divide-y divide-gray-200">
          <AppRowWhite
              v-for="app in group"
              :key="app.app_id"
              :app="app"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import AppRowWhite from '@/components/AppRowWhite.vue'

export default {
  components: { AppRowWhite },

  props: {
    apps: {
      type: Array,
      required: true
    }
  },

  data() {
    return {
      currentPage: 0,
      touchStartX: 0,
      touchEndX: 0
    }
  },

  computed: {
    pages() {
      const size = 3
      const result = []
      for (let i = 0; i < this.apps.length; i += size)
        result.push(this.apps.slice(i, i + size))
      return result
    }
  },

  methods: {
    startSwipe(e) {
      this.touchStartX = e.touches[0].clientX
    },

    moveSwipe(e) {
      this.touchEndX = e.touches[0].clientX
    },

    endSwipe() {
      const delta = this.touchStartX - this.touchEndX

      if (delta > 50 && this.currentPage < this.pages.length - 1) {
        this.currentPage++
      }

      if (delta < -50 && this.currentPage > 0) {
        this.currentPage--
      }
    }
  }
}
</script>
