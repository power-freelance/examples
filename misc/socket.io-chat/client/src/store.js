import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    messages: []
  },
  mutations: {
    pushMessage: (state, message) => state.messages.push(message)
  },
  actions: {
    'SOCKET_message': ({commit}, message) => {
      commit('pushMessage', message);
    }
  }
})
