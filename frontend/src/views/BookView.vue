<template>
    <!-- <FlightDetail :flight="selectedBookFlights" /> -->
    <div class="w-full bg-blue-100 py-8 text-center rounded-md mr-5 ">

        <p class="font-bold text-3xl">Enter Passenger Details</p>
    </div>
    <PassengerForm />
</template>

<script>
import { useFlightStore } from '@/stores/flight';
import { mapState } from 'pinia';
import FlightDetail from '@/components/FlightDetail.vue';
import PassengerForm from '@/components/PassengerForm.vue';

export default {
    name: "BookView",
    components: {
        // FlightDetail,
        PassengerForm
    },
    computed: {
        ...mapState(useFlightStore, [
            "selectedBookFlights"
        ])
    },
    mounted() {
        console.log(this.selectedBookFlights);

    }
}
</script>

<style scoped></style>