import axiosInstance from "./interceptor";

export function getStationApi() {
  return axiosInstance.get("/station");
}

export function searchFlightApi(source, destination, date) {
  return axiosInstance.get("/flights", {
    params: {
      source,
      destination,
      date,
    },
  });
}
