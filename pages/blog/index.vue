<script setup lang="ts">
import ArticleCard from "~/components/article-card.vue";

const { data: articles } = useAsyncData('blog-articles', async () => {
  const posts = await queryCollection('blog').all();

  return posts.sort((a, b) => new Date(b.date) - new Date(a.date));
});

console.log(articles.value);
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