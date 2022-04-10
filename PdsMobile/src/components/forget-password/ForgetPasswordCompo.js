import React from "react";
import ForgetPasswordView from "./ForgetPasswordView";
import ForgetPasswordContainer from "./ForgetPasswordContainer";
import ForgetPasswordRepo from "./ForgetPasswordRepo";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {alert} from "../../utils";

const ForgetPasswordCompo = ({visible, setVisible}) => {
  const forgetPasswordRepo = ForgetPasswordRepo(axios, useFetchCall);
  return (
    <ForgetPasswordContainer {...forgetPasswordRepo}>
      <ForgetPasswordView
        visible={visible}
        setVisible={setVisible}
        alert={alert}
      />
    </ForgetPasswordContainer>
  );
};

export default ForgetPasswordCompo;
