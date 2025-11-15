<template>
  <div class="search-layout">
    <!-- Хедер с поиском -->
    <header class="bg-white shadow-sm border-b border-gray-200 sticky top-0 z-50">
      <div class="w-full px-4 py-3">
        <div class="flex items-center gap-3">
          <!-- Кнопка назад -->
          <button
              @click="goBack"
              class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full hover:bg-gray-100 transition-colors"
          >
            <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
            </svg>
          </button>

          <!-- Поле поиска -->
          <div class="flex-1 relative">
            <input
                ref="searchInput"
                v-model="searchQuery"
                @keyup.enter="performSearch"
                type="text"
                placeholder="Поиск приложений..."
                class="w-full px-4 py-3 bg-gray-100 rounded-2xl border-0 focus:ring-2 focus:ring-blue-500 focus:bg-white transition-colors outline-none"
            >

            <button
                v-if="searchQuery"
                @click="clearSearch"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </div>

          <button
              @click="goToProfile"
              class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full hover:bg-gray-100 transition-colors"
          >
            <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
          </button>
        </div>
      </div>
    </header>

    <!-- Основной контент на ВЕСЬ ЭКРАН -->
    <main class="main-content">
      <div :class="isMobile ? 'w-full' : 'w-full max-w-7xl mx-auto px-4'">
        <slot
            :search-query="searchQuery"
            :is-loading="isLoading"
            :categories="categories"
            :search-results="searchResults"
            :perform-search="performSearch"
            :clear-search="clearSearch"
            :is-mobile="isMobile"
        ></slot>
      </div>
    </main>
  </div>
</template>

<script>
export default {
  name: 'SearchLayout',
  data() {
    return {
      searchQuery: '',
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
    }
  },
  mounted() {
    this.isMobile = window.innerWidth < 768
    if (this.isMobile && this.$refs.searchInput) {
      setTimeout(() => this.$refs.searchInput?.focus(), 300)
    }
  },
  methods: {
    goBack() { this.$router.push('/') },
    goToProfile() { alert('Профиль') },
    performSearch() {
      if (!this.searchQuery.trim()) {
        this.clearSearch()
        return
      }
      this.isLoading = true
      setTimeout(() => {
        this.searchResults = [
          { app_id: 1, title: `${this.searchQuery} Game`, developer: "Game Studio", category: "Игры" },
          { app_id: 2, title: `${this.searchQuery} App`, developer: "App Corp", category: "Утилиты" },
        ]
        this.isLoading = false
      }, 500)
    },
    clearSearch() { this.searchQuery = ''; this.searchResults = [] }
  }
}
</script>

<style scoped>
.search-layout {
  width: 100vw;
  min-height: 100vh;
  background: #f9fafb;
  display: flex;
  flex-direction: column;
}

.main-content {
  flex: 1;
  width: 100%;
  /* НЕТ overflow - скроллит body */
}

/* Для мобильных - растягиваем на всю ширину */
@media (max-width: 767px) {
  .search-layout {
    width: 100%;
  }
}
</style>