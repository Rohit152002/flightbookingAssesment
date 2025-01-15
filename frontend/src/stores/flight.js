import { getStationApi } from "@/api/flight";
import { defineStore } from "pinia";

export const useFlightStore = defineStore("flightStation", {
  state: () => ({
    stations: [],
    flights: [],
    bookedFlight: {},
    departureFlights: [],
    returnFlights: [],
    selectedDepartureFlights: {},
    selectedReturnFlights: {},
    selectedJourneyPrice: 0,
    selectedReturnPrice: 0,
    selectedBookFlights: {},
  }),
  actions: {
    async getStations() {
      try {
        const res = await getStationApi();
        console.log(res);
        this.stations = await res.data.data;
        return this.stations;
      } catch (err) {
        console.log(err);
        this.error = err;
      }
    },

    setFlight(value) {
      this.flights = value;
    },
    setDepartureFlights(value) {
      this.selectedDepartureFlights = value;
    },
    setReturnFlights(value) {
      this.selectedReturnFlights = value;
    },
    setJourneyPrice(value) {
      this.selectedJourneyPrice = value;
    },
    setReturnPrice(value) {
      this.selectedReturnPrice = value;
    },
    setSelectedBookFlights(value) {
      this.selectedBookFlights = value;
    },
  },
});
