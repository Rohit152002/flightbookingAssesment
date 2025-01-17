<template>
    <div
        class="w-full h-52 border border-gray-300 rounded-lg bg-white text-center flex justify-center gap-4 items-center p-4 ">
        <div v-if="selectedDepartureFlights === null">
            No Flight Available
        </div>
        <!-- Departure Flight Details -->
        <div v-else
            class="flex flex-col justify-start items-start border border-black rounded-md px-10 py-6 text-gray-800 space-y-2">
            <h2 class="text-lg font-semibold text-indigo-600">Departure</h2>
            <p class="text-md font-medium">{{ selectedDepartureFlights.Source }} → {{
                selectedDepartureFlights.Destination }}</p>
            <p class="text-sm text-gray-600">Price: ₹{{ selectedDepartureFlights.PriceSaver }} (Saver) / ₹{{
                selectedDepartureFlights.PriceFlexi }} (Flexi)</p>
        </div>

        <!-- Return Flight Details -->
        <div v-if="selectedReturnFlights === null">
            No Flights Available
        </div>
        <div v-else
            class="flex flex-col justify-center items-start border border-black rounded-md px-10 py-6 text-gray-800 space-y-2">
            <h2 class="text-lg font-semibold text-indigo-600">Return</h2>
            <p class="text-md font-medium">{{ selectedReturnFlights.Source }} → {{ selectedReturnFlights.Destination }}
            </p>
            <p class="text-sm text-gray-600">Price: ₹{{ selectedReturnFlights.PriceSaver }} (Saver) / ₹{{
                selectedReturnFlights.PriceFlexi }} (Flexi)</p>
        </div>

        <!-- Total Price Details -->
        <div
            class="flex flex-col h-fit border border-black rounded-md px-10 py-10 justify-center items-center text-gray-800 space-y-2">
            <h2 class="text-lg font-semibold text-indigo-600">Total Price</h2>
            <p class="text-sm text-gray-600">Price: ₹{{ selectedJourneyPrice + selectedReturnPrice }} </p>
        </div>
        <button
            class="mt-4 px-6 py-2 bg-indigo-600 text-white font-semibold rounded-lg shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
            Book Now
        </button>
    </div>


    <div class="flex w-full justify-center pb-60 pt-20 gap-5">
        <div>
            <h1 class="text-center text-xl px-4 py-2 underline">Departure</h1>
            <FlightSearch :flights="departureFlights" type="Departure" />
        </div>
        <div>
            <h1 class="text-center text-xl px-4 py-2 underline">Return</h1>
            <FlightSearch :flights="returnFlights" type="Return" />
        </div>
    </div>
</template>

<script>
import { mapState } from 'pinia';
import FlightSearch from './FlightSearch.vue';
import { useFlightStore } from '@/stores/flight';



export default {
    name: 'RoundTripComponent',
    data() {
        return {
            hello: "lkfjasdf"
        };
    },

    props: {
        departureFlights: {
            type: Array,
            required: true,
        },
        returnFlights: {
            type: Array,
            required: true
        }
    },

    emits: {},

    components: {
        FlightSearch
    },

    created() { },

    mounted() {

    },

    methods: {
        bookFlights(){
         this.$router.push('/book')
        }
    },

    computed: {
        ...mapState(useFlightStore, ['selectedDepartureFlights', 'selectedReturnFlights', 'selectedJourneyPrice', 'selectedReturnPrice'])
    },

    watch: {},

    filters: {},

    directives: {}
};
</script>

<style lang="scss" scoped></style>
