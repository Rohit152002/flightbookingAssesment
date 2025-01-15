<template>
    <div
        class="max-w-screen-md mx-auto p-6 bg-white rounded-xl shadow-sm hover:shadow-lg transition-shadow duration-200 ease-in-out">
        <div class="space-y-4">
            <p class="text-xl font-semibold text-gray-800">{{ flight.Source }} ↔ {{ flight.Destination }}</p>
            <p class="text-lg text-gray-600">{{ formattedDate }}</p>

            <div class="flex justify-between text-sm text-gray-500">
                <div class="flex gap-4 text-lg">
                    <p><strong>Departure: </strong></p>
                    <p>{{ formattedDepartureTime }}</p>
                </div>
                <div class="flex gap-4 text-lg">
                    <p><strong>Arrival: </strong></p>
                    <p>{{ formattedArrivalTime }}</p>
                </div>
            </div>

            <div class="border-t  border-gray-200 pt-4">
                <div class="flex justify-start gap-2 items-center text-lg">
                    <p class="text-lg font-medium text-gray-800">Airline :</p>
                    <p class="text-lg text-gray-600">{{ flight.Airline }}</p>
                </div>

                <!-- Default Price -->
                <div class="mt-4 flex items-center gap-2">
                    <p class="text-lg font-medium text-gray-800"> Price :- </p>
                    <p class="text-lg text-black ">₹{{ selectedPriceValue }}</p>
                </div>
                <div class="mt-4 flex gap-2 items-center">
                    <p class="text-lg font-medium text-gray-800">Passenger No. :- </p>
                    <button @click="decreasePassenger" :disabled="flight.passengers <= 1"
                        class="text-xl border border-black font-sans px-4 py-2"> -
                    </button>
                    <p class="text-lg text-black font-semibold border border-black rounded-md px-4 py-2">{{
                        flight.passengers }}</p>
                    <button @click="increasePassenger" :disabled="flight.passengers === 10"
                        class="text-xl border border-black font-sans px-4 py-2">+</button>
                </div>

                <div class="mt-4">
                    <p class="text-lg font-medium text-gray-800">Select Price Option</p>

                    <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                        <button @click="selectPrice('Saver', index)" :class="[
                            'p-4 rounded-lg border transition-all',
                            selectedPriceType === 'Saver'
                                ? 'border-indigo-600 bg-indigo-50 text-indigo-600'
                                : 'border-gray-200 hover:border-indigo-600',
                        ]">
                            <p class="text-sm text-gray-500">Saver Fare</p>
                            <p class="text-xl font-semibold">
                                ₹ {{ flight.PriceSaver * flight.passengers }}
                            </p>
                            <p class="text-xs text-gray-500 mt-1">Non-refundable</p>
                        </button>

                        <button @click="selectPrice('Flexi', index)" :class="[
                            'p-4 rounded-lg border transition-all',
                            selectedPriceType === 'Flexi'
                                ? 'border-indigo-600 bg-indigo-50 text-indigo-600'
                                : 'border-gray-200 hover:border-indigo-600',
                        ]">
                            <p class="text-sm text-gray-500">Flexi Fare</p>
                            <p class="text-xl font-semibold">
                                ₹ {{ flight.PriceFlexi * flight.passengers }}
                            </p>
                            <p class="text-xs text-gray-500 mt-1">Fully refundable</p>
                        </button>
                    </div>
                    <!-- <div class="mt-2">
                        <p class="text-sm text-gray-600">Selected Price: ₹{{ selectedPriceValue }}</p>
                    </div> -->
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { format } from 'date-fns';

export default {
    name: "FlightDetail",
    props: {
        flight: {
            type: Object,
            required: true
        }
    },
    data() {
        return {
            selectedPriceType: "",
            selectedPrice: '',  // Default selected price option
        };
    },
    computed: {
        formattedDate() {
            return format(new Date(this.flight.Date), 'yyyy-MM-dd');
        },
        formattedDepartureTime() {
            return format(new Date(this.flight.DepartureTime), 'HH:mm');
        },
        formattedArrivalTime() {
            return format(new Date(this.flight.ArrivalTime), 'HH:mm');
        },
        selectedPriceValue() {
            if (this.selectedPrice === "PriceSaver") {
                return this.flight.PriceSaver * this.flight.passengers
            } else if (this.selectedPrice === "PriceFlexi") {
                return this.flight.PriceFlexi * this.flight.passengers
            }
            return this.flight.defaultPrice * this.flight.passengers
        }
    },
    methods: {
        increasePassenger() {
            this.flight.passengers++
        },
        decreasePassenger() {
            this.flight.passengers--
        },
        selectPrice(priceType, index) {
            if (priceType === 'Saver') {
                this.selectedPriceType = 'Saver'
                this.selectedPrice = 'PriceSaver'
                console.log(this.selectedPrice);

            } else if (priceType === 'Flexi') {
                this.selectedPriceType = 'Flexi'
                this.selectedPrice = 'PriceFlexi'
                console.log(this.selectedPrice);

            }
        }
    },
    mounted() {
        if (this.flight.defaultPrice === this.flight.PriceSaver) {
            this.selectedPriceType = "Saver"
            this.selectedPrice = "Saver"
        } else {
            this.selectedPriceType = "Flexi"
            this.selectedPrice = "Flexi"
        }
    }
}
</script>

<style scoped></style>
