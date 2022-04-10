import React from "react";
export const AuthContext = React.createContext(null);
export const useAuth = () => {
  return React.useContext(AuthContext);
};
