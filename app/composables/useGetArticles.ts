export const useGetArticles = (limit?: number) => {
    const {data} = useAsyncData('blog-articles', async () => {
        const posts = await queryCollection('blog').all();

        const sortedPosts = posts.sort((a, b) => new Date(b.date).getTime() - new Date(a.date).getTime());
        
        return limit ? sortedPosts.slice(0, limit) : sortedPosts;
    });

    return data;
};