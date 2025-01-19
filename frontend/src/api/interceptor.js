import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://localhost:8080",
  headers: {
    Accept: "application/pdf",
  },
  //   responseType: "blob", // Add this for PDF data
});

export default axiosInstance;
