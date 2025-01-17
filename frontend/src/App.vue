<script>
import { RouterLink, RouterView } from 'vue-router'
export default {

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
// import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <!-- <button v-if="currentRouteName !== '/'" @click="goBackFunction"
    class="bg-blue-600 px-4 py-2 rounded-md text-white ml-5 my-4">Back</button> -->

  <RouterView />
</template>

<style scoped>
.activeLink {
  background-color: white;
  border-radius: 10px;


}
</style>
