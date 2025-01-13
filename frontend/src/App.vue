<script setup>
import { RouterLink, RouterView } from 'vue-router'
// import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <header>


    <nav class="bg-blue-400 px-4 py-2 flex gap-2 items-center">
      <RouterLink to="/" class="text-xl font-sans px-4 py-2" active-class="activeLink">Flight</RouterLink>
      <RouterLink aria-disabled="true" to="/" class="text-xl font-sans px-4 py-2  cursor-not-allowed">Hotel
      </RouterLink>
    </nav>

  </header>

  <RouterView />
</template>

<style scoped>
.activeLink {
  background-color: white;
  border-radius: 10px;


}
</style>
