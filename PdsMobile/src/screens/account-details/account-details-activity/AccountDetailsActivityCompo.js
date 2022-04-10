import React from "react";
import {axios} from "../../../config";

import {useFetchCall} from "../../../utils";
import AccountDetailsActivityContainer from "./AccountDetailsActivityContainer";
import AccountDetailsActivityRepo from "./AccountDetailsActivityRepo";
import AccountDetailsActivityView from "./AccountDetailsActivityView";

const AccountDetailsActivityCompo = () => {
  const accountDetailsActivityRepo = AccountDetailsActivityRepo(
    axios,
    useFetchCall,
  );
  return (
    <AccountDetailsActivityContainer {...accountDetailsActivityRepo}>
      <AccountDetailsActivityView />
    </AccountDetailsActivityContainer>
  );
};

export default AccountDetailsActivityCompo;
