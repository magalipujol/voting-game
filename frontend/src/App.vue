<script setup>

import { watch, computed } from 'vue'
import {useQuery} from '@vue/apollo-composable'
import gql from 'graphql-tag'
console.log("AAAAAAAAAAAAAAAAAAA")

const allRoomsQuery = gql`
  query {
    rooms {
      id
      players {
        name
      }
      votes
      voting
    }
  }
`;

const { result } = useQuery(allRoomsQuery);
console.log("BBBBBBBBBBBBBBBB")
const rooms = computed(() => result ?? {})
console.log(rooms)

 watch(rooms, value => {
   console.log("Mirando")
      console.log(value)
    })
console.log("CCCCCCCCCCCC")

</script>

<template>
</template>

<style>

</style>
