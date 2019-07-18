<template>
  <v-layout column fill-height>
    <v-flex grow>
      <MessageList/>
    </v-flex>
    <v-flex shrink>
      <MessageForm/>
    </v-flex>
  </v-layout>
</template>

<script>

  import MessageList from "../components/MessageList";
  import MessageForm from "../components/MessageForm";

  export default {

    components: {
      MessageForm,
      MessageList
    }

  }
</script>
