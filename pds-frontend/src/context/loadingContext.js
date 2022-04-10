import { useState, useContext, createContext } from "react";
export const LoadingContext = createContext(null);

export const useLoading = () => {
  return useContext(LoadingContext);
};

export const LoadingProvider = ({ children, _useState = useState }) => {
  let [isLoading, setIsLoading] = useState();

  let value = { isLoading, setIsLoading };

  return (
    <LoadingContext.Provider value={value}>{children}</LoadingContext.Provider>
  );
};
