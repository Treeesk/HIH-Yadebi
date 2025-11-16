<template>
  <div class="w-full min-h-screen bg-white p-4">

    <!-- HEADER -->
    <div class="flex justify-between items-center mb-4">
      <button @click="$router.back()" class="text-3xl">⬅️</button>
      <button class="text-3xl">↗️</button>
    </div>

    <!-- Состояние загрузки -->
    <div v-if="isLoading" class="flex justify-center items-center py-12">
      <div class="animate-spin rounded-full h-8 w-8 border-b-2 border-blue-500"></div>
      <span class="ml-3 text-gray-600">Загрузка...</span>
    </div>

    <!-- Контент приложения -->
    <div v-else-if="app && app.app_id">
      <!-- TOP BLOCK -->
      <div class="flex gap-4">
        <!-- ICON -->
        <div class="w-24 h-24 rounded-3xl bg-gray-200 flex justify-center items-center text-5xl overflow-hidden">
          <img
              v-if="app.app_link_large_icon"
              :src="`https://${app.app_link_large_icon}`"
              :alt="app.app_title || app.app_name"
              class="w-full h-full object-cover"
          >
          <span v-else>🤖</span>
        </div>

        <!-- TITLE -->
        <div class="flex flex-col justify-center">
          <h1 class="text-2xl font-bold text-gray-900">{{ app.app_title || app.app_name }}</h1>
          <p class="text-gray-700 text-lg font-medium">{{ app.app_category }}</p>
        </div>
      </div>

      <!-- INFO ROW -->
      <div class="flex justify-between mt-6 text-sm text-gray-800">
        <div class="text-center">
          <div class="font-semibold text-lg">{{ app.rating || "4.5" }} ⭐</div>
          <div>отзывов: {{ app.app_count_downloads || app.app_count_downloas || 1200 }}</div>
        </div>

        <div class="text-center">
          <div class="font-semibold text-lg">{{ formatSize(app.app_size) }}</div>
          <div>размер</div>
        </div>

        <div class="text-center">
          <div class="font-semibold text-lg">{{ app.app_age_rating || "3+" }}+</div>
          <div>возраст</div>
        </div>
      </div>

      <!-- DOWNLOAD BUTTON -->
      <div class="mt-6">
        <button
            @click="downloadApp"
            class="w-full py-4 bg-blue-500 text-white text-lg font-semibold rounded-xl hover:bg-blue-600 active:scale-95 transition-all shadow-lg"
        >
          📥 Скачать приложение
        </button>
      </div>

      <!-- GALLERY -->
      <div class="mt-8">
        <h2 class="text-xl font-bold text-gray-900 mb-3">📸 Скриншоты</h2>
        <div class="flex gap-3 overflow-x-auto pb-2">
          <div
              v-for="n in 3"
              :key="n"
              class="w-40 h-56 rounded-xl bg-gray-200 flex items-center justify-center text-4xl shrink-0"
          >
            🖼️
          </div>
        </div>
      </div>

      <!-- DESCRIPTION -->
      <div class="mt-8">
        <h2 class="text-xl font-bold text-gray-900 mb-3">📝 Описание</h2>
        <p class="text-gray-800 mt-2 leading-relaxed">
          {{ app.app_description || "Описание приложения будет здесь..." }}
        </p>
      </div>

      <!-- ADDITIONAL INFO -->
      <div class="mt-6 grid grid-cols-2 gap-4 text-sm text-gray-800 bg-gray-50 p-4 rounded-xl">
        <div>
          <span class="font-semibold">Версия:</span> {{ app.app_version || "1.0.0" }}
        </div>
        <div>
          <span class="font-semibold">Загрузок:</span> {{ app.app_count_downloads || app.app_count_downloas || 1200 }}
        </div>
        <div>
          <span class="font-semibold">Разработчик:</span> ID {{ app.app_developer_id }}
        </div>
        <div>
          <span class="font-semibold">Дата выхода:</span> {{ formatDate(app.app_created_at) }}
        </div>
      </div>

      <!-- REVIEWS -->
      <div class="mt-8 mb-20">
        <h2 class="text-xl font-bold text-gray-900 mb-4">⭐ Отзывы</h2>

        <div class="flex items-center gap-4 mb-2">
          <div class="text-4xl font-bold text-gray-900">{{ app.rating || "4.5" }}</div>
          <div class="text-yellow-500 text-xl">★★★★★</div>
        </div>

        <p class="text-gray-700 font-medium">оценок: {{ app.reviews || 2000 }}</p>

        <!-- Пример отзыва -->
        <div class="mt-4 p-4 bg-gray-50 rounded-xl">
          <div class="flex items-center gap-2 mb-2">
            <div class="w-8 h-8 bg-blue-500 rounded-full flex items-center justify-center text-white text-sm">А</div>
            <span class="font-semibold text-gray-900">Алексей</span>
          </div>
          <div class="text-yellow-500 text-sm mb-1">★★★★★</div>
          <p class="text-gray-800 text-sm">Отличное приложение! Все работает быстро и без багов.</p>
        </div>
      </div>

    </div>

    <!-- Сообщение об ошибке -->
    <div v-else class="text-center py-12">
      <div class="text-gray-400 text-4xl mb-3">😕</div>
      <h3 class="text-gray-800 font-semibold text-lg mb-2">Приложение не найдено</h3>
      <p class="text-gray-700">Попробуйте выбрать другое приложение</p>
    </div>

  </div>
</template>

<script>
export default {
  data() {
    return {
      app: null,
      isLoading: false
    }
  },

  mounted() {
    this.fetchAppData()
  },

  methods: {
    async fetchAppData() {
      const id = this.$route.params.id
      console.log('Loading app with ID:', id)

      if (!id) {
        console.error('No app ID provided')
        return
      }

      this.isLoading = true

      try {
        // ИСПОЛЬЗУЕМ ТОЛЬКО ЗАГЛУШКИ - НЕ ДЕЛАЕМ РЕАЛЬНЫЕ ЗАПРОСЫ
        await new Promise(resolve => setTimeout(resolve, 500)) // Имитация загрузки

        this.app = this.getMockAppData(id)
        console.log('App data loaded:', this.app)

      } catch (error) {
        console.error('Error loading app:', error)
        this.app = this.getMockAppData(id)
      } finally {
        this.isLoading = false
      }
    },

    getMockAppData(id) {
      const mockApps = {
        '1': {
          app_id: 1,
          app_title: "Tinkoff",
          app_name: "Tinkoff",
          app_category: "Банки",
          app_description: "Мобильное приложение Тинькофф Банка с удобным интерфейсом для управления финансами, платежами и инвестициями.",
          app_size: 145.8,
          app_age_rating: 12,
          app_developer_id: 5,
          app_count_downloads: 15000000,
          app_version: "5.24.1",
          app_link_apk: "disk.yandex.ru/client/d/tinkoff_apk",
          app_link_large_icon: "api.dicebear.com/7.x/icons/svg?seed=tinkoff&scale=90&size=200&radius=20&backgroundColor=ff4444",
          app_created_at: "15 марта 2015",
          rating: 4.8,
          reviews: 450000
        },
        '2': {
          app_id: 2,
          app_title: "VK",
          app_name: "VK",
          app_category: "Социальные сети",
          app_description: "Самая популярная социальная сеть в России с мессенджером, группами, видео и музыкой.",
          app_size: 89.3,
          app_age_rating: 16,
          app_developer_id: 3,
          app_count_downloads: 25000000,
          app_version: "8.65.1",
          app_link_apk: "disk.yandex.ru/client/d/vk_apk",
          app_link_large_icon: "api.dicebear.com/7.x/icons/svg?seed=vk&scale=90&size=200&radius=20&backgroundColor=4c75a3",
          app_created_at: "10 октября 2008",
          rating: 4.3,
          reviews: 3200000
        },
        '3': {
          app_id: 3,
          app_title: "YouTube",
          app_name: "YouTube",
          app_category: "Видео",
          app_description: "Крупнейшая в мире платформа для просмотра и загрузки видео. Миллионы роликов на любой вкус.",
          app_size: 67.2,
          app_age_rating: 16,
          app_developer_id: 1,
          app_count_downloads: 5000000000,
          app_version: "18.45.43",
          app_link_apk: "disk.yandex.ru/client/d/youtube_apk",
          app_link_large_icon: "api.dicebear.com/7.x/icons/svg?seed=youtube&scale=90&size=200&radius=20&backgroundColor=ff0000",
          app_created_at: "14 февраля 2005",
          rating: 4.4,
          reviews: 150000000
        },
        '101': {
          app_id: 101,
          app_title: "Clash Royale",
          app_name: "Clash Royale",
          app_category: "Игры",
          app_description: "Новая мега топовая игра от нашумевших разработчиков Clash of Clans уже в сети. Скачивай по моей ссылке в описании и упей урвать скин на реактивный истребитель и 20 тысяч серебра",
          app_size: 205.3,
          app_age_rating: 12,
          app_developer_id: 12,
          app_count_downloads: 1200000,
          app_version: "1.0.0",
          app_link_apk: "disk.yandex.ru/client/d/asf45sadf67as",
          app_link_large_icon: "api.dicebear.com/7.x/icons/svg?seed=clashroyale&scale=90&size=200&radius=20&backgroundColor=b6e3f4",
          app_link_screenshots: "disk.yandex.ru/client/d/asf45sadf67as",
          app_created_at: "1 ноября 2011",
          rating: 4.7
        },
        '102': {
          app_id: 102,
          app_title: "Clash of Clans",
          app_name: "Clash of Clans",
          app_category: "Игры",
          app_description: "Стратегическая игра, где вы строите деревню, тренируете армию и сражаетесь с другими игроками.",
          app_size: 180.5,
          app_age_rating: 10,
          app_developer_id: 12,
          app_count_downloads: 2500000,
          app_version: "15.83.4",
          app_link_large_icon: "api.dicebear.com/7.x/icons/svg?seed=clashofclans&scale=90&size=200&radius=20&backgroundColor=ffdfbf",
          rating: 4.8
        },
        '201': {
          app_id: 201,
          app_title: "Minecraft",
          app_name: "Minecraft",
          app_category: "Игры",
          app_description: "Исследуйте бесконечные миры, стройте от простых домов до величественных замков.",
          app_size: 350.2,
          app_age_rating: 7,
          app_developer_id: 8,
          app_count_downloads: 3000000,
          app_version: "1.20.1",
          app_link_large_icon: "api.dicebear.com/7.x/icons/svg?seed=minecraft&scale=90&size=200&radius=20&backgroundColor=c0aede",
          rating: 4.9
        }
      }

      return mockApps[id] || this.createFallbackApp(id)
    },

    createFallbackApp(id) {
      return {
        app_id: parseInt(id),
        app_title: `Приложение #${id}`,
        app_name: `App ${id}`,
        app_category: "Другое",
        app_description: "Описание приложения временно недоступно.",
        app_size: 100.0,
        app_age_rating: 12,
        app_developer_id: 0,
        app_count_downloads: 1000,
        app_version: "1.0.0",
        app_link_large_icon: `https://api.dicebear.com/7.x/icons/svg?seed=app${id}&scale=90&size=200&radius=20&backgroundColor=b6e3f4`,
        app_created_at: "2024",
        rating: 4.0,
        reviews: 100
      }
    },

    downloadApp() {
      if (this.app.app_link_apk) {
        console.log("Скачивание приложения:", this.app.app_title || this.app.app_name)
        alert(`Начинается скачивание: ${this.app.app_title || this.app.app_name}`)
      } else {
        alert("Ссылка для скачивания недоступна")
      }
    },

    formatSize(size) {
      if (!size) return "200 МБ"
      return typeof size === 'number' ? `${size} МБ` : size
    },

    formatDate(date) {
      if (!date) return "1 ноября 2011"
      return date
    }
  },

  watch: {
    '$route.params.id': {
      handler(newId) {
        if (newId) {
          this.fetchAppData()
        }
      },
      immediate: true
    }
  }
}
</script>