<script>
import { computed } from 'vue'
import { useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'
import { provide } from 'vue'
import { InMemoryCache } from 'apollo-cache-inmemory'
import ApolloClient from 'apollo-boost'
import { DefaultApolloClient } from '@vue/apollo-composable'

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
         rooms: allRoomsQuery
     },
     data() {
         return {
             rooms: Array
        }
    },
    // setup() {
    //     provide(DefaultApolloClient, apolloClient)
    //     const { result } = useQuery(allRoomsQuery);
    //     rooms = computed(() => result.value?.rooms ?? ["error"]);
    // }
    // computed: {
    //     rooms: rooms
    // },
    // methods: {
    //     joinRoom(roomId) {
    //         this.$router.push({
    //             name: 'room',
    //             params: {
    //                 roomId: roomId
    //             }
    //         })
    //     }
    // }
 }
</script>
<template>
    <ul>
        <li v-for="room in rooms" :key="room.id">
            {{ room.id }}
        </li>
    </ul>
</template>