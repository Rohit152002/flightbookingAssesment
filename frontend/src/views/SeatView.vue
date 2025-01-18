<template>
  <div class="min-h-screen bg-gray-50 p-8">
    <div class="max-w-6xl   mx-auto relative">
      <!-- Back button -->
      <button @click="goBack"
        class="absolute -left-4 top-0 p-2 text-gray-600 hover:text-gray-900 transition-colors">
        &lt; Back
      </button>

      <!-- Header -->
      <h1 class="text-3xl font-light text-gray-800 text-center mb-12">
        Select seats for {{ selectedBookFlights.passengers }}
        {{ selectedBookFlights.passengers === 1 ? 'passenger' : 'passengers' }}
      </h1>

      <div class="flex justify-between items-start gap-16">
        <!-- Legend -->
        <div class="space-y-6">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded bg-red-100 border border-red-200" />
            <span class="text-sm text-gray-600">Reserved</span>
          </div>
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded bg-emerald-100 border border-emerald-200" />
            <span class="text-sm text-gray-600">Selected</span>
          </div>
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded bg-white border border-gray-200" />
            <span class="text-sm text-gray-600">Available</span>
          </div>
        </div>

        <!-- Seat Grid -->
        <div class="">
          <div class="grid grid-cols-2 gap-8">
            <div v-for="rowIndex in rows" :key="rowIndex" class="flex gap-4 justify-center">
              <button v-for="seatIndex in column"
                :key="'row' + rowIndex + '-col-' + seatIndex"
                @click="selectSeat(rowIndex, seatIndex)"
                :disabled="seatStatus(rowIndex, seatIndex) === 'reserved' ||
                          (selectedBookFlights.passengers <= selectedSeats.length &&
                           !selectedSeats.includes((rowIndex - 1) * 3 + seatIndex))"
                :class="[
                  'w-10 h-10 rounded text-sm transition-all',
                  {
                    'bg-red-100 border border-red-200 cursor-not-allowed':
                      seatStatus(rowIndex, seatIndex) === 'reserved',
                    'bg-emerald-100 border border-emerald-200 hover:bg-emerald-200':
                      seatStatus(rowIndex, seatIndex) === 'selected',
                    'bg-white border border-gray-200 hover:border-gray-400':
                      seatStatus(rowIndex, seatIndex) === 'available'
                  }
                ]"
              >
                {{ (rowIndex - 1) * 3 + seatIndex }}
              </button>
            </div>
          </div>
        </div>
      <div class="mt-12 flex w-fit gap-10 justify-between items-center bg-white p-6 rounded-lg shadow-sm">
        <div class="text-gray-800">
          <span class="text-sm">Total Price</span>
          <p class="text-2xl font-light">{{ priceBySelected.toLocaleString() }}</p>
        </div>
        <button
          @click="seatContinue"
          :disabled="selectedSeats.length!= selectedBookFlights.passengers"
          :class="[
            'px-6 py-3 rounded-lg text-sm font-medium transition-colors',
            selectedSeats.length == selectedBookFlights.passengers
              ? 'bg-blue-600 text-white hover:bg-blue-700'
              : 'bg-gray-100 text-gray-400 cursor-not-allowed'
          ]"
        >
          Continue
        </button>
      </div>
      </div>

      <!-- Price and Continue -->

      <!-- Toast Notification -->
      <div v-if="toast"
        class="fixed top-4 right-4 bg-emerald-500 text-white px-4 py-3 rounded-lg shadow-lg animate-fade-in">
        {{ message }}
      </div>
    </div>
  </div>
</template>

<script>
import { useFlightStore } from '@/stores/flight';
import { mapState } from 'pinia';

export default {
  name: "MinimalistSeatView",
  data() {
    return {
      toast: false,
      message: "All seats selected successfully",
      rows: 60,
      column: 3,
      selectedSeats: [],
    };
  },
  computed: {
    ...mapState(useFlightStore, ["selectedBookFlights"]),
    priceBySelected() {
      return this.selectedBookFlights.defaultPrice +
        (this.selectedSeats.length * 300);
    }
  },
  methods: {
    seatContinue() {
      this.selectedBookFlights.selectedSeats = this.selectedSeats;
      this.$router.push('/payment');
    },
    goBack() {
      this.$router.go(-1);
    },
    seatStatus(row, seatIndex) {
      const seatNumber = (row - 1) * 3 + seatIndex;
      const reservedSeats = [1, 2, 4, 5, 7, 8, 10, 11];

      if (reservedSeats.includes(seatNumber)) return 'reserved';
      if (this.selectedSeats.includes(seatNumber)) return 'selected';
      return 'available';
    },
    selectSeat(row, seatIndex) {
      const seatNumber = (row - 1) * 3 + seatIndex;

      if (this.selectedSeats.includes(seatNumber)) {
        const index = this.selectedSeats.indexOf(seatNumber);
        if (index > -1) {
          this.selectedSeats.splice(index, 1);
          this.toast = false;
        }
      } else if (this.selectedSeats.length < this.selectedBookFlights.passengers) {
        this.selectedSeats.push(seatNumber);

        if (this.selectedSeats.length == this.selectedBookFlights.passengers) {
          this.message = `${this.selectedBookFlights.passengers} passenger${
            this.selectedBookFlights.passengers > 1 ? 's have' : ' has'
          } been selected`;
          this.toast = true;
        }
      }
    }
  }
};
</script>

<style>
.animate-fade-in {
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(-10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
