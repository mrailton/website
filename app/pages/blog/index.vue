<script setup lang="ts">
const articles = useGetArticles();
</script>

<template>
  <section>
    <div class="flex flex-col items-center justify-center">
      <h1 class="text-3xl font-semibold text-gray-800 mb-8">Blog Articles</h1>
      <ul>
        <li v-for="article in articles" :key="article.stem">
          <ArticleCard :article="article" />
        </li>
      </ul>
    </div>
  </section>
</template>

<style scoped>

</style>