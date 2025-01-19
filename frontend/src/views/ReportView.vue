<template>
  <div class="min-h-screen bg-gray-50 py-12">
    <div class="max-w-4xl mx-auto px-4 flex flex-col gap-8">
      <div class="flex justify-between">
        <button
          type="button"
          @click="goBack"
          class="bg-gray-100 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-200 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
        >
          Back
        </button>
        <div class="flex gap-4">
          <button
            type="button"
            @click="goBack"
            class="flex-1 bg-green-400 text-gray-700 py-2 px-4 rounded-lg font-medium hover:bg-gray-200 transition-colors duration-200 focus:outline-none focus:ring-2 focus:ring-gray-500 focus:ring-offset-2"
          >
            Send Email
          </button>
          <button
            @click="downloadTicket"
            class="inline-flex items-center px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors duration-200"
          >
            <svg
              class="w-4 h-4 mr-2"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4 4m0 0l-4-4m4 4V4"
              />
            </svg>
            Download Ticket
          </button>
        </div>
      </div>
      <div
        v-for="passenger in bookingData.Passengers"
        :key="passenger.ID"
        class="mb-8"
      >
        <div class="bg-white rounded-xl shadow-lg overflow-hidden">
          <!-- Header with gradient -->
          <div class="bg-gradient-to-r from-blue-600 to-blue-800 p-6">
            <div class="flex justify-between items-center">
              <h2 class="text-xl font-bold text-white">Flight Ticket</h2>
              <span class="text-sm text-blue-100"
                >PNR: {{ bookingData.PNRnumber }}</span
              >
            </div>
          </div>

          <!-- Main content -->
          <div class="p-6 space-y-6">
            <!-- Flight ID -->
            <div class="flex justify-between items-center">
              <span class="text-sm text-gray-500"
                >Flight ID: {{ bookingData.FlightID }}</span
              >
            </div>

            <!-- Flight Route -->
            <div class="flex items-center justify-between py-4">
              <div class="text-center flex-1">
                <h3 class="text-sm font-medium text-gray-500 mb-1">From</h3>
                <p class="text-xl font-bold text-gray-900">
                  {{ bookingData.Flight.Source }}
                </p>
                <p class="text-sm font-medium text-gray-600">
                  {{ formatTime(bookingData.Flight.DepartureTime) }}
                </p>
              </div>

              <div class="flex-1 px-4">
                <div class="relative">
                  <div class="border-t-2 border-dashed border-gray-300"></div>
                  <div
                    class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"
                  >
                    <svg
                      class="w-8 h-8 text-blue-600"
                      fill="none"
                      stroke="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M5 12h14M12 5l7 7-7 7"
                      />
                    </svg>
                  </div>
                </div>
              </div>

              <div class="text-center flex-1">
                <h3 class="text-sm font-medium text-gray-500 mb-1">To</h3>
                <p class="text-xl font-bold text-gray-900">
                  {{ bookingData.Flight.Destination }}
                </p>
                <p class="text-sm font-medium text-gray-600">
                  {{ formatTime(bookingData.Flight.ArrivalTime) }}
                </p>
              </div>
            </div>

            <!-- Passenger Details -->
            <div class="border-t flex justify-between border-gray-200 pt-4">
              <div class="bg-gray-50 rounded-lg p-4">
                <h3 class="text-sm font-medium text-gray-500 mb-2">
                  Passenger Details
                </h3>
                <p class="text-lg font-bold text-gray-900">
                  {{ passenger.FirstName }} {{ passenger.LastName }}
                </p>
              </div>
              <div class="bg-gray-50 rounded-lg p-4">
                <h3 class="text-sm font-medium text-gray-500 mb-2">
                  Seat Number:
                </h3>
                <p class="text-lg font-bold text-gray-900">
                  {{ passenger.SeatNumber }}
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { downloadTickets, getFlightsByReferenceNo } from "@/api/flight";
import { format } from "date-fns";
import FileSaver from "file-saver";

export default {
  name: "ReportView",
  data() {
    return {
      bookingData: {},
    };
  },
  props: {},

  components: {},

  created() {},

  async mounted() {
    try {
      const result = await getFlightsByReferenceNo(
        this.$route.params.reference
      );
      console.log(result);

      this.bookingData = result.data.result;
      console.log(this.bookingData);
    } catch (error) {
      console.error("Error fetching booking data:", error);
    }
  },

  beforeDestroy() {},

  methods: {
    formatDate(date) {
      return format(new Date(date), "MMM dd, yyyy");
    },
    // Format time as "hh:mm a"
    formatTime(time) {
      return format(new Date(time), "hh:mm a");
    },
    goBack() {
      this.$router.back();
    },

     async downloadTicket() {
      try {
        const result = await downloadTickets(this.$route.params.reference);
        console.log(result);
        // const pdfBlob=new Blob([result.data],{type:"application/pdf"})
        const blob = new Blob([result.data], { type: 'application/pdf' });

    // Create blob URL
    const url = window.URL.createObjectURL(blob);

    // Create temporary link element
    const link = document.createElement('a');
    link.href = url;
    link.download = `ticket_${this.$route.params.reference}.pdf`;

    // Append to document, click and cleanup
    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);

    // Free up memory by revoking the URL
    window.URL.revokeObjectURL(url);

    } catch (error) {
        console.error("Error fetching booking data:", error);
      }
    }

  },

  computed: {},

  watch: {},

  directives: {},

  filters: {},
};
</script>

<style scoped lang="scss"></style>
