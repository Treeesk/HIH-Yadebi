<template>
  <div class="w-full min-h-screen bg-white px-4 pt-4 pb-20">

    <!-- ========================
         HEADER + PROFILE BUTTON
       ======================== -->
    <div class="w-full flex justify-between items-center mb-5">
      <h1 class="text-xl font-semibold text-gray-900">Приложения</h1>

      <button
          @click="showProfileToast"
          class="w-10 h-10 rounded-full bg-gray-200 flex items-center justify-center text-xl active:scale-95 transition"
      >
        👤
      </button>
    </div>

    <!-- ===============
         ТВОЙ СТАРЫЙ ПОИСК
       =============== -->


    <!-- ДАЛЕЕ ТВОЙ КОНТЕНТ (оставляем как есть) -->
    <Banner class="mt-6" text="MAX — приложение дня" />

    <div
        v-for="cat in filteredCategories"
        :key="cat.category_name"
        class="mt-10"
    >
      <SectionBlock :title="cat.category_name">
        <div v-if="cat.category_name.toLowerCase().includes('банк')">
          <div class="flex flex-col divide-y divide-gray-100">
            <AppRowWhite
                v-for="app in cat.apps"
                :key="app.app_id"
                :app="app"
            />
          </div>
        </div>

        <div v-else>
          <SwipeCarousel :apps="cat.apps" />
        </div>
      </SectionBlock>
    </div>

    <!-- =========================
         ТОСТ (IOS / VK STYLE)
       ========================= -->
    <div
        v-if="toastVisible"
        class="fixed top-4 left-1/2 transform -translate-x-1/2
             bg-gray-900 text-white px-5 py-3 rounded-2xl shadow-lg
             animate-toast z-50 text-sm font-medium"
    >
      ✔ Вход выполнен
    </div>

  </div>
</template>

<script>
import Banner from "@/components/UI/Banner.vue";
import SectionBlock from "@/components/UI/SectionBlock.vue";
import AppRowWhite from "@/components/AppRowWhite.vue";
import SwipeCarousel from "@/components/SwipeCarousel.vue";
import json from "@/data/apps.json";

export default {
  components: { Banner, SectionBlock, AppRowWhite, SwipeCarousel },

  data() {
    return {
      toastVisible: false,
      query: "",
      categories: json.categories
    };
  },

  methods: {
    showProfileToast() {
      this.toastVisible = true;
      setTimeout(() => {
        this.toastVisible = false;
      }, 2000);
    }
  },

  computed: {
    filteredCategories() {
      const q = this.query.trim().toLowerCase();
      if (!q) return this.categories;

      return this.categories
          .map(cat => ({
            ...cat,
            apps: cat.apps.filter(
                app =>
                    app.app_name.toLowerCase().includes(q) ||
                    app.app_category.toLowerCase().includes(q)
            )
          }))
          .filter(cat => cat.apps.length > 0);
    }
  }
};
</script>

<style scoped>
@keyframes toast-slide {
  0% {
    opacity: 0;
    transform: translate(-50%, -20px);
  }
  10% {
    opacity: 1;
    transform: translate(-50%, 0);
  }
  90% {
    opacity: 1;
    transform: translate(-50%, 0);
  }
  100% {
    opacity: 0;
    transform: translate(-50%, -20px);
  }
}

.animate-toast {
  animation: toast-slide 2s ease forwards;
}
</style>