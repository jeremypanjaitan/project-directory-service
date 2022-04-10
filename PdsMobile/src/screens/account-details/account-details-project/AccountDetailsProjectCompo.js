import React from "react";
import {axios} from "../../../config";
import {useFetchCall} from "../../../utils";
import AccountDetailsProjectContainer from "./AccountDetailsProjectContainer";
import AccountDetailsProjectRepo from "./AccountDetailsProjectRepo";
import AccountDetailsProjectView from "./AccountDetailsProjectView";

const AccountDetailsCompo = () => {
  const accountDetailsProjectRepo = AccountDetailsProjectRepo(
    axios,
    useFetchCall,
  );
  return (
    <AccountDetailsProjectContainer {...accountDetailsProjectRepo}>
      <AccountDetailsProjectView />
    </AccountDetailsProjectContainer>
  );
};

export default AccountDetailsCompo;
