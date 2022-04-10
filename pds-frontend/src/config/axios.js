import axios from "axios";
import { API_AUTH, API_REFRESH_TOKEN } from "./apiRoutes";
import { customHistory, LOGIN } from "../router";
import {
  getAuthorizationToken,
  getRefreshToken,
  setRefreshToken,
  setToken,
} from "../utils";
const instance = axios.create({
  baseURL: process.env.REACT_APP_API_BASE_URL,
});
instance.interceptors.request.use(
  async (config) => {
    const token = getAuthorizationToken();
    if (token) {
      config.headers["Authorization"] = token;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);
instance.interceptors.response.use(
  (response) => {
    return response;
  },
  async (error) => {
    const originalRequest = error.config;

    if (
      error.response.status === 401 &&
      originalRequest.url === API_REFRESH_TOKEN
    ) {
      customHistory.push(LOGIN);
      return Promise.reject(error);
    }
    if (error.response.status === 401 && originalRequest.url !== API_AUTH) {
      return instance
        .post(API_REFRESH_TOKEN, {
          refreshToken: getRefreshToken(),
        })
        .then((res) => {
          if (res.status === 200) {
            setRefreshToken(res.data.data.newRefreshToken);
            setToken(res.data.data.accessToken);
            instance.defaults.headers.common["Authorization"] =
              getAuthorizationToken();
            return instance(originalRequest);
          }
        });
    }
    return Promise.reject(error);
  }
);

export default instance;
