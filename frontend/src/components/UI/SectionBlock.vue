<template>
  <section class="mb-10">
    <h2 class="font-semibold text-xl mb-4 text-gray-900">
      {{ title }}
    </h2>
    <slot />
  </section>
</template>

<script>
export default { props: ['title'] }
</script>
