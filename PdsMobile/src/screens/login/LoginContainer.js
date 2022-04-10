import React, {useEffect} from "react";

const LoginContainer = ({children, useLogin, useAuth}) => {
  const login = useLogin();
  const auth = useAuth();
  const handleSubmit = credential => {
    login.start(credential);
  };
  useEffect(() => {
    if (login.isSuccess) {
      auth.setUserData({data: login.data.userData});
    }
  }, [login.isSuccess]);
  return React.cloneElement(children, {
    handleSubmit,
    isLoginLoading: login.isLoading,
    isLoginError: login.isError,
    isLoginSuccess: login.isSuccess,
    loginData: login.data,
    loginError: login.error,
    loginReset: login.reset,
  });
};

export default LoginContainer;
