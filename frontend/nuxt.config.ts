import { defineNuxtConfig } from "nuxt/config";

// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-04-03',
  devtools: { enabled: true },
  ssr: false,

  modules: ['@nuxtjs/tailwindcss'],

  runtimeConfig: {
    public: {
      backendUrl: 'http://localhost:3001',
      hasuraUrl: 'http://localhost:8080/v1/graphql',
      graphqlEndpoint: 'http://localhost:8080/v1/graphql',
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
    transpile: ['@vue/apollo-composable', '@apollo/client', 'tslib'],
  },

  nitro: {
    externals: {
      inline: ['tslib'],
    },
  },
})
