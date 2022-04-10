import React, { useEffect } from "react";
import { useLocation, Navigate } from "react-router-dom";
import { AuthContext } from "../context";
import { LOGIN, DEFAULT } from "../router";
import { axios, API_AUTH, API_PROFILE, API_LOGOUT } from "../config";
import { useFetchCall } from "../hooks";
import {
  setToken,
  deleteToken,
  deleteRefreshToken,
  setRefreshToken,
  getToken,
} from "../utils";

export const useAuth = () => {
  return React.useContext(AuthContext);
};

export const AuthProvider = ({ children }) => {
  let [user, setUser] = React.useState();

  let value = { user, setUser };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

export const RequireAuth = ({
  children,
  _useAuth = useAuth,
  _useLocation = useLocation,
}) => {
  let auth = _useAuth();
  let location = _useLocation();

  if (!auth?.user) {
    // Redirect them to the /login page, but save the current location they were
    // trying to go to when they were redirected. This allows us to send them
    // along to that page after they login, which is a nicer user experience
    // than dropping them off on the home page.
    return <Navigate to={LOGIN} state={{ from: location }} replace />;
  }

  return children;
};

export const AlreadyAuth = ({
  children,
  _useAuth = useAuth,
  _useLocation = useLocation,
}) => {
  let auth = _useAuth();
  let location = _useLocation();
  let from = location.state?.from?.pathname || DEFAULT;
  if (auth?.user && getToken()) {
    //if user is authenticated then
    //redirect to user previous page or to home page
    return <Navigate to={from} state={{ from: location }} replace />;
  }
  return children;
};

export const useCheckAuth = () => {
  const login = async () => {
    try {
      const res = await axios.get(API_PROFILE);
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(login);
};

export const CheckAuth = ({
  children,
  _useCheckAuth = useCheckAuth,
  _useAuth = useAuth,
}) => {
  const checkAuth = _useCheckAuth();
  const auth = _useAuth();
  useEffect(() => {
    checkAuth.start();
    //eslint-disable-next-line
  }, []);
  useEffect(() => {
    if (checkAuth.isSuccess) {
      auth.setUser(checkAuth.data.data);
    }
    //eslint-disable-next-line
  }, [checkAuth.isSuccess]);
  return checkAuth.isLoading ? null : children;
};

export const useLogin = () => {
  const login = async (credential) => {
    try {
      const res = await axios.post(API_AUTH, credential);
      const data = res.data;
      setToken(data.data.tokenData.accessToken);
      setRefreshToken(data.data.tokenData.refreshToken);
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(login);
};

export const useLogout = () => {
  const login = async () => {
    try {
      const res = await axios.post(API_LOGOUT, null);
      const data = res.data;
      deleteToken();
      deleteRefreshToken();
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(login);
};
