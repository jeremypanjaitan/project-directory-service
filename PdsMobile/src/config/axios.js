import axios from "axios";
import {API_BASE_URL} from "@env";
import {getAuthToken} from "../utils";

const instance = axios.create({
  baseURL: API_BASE_URL,
});
instance.interceptors.request.use(
  async config => {
    const token = await getAuthToken();
    if (token) {
      config.headers["Authorization"] = "Bearer " + token; // for Spring Boot back-end
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  },
);

export default instance;
