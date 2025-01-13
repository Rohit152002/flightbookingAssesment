<template>
    <div class="w-screen mx-auto">
        <form @submit.prevent="SearchFlight" class="flex flex-wrap gap-6 border border-gray-300 rounded-lg p-6">

            <div class="flex gap-6 w-full mb-4">
                <label class="flex items-center space-x-2">
                    <input type="radio" name="tripType" value="one-way" v-model="tripType" checked>
                    <span>One Way</span>
                </label>
                <label class="flex items-center space-x-2">
                    <input type="radio" name="tripType" value="round-trip" v-model="tripType">
                    <span>Round Trip</span>
                </label>
            </div>

            <div class="flex flex-wrap gap-4 w-full mb-4 items-end">
                <div class="w-full  flex justify-around items-end gap-4">
                    <div class="flex flex-col relative w-full  sm:w-[280px] space-y-4">

                        <label for="source" class="text-lg font-medium text-gray-700">Source:</label>
                        <input v-model="source" @input="filterSource" type="text" name="source" id="source"
                            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
                            placeholder="Search for a station..." required @keydown.down="navigateDown"
                            @keydown.up="navigateUp" @keydown.enter="selectStationFromKeyboard" />

                        <!-- Floating Station List -->
                        <div v-if="source && filteredStations.length && !sourceSelected"
                            class="absolute  top-full w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg z-10">
                            <div class="space-y-2">
                                <p v-for="(station, i) in filteredStations" :key="i"
                                    :class="{ 'bg-indigo-100 text-indigo-700': i === highlightedIndex }"
                                    @click="selectSourceStation(station)"
                                    class="px-4 py-2 cursor-pointer hover:bg-indigo-100 hover:text-indigo-700 transition-colors duration-200">
                                    {{ station }}
                                </p>
                            </div>
                        </div>
                    </div>

                    <div class="flex flex-col relative  w-full sm:w-[280px] space-y-4">

                        <label for="source" class="text-lg font-medium text-gray-700">Destination:</label>
                        <input v-model="destination" @input="filterDestination" type="text" name="source" id="source"
                            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
                            placeholder="Search for a station..." required @keydown.down="navigateDown"
                            @keydown.up="navigateUp" @keydown.enter="selectDestinationFromKeyboard" />

                        <!-- Floating Station List -->
                        <div v-if="destination && filteredStations.length && !destinationSelected"
                            class="absolute  top-full w-full mt-1 bg-white border border-gray-300 rounded-lg shadow-lg z-10">
                            <div class="space-y-2">
                                <p v-for="(station, i) in filteredStations" :key="i"
                                    :class="{ 'bg-indigo-100 text-indigo-700': i === highlightedIndex }"
                                    @click="selectDestinationStation(station)"
                                    class="px-4 py-2 cursor-pointer hover:bg-indigo-100 hover:text-indigo-700 transition-colors duration-200">
                                    {{ station }}
                                </p>
                            </div>
                        </div>
                    </div>
                    <div class="flex flex-col relative w-full sm:w-[280px] space-y-4">
                        <label for="journey" class="text-lg font-medium text-gray-700">Date of Journey:</label>
                        <input type="date"
                            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
                            v-model="journeyDate" :min="today" required />
                    </div>



                    <div class="flex flex-col relative w-full sm:w-[280px] space-y-4" v-if="tripType === 'round-trip'">
                        <label for="return" class="text-lg font-medium text-gray-700">Date of Return:</label>
                        <input type="date"
                            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
                            required v-model="returnDate" :min="journeyDate" />
                    </div>
                    <div class="flex flex-col relative w-full sm:w-[280px] space-y-4">
                        <label for="passengers" class="text-lg font-medium text-gray-700">No. of Passengers:</label>
                        <select
                            class="border relative border-gray-300 rounded-lg px-4 py-3 text-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition duration-200"
                            v-model="passengers" required name="passengers" id="passengers">
                            <option value=1>1</option>
                            <option value=2>2</option>
                            <option value=3>3</option>
                            <option value=4>4</option>
                            <option value=5>5</option>
                            <option value=6>6</option>
                            <option value=7>7</option>
                            <option value=8>8</option>
                            <option value=9>9</option>
                            <option value=10>10</option>
                        </select>
                    </div>


                    <button type="submit"
                        class="bg-blue-500 text-white py-4 px-6 rounded hover:bg-blue-600">Search</button>

                </div>
            </div>
        </form>
        <FlightSearch :flights="flights" :passengers="passengers" />
    </div>
</template>

<script>
import { Button } from '@/vue/ui/button'
import { Input } from '@/vue/ui/input'
import { useFlightStore } from '@/stores/flight'
import { searchFlightApi } from '@/api/flight'
import FlightSearch from '@/components/FlightSearch.vue'

export default {
    name: "SearchComponent",
    components: {
        Button,
        Input,
        FlightSearch

    },
    data() {
        return {
            flightStore: useFlightStore(),
            passengers: 1,
            flights: null,
            source: "",
            sourceSelected: false,
            destinationSelected: false,
            destination: "",
            stations: [],
            tripType: "one-way",
            journeyDate: "",
            returnDate: "",
            filteredStations: [],
            highlightedIndex: -1
        };
    },
    methods: {
        filterSource() {
            if (this.source) {
                this.sourceSelected = false
                this.filteredStations = this.stations.filter(station =>
                    station.toLowerCase().includes(this.source.toLowerCase())
                );
                this.highlightedIndex = -1

            } else {
                this.filteredStations = this.stations;
            }
        },
        filterDestination() {
            if (this.destination) {
                this.destinationSelected = false
                this.filteredStations = this.stations.filter(station =>
                    station.toLowerCase().includes(this.destination.toLowerCase())

                );

            } else {
                this.filteredStations = this.stations;
            }
        },
        selectSourceStation(value) {
            this.source = value
            this.sourceSelected = true
            this.highlightedIndex = -1
        },
        selectDestinationStation(value) {
            this.destination = value
            this.destinationSelected = true
        },
        navigateDown(event) {
            if (this.highlightedIndex < this.filteredStations.length - 1) {
                this.highlightedIndex++;
            }
        },

        navigateUp(event) {
            if (this.highlightedIndex > 0) {
                this.highlightedIndex--
            }
        },
        selectStationFromKeyboard(event) {
            if (this.highlightedIndex !== -1) {
                this.selectSourceStation(this.filteredStations[this.highlightedIndex])
            }
        },
        selectDestinationFromKeyboard(event) {
            if (this.highlightedIndex !== -1) {
                this.selectDestinationStation(this.filteredStations[this.highlightedIndex])
            }
        },
        async SearchFlight() {
            console.log('submit');
            console.log(this.source);
            console.log(this.destination);
            console.log(this.journeyDate);
            console.log(this.passengers);

            const result = await searchFlightApi(this.source, this.destination, this.journeyDate)

            this.flightStore.setFlight(result.data.data);
            this.flights = result.data.data
            console.log(result.data.data);
            console.log(this.flights);

        }

    },
    computed: {
        today() {
            const date = new Date();
            const day = String(date.getDate()).padStart(2, "0");
            const month = String(date.getMonth() + 1).padStart(2, "0");
            const year = date.getFullYear();
            return `${year}-${month}-${day}`;

        }
    },
    async mounted() {
        const result = await this.flightStore.getStations()
        this.stations = result
        this.filteredStations = result
        console.log(result);

    }

}
</script>