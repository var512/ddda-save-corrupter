<template>
  <h2>Export file</h2>
  <form class="row">
    <div class="col-12">
      Generate and download a modified <span class="text-black-50">.sav</span> or <span
        class="text-black-50"
      >.xml</span>
    </div>
    <div class="col-12 mt-2">
      <a
        :href="BACKEND_URL + '/files/sav/export'"
        class="btn btn-primary me-2"
        target="_blank"
        rel="noreferrer"
        @click="showWarning"
      >
        <font-awesome-icon icon="download" /> Export .sav
      </a>
      <a
        :href="BACKEND_URL + '/files/xml/export'"
        class="btn btn-primary"
        target="_blank"
        rel="noreferrer"
        @click="showWarning"
      >
        <font-awesome-icon icon="download" /> Export .xml
      </a>
    </div>
  </form>
</template>

<script>
import { BACKEND_URL } from '@/constants/app';

export default {
  data() {
    return {
      BACKEND_URL,
    };
  },
  methods: {
    // needs a server side api and a hook
    // golang broken pipe
    showWarning() {
      this.$swal('Generating data', 'Download your file before continuing...');
    },
  },
};
</script>
