<script>
import SearchComponent from '../components/SearchComponent.vue'
import { RouterLink, RouterView } from 'vue-router'

export default {
  components: {
    SearchComponent,

  },
  methods: {
    goBackFunction() {
      if (this.$router) {
        this.$router.back()
      }
    }

  }, computed: {
    currentRouteName() {
      return this.$route.path
    }
  }
}

</script>

<template>
  <main>
    <header>


      <nav class="bg-blue-400 px-4 py-2 flex gap-2 items-center">
        <RouterLink to="/" class="text-xl font-sans px-4 py-2" active-class="activeLink">Flight</RouterLink>
        <RouterLink aria-disabled="true" to="/" class="text-xl font-sans px-4 py-2  cursor-not-allowed">Hotel
        </RouterLink>
      </nav>

      <button v-if="currentRouteName !== '/'" @click="goBackFunction"
        class="bg-blue-600 px-4 py-2 rounded-md text-white ml-5 my-4">Back</button>

    </header>
    <SearchComponent />
  </main>
</template>

<style>
.activeLink {
  background-color: white;
  border-radius: 10px;


}
</style>
