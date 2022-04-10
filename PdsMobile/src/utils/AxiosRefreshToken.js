import {useEffect} from "react";
import {API_AUTH, API_REFRESH_TOKEN, axios} from "../config";
import {useAuth} from "../services";
import {
  deleteAuthToken,
  deleteRefreshToken,
  getAuthToken,
  getRefreshToken,
  saveAuthToken,
  saveRefreshToken,
} from "./authToken";
const AxiosRefreshToken = ({children}) => {
  const auth = useAuth();
  useEffect(() => {
    axios.interceptors.response.use(
      response => {
        return response;
      },
      async error => {
        const originalRequest = error.config;

        if (
          error.response.status === 401 &&
          originalRequest.url === API_REFRESH_TOKEN
        ) {
          auth.setUserData();
          await deleteAuthToken();
          await deleteRefreshToken();
          return Promise.reject(error);
        }
        const refreshToken = await getRefreshToken();
        if (error.response.status === 401 && originalRequest.url !== API_AUTH) {
          return axios
            .post(API_REFRESH_TOKEN, {
              refreshToken: refreshToken,
            })
            .then(async res => {
              if (res.status === 200) {
                await saveRefreshToken(res.data.data.newRefreshToken);
                await saveAuthToken(res.data.data.accessToken);
                const newRefreshToken = await getAuthToken();
                axios.defaults.headers.common["Authorization"] =
                  newRefreshToken;
                return axios(originalRequest);
              }
            });
        }
        return Promise.reject(error);
      },
    );
  }, []);

  return children;
};

export default AxiosRefreshToken;
