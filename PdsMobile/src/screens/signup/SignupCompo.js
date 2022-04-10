import SignupContainer from "./SignupContainer";
import SignupView from "./SignupView";
import SignupRepo from "./SignupRepo";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {alert} from "../../utils/alert";
import {goToLogin} from "../../navigator";
import React from "react";

const SignupCompo = () => {
  const signupRepo = SignupRepo(axios, useFetchCall);
  return (
    <SignupContainer {...signupRepo}>
      <SignupView alert={alert} goToLogin={goToLogin} />
    </SignupContainer>
  );
};

export default SignupCompo;
