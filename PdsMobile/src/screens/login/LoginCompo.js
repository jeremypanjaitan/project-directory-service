import React from "react";
import LoginContainer from "./LoginContainer";
import LoginView from "./LoginView";
import LoginRepo from "./LoginRepo";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {useAuth} from "../../services";
import {alert} from "../../utils/alert";
import {goToRegister} from "../../navigator/NavigationHelper";

const LoginCompo = () => {
  const {useLogin} = LoginRepo(axios, useFetchCall);
  return (
    <LoginContainer useLogin={useLogin} useAuth={useAuth}>
      <LoginView alert={alert} goToRegister={goToRegister} />
    </LoginContainer>
  );
};

export default LoginCompo;
