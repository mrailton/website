<script setup lang="ts">
defineProps(['article']);
const {formatDate} = useDateFormatter();
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