import { ApolloClient, InMemoryCache, HttpLink, ApolloLink } from '@apollo/client/core'
import { provideApolloClient } from '@vue/apollo-composable'

export default defineNuxtPlugin((nuxtApp) => {
  const config = useRuntimeConfig()
  
  const httpLink = new HttpLink({
    uri: config.public.graphqlEndpoint,
  })

  const authLink = new ApolloLink((operation, forward) => {
    const token = useCookie('auth_token').value
    
    // Check if operation explicitly requests anonymous role
    const context = operation.getContext()
    const forceAnonymous = context.forceAnonymous === true
    
    if (forceAnonymous) {
      // Explicitly use anonymous role (for public browsing)
      operation.setContext({
        headers: {
          'x-hasura-role': 'anonymous',
        },
      })
    } else if (token) {
      // Use JWT authentication for user-specific queries
      operation.setContext({
        headers: {
          authorization: `Bearer ${token}`,
        },
      })
    } else {
      // Default to anonymous when not logged in
      operation.setContext({
        headers: {
          'x-hasura-role': 'anonymous',
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
