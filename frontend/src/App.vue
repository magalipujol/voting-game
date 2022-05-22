<script setup>

import { watch } from 'vue'
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

 watch(result, value => {
   console.log("Mirando")
      console.log(value)
    })

</script>

<template>
</template>

<style>

</style>
