import { TOKEN, REFRESH_TOKEN } from "../constants";
export const setToken = (token) => {
  window.localStorage.setItem(TOKEN, token);
};
export const setRefreshToken = (refreshToken) => {
  window.localStorage.setItem(REFRESH_TOKEN, refreshToken);
};

export const getToken = () => {
  return window.localStorage.getItem(TOKEN);
};

export const getAuthorizationToken = () => {
  return `Bearer ${getToken()}`;
};

export const deleteToken = () => {
  return window.localStorage.removeItem(TOKEN);
};
export const deleteRefreshToken = () => {
  return window.localStorage.removeItem(REFRESH_TOKEN);
};
export const getRefreshToken = () => {
  return window.localStorage.getItem(REFRESH_TOKEN);
};
