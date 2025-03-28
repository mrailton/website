<script setup lang="ts">
const slug = useRoute().params.slug

const {data: article} = await useAsyncData(`blog-${slug}`, () => {
  return queryCollection('blog').path(`/blog/${slug}`).first()
});

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
