<template>
  <div
      class="w-14 h-14 rounded-2xl flex items-center justify-center text-3xl font-bold"
      :style="{ background: bgColor }"
  >
    {{ emoji }}
  </div>
</template>

<script>
export default {
  props: {
    emoji: { type: String, default: "🤖" }
  },
  computed: {
    bgColor() {
      const colors = ["#FFE082", "#FFAB91", "#80DEEA", "#CF93D9", "#A5D6A7"];
      return colors[Math.floor(Math.random() * colors.length)];
    }
  }
}
</script>