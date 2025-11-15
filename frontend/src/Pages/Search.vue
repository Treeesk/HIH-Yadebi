<template>
  <SearchLayout
      @search-completed="handleSearchCompleted"
      @search-cleared="handleSearchCleared"
      v-slot="slotProps"
  >
    <!-- Состояние загрузки -->
    <div v-if="slotProps.isLoading" class="flex justify-center items-center py-8">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
    </div>

    <!-- Результаты поиска -->
    <div v-else-if="slotProps.searchResults.length > 0" class="pt-4 pb-8">
      <h2 class="text-lg font-semibold text-gray-800 mb-4">
        Результаты поиска: "{{ slotProps.searchQuery }}"
      </h2>
      <div class="grid grid-cols-2 gap-4">
        <div
            v-for="app in slotProps.searchResults"
            :key="app.app_id"
            class="bg-white rounded-xl shadow-sm p-4 hover:shadow-md transition-shadow cursor-pointer border border-gray-100"
        >
          <div class="aspect-[4/3] bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg mb-3 flex items-center justify-center">
            <span class="text-blue-600 text-sm font-medium">APP</span>
          </div>
          <h3 class="font-semibold text-gray-800 text-sm truncate">{{ app.title }}</h3>
          <p class="text-xs text-gray-500 mt-1 truncate">{{ app.developer }}</p>
        </div>
      </div>
    </div>

    <!-- Категории (когда нет поиска) -->
    <div v-else class="pt-4 pb-8">
      <div class="grid grid-cols-2 gap-4">
        <div
            v-for="category in slotProps.categories"
            :key="category.category_id"
            class="bg-white rounded-xl shadow-sm overflow-hidden hover:shadow-md transition-all cursor-pointer border border-blue-100"
            @click="selectCategory(category)"
        >
          <!-- Высокий прямоугольник с эмодзи -->
          <div class="aspect-[3/4] bg-gradient-to-br from-blue-500 to-blue-600 flex items-center justify-center">
            <span class="text-white text-3xl">{{ getCategoryEmoji(category.category_title) }}</span>
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

    <!-- Сообщение, когда нет результатов -->
    <div
        v-if="!slotProps.isLoading && slotProps.searchQuery && slotProps.searchResults.length === 0"
        class="text-center py-12"
    >
      <div class="text-gray-400 text-4xl mb-3">🔍</div>
      <h3 class="text-gray-600 font-semibold text-lg mb-2">Ничего не найдено</h3>
      <p class="text-gray-500">Попробуйте изменить запрос поиска</p>
    </div>
  </SearchLayout>
</template>

<script>
import SearchLayout from '@/components/Layout/SearchLayout.vue'

export default {
  name: 'SearchView',
  components: {
    SearchLayout
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
/* Дополнительные стили не требуются */
</style>