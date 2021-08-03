<template>
  <div v-if="pawn === null">
    <h2>Pawn ({{ category }})</h2>
    <p class="text-black-50">
      Pawn data is empty.
    </p>
  </div>
  <pawn-information
    :pawn-data="pawn"
    :category="category"
  />
</template>

<script>
import { BACKEND_URL } from '@/constants/app';
import axios from 'axios';
import GetAxiosErrorMessage from '@/helpers/GetAxiosErrorMessage';
import PawnInformation from '@/views/PawnInformation.vue';

export default {
  components: { PawnInformation },
  data() {
    return {
      BACKEND_URL,
      pawn: {},
    };
  },
  computed: {
    category() {
      return this.$route.params.category;
    },
  },
  created() {
    this.loadPawn();
  },
  methods: {
    async loadPawn() {
      try {
        this.$store.commit('isLoading', true);
        const r = await axios.get(`pawns/${this.category}`);
        if (r.status === 200) {
          this.pawn = r.data.pawn;
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
