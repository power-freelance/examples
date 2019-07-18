<template>
  <v-list>
    <v-list-tile v-for="(message, index) in messages" :key="index">
      <v-list-tile-title>
        {{message}}
      </v-list-tile-title>
    </v-list-tile>
  </v-list>
</template>

<script>
  import {mapState} from 'vuex'

  export default {

    computed: {
      ...mapState({
        messages: state => state.messages
      })
    },

  }
</script>
