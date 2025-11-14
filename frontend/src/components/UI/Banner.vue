<template>
  <div
      class="w-full h-40 rounded-3xl bg-gradient-to-r from-indigo-500 to-blue-500
           flex items-center justify-center shadow-md mt-6"
  >
    <p class="text-white text-lg font-semibold">{{ text }}</p>
  </div>
</template>

<script>
export default {
  props: ["text"]
};
</script>
