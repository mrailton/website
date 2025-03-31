<script setup lang="ts">
const slug = useRoute().params.slug

const {data: article} = await useAsyncData(`blog-${slug}`, () => {
  return queryCollection('blog').path(`/blog/${slug}`).first()
});

const { formatDate } = useDateFormatter();

</script>

<template>
  <div class="max-w-4xl px-6 pb-20 mx-auto">
    <div v-if="article">
      <h1 class="text-3xl font-semibold text-gray-800 mb-4">{{ article.title }}</h1>
      <span class="block text-gray-600 font-light text-sm mb-8">Posted: {{ formatDate(article.date) }}</span>
      <div class="prose lg:prose-xl">
        <ContentRenderer :value="article"/>
      </div>
    </div>
    <div v-else>
      Content Not Found
    </div>
  </div>
</template>
