<template>
  <!-- TODO category === 'main' ? <AppearanceOverride data={{ pawn: props.data.pawn }} /> : null -->
  <div id="pawn-information">
    <h2>Pawn information</h2>
    <table class="table table-striped table-bordered">
      <tbody>
        <tr>
          <th>
            Actions
          </th>
          <td>
            <!-- TODO /pawns/category/replace -->
            <a
              v-if="false"
              :href="BACKEND_URL + '/pawns/' + category + '/replace'"
              class="btn btn-primary btn-sm me-2"
              target="_blank"
              rel="noreferrer"
            >
              <font-awesome-icon icon="people-arrows" />
              Replace
            </a>
            <a
              :href="BACKEND_URL + '/pawns/' + category + '/export'"
              class="btn btn-primary btn-sm"
              target="_blank"
              rel="noreferrer"
            >
              <font-awesome-icon icon="download" />
              Export .xml
            </a>
          </td>
        </tr>
        <tr>
          <th>
            Name
          </th>
          <td>
            {{ pawn?.data?.Attributes?.Name }}
          </td>
        </tr>
        <tr>
          <th>
            Gender
          </th>
          <td>
            {{ pawn?.data?.Attributes?.Gender?.Value === 0 ? 'Male' : 'Female' }}
            ({{ pawn?.data?.Attributes?.Gender?.Value }})
          </td>
        </tr>
        <tr>
          <th>
            Nickname
          </th>
          <td>
            {{ pawn?.data?.Attributes?.Nickname?.Value }}
          </td>
        </tr>
      </tbody>
    </table>

    <pre class="mt-3">
      <code class="language-css">
        {{ pawn?.dataXML }}
      </code>
    </pre>
  </div>
</template>

<script>
import { BACKEND_URL } from '@/constants/app';

export default {
  props: {
    pawnData: {
      type: Object,
      required: true,
      default() {
        return {};
      },
    },
  },
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
  watch: {
    pawnData(p) {
      this.pawn = p;
    },
  },
};
</script>
