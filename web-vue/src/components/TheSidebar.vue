<template>
  <nav
    id="sidebar-menu"
    class="col-md-4 col-lg-3 col-xl-3 d-md-block sidebar"
  >
    <div class="sidebar-brand">
      <span>DDDA Save Corrupter</span>
    </div>
    <ul class="nav nav-tabs flex-column">
      <li class="nav-item">
        <router-link
          :to="{ name: 'import-file' }"
          class="nav-link"
        >
          <font-awesome-icon icon="file" />
          Import file
        </router-link>
      </li>
      <li v-show="$store.state.hasUserFile === true">
        <router-link
          :to="{ name: 'pawns', params: { category: 'main'}}"
          class="nav-link"
        >
          <font-awesome-icon icon="chess-knight" />
          Main Pawn
        </router-link>
      </li>
      <li v-show="$store.state.hasUserFile === true">
        <router-link
          :to="{ name: 'pawns', params: { category: 'first'}}"
          class="nav-link"
        >
          <font-awesome-icon icon="chess-pawn" />
          First Pawn
        </router-link>
      </li>
      <li v-show="$store.state.hasUserFile === true">
        <router-link
          :to="{ name: 'pawns', params: { category: 'second'}}"
          class="nav-link"
        >
          <font-awesome-icon icon="chess-pawn" />
          Second Pawn
        </router-link>
      </li>
      <li v-show="$store.state.hasUserFile === true">
        <router-link
          :to="{ name: 'export-file' }"
          class="nav-link"
        >
          <font-awesome-icon icon="download" />
          Export file
        </router-link>
      </li>
    </ul>
  </nav>
</template>
