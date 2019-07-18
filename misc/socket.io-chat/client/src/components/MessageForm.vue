<template>
  <v-layout row>
    <v-flex grow>
      <v-text-field
        solo
        hide-details
        single-line
        v-model="message"
      />
    </v-flex>
    <v-flex shrink>
      <v-btn color="primary" @click="send">
        SEND
      </v-btn>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {

    data: () => ({
      message: ''
    }),

    methods: {
      send() {
        this.$socket.emit('message', this.message);
        this.message = '';
      }
    }
  }
</script>

<style scoped>

</style>