import AccountSettingForm from "../account-setting-form/AccountSettingFormContainer";

const AccountSettingUpdateView = ({
  defaultData,
  handleSubmitData,
  handleSubmitPassword,
  handleSubmitOldPassword,
  passwordVisible,
  setPasswordVisible,
  isError,
  isSuccess,
  loginReset,
}) => {
  return (
    <>
      <AccountSettingForm
        handleSubmitData={handleSubmitData}
        handleSubmitPassword={handleSubmitPassword}
        handleSubmitOldPassword={handleSubmitOldPassword}
        defaultData={defaultData}
        passwordVisible={passwordVisible}
        setPasswordVisible={setPasswordVisible}
        isSuccess={isSuccess}
        isError={isError}
        loginReset={loginReset}
      />
    </>
  );
};

export default AccountSettingUpdateView;
