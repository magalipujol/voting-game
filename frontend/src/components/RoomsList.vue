<script>
import gql from 'graphql-tag'
import { InMemoryCache } from 'apollo-cache-inmemory'
import ApolloClient from 'apollo-boost'

const cache = new InMemoryCache()


const apolloClient = new ApolloClient({
    uri: 'http://localhost:8080/graphql',
    cache,
})

// create and excecute a query
const allRoomsQuery = gql`
  query {
      rooms {
          id
      players {
          name
      }
      voting
    }
  }
`;

export default {
    apollo: {
        rooms: {
            query: allRoomsQuery,
            error() {
                this.error = true
            }
        }
    },
    data() {
        return {
            rooms: Array,
            error: false,
        }
    },
}
</script>
<template>
    <p v-if="$apollo.queries.rooms.loading">Loading...</p>
    <ul>
        <li v-for="room in rooms" :key="room.id">
            {{ room.id }}
        </li>
    </ul>
</template>