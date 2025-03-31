// https://nuxt.com/docs/api/configuration/nuxt-config
import tailwindcss from "@tailwindcss/vite";

export default defineNuxtConfig({
    compatibilityDate: '2025-03-29',
    future: {
        compatibilityVersion: 4,
    },
    devtools: {
        enabled: true,
    },
    css: [
        '~/assets/css/main.css',
    ],
    vite: {
        plugins: [
            tailwindcss(),
        ],
    },
    modules: [
        '@nuxt/content',
    ],
    content: {
        build: {
            markdown: {
                highlight: {
                    theme: {
                        default: 'dracula'
                    }
                }
            }
        }
    }
});
