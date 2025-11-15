<template>
  <div class="search-layout min-h-screen bg-gray-50">
    <!-- Хедер с поиском -->
    <header class="bg-white shadow-sm border-b border-gray-200">
      <div class="w-full px-4 py-3">
        <div class="flex items-center gap-3">
          <!-- Кнопка назад -->
          <button
              @click="goBack"
              class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full hover:bg-gray-100 transition-colors active:bg-gray-200"
              aria-label="Назад"
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
                class="w-full px-4 py-3 bg-gray-100 rounded-2xl border-0 focus:ring-2 focus:ring-blue-500 focus:bg-white transition-colors outline-none text-gray-800 placeholder-gray-500"
            >

            <!-- Кнопка очистки -->
            <button
                v-if="searchQuery"
                @click="clearSearch"
                class="absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400 hover:text-gray-600 transition-colors"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
              </svg>
            </button>
          </div>

          <!-- Иконка пользователя -->
          <button
              @click="goToProfile"
              class="flex-shrink-0 w-10 h-10 flex items-center justify-center rounded-full hover:bg-gray-100 transition-colors active:bg-gray-200"
              aria-label="Профиль"
          >
            <svg class="w-6 h-6 text-gray-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"/>
            </svg>
          </button>
        </div>
      </div>
    </header>

    <!-- Основной контент -->
    <main class="w-full h-[calc(100vh-80px)] overflow-y-auto">
      <div :class="isMobile ? 'w-full' : 'container mx-auto px-4'">
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
    // Определяем тип устройства
    this.isMobile = this.checkIsMobile()

    // Фокус на поле ввода на мобильных устройствах
    if (this.isMobile && this.$refs.searchInput) {
      setTimeout(() => {
        this.$refs.searchInput.focus()
        setTimeout(() => {
          this.$refs.searchInput.focus()
        }, 100)
      }, 300)
    }
  },
  methods: {
    checkIsMobile() {
      return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ||
          window.innerWidth < 768
    },

    goBack() {
      this.$router.push('/')
    },

    goToProfile() {
      alert('Переход в профиль пользователя')
    },

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
          { app_id: 3, title: `${this.searchQuery} Tool`, developer: "Tool Inc", category: "Инструменты" },
          { app_id: 4, title: `${this.searchQuery} Messenger`, developer: "Message Corp", category: "Мессенджеры" },
          { app_id: 5, title: `${this.searchQuery} Bank`, developer: "Bank Inc", category: "Банки" },
          { app_id: 6, title: `${this.searchQuery} Music`, developer: "Music Corp", category: "Музыка" }
        ]
        this.isLoading = false

        this.$emit('search-completed', {
          query: this.searchQuery,
          results: this.searchResults
        })
      }, 500)
    },

    clearSearch() {
      this.searchQuery = ''
      this.searchResults = []
      this.$emit('search-cleared')
    }
  },
  watch: {
    searchQuery(newQuery) {
      if (!newQuery.trim()) {
        this.clearSearch()
      }
    }
  },
  emits: ['search-completed', 'search-cleared']
}
</script>

<style scoped>
.search-layout {
  min-height: 100vh;
}

/* Убираем ВСЕ отступы на мобильных */
@media (max-width: 767px) {
  .search-layout {
    padding-left: 0 !important;
    padding-right: 0 !important;
    margin-left: 0 !important;
    margin-right: 0 !important;
  }

  main {
    padding-left: 0 !important;
    padding-right: 0 !important;
  }
}

.h-\[calc\(100vh-80px\)\]::-webkit-scrollbar {
  width: 4px;
}

.h-\[calc\(100vh-80px\)\]::-webkit-scrollbar-track {
  background: #f1f5f9;
  border-radius: 2px;
}

.h-\[calc\(100vh-80px\)\]::-webkit-scrollbar-thumb {
  background: #cbd5e1;
  border-radius: 2px;
}

.h-\[calc\(100vh-80px\)\]::-webkit-scrollbar-thumb:hover {
  background: #94a3b8;
}
</style>