<template>
  <div class="search-layout min-h-screen bg-gray-50">

    <!-- HEADER -->
    <header class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-50">
      <div class="w-full px-4 py-3">
        <div class="flex items-center gap-3 transition-all duration-200">

          <!-- BACK BUTTON -->
          <button
              @click="goBack"
              class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full hover:bg-gray-100 active:bg-gray-200 transition-colors"
              aria-label="Назад"
          >
            <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
          </button>

          <!-- SEARCH FIELD -->
          <div class="flex-1 relative">
            <input
                ref="searchInput"
                v-model="searchQuery"
                @input="onTyping"
                type="text"
                placeholder="Поиск приложений..."
                class="search-input w-full px-4 py-3 bg-gray-100 rounded-2xl border-0 outline-none text-gray-800 placeholder-gray-500 transition-all duration-200 focus:ring-2 focus:ring-blue-500 focus:bg-white"
            >

            <!-- CLEAR BUTTON -->
            <button
                v-if="searchQuery"
                @click="clearSearch"
                class="absolute right-3 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600 transition"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </div>

        </div>
      </div>
    </header>

    <!-- CONTENT -->
    <main class="w-full h-[calc(100vh-80px)] overflow-y-auto">
      <div :class="isMobile ? 'w-full px-4' : 'container mx-auto px-4'">
        <slot
            :search-query="searchQuery"
            :is-loading="isLoading"
            :categories="categories"
            :search-results="searchResults"
            :is-mobile="isMobile"
        />
      </div>
    </main>

  </div>
</template>

<script>
export default {
  name: "SearchLayout",

  data() {
    return {
      searchQuery: "",
      isLoading: false,
      categories: [
        { category_id: 1, category_title: "Шутеры" },
        { category_id: 2, category_title: "Банки" },
        { category_id: 3, category_title: "Мессенджеры" },
        { category_id: 4, category_title: "Игры" },
        { category_id: 5, category_title: "Образование" },
        { category_id: 6, category_title: "Музыка" },
        { category_id: 7, category_title: "Социальные сети" },
        { category_id: 8, category_title: "Утилиты" },
        { category_id: 9, category_title: "Карты" },
        { category_id: 10, category_title: "Погода" },
        { category_id: 11, category_title: "Еда" },
        { category_id: 12, category_title: "Транспорт" },
        { category_id: 13, category_title: "Здоровье" },
        { category_id: 14, category_title: "Спорт" },
        { category_id: 15, category_title: "Новости" },
        { category_id: 16, category_title: "Книги" },
        { category_id: 17, category_title: "Бизнес" },
        { category_id: 18, category_title: "Финансы" },
        { category_id: 19, category_title: "Фото" },
        { category_id: 20, category_title: "Видео" }
      ],
      searchResults: [],
      isMobile: false
    };
  },

  mounted() {
    this.isMobile = this.checkMobile();

    // стабилизированный автофокус
    setTimeout(() => {
      this.$refs.searchInput?.focus();
    }, 120);
  },

  methods: {
    checkMobile() {
      return window.innerWidth < 768 || /iPhone|Android/i.test(navigator.userAgent);
    },

    goBack() {
      this.$router.push("/");
    },

    onTyping() {
      this.isLoading = true;

      clearTimeout(this._debounce);
      this._debounce = setTimeout(() => {
        this.performSearch();
      }, 250);
    },

    performSearch() {
      if (!this.searchQuery.trim()) {
        this.clearSearch();
        return;
      }

      this.isLoading = true;

      setTimeout(() => {
        this.searchResults = [
          { app_id: 1, title: `${this.searchQuery} Game`, developer: "Game Studio" },
          { app_id: 2, title: `${this.searchQuery} App`, developer: "App Corp" },
          { app_id: 3, title: `${this.searchQuery} Tool`, developer: "Tools Inc" }
        ];

        this.isLoading = false;
      }, 200);
    },

    clearSearch() {
      this.searchQuery = "";
      this.searchResults = [];
      this.$emit("search-cleared");
      setTimeout(() => this.$refs.searchInput?.focus(), 50);
    }
  }
};
</script>

<style scoped>
/* УБИРАЕМ ОТСТУПЫ НА МОБИЛКАХ */
@media (max-width: 767px) {
  .search-layout,
  main {
    padding: 0 !important;
    margin: 0 !important;
  }
}

/* Мягкая анимация поля */
.search-input {
  transition: all 0.2s ease;
}
.search-input:focus {
  transform: scale(1.01);
}
</style>