<template>
  <!-- <FlightDetail :flight="selectedBookFlights" /> -->
  <div class="w-full bg-blue-100 py-8 text-center rounded-md mr-5">
    <button
      @click="goBack"
      class="absolute left-4 top-0 p-2 text-gray-600 hover:text-gray-900 transition-colors"
    >
      &lt; Back
    </button>
    <p class="font-bold text-3xl">Enter Passenger Details</p>
  </div>
  <PassengerForm />
</template>

<script>
import { useFlightStore } from "@/stores/flight";
import { mapState } from "pinia";
import FlightDetail from "@/components/FlightDetail.vue";
import PassengerForm from "@/components/PassengerForm.vue";

export default {
  name: "BookView",
  components: {
    // FlightDetail,
    PassengerForm,
  },
  computed: {
    ...mapState(useFlightStore, ["selectedBookFlights"]),
  },
  methods: {
    goBack() {
      this.$router.go(-1);
    },
  },
  mounted() {
    console.log(this.selectedBookFlights);
  },
};
</script>

<style scoped></style>
