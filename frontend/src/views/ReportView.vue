<template>
     <div class="min-h-screen bg-gray-50 py-12">
        <div v-show="showEmailInput" class="fixed inset-0 flex justify-center items-center bg-gray-800 bg-opacity-50">
    <div  class="absolute p-6 bg-white rounded-lg shadow-lg w-96">

        <input type="email" @input="sendEmail" class="w-full p-4 border rounded-lg shadow-md mb-4" placeholder="Enter your email">

<div v-if="loading===false" class="flex gap-4">

    <button @click="sendApi" class="w-full p-4 bg-blue-500 text-white rounded-lg shadow-md hover:bg-blue-600">
        Send
    </button>
    <button @click="showEmailInput=false" class="w-full p-4 bg-red-500 text-white rounded-lg shadow-md hover:bg-red-600">
        Cancel
    </button>
</div>
<div v-else>
    <LoadingComponent  />
</div>
</div>
</div>
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
            @click="emailButton"
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
          <div v-if="downloading">
            <LoadingComponent/>
          </div>
        </div>


      </div>
  <div v-if="display" class="bg-white border w-[40rem]  rounded-lg shadow-md p-8">
    <!-- Header Section -->
    <div class="border-b border-gray-200 pb-6">
      <h1 class="text-2xl font-bold text-gray-900 mb-4">Flight Booking Report</h1>
      <div class="grid grid-cols-2 gap-4">
        <div>
          <p class="text-sm text-gray-600">Booking PNR Number:</p>
          <p class="text-lg font-semibold text-gray-900">{{ bookingData.PNRnumber }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-600">Date of Booking:</p>
          <p class="text-lg font-semibold text-gray-900">{{ formatDate(bookingData.CreatedAt) }}</p>
        </div>
      </div>
    </div>

    <!-- Flight Details -->
    <div class="py-6 border-b border-gray-200">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">Flight Details</h2>
      <div class="grid grid-cols-2 gap-6">
        <div>
          <p class="text-sm text-gray-600">Date of Travel:</p>
          <p class="text-base font-medium text-gray-900">{{ formatDate(bookingData.Flight.DepartureTime) }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-600">Flight Number:</p>
          <p class="text-base font-medium text-gray-900">{{ bookingData.FlightID }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-600">From:</p>
          <p class="text-base font-medium text-gray-900">{{ bookingData.Flight.Source }}</p>
          <p class="text-sm text-gray-500">{{ formatTime(bookingData.Flight.DepartureTime) }}</p>
        </div>
        <div>
          <p class="text-sm text-gray-600">To:</p>
          <p class="text-base font-medium text-gray-900">{{ bookingData.Flight.Destination }}</p>
          <p class="text-sm text-gray-500">{{ formatTime(bookingData.Flight.ArrivalTime) }}</p>
        </div>
        <div>
            <p class="text-sm text-gray-600">Price:</p>
          <p class="text-base font-medium text-gray-900">Rs {{ bookingData.Price*bookingData.Passengers.length }}</p>

        </div>
      </div>
    </div>

    <!-- Passenger Details -->
    <div class="pt-6">
      <h2 class="text-lg font-semibold text-gray-900 mb-4">Passenger Details</h2>
      <div class="space-y-4">
        <div v-for="(passenger, index) in bookingData.Passengers" :key="passenger.ID"
             class="flex justify-between items-start p-4 bg-gray-50 rounded-lg">
          <div class="flex-1">
            <p class="text-sm text-gray-600 mb-1">Passenger {{ index + 1 }}</p>
            <p class="text-base font-medium text-gray-900">
              {{ passenger.FirstName }} {{ passenger.LastName }}
            </p>
          </div>
          <div class="text-right">
            <p class="text-sm text-gray-600 mb-1">Seat Number</p>
            <p class="text-base font-medium text-gray-900">{{ passenger.SeatNumber }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
<div v-if="toast"
        class="fixed top-4 right-4 bg-emerald-500 text-white px-4 py-3 rounded-lg shadow-lg animate-fade-in">
        {{ message }}
      </div>
    </div>
</template>

<script>
import { downloadTickets, getFlightsByReferenceNo, sendEmailApi } from '@/api/flight';
import LoadingComponent from '@/components/LoadingComponent.vue';
import { format } from 'date-fns';

export default {
    name: 'ReportView',
    data() {
        return {
            display:false,
            bookingData: {},
            email:"",
            showEmailInput:false,
            loading:false,
            toast:false,
            message:"dummy message",
            downloading:false,

        };
    },
    props: {},

    components: {
        LoadingComponent
    },

    created() {},

    async mounted() {
        try {
      const result = await getFlightsByReferenceNo(
        this.$route.params.reference
      );
      console.log(result);

      this.bookingData = result.data.result;
      this.display=true
      console.log(this.bookingData);
    } catch (error) {
      console.error("Error fetching booking data:", error);
    }
    },

    beforeDestroy() {},

    methods: {
        async sendApi(){
            try{

                this.loading=true
                const result = await sendEmailApi(this.email,this.$route.params.reference)
                if(result.status===200){
                    this.toast=true
                    this.message=result.data.message
                    this.loading=false
                    this.showEmailInput=false
                }
            }catch(err)
            {
                alert(err)
            }

        },
        emailButton(){
            this.showEmailInput=true;
        },
        sendEmail(event){
            this.email=event.target.value
        },
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
        this.downloading=true
        const result = await downloadTickets(this.$route.params.reference);
        console.log(result);
        if(result.status===200){
            this.downloading=false
        }

        const blob = new Blob([result.data], { type: 'application/pdf' });


    const url = window.URL.createObjectURL(blob);


    const link = document.createElement('a');
    link.href = url;
    link.download = `ticket_${this.$route.params.reference}.pdf`;


    document.body.appendChild(link);
    link.click();
    document.body.removeChild(link);


    window.URL.revokeObjectURL(url);

    } catch (error) {
        console.error("Error fetching booking data:", error);
      }
    }

    },

    computed: {},

    watch: {},

    directives: {},

    filters: {}
};
</script>

<style scoped lang="scss"></style>
