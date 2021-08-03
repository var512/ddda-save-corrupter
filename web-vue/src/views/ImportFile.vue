<template>
  <h2>Import file</h2>
  <form
    class="row"
    @submit.prevent="onSubmit"
  >
    <div class="col-12">
      <label
        for="userfile"
        class="form-label"
      >
        Select a <span class="text-black-50">.sav</span> file
      </label>
    </div>
    <div class="col-sm-12 col-md-10 col-lg-8">
      <input
        id="userfile"
        type="file"
        class="form-control-file"
        name="userfile"
        @change="onChange"
      >
    </div>
    <div class="col-12 mt-3">
      <button
        type="submit"
        class="btn btn-primary"
      >
        <font-awesome-icon icon="file" /> Import
      </button>
    </div>
    <div class="col-12 mt-3">
      <p class="text-primary small">
        <b>Warning:</b> backup your .sav files before using this tool
      </p>
    </div>
  </form>
</template>

<script>
import axios from 'axios';
import GetAxiosErrorMessage from '@/helpers/GetAxiosErrorMessage';

export default {
  data() {
    return {
      file: null,
    };
  },
  methods: {
    onChange(e) {
      if (e.target.files.length !== 1) {
        this.$swal('Error', 'e.target.files.length !== 1');
        return;
      }
      this.file = e.target.files;
    },
    onSubmit() {
      if (this.file === null) {
        this.$swal('Error', 'Select a .sav file');
        return;
      }
      this.$store.commit('isLoading', true);
      const formData = new FormData();
      formData.append('userfile', this.file[0]);
      this.postFile(formData);
    },
    async postFile(formData) {
      try {
        const r = await axios.post('files', formData, {
          headers: {
            'content-type': 'multipart/form-data',
          },
        });

        if (r.status === 200) {
          this.$store.commit('hasUserFile', true);
          this.$swal(r.data);
        }
      } catch (err) {
        this.$swal('Error', GetAxiosErrorMessage(err));
      } finally {
        this.$store.commit('isLoading', false);
      }
    },
  },
};
</script>
