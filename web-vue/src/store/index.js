import { createStore } from 'vuex';

export default createStore({
  state: {
    isLoading: false,
    hasUserFile: false,
  },
  mutations: {
    isLoading(state, p) {
      state.isLoading = p;
    },
    hasUserFile(state, p) {
      state.hasUserFile = p;
    },
  },
  actions: {
  },
  modules: {
  },
});
