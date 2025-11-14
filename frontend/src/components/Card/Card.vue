<template>
  <div class="bg-white rounded-xl border border-gray-200 shadow-sm p-4
              flex items-center gap-4 hover:shadow-md transition cursor-pointer">

    <CardImage :name="app.app_name" />

    <div class="flex flex-col">
      <CardTitle :text="app.app_name" />
      <CardCategory :category="app.app_categorie" />
    </div>

  </div>
</template>

<script>
import CardTitle from "./CardTitle.vue";
import CardCategory from "./CardCategory.vue";
import CardImage from "./CardImage.vue";

export default {
  props: { app: Object },
  components: { CardTitle, CardCategory, CardImage }
};
</script>
