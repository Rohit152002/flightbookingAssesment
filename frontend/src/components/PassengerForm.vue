<template>
    <div class="max-w-4xl mx-auto p-4">
        <form @submit.prevent="submitForm" class="space-y-6">
            <div v-for="(passenger, index) in selectedBookFlights.passengers" :key="index"
                class="border p-4 rounded-lg shadow-sm space-y-4">
                <h3 class="text-lg font-semibold text-gray-700">Passenger {{ index + 1 }}</h3>

                <div>
                    <label for="firstName" class="block text-sm font-medium text-gray-600">First Name:</label>
                    <input type="text" v-model="passenger.firstName" :name="`firstName-${index}`"
                        placeholder="Enter first name"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                        required />
                </div>

                <div>
                    <label for="lastName" class="block text-sm font-medium text-gray-600">Last Name:</label>
                    <input type="text" v-model="passenger.lastName" :name="`lastName-${index}`"
                        placeholder="Enter last name"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                        required />
                </div>

                <div>
                    <label for="email" class="block text-sm font-medium text-gray-600">Email:</label>
                    <input type="email" v-model="passenger.email" :name="`email-${index}`" placeholder="Enter email"
                        class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                        required />
                </div>
            </div>

            <button type="submit"
                class="w-full py-3 mt-4 bg-blue-500 text-white font-semibold rounded-md shadow-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50">
                Submit
            </button>
        </form>
    </div>
</template>

<script>
import { useFlightStore } from '@/stores/flight';
import { mapState } from 'pinia';

export default {
    name: "PassengerForm",
    computed: {
        ...mapState(useFlightStore, ["selectedBookFlights"])
    },
    methods: {
        submitForm() {

            console.log('Form Submitted with:', this.selectedBookFlights.passengers);
        }
    },
    mounted() {
        console.log(this.selectedBookFlights.passengers);
    }
}
</script>

<style scoped></style>