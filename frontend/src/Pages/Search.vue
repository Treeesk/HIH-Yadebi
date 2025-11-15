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
    <div v-else-if="slotProps.searchResults.length > 0" :class="slotProps.isMobile ? 'pt-4 pb-8' : 'py-8'">
      <h2 class="text-lg font-semibold text-gray-800 mb-4" :class="slotProps.isMobile ? 'px-4' : ''">
        Результаты поиска: "{{ slotProps.searchQuery }}"
      </h2>

      <div :class="slotProps.isMobile ? 'grid grid-cols-2 gap-3 px-2' : 'grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 gap-4'">
        <div
            v-for="app in slotProps.searchResults"
            :key="app.app_id"
            class="bg-white rounded-xl shadow-sm p-3 hover:shadow-md transition-shadow cursor-pointer border border-gray-100"
        >
          <div class="aspect-square bg-gradient-to-br from-blue-50 to-blue-100 rounded-lg mb-2 flex items-center justify-center">
            <span class="text-blue-600 text-sm font-medium">APP</span>
          </div>
          <h3 class="font-semibold text-gray-800 text-sm truncate">{{ app.title }}</h3>
        </div>
      </div>
    </div>

    <!-- Категории (когда нет поиска) - РАСТЯНУТЫ НА ВЕСЬ ЭКРАН -->
    <div v-else :class="slotProps.isMobile ? 'pt-4 pb-8' : 'py-8'">
      <!-- Desktop версия - МНОГО КОЛОНОК -->
      <div v-if="!slotProps.isMobile" class="grid grid-cols-3 md:grid-cols-4 lg:grid-cols-5 xl:grid-cols-6 2xl:grid-cols-7 gap-4">
        <div
            v-for="category in slotProps.categories"
            :key="category.category_id"
            class="bg-white rounded-xl shadow-sm overflow-hidden hover:shadow-md transition-all cursor-pointer border border-blue-100"
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

      <!-- Mobile версия - 2 колонки на ВСЮ ШИРИНУ -->
      <div v-else class="grid grid-cols-2 gap-3 px-2">
        <div
            v-for="category in slotProps.categories"
            :key="category.category_id"
            class="bg-white rounded-xl shadow-sm overflow-hidden hover:shadow-md transition-all cursor-pointer border border-blue-100"
            @click="selectCategory(category)"
        >
          <div class="aspect-[3/4] bg-gradient-to-br from-blue-500 to-blue-600 flex items-center justify-center">
            <span class="text-white text-2xl">{{ getCategoryEmoji(category.category_title) }}</span>
          </div>
          <div class="p-3">
            <h3 class="font-semibold text-gray-800 text-center text-sm truncate">
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
        :class="slotProps.isMobile ? 'px-4' : ''"
    >
      <div class="text-gray-400 text-4xl mb-3">🔍</div>
      <h3 class="text-gray-600 font-semibold text-lg mb-2">Ничего не найдено</h3>
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