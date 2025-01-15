<template>
    <!-- <FlightDetail :flight="selectedBookFlights" /> -->
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