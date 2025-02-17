<template>
  <div class="h-full mx-auto">
    <form
      @submit.prevent="SearchFlight"
      class="flex flex-wrap gap-6 border border-gray-300 rounded-lg p-6"
    >
      <div class="flex gap-6 w-full mb-4">
        <label class="flex items-center space-x-2">
          <input
            type="radio"
            name="tripType"
            value="one-way"
            autocomplete="off"
            v-model="tripType"
            @click="flights = null"
            checked
          />
          <span>One Way</span>
        </label>
        <label class="flex items-center space-x-2">
          <input
            type="radio"
            name="tripType"
            value="round-trip"
            autocomplete="off"
            @click="flights = null"
            v-model="tripType"
          />
          <span>Round Trip</span>
        </label>
      </div>

      <div class="w-full flex gap-4 justify-around items-end">
        <div class="flex flex-col relative w-full space-y-4">
          <label for="source" class="text-lg font-medium text-gray-700"
            >Source:</label
          >
          <input
            v-model="source"
            @input="filterSource"
            type="text"
            name="source"
            id="source"
            class="border relative w-full flex-1 border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
            placeholder="Search for a station..."
            required
            @keydown.down="navigateDown"
            @keydown.up="navigateUp"
            autocomplete="off"
            @keydown.enter="selectStationFromKeyboard"
          />

          <div
            v-if="source && filteredStations.length && !sourceSelected"
            class="absolute top-full w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg z-10"
          >
            <div class="space-y-2">
              <p
                v-for="(station, i) in filteredStations"
                :key="i"
                :class="{
                  'bg-indigo-100 text-indigo-700': i === highlightedIndex,
                }"
                @click="selectSourceStation(station)"
                class="px-4 py-2 cursor-pointer hover:bg-indigo-100 hover:text-indigo-700 transition-colors duration-200"
              >
                {{ station }}
              </p>
            </div>
          </div>
        </div>

        <div class="flex flex-col relative w-full space-y-4">
          <label for="source" class="text-lg font-medium text-gray-700"
            >Destination:</label
          >
          <input
            v-model="destination"
            @input="filterDestination"
            type="text"
            name="source"
            id="source"
            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
            placeholder="Search for a station..."
            required
            @keydown.down="navigateDown"
            @keydown.up="navigateUp"
            autocomplete="off"
            @keydown.enter="selectDestinationFromKeyboard"
          />

          <div
            v-if="
              destination && filteredStations.length && !destinationSelected
            "
            class="absolute top-full w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg z-10"
          >
            <div class="space-y-2">
              <p
                v-for="(station, i) in filteredStations"
                :key="i"
                :class="{
                  'bg-indigo-100 text-indigo-700': i === highlightedIndex,
                }"
                @click="selectDestinationStation(station)"
                class="px-4 py-2 cursor-pointer hover:bg-indigo-100 hover:text-indigo-700 transition-colors duration-200"
              >
                {{ station }}
              </p>
            </div>
          </div>
        </div>

        <div class="flex flex-col relative w-full space-y-4">
          <label for="journey" class="text-lg font-medium text-gray-700"
            >Date of Journey:</label
          >
          <input
            type="date"
            class="border w-full relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
            autocomplete="off"
            v-model="journeyDate"
            :min="today"
            required
          />
        </div>

        <div
          class="flex flex-col relative w-full space-y-4"
          v-if="tripType === 'round-trip'"
        >
          <label for="return" class="text-lg font-medium text-gray-700"
            >Date of Return:</label
          >
          <input
            type="date"
            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
            required
            v-model="returnDate"
            :min="journeyDate"
          />
        </div>

        <div class="flex flex-col relative w-full space-y-4">
          <label for="passengers" class="text-lg font-medium text-gray-700"
            >No. of Passengers:</label
          >
          <select
            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
            v-model="passengers"
            required
            name="passengers"
            id="passengers"
          >
            <option value="1">1</option>
            <option value="2">2</option>
            <option value="3">3</option>
            <option value="4">4</option>
            <option value="5">5</option>
            <option value="6">6</option>
            <option value="7">7</option>
            <option value="8">8</option>
            <option value="9">9</option>
            <option value="10">10</option>
          </select>
        </div>

        <button
          type="submit"
          class="bg-blue-500 text-white py-4 px-6 rounded hover:bg-blue-600 w-full"
        >
          Search
        </button>
      </div>
    </form>
    <div v-show="loading" class="h-80 flex items-center justify-center">
      <LoadingComponent />
    </div>
    <div class="mt-20 mb-20">
      <FlightSearch
        v-if="!loading && tripType === 'one-way' && flights != null"
        :flights="flights"
        :passengers="passengers"
      />

      <RoundTripComponent
        v-else-if="!loading && tripType === 'round-trip' && flights != null"
        :departure-flights="flights.departureFlights"
        :return-flights="flights.returnFlights"
      />
      <div
        v-else-if="flights === null && search === true"
        class="h-56 text-xl font-bold flex items-center justify-center"
      >
        There is not a flight for this route
      </div>
      <div
        v-else
        class="h-56 flex items-center text-xl font-bold justify-center"
      >
        Search A flight
      </div>
    </div>
  </div>
</template>

<script>
import { Button } from "@/vue/ui/button";
import { Input } from "@/vue/ui/input";
import { useFlightStore } from "@/stores/flight";
import { searchFlightApi, searchFlightWithReturnDate } from "@/api/flight";
import FlightSearch from "@/components/FlightSearch.vue";

import LoadingComponent from "./LoadingComponent.vue";
import RoundTripComponent from "./RoundTripComponent.vue";

export default {
  name: "SearchComponent",
  components: {
    Button,
    Input,
    FlightSearch,
    LoadingComponent,
    RoundTripComponent,
  },
  data() {
    return {
      loading: false,
      flightStore: useFlightStore(),
      passengers: 1,
      flights: null,
      search: false,
      source: "",
      sourceSelected: false,
      destinationSelected: false,
      destination: "",
      stations: [],
      tripType: "one-way",
      journeyDate: "",
      returnDate: "",
      filteredStations: [],
      highlightedIndex: -1,
    };
  },
  methods: {
    filterSource() {
      if (this.source) {
        this.sourceSelected = false;
        this.filteredStations = this.stations.filter((station) =>
          station.toLowerCase().includes(this.source.toLowerCase())
        );
        this.highlightedIndex = -1;
      } else {
        this.filteredStations = this.stations;
      }
    },
    filterDestination() {
  if (this.destination) {
    this.destinationSelected = false;

    this.filteredStations = this.stations.filter(
      (station) =>
        station !== this.source &&
        station.toLowerCase().includes(this.destination.toLowerCase())
    );
  } else {
    this.filteredStations = this.stations.filter(
      (station) => station !== this.source
    );
  }
},
    selectSourceStation(value) {
      this.source = value;
      this.sourceSelected = true;
      this.highlightedIndex = -1;
    },
    selectDestinationStation(value) {
      this.destination = value;
      this.destinationSelected = true;
    },
    navigateDown(event) {
      if (this.highlightedIndex < this.filteredStations.length - 1) {
        this.highlightedIndex++;
      }
    },

    navigateUp(event) {
      if (this.highlightedIndex > 0) {
        this.highlightedIndex--;
      }
    },
    selectStationFromKeyboard(event) {
      if (this.highlightedIndex !== -1) {
        this.selectSourceStation(this.filteredStations[this.highlightedIndex]);
      }
    },
    selectDestinationFromKeyboard(event) {
      if (this.highlightedIndex !== -1) {
        this.selectDestinationStation(
          this.filteredStations[this.highlightedIndex]
        );
      }
    },
    async SearchFlight() {
      this.loading = true;
      if (this.tripType == "round-trip") {
        const result = await searchFlightWithReturnDate(
          this.source,
          this.destination,
          this.journeyDate,
          this.returnDate
        );
        if (result.data == null) {
          this.flights = null;
          this.search = true;
          this.loading = false;
          return;
        }

        this.flightStore.setFlight(result.data);
        this.flights = result.data;
        if (this.flights.departureFlights === null) {
          this.flightStore.selectedDepartureFlights = null;
          this.loading = false;
          return;
        }
        if (this.flights.returnFlights === null) {
          this.flightStore.selectedReturnFlights = null;
          this.loading = false;
          return;
        }
        this.flightStore.selectedDepartureFlights =
          this.flights.departureFlights[0];
        this.flightStore.selectedReturnFlights = this.flights.returnFlights[0];
        this.flightStore.selectedJourneyPrice =
          this.flights.departureFlights[0].PriceSaver;
        this.flightStore.selectedReturnPrice =
          this.flights.returnFlights[0].PriceSaver;

        console.log(result);
      } else {
        console.log("submit");
        console.log(this.source);
        console.log(this.destination);
        console.log(this.journeyDate);
        console.log(this.passengers);

        const result = await searchFlightApi(
          this.source,
          this.destination,
          this.journeyDate
        );

        if (result.data.data === null) {
          this.flights === null;
          this.search = true;
          this.loading = false;
          return;
        }
        this.flightStore.setFlight(result.data.data);

        this.flights = result.data.data;
        console.log(result.data.data);
        console.log(this.flights);
      }
      this.loading = false;
    },
  },
  computed: {
    today() {
      const date = new Date();
      const day = String(date.getDate()).padStart(2, "0");
      const month = String(date.getMonth() + 1).padStart(2, "0");
      const year = date.getFullYear();
      return `${year}-${month}-${day}`;
    },
  },
  async mounted() {
    const result = await this.flightStore.getStations();
    this.stations = result;
    this.filteredStations = result;
    console.log(result);
  },
};
</script>
