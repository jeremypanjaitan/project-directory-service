import AccountSettingsView from "./AccountSettingsView";
import AccountSettingsContainer from "./AccountSettingsContainer";
import AccountSettingsRepo from "./AccountSettingsRepo";
import {alert, confirmAlert, useCloudStorage} from "../../utils";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {useAuth} from "../../services";

import React from "react";
import LoginRepo from "../login/LoginRepo";

const AccountSettingsCompo = props => {
  const accountSettingsRepo = AccountSettingsRepo(axios, useFetchCall);
  const {useLogin} = LoginRepo(axios, useFetchCall);

  return (
    <AccountSettingsContainer
      useCloudStorage={useCloudStorage}
      useAuth={useAuth}
      {...props}
      {...accountSettingsRepo}
      useLogin={useLogin}>
      <AccountSettingsView alert={alert} confirmAlert={confirmAlert} />
    </AccountSettingsContainer>
  );
};

export default AccountSettingsCompo;
