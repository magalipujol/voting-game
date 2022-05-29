import { createApp, h, provide } from 'vue'
import App from './App.vue'
import { InMemoryCache } from 'apollo-cache-inmemory'
import ApolloClient from 'apollo-boost'
import { DefaultApolloClient } from '@vue/apollo-composable'
import { createApolloProvider } from '@vue/apollo-option'

const cache = new InMemoryCache()


const apolloClient = new ApolloClient({
    uri: 'http://localhost:8080/graphql',
    cache,
})

const apolloProvider = createApolloProvider({
    defaultClient: apolloClient,
})

const app = createApp({
    render: () => h(App)
})

app.use(apolloProvider)
app.mount('#app')
