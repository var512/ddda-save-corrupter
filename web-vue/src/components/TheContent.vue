<template>
  <main class="col-md-8 col-lg-9 col-xl-9 ms-sm-auto px-md-4">
    <div id="content">
      <router-view :key="$route.path" />
    </div>
  </main>
</template>
