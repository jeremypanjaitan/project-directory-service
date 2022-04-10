import React from "react";
export const GlobalContext = React.createContext(null);
export const useGlobal = () => {
  return React.useContext(GlobalContext);
};
