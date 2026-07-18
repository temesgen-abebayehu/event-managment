import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from '@apollo/client/core'
import { provideApolloClient } from '@vue/apollo-composable'

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig()
  
  const httpLink = new HttpLink({
    uri: config.public.graphqlEndpoint,
  })

  const authLink = new ApolloLink((operation, forward) => {
    const token = useCookie('auth_token').value
    if (token) {
      operation.setContext({
        headers: {
          authorization: `Bearer ${token}`,
        },
      })
    }
    return forward(operation)
  })

  const apolloClient = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: new InMemoryCache(),
  })

  provideApolloClient(apolloClient)
  
  return {
    provide: {
      apolloClient,
    },
  }
})
