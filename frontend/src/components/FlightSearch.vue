<template>
    <div v-if="flights" class="space-y-4 w-full max-w-3xl mx-auto">
        <div v-for="(flight, index) in flights" :key="flight.ID"
            class="bg-white rounded-lg shadow-sm border border-gray-200 overflow-hidden">
            <!-- Header - Always Visible -->
            <div @click="toggleFlight(index)"
                class="p-4 flex items-center justify-between cursor-pointer hover:bg-gray-50">
                <div class="flex items-center space-x-4">
                    <div class="p-2 bg-indigo-50 rounded-full">
                        <!-- Plane Icon -->
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-indigo-600" viewBox="0 0 24 24"
                            fill="none" stroke="currentColor" stroke-width="2">
                            <path
                                d="M22 16.92v3a2 2 0 0 1-2.18 2 19.79 19.79 0 0 1-8.63-3.07 19.5 19.5 0 0 1-6-6 19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72 12.84 12.84 0 0 0 .7 2.81 2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45 12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z" />
                        </svg>
                    </div>
                    <div>
                        <h3 class="font-medium text-gray-900">{{ flight.Airline }}</h3>
                        <p class="text-sm text-gray-500">{{ flight.Source }} → {{ flight.Destination }}</p>
                        <!-- Flight Details - Only shown when NOT expanded -->
                        <div v-if="selectedFlight !== index"
                            class="mt-2 flex items-center space-x-4 text-sm text-gray-500">
                            <span>{{ formatDate(flight.DepartureTime) }}</span>
                            <span>→</span>
                            <span>{{ formatDate(flight.ArrivalTime) }}</span>
                            <span v-if="flight.layover" class="text-amber-600">
                                ({{ flight.layover }} layover)
                            </span>
                        </div>
                    </div>
                </div>
                <div class="flex items-center space-x-4">
                    <div class="text-right">
                        <p class="text-sm text-gray-600">From</p>
                        <p class="text-lg font-semibold text-indigo-600">
                            ₹ {{ selectedPrice === "Saver" ? flight.PriceSaver : flight.PriceFlexi }}
                        </p>
                    </div>
                    <!-- Chevron Icon -->
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-gray-400 transform transition-transform"
                        :class="{ 'rotate-180': selectedFlight === index }" viewBox="0 0 24 24" fill="none"
                        stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7" />
                    </svg>
                </div>
            </div>

            <!-- Expanded Content -->
            <div v-show="selectedFlight === index" class="border-t border-gray-200">
                <div class="p-4 space-y-4">
                    <!-- Price Selection -->
                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <button @click="selectPrice('Saver')" :class="[
                            'p-4 rounded-lg border transition-all',
                            selectedPrice === 'Saver'
                                ? 'border-indigo-600 bg-indigo-50 text-indigo-600'
                                : 'border-gray-200 hover:border-indigo-600'
                        ]">
                            <p class="text-sm text-gray-500">Saver Fare</p>
                            <p class="text-xl font-semibold">₹ {{ flight.PriceSaver * passengers }}</p>
                            <p class="text-xs text-gray-500 mt-1">Non-refundable</p>
                        </button>

                        <button @click="selectPrice('Flexi')" :class="[
                            'p-4 rounded-lg border transition-all',
                            selectedPrice === 'Flexi'
                                ? 'border-indigo-600 bg-indigo-50 text-indigo-600'
                                : 'border-gray-200 hover:border-indigo-600'
                        ]">
                            <p class="text-sm text-gray-500">Flexi Fare</p>
                            <p class="text-xl font-semibold">₹ {{ flight.PriceFlexi * passengers }}</p>
                            <p class="text-xs text-gray-500 mt-1">Fully refundable</p>
                        </button>
                    </div>

                    <!-- Layover Info -->
                    <div v-if="flight.layover" class="bg-amber-50 border-l-4 border-amber-400 p-4 rounded-r-lg">
                        <p class="font-medium text-amber-700">
                            Layover: {{ flight.layover }}
                        </p>
                        <div class="mt-2 space-y-1 text-sm text-amber-600">
                            <p>Combined Saver Price: ₹ {{ flight.combined_price_saver }}</p>
                            <p>Combined Flexi Price: ₹ {{ flight.combined_price_flexi }}</p>
                        </div>
                    </div>

                    <!-- Book Button -->
                    <div class="flex justify-end pt-2">
                        <button
                            class="px-6 py-2 bg-indigo-600 text-white font-semibold rounded-lg shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2">
                            Book Now
                        </button>
                    </div>
                </div>
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
        flights: {
            type: Array,
            required: true
        },
        passengers: {
            type: Number,
            default: 1
        }
    },
    data() {
        return {

            selectedPrice: "Saver",
            selectedFlight: null
        }
    },

    methods: {
        formatDate(date) {

            return format(new Date(date), 'dd-MM-yyyy HH:mm aa');

        },
        selectPrice(priceType) {
            this.selectedPrice = this.selectedPrice === priceType ? null : priceType
        },
        toggleFlight(index) {
            this.selectedFlight = this.selectedFlight === index ? null : index
        },
    }
};
</script>
