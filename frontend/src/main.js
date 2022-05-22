import { createApp, h, provide } from 'vue'
import App from './App.vue'
import { createHttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'
import ApolloClient from 'apollo-boost'
import { DefaultApolloClient } from '@vue/apollo-composable'


const cache = new InMemoryCache()

const apolloClient = new ApolloClient({
  uri: 'https://localhost:8080/',
  cache,
})


const app = createApp({
    setup() {
        provide(DefaultApolloClient, apolloClient)
    },
    render: () => h(App)
})

app.mount('#app')