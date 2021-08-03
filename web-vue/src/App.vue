<template>
  <div
    v-if="$store.state.isLoading === true"
    class="d-flex justify-content-center app__spinner"
  >
    <div class="spinner-border" />
  </div>
  <div
    v-show="$store.state.isLoading === false"
    class="container-fluid"
  >
    <div class="row">
      <the-sidebar />
      <the-content />
    </div>
  </div>
</template>

<script>
import TheSidebar from './components/TheSidebar.vue';
import TheContent from './components/TheContent.vue';

export default {
  components: {
    TheSidebar,
    TheContent,
  },
};
</script>

<style>
.app__spinner {
  margin-top: 2rem;
}
</style>
