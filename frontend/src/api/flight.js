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

export function searchFlightWithReturnDate(
  source,
  destination,
  departureDate,
  returnDate
) {
  return axiosInstance.get("/flights/two-way", {
    params: {
      source,
      destination,
      "departure-date": departureDate,
      "return-date": returnDate,
    },
  });
}
