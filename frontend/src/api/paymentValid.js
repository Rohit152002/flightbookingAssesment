import axiosInstance from "./interceptor";

export function paymentValidation(params) {
  return axiosInstance.post("/verification", params);
}
