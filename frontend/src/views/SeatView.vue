<template>
    <div class="w-screen h-screen flex flex-col justify-center gap-4 p-4 bg-blue-100">
        <button @click="goBack"
            class="absolute top-4 left-4 px-4 py-2 bg-gray-300 text-gray-700 rounded-lg shadow-md hover:bg-gray-400 transition duration-200 focus:outline-none">
            &lt; Back
        </button>
        <h1 class="text-center text-2xl font-semibold mb-6">Select seats for {{ selectedBookFlights.passengers }}
            Passengers</h1>

        <div class="w-full mt-16 flex justify-center">
            <div class="flex flex-col gap-4">
                <div class="flex items-center font-semibold gap-4">
                    <button
                        class="w-12 h-12 bg-red-300 rounded-lg text-center flex items-center justify-center hover:bg-gray-400 focus:outline-none"></button>
                    <label for="">Reserved</label>
                </div>
                <div class="flex items-center font-semibold gap-4">
                    <button
                        class="w-12 h-12 bg-green-300 rounded-lg text-center flex items-center justify-center hover:bg-gray-400 focus:outline-none"></button>
                    <label for="">Selected</label>
                </div>
                <div class="flex items-center font-semibold gap-4">
                    <button
                        class="w-12 h-12 bg-gray-300 rounded-lg text-center flex items-center justify-center hover:bg-gray-400 focus:outline-none"></button>
                    <label for="">Available</label>
                </div>
            </div>

            <div class="grid grid-cols-2 gap-4 w-[500px]">
                <div v-for="rowIndex in rows" :key="rowIndex" class="flex gap-8 justify-center">
                    <div class="flex gap-2">
                        <button v-for="seatIndex in column" :key="'row' + rowIndex + '-col1-' + seatIndex"
                            class="w-12 h-12 bg-gray-300 rounded-lg text-center flex items-center justify-center hover:bg-gray-400 focus:outline-none"
                            :class="{
                                'cursor-not-allowed bg-red-300': seatStatus(rowIndex, seatIndex) === 'reserved',
                                'bg-blue-500 hover:bg-blue-500': seatStatus(rowIndex, seatIndex) === 'available',
                                'bg-green-300 hover:bg-green-300': seatStatus(rowIndex, seatIndex) === 'selected'
                            }" @click="selectSeat(rowIndex, seatIndex)"
                            :disabled="seatStatus(rowIndex, seatIndex) === 'reserved' || (selectedBookFlights.passengers <= selectedSeats.length && checkSeats(rowIndex, seatIndex))">
                            {{ (rowIndex - 1) * 3 + seatIndex }}
                        </button>
                    </div>
                </div>

            </div>
        </div>
        <div
            class="flex justify-between items-center gap-4 w-full max-w-[500px] mx-auto mt-6 border border-green-400 rounded-lg px-4 py-2">
            <p class="font-sans text-xl font-semibold text-gray-800">
                Price: <span class="text-lg text-blue-600">{{ priceBySelected }}</span>
            </p>

            <button :disabled="selectedSeats.length > selectedBookFlights.passengers"
                class="px-6 py-3 text-lg font-semibold rounded-lg text-center flex items-center justify-center focus:outline-none transition duration-300"
                :class="giveClassbyNoOfPassenger" @click="seatContinue">
                Continue
            </button>
        </div>
        <div v-show="toast"
            class="animate-in animate-out absolute top-32 right-5 bg-green-400 text-white px-4 py-2 rounded-md">
            {{ message }}
        </div>
    </div>
</template>


<script>
import { useFlightStore } from '@/stores/flight';
import { mapState } from 'pinia';

export default {
    name: "SeatView",
    data() {
        return {
            toast: false,
            message: "some message happend",
            rows: 10,
            column: 3,
            selectedSeats: [],
        };
    },
    computed: {
        ...mapState(useFlightStore, ["selectedBookFlights"]),
        priceBySelected() {
            return this.selectedBookFlights.defaultPrice + (this.selectedSeats.length === 0 ? 0 : this.selectedSeats.length * 300)
        },
        giveClassbyNoOfPassenger() {
            console.log(this.selectedSeats.length);
            console.log(this.selectedBookFlights.passengers);


            if (this.selectedSeats.length == this.selectedBookFlights.passengers) {
                console.log('equal');


                return "bg-blue-600 text-white hover:bg-blue-700"
            } else {
                return "bg-gray-400 cursor-not-allowed"
            }
        }
    },
    methods: {
        seatContinue() {
            this.$router.push('/payment')
        },
        goBack() {

            this.$router.go(-1);
        },
        submitSelectedSeats() {
            this.selectedBookFlights.selectedSeats = this.selectedSeats
            // console.log('submitted')
        },
        seatStatus(row, seatIndex) {
            const reservedSeats = [
                1, 2, 4, 5, 7, 8, 10, 11
            ];
            return reservedSeats.includes((row - 1) * 3 + seatIndex) ? 'reserved' : this.selectedSeats.includes((row - 1) * 3 + seatIndex) ? "selected" : 'available';
        },
        checkSeats(row, seatIndex) {
            const seat = (row - 1) * 3 + seatIndex;
            if (!this.selectedSeats.includes(seat)) {

                return true
            } else {
                return false
            }

        },
        selectSeat(row, seatIndex) {
            const seat = (row - 1) * 3 + seatIndex;
            if (!this.selectedSeats.includes(seat)) {
                this.selectedSeats.push(seat);
                if (this.selectedSeats.length == this.selectedBookFlights.passengers) {
                    this.message = `${this.selectedBookFlights.passengers} passenger has been selected`
                    this.toast = true
                }
            } else {
                const index = this.selectedSeats.indexOf(seat);
                if (index > -1) {
                    this.selectedSeats.splice(index, 1);
                    this.toast = false
                }
            }
        }
    }
};
</script>
