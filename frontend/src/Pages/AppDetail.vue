<template>
  <div class="w-full min-h-screen bg-white p-4">

    <!-- HEADER -->
    <div class="flex justify-between items-center mb-4">
      <button @click="$router.back()" class="text-3xl">⬅️</button>
      <button class="text-3xl">↗️</button>
    </div>

    <!-- TOP BLOCK -->
    <div class="flex gap-4">

      <!-- ICON -->
      <div class="w-24 h-24 rounded-3xl bg-gray-200 flex justify-center items-center text-5xl">
        {{ app.emoji || "🤖" }}
      </div>

      <!-- TITLE -->
      <div class="flex flex-col justify-center">
        <h1 class="text-2xl font-bold">{{ app.app_name }}</h1>
        <p class="text-gray-500 text-lg">{{ app.app_category }}</p>
      </div>

    </div>

    <!-- INFO ROW -->
    <div class="flex justify-between mt-6 text-sm text-gray-700">
      <div class="text-center">
        <div class="font-semibold text-lg">{{ app.rating || "4.5" }} ⭐</div>
        <div>отзывов: {{ app.reviews || 1200 }}</div>
      </div>

      <div class="text-center">
        <div class="font-semibold text-lg">{{ app.size || "200 МБ" }}</div>
        <div>размер</div>
      </div>

      <div class="text-center">
        <div class="font-semibold text-lg">{{ app.age || "3+" }}</div>
        <div>возраст</div>
      </div>
    </div>

    <!-- GALLERY -->
    <div class="flex gap-3 mt-8 overflow-x-auto pb-2">
      <div
          v-for="n in 3"
          :key="n"
          class="w-40 h-56 rounded-xl bg-gray-200 flex items-center justify-center text-4xl shrink-0"
      >
        🖼️
      </div>
    </div>

    <!-- DESCRIPTION -->
    <div class="mt-8">
      <h2 class="text-xl font-bold">Описание</h2>
      <p class="text-gray-600 mt-2 leading-relaxed">
        {{ app.description || "Описание приложения будет здесь..." }}
      </p>
    </div>

    <!-- REVIEWS -->
    <div class="mt-8 mb-20">
      <h2 class="text-xl font-bold mb-2">Отзывы</h2>

      <div class="flex items-center gap-4">
        <div class="text-4xl font-bold">{{ app.rating || "4.5" }}</div>
        <div class="text-yellow-500 text-xl">★★★★★</div>
      </div>

      <p class="text-gray-500 mt-1">оценок: {{ app.reviews || 2000 }}</p>
    </div>

  </div>
</template>

<script>
import json from "@/data/apps.json"

export default {
  data() {
    return {
      app: {}
    }
  },

  mounted() {
    const id = this.$route.params.id
    this.app = this.findAppById(id)
  },

  methods: {
    findAppById(id) {
      for (const category of json.categories) {
        const found = category.apps.find(a => a.app_id == id)
        if (found) return found
      }
      return {}
    }
  }
}
</script>