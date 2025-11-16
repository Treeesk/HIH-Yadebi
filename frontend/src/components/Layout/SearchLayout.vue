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
                @keyup.enter="handleEnter"
                @blur="handleInputBlur"
                type="text"
                placeholder="Поиск приложений..."
                class="w-full px-4 py-3 bg-gray-100 rounded-2xl border-0 focus:ring-2 focus:ring-blue-500 focus:bg-white transition-colors outline-none text-gray-900 placeholder-gray-500"
                inputmode="search"
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
      isMobile: false,
      keyboardInterval: null,
      shouldKeepKeyboardOpen: true
    }
  },
  mounted() {
    this.isMobile = this.checkIsMobile()

    // ВОССТАНАВЛИВАЕМ СОСТОЯНИЕ ПОИСКА ИЗ sessionStorage
    this.restoreSearchState()

    // АВТОМАТИЧЕСКОЕ ОТКРЫТИЕ КЛАВИАТУРЫ НА МОБИЛЬНЫХ
    if (this.isMobile) {
      this.$nextTick(() => {
        this.keepKeyboardOpen()
      })
    }
  },
  beforeUnmount() {
    // Очищаем интервал при размонтировании компонента
    if (this.keyboardInterval) {
      clearInterval(this.keyboardInterval)
    }
  },
  methods: {
    checkIsMobile() {
      return /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(navigator.userAgent) ||
          window.innerWidth < 768
    },

    // ОБРАБОТКА НАЖАТИЯ ENTER
    handleEnter() {
      // Выполняем поиск
      this.performSearch()

      // Скрываем клавиатуру на мобильных устройствах
      if (this.isMobile) {
        this.hideKeyboard()
      }
    },

    // СКРЫТИЕ КЛАВИАТУРЫ
    hideKeyboard() {
      const input = this.$refs.searchInput
      if (!input) return

      console.log('Hiding keyboard...')

      // Отключаем автоматическое открытие клавиатуры
      this.shouldKeepKeyboardOpen = false

      // Очищаем интервал автоматического фокуса
      if (this.keyboardInterval) {
        clearInterval(this.keyboardInterval)
        this.keyboardInterval = null
      }

      // Убираем фокус с инпута
      input.blur()

      // Для iOS - дополнительно скрываем клавиатуру
      if (/iPhone|iPad|iPod/.test(navigator.userAgent)) {
        // На iOS можно попробовать изменить тип инпута
        const originalType = input.type
        input.type = 'text'
        setTimeout(() => {
          input.type = originalType
        }, 100)
      }

      // Для Android - просто blur обычно достаточно
      if (/Android/.test(navigator.userAgent)) {
        // Дополнительный blur для надежности
        setTimeout(() => {
          input.blur()
        }, 100)
      }
    },

    // ВОССТАНОВЛЕНИЕ СОСТОЯНИЯ ПОИСКА
    restoreSearchState() {
      const savedState = sessionStorage.getItem('searchState')
      if (savedState) {
        try {
          const state = JSON.parse(savedState)
          this.searchQuery = state.query || ''
          this.searchResults = state.results || []
          console.log('Search state restored:', state)
        } catch (error) {
          console.error('Error restoring search state:', error)
        }
      }
    },

    // СОХРАНЕНИЕ СОСТОЯНИЯ ПОИСКА
    saveSearchState() {
      const state = {
        query: this.searchQuery,
        results: this.searchResults
      }
      sessionStorage.setItem('searchState', JSON.stringify(state))
    },

    // ОЧИСТКА СОХРАНЕННОГО СОСТОЯНИЯ
    clearSearchState() {
      sessionStorage.removeItem('searchState')
    },

    // МЕТОД ДЛЯ ПОСТОЯННОГО УДЕРЖАНИЯ КЛАВИАТУРЫ
    keepKeyboardOpen() {
      const input = this.$refs.searchInput
      if (!input) {
        console.log('Input not found')
        return
      }

      console.log('Keeping keyboard open...')

      // Функция для фокусировки
      const focusInput = () => {
        if (document.activeElement !== input && this.shouldKeepKeyboardOpen) {
          input.focus()
          input.setSelectionRange(input.value.length, input.value.length)
        }
      }

      // Сразу фокусируем
      focusInput()

      // Для iOS - специальная обработка
      if (/iPhone|iPad|iPod/.test(navigator.userAgent)) {
        this.handleIOSKeyboard(input, focusInput)
      } else {
        // Для Android и других устройств
        this.handleAndroidKeyboard(input, focusInput)
      }
    },

    // ОБРАБОТКА ДЛЯ iOS
    handleIOSKeyboard(input, focusInput) {
      // Метод 1: Множественный фокус с задержками
      setTimeout(focusInput, 100)
      setTimeout(focusInput, 300)
      setTimeout(focusInput, 500)
      setTimeout(focusInput, 1000)

      // Метод 2: Интервал для постоянного фокуса
      this.keyboardInterval = setInterval(focusInput, 2000)
    },

    // ОБРАБОТКА ДЛЯ ANDROID
    handleAndroidKeyboard(input, focusInput) {
      // Метод 1: Быстрое переключение фокуса
      setTimeout(() => {
        input.blur()
        setTimeout(focusInput, 50)
      }, 200)

      // Метод 2: Множественный фокус
      setTimeout(focusInput, 400)
      setTimeout(focusInput, 800)
      setTimeout(focusInput, 1200)

      // Метод 3: Интервал для постоянного фокуса
      this.keyboardInterval = setInterval(focusInput, 1500)
    },

    // ОБРАБОТЧИК ПОТЕРИ ФОКУСА
    handleInputBlur() {
      if (this.isMobile && this.shouldKeepKeyboardOpen) {
        console.log('Input lost focus, refocusing...')
        // Немедленно возвращаем фокус
        setTimeout(() => {
          const input = this.$refs.searchInput
          if (input && document.activeElement !== input && this.shouldKeepKeyboardOpen) {
            input.focus()
          }
        }, 100)
      }
    },

    goBack() {
      // Очищаем интервал перед уходом
      if (this.keyboardInterval) {
        clearInterval(this.keyboardInterval)
      }
      this.$router.push('/')
    },

    goToProfile() {
      alert('Профиль')
    },

    async performSearch() {
      if (!this.searchQuery.trim()) {
        this.clearSearch()
        return
      }

      this.isLoading = true

      try {
        // ЗАГЛУШКА ДЛЯ ДЕМОНСТРАЦИИ - ЗАМЕНИТЕ НА РЕАЛЬНЫЙ API
        console.log('Searching for:', this.searchQuery)

        // Имитация API запроса
        await new Promise(resolve => setTimeout(resolve, 1000))

        // ЗАГЛУШЕЧНЫЕ ДАННЫЕ ДЛЯ ТЕСТИРОВАНИЯ
        this.searchResults = [
          {
            app_id: 1,
            app_name: `${this.searchQuery} Game`,
            app_category: "Игры",
            app_little_icon_link: "https://api.dicebear.com/7.x/icons/svg?seed=game&scale=90&size=200&radius=20&backgroundColor=b6e3f4"
          },
          {
            app_id: 2,
            app_name: `${this.searchQuery} App`,
            app_category: "Утилиты",
            app_little_icon_link: "https://api.dicebear.com/7.x/icons/svg?seed=app&scale=90&size=200&radius=20&backgroundColor=ffdfbf"
          },
          {
            app_id: 3,
            app_name: `${this.searchQuery} Tool`,
            app_category: "Инструменты",
            app_little_icon_link: "https://api.dicebear.com/7.x/icons/svg?seed=tool&scale=90&size=200&radius=20&backgroundColor=c0aede"
          }
        ]

        // СОХРАНЯЕМ РЕЗУЛЬТАТЫ ПОИСКА
        this.saveSearchState()

        this.$emit('search-completed', {
          query: this.searchQuery,
          results: this.searchResults
        })

      } catch (error) {
        console.error('Ошибка при поиске:', error)
        this.searchResults = []
      } finally {
        this.isLoading = false
      }
    },

    clearSearch() {
      this.searchQuery = '';
      this.searchResults = []
      // ВКЛЮЧАЕМ АВТОМАТИЧЕСКОЕ ОТКРЫТИЕ КЛАВИАТУРЫ ПРИ ОЧИСТКЕ
      this.shouldKeepKeyboardOpen = true
      // ОЧИЩАЕМ СОХРАНЕННОЕ СОСТОЯНИЕ ПРИ ОЧИСТКЕ ПОИСКА
      this.clearSearchState()
      this.$emit('search-cleared')
    }
  },
  emits: ['search-completed', 'search-cleared']
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
}

/* Для мобильных - растягиваем на всю ширину */
@media (max-width: 767px) {
  .search-layout {
    width: 100%;
  }

  header {
    position: sticky;
    top: 0;
    z-index: 50;
  }
}

/* Важные стили для работы клавиатуры */
input {
  font-size: 16px !important;
  transform: translateZ(0);
}

/* Убираем outline для мобильных чтобы не мешал */
@media (max-width: 767px) {
  input:focus {
    outline: none;
    box-shadow: none;
  }
}

/* Для Safari на iOS */
@supports (-webkit-touch-callout: none) {
  input {
    font-size: 16px !important;
  }
}
</style>