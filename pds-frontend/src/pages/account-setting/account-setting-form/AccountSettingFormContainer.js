import { useEffect, useState } from "react";
import { useLoading } from "../../../context";
import { useDivisionData, useRoleData } from "./AccountSettingFormRepo";
import AccountSettingFormView from "./AccountSettingFormView";

const AccountSettingFormContainer = ({
  handleSubmitData,
  handleSubmitPassword,
  handleSubmitOldPassword,
  defaultData,
  passwordVisible,
  setPasswordVisible,
  isError,
  isSuccess,
  loginReset,
}) => {
  const roleData = useRoleData();
  const divisionData = useDivisionData();
  const [division, setDivision] = useState();
  const [role, setRole] = useState();
  const { setIsLoading } = useLoading();
  useEffect(() => {
    setIsLoading(true);
    roleData.start();
    divisionData.start();
    //eslint-disable-next-line
  }, []);

  useEffect(() => {
    if (roleData.isSuccess && divisionData.isSuccess) {
      setRole(roleData.data);
      setDivision(divisionData.data);
      roleData.reset();
      divisionData.reset();
      setIsLoading(false);
    }
    //eslint-disable-next-line
  }, [roleData.isSuccess, divisionData.isSuccess, setIsLoading]);

  return (
    <>
      <AccountSettingFormView
        handleSubmitData={handleSubmitData}
        handleSubmitPassword={handleSubmitPassword}
        handleSubmitOldPassword={handleSubmitOldPassword}
        defaultData={defaultData}
        roles={role || []}
        divisions={division || []}
        passwordVisible={passwordVisible}
        setPasswordVisible={setPasswordVisible}
        isSuccess={isSuccess}
        isError={isError}
        loginReset={loginReset}
      />
    </>
  );
};

export default AccountSettingFormContainer;
