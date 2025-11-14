<template>
  <div
      class="w-full h-40 rounded-2xl bg-gradient-to-r from-indigo-500 to-blue-500
           flex items-center justify-center text-center shadow-lg mb-8 px-4"
  >
    <p class="text-white text-xl font-semibold leading-snug">
      {{ text }}
    </p>
  </div>
</template>

<script>
export default {
  props: ["text"]
};
</script>
