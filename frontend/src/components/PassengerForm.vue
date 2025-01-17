<template>
    <div class="max-w-4xl mx-auto p-4">

        <form @submit.prevent="submitForm" class="space-y-6">
            <!-- Loop through the number of passengers -->
            <div v-for="index in parseInt(selectedBookFlights.passengers)" :key="index"
                class="border p-4 rounded-lg shadow-sm space-y-4">


                <h3 class="text-lg font-semibold text-gray-700">Passenger {{ index }}</h3>

                <!-- First Name Input -->
                <div class="flex  gap-4 w-full items-stretch">


                    <div class="w-full">
                        <label for="title" class="block text-sm font-medium text-gray-600">Title:</label>
                        <select :name="`title-${index}`" class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                            @change="updatePassengerData(index - 1, 'title', $event.target.value)" required>
                            <option value="" disabled selected>Select title</option>
                            <option value="Mr">Mr</option>
                            <option value="Ms">Ms</option>
                        </select>
                    </div>
                    <div class="w-full">
                        <label for="firstName" class="block text-sm font-medium text-gray-600">First Name:</label>
                        <input type="text" :value="passengerData[index - 1]?.firstName || ''"
                            :name="`firstName-${index}`" placeholder="Enter first name"
                            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                            @input="updatePassengerData(index - 1, 'firstName', $event.target.value)" required />
                    </div>

                    <!-- Last Name Input -->
                    <div class="w-full">
                        <label for="lastName" class="block text-sm font-medium text-gray-600">Last Name:</label>
                        <input type="text" :value="passengerData[index - 1]?.lastName || ''" :name="`lastName-${index}`"
                            placeholder="Enter last name"
                            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                            @input="updatePassengerData(index - 1, 'lastName', $event.target.value)" required />
                    </div>

                    <!-- Email Input -->
                    <div class="w-full">
                        <label for="email" class="block text-sm font-medium text-gray-600">Email:</label>
                        <input type="email" :value="passengerData[index - 1]?.email || ''" :name="`email-${index}`"
                            placeholder="Enter email"
                            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500"
                            @input="updatePassengerData(index - 1, 'email', $event.target.value)" required />
                    </div>
                </div>

            </div>

            <!-- Submit Button -->
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
    data() {
        return {
            passengerData: []
        };
    },
    methods: {
        // Handle form submission
        submitForm() {
            console.log('Form Submitted with:', this.passengerData);
            this.selectedBookFlights.passengerDetails=this.passengerData
            this.$router.push("/seat")
        },
        addTitle(event){
            this.passengerData.title=event.target.value
        },

        // Update passenger data
        updatePassengerData(index, field, value) {
            // Ensure the passengerData array is large enough
            if (!this.passengerData[index]) {
                this.passengerData[index] = {};
            }

            this.passengerData[index][field] = value;
        }
    },



    mounted() {
        // Initialize the passengerData array when the component is mounted
        this.passengerData = Array.from({ length: this.selectedBookFlights.passengers }, () => ({
            firstName: '',
            lastName: '',
            email: ''
        }));
    }
}
</script>

<style scoped></style>
