import React, {useState} from "react";
import {GlobalContext} from "./globalContext";
const GlobalProvider = ({children}) => {
  const [totalComment, setTotalComment] = useState();

  let value = {totalComment, setTotalComment};

  return (
    <GlobalContext.Provider value={value}>{children}</GlobalContext.Provider>
  );
};
export default GlobalProvider;
