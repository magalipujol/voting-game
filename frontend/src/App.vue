<script setup>

import { computed, watchEffect } from 'vue'
import { useMutation, useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'

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
const { result } = useQuery(allRoomsQuery);
const rooms = computed(() => result.value?.rooms ?? ["error"]);



</script>

<template>
  <ul>
    <li v-for="room in rooms" :key="room.id">
      {{ room.id }}
    </li>
  </ul>
</template>

<style>
</style>
