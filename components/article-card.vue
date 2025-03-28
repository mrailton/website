<script setup lang="ts">
defineProps(['article']);

const formatDate = (dateString) => {
  const date = new Date(dateString);
  return new Intl.DateTimeFormat("en-GB", {
    day: "numeric",
    month: "long",
    year: "numeric",
  }).format(date).replace(/\b(\d{1,2})\b/, (d) => `${d}${getOrdinal(d)}`);
};

const getOrdinal = (day) => {
  const suffixes = ["th", "st", "nd", "rd"];
  const v = day % 100;
  return suffixes[(v - 20) % 10] || suffixes[v] || suffixes[0];
};

console.log('article')
</script>

<template>
  <NuxtLink class="mb-8 max-w-3xl w-full block bg-white shadow-md border-t-4 border-indigo-600" :to="article.stem">
    <div class="flex items-center justify-between px-4 py-2">
      <h3 class="text-lg font-medium text-gray-700">
        {{ article.title }}
      </h3>

      <span class="block text-gray-600 font-light text-sm">
            {{ formatDate(article.date) }}
        </span>
    </div>
  </NuxtLink>
</template>

<style scoped>

</style>