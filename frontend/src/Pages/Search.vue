<template>
  <SearchLayout
      @search-completed="handleSearchCompleted"
      @search-cleared="handleSearchCleared"
      v-slot="slotProps"
  >
    <!-- Состояние загрузки -->
    <div v-if="slotProps.isLoading" class="flex justify-center items-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
      <span class="ml-3 text-gray-600">Поиск...</span>
    </div>

    <!-- Результаты поиска -->
    <div v-else-if="slotProps.searchResults.length > 0" :class="slotProps.isMobile ? 'pt-4 pb-8' : 'py-8'">
      <h2 class="text-lg font-semibold text-gray-800 mb-4" :class="slotProps.isMobile ? 'px-4' : ''">
        Результаты поиска: "{{ slotProps.searchQuery }}"
      </h2>

      <!-- Используем AppRowWhite для отображения результатов -->
      <div :class="slotProps.isMobile ? 'space-y-3 px-4' : 'space-y-2'">
        <AppRowWhite
            v-for="app in slotProps.searchResults"
            :key="app.app_id"
            :app="app"
            :class="!slotProps.isMobile ? 'desktop-app-row' : ''"
        />
      </div>
    </div>

    <!-- Категории (когда нет поиска) -->
    <div v-else :class="slotProps.isMobile ? 'pt-4 pb-8' : 'py-8'">
      <!-- Desktop версия -->
      <div v-if="!slotProps.isMobile" class="grid grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 2xl:grid-cols-7 gap-4">
        <div
            v-for="category in slotProps.categories"
            :key="category.category_id"
            class="bg-white rounded-xl shadow-sm overflow-hidden hover:shadow-md transition-all cursor-pointer border border-blue-100 desktop-category-card"
            @click="selectCategory(category)"
        >
          <div class="aspect-square bg-gradient-to-br from-blue-500 to-blue-600 flex items-center justify-center">
            <span class="text-white text-2xl">{{ getCategoryEmoji(category.category_title) }}</span>
          </div>
          <div class="p-3">
            <h3 class="font-semibold text-gray-800 text-center text-sm truncate">
              {{ category.category_title }}
            </h3>
          </div>
        </div>
      </div>

      <!-- Mobile версия - 2 КАТЕГОРИИ В СТРОКЕ -->
      <div v-else class="grid grid-cols-2 gap-5 px-4">
        <div
            v-for="category in slotProps.categories"
            :key="category.category_id"
            class="bg-white rounded-2xl shadow-sm overflow-hidden hover:shadow-md transition-all cursor-pointer border border-blue-100"
            @click="selectCategory(category)"
        >
          <!-- Широкая и низкая карточка -->
          <div class="aspect-[4/3] bg-gradient-to-br from-blue-500 to-blue-600 flex items-center justify-center relative">
            <span class="text-white text-3xl">{{ getCategoryEmoji(category.category_title) }}</span>
            <!-- Градиентный оверлей для красоты -->
            <div class="absolute inset-0 bg-gradient-to-r from-blue-600/20 to-blue-500/20"></div>
          </div>
          <!-- Подпись категории -->
          <div class="p-3">
            <h3 class="font-semibold text-gray-800 text-center text-sm">
              {{ category.category_title }}
            </h3>
          </div>
        </div>
      </div>
    </div>

    <!-- Сообщение "Ничего не найдено" ТОЛЬКО для мобильных -->
    <div
        v-if="!slotProps.isLoading && slotProps.searchQuery && slotProps.searchResults.length === 0 && slotProps.isMobile"
        class="text-center py-12 px-4"
    >
      <div class="text-gray-400 text-4xl mb-3">🔍</div>
      <h3 class="text-gray-600 font-semibold text-lg mb-2">Ничего не найдено</h3>
    </div>
  </SearchLayout>
</template>

<script>
import SearchLayout from '@/components/Layout/SearchLayout.vue'
import AppRowWhite from '@/components/AppRowWhite.vue'

export default {
  name: 'SearchView',
  components: {
    SearchLayout,
    AppRowWhite
  },
  methods: {
    handleSearchCompleted(searchData) {
      console.log('Search completed:', searchData)
    },

    handleSearchCleared() {
      console.log('Search cleared')
    },

    selectCategory(category) {
      console.log('Selected category:', category)
    },

    getCategoryEmoji(categoryTitle) {
      const emojiMap = {
        'Шутеры': '🎯',
        'Банки': '🏦',
        'Мессенджеры': '💬',
        'Игры': '🎮',
        'Образование': '📚',
        'Музыка': '🎵',
        'Социальные сети': '👥',
        'Утилиты': '🛠️',
        'Карты': '🗺️',
        'Погода': '🌤️',
        'Еда': '🍕',
        'Транспорт': '🚗',
        'Здоровье': '🏥',
        'Спорт': '⚽',
        'Новости': '📰',
        'Книги': '📖',
        'Бизнес': '💼',
        'Финансы': '💰',
        'Фото': '📸',
        'Видео': '🎥'
      }
      return emojiMap[categoryTitle] || '📱'
    }
  }
}
</script>

<style scoped>
/* Дополнительные анимации для плавности */
.bg-gradient-to-br {
  transition: all 0.3s ease;
}

/* Эффект при нажатии на мобильных */
@media (max-width: 767px) {
  .cursor-pointer:active {
    transform: scale(0.98);
    transition: transform 0.1s ease;
  }
}

/* Стили для десктопа */
@media (min-width: 768px) {
  /* Выделение при наведении на приложения */
  :deep(.desktop-app-row) {
    transition: all 0.2s ease-in-out;
    border-radius: 12px;
  }

  :deep(.desktop-app-row:hover) {
    background-color: #f8fafc;
    transform: translateY(-2px);
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }

  /* Выделение при наведении на категории */
  .desktop-category-card {
    transition: all 0.2s ease-in-out;
  }

  .desktop-category-card:hover {
    transform: translateY(-4px);
    box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
    border-color: #3b82f6;
  }
}
</style>