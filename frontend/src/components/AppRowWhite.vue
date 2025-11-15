<template>
  <div class="flex items-center py-3 w-full active:opacity-70" @click="openApp">

    <!-- ICON -->
    <div class="w-14 h-14 overflow-hidden rounded-2xl bg-gray-200 flex-shrink-0">
      <img :src="iconUrl" class="w-full h-full object-cover" />
    </div>

    <!-- TEXT -->
    <div class="flex flex-col flex-grow px-3 overflow-hidden">
      <h3 class="text-base font-semibold text-gray-900 truncate">
        {{ app.app_name }}
      </h3>
      <p class="text-sm text-gray-500 truncate">
        {{ app.app_category }}
      </p>
    </div>

    <!-- BUTTON -->
    <button
        @click.stop="download"
        class="px-4 py-1.5 rounded-full bg-blue-500 text-white text-sm font-semibold whitespace-nowrap hover:bg-blue-600 active:scale-95 transition"
    >
      Загрузить
    </button>
  </div>
</template>

<script>
export default {
  props: { app: Object },

  computed: {
    iconUrl() {
      return `https://api.dicebear.com/7.x/icons/svg?seed=${encodeURIComponent(this.app.app_name)}&scale=90&size=200&radius=20&backgroundColor=b6e3f4,ffdfbf,c0aede,d1d4f9`;
    }
  },

  methods: {
    openApp() {
      this.$router.push({
        name: "app-detail",
        params: { id: this.app.app_id }
      });
    },
    download() {
      console.log("Download:", this.app.app_name);
    }
  }
}
</script>