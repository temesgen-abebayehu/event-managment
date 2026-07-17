// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },

  modules: ['@nuxtjs/tailwindcss'],

  runtimeConfig: {
    public: {
      backendUrl: process.env.BACKEND_URL || 'http://localhost:3001',
      hasuraUrl: process.env.HASURA_GRAPHQL_URL || 'http://localhost:8080/v1/graphql',
    },
  },

  app: {
    pageTransition: { name: 'page', mode: 'out-in' },
    layoutTransition: { name: 'layout', mode: 'out-in' },
    head: {
      title: 'EventHub Ethiopia',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        {
          name: 'description',
          content: 'Discover and manage amazing events in Ethiopia',
        },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
      ],
    },
  },

  css: ['~/assets/css/main.css'],

  build: {
    transpile: ['@vue/apollo-composable', '@apollo/client'],
  },
})
