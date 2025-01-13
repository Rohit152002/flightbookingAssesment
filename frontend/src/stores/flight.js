import { getStationApi } from "@/api/flight";
import { defineStore } from "pinia";

export const useFlightStore = defineStore("flightStation", {
  state: () => ({
    stations: [],
    flights: [],
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
  },
});
