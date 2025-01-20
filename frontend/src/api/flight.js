import axios from "axios";
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

export function bookFlights(bookFlightObject) {
  return axiosInstance.post("/books", {
    ...bookFlightObject,
  });
}

export function getFlightsByReferenceNo(referenceNo) {
  return axiosInstance.get(`/books/${referenceNo}`);
}

export function downloadTickets(referenceNo) {
  return axios.get(`http://localhost:8080/pdf/${referenceNo}`, {
    responseType: "blob",
    headers: {
      Accept: "application/pdf",
    },
  });
}

export function sendEmailApi(email, referenceNo) {
  return axiosInstance.post("/email", {
    email,
    referenceNo,
  });
}
