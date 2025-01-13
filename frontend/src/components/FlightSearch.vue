<template>
    <div v-if="flights && flights.length" class="space-y-8 my-6 mx-auto px-4 md:px-16 lg:px-32">

        <div v-for="flight in flights" :key="flight.ID"
            class="bg-white p-6 rounded-xl shadow-xl border border-gray-200 transition-transform duration-300 transform hover:scale-100">

            <div class="flex justify-between items-center mb-6">
                <div>
                    <h2 class="text-2xl font-semibold text-indigo-600">{{ flight.Airline }}</h2>
                    <p class="text-lg text-gray-500">{{ flight.Source }} &rarr; {{ flight.Destination }}</p>
                </div>
                <div class="flex gap-2 text-sm text-gray-400">
                    <span>{{ formatDate(flight.DepartureTime) }}</span>
                    <span>&#8226;</span>
                    <span>{{ formatDate(flight.ArrivalTime) }}</span>
                </div>
            </div>

            = <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6 mb-4">

                <div>
                    <p class="text-sm text-gray-500">Departure</p>
                    <p class="text-lg font-medium text-gray-800">{{ formatDate(flight.DepartureTime) }}</p>
                </div>

                <div>
                    <p class="text-sm text-gray-500">Arrival</p>
                    <p class="text-lg font-medium text-gray-800">{{ formatDate(flight.ArrivalTime) }}</p>
                </div>

                <div class="flex flex-col gap-2">
                    <p class="text-sm text-gray-500">Price (Saver)</p>
                    <p class="text-xl font-medium text-green-600">₹ {{ flight.PriceSaver }}</p>
                    <p class="text-sm text-gray-500">Price (Flexi)</p>
                    <p class="text-xl font-medium text-blue-600">₹ {{ flight.PriceFlexi }}</p>
                </div>
            </div>

            <div v-if="flight.layover"
                class="bg-yellow-50 border-l-4 border-yellow-500 p-4 mt-6 rounded-lg shadow-inner">
                <p class="font-medium text-yellow-600">Layover: {{ flight.layover }}</p>
                <p class="text-sm text-gray-600">Combined Price (Saver): ₹ {{ flight.combined_price_saver }}</p>
                <p class="text-sm text-gray-600">Combined Price (Flexi): ₹ {{ flight.combined_price_flexi }}</p>
            </div>

            <div class="mt-6 flex justify-end">
                <button
                    class="px-6 py-2 bg-indigo-600 text-white font-semibold rounded-lg shadow-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500">
                    Book Now
                </button>
            </div>
        </div>
    </div>

    <div v-else class="text-center text-gray-500">
        <p>No flights available for the selected route.</p>
    </div>
</template>

<script>
import { format } from 'date-fns';

export default {
    name: "FlightSearch",
    props: {
        flights: Array,
    },

    methods: {
        formatDate(date) {
            return format(new Date(date), 'dd-MM-yyyy HH:mm aa');

        }
    }
};
</script>
