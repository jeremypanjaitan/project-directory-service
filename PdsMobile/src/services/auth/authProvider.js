import React, {useState} from "react";
import {AuthContext} from "./authContext";
const AuthProvider = ({children}) => {
  const [userData, setUserData] = useState(false);

  let value = {userData, setUserData};

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};
export default AuthProvider;
