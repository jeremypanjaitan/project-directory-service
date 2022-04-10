import React, {useEffect, useState} from "react";
import {useIsFocused} from "@react-navigation/native";
import {en} from "../../shared";
import {confirmAlert} from "../../utils";

const AccountSettingsContainer = ({
  children,
  useCloudStorage,
  useRoleData,
  useDivisionData,
  useDetailAccountData,
  useAuth,
  useLogin,
  useUpdateAccountData,
  useSendChangePasswordLink,
}) => {
  const detailAccountData = useDetailAccountData();
  const cloudStorage = useCloudStorage();
  const roleData = useRoleData();
  const divisionData = useDivisionData();
  const isFocused = useIsFocused();
  const [visibleModal, setVisibleModal] = useState(false);
  const showModal = () => setVisibleModal(true);
  const hideModal = () => setVisibleModal(false);

  const updateAccountData = useUpdateAccountData();
  const sendChangePasswordLink = useSendChangePasswordLink();
  const login = useLogin();
  const auth = useAuth();
  const handleSendChangePassword = () => {
    sendChangePasswordLink.start(auth?.userData?.data?.email);
  };
  const handleSubmitLogin = value => {
    confirmAlert(en.warning, en.confirmChangePassword, () =>
      login.start({
        email: auth?.userData?.data?.email,
        password: value.password,
      }),
    );
  };

  if (login.isSuccess) {
    console.log("sucess");
    setVisibleModal(false);
    sendChangePasswordLink.start(auth?.userData?.data?.email);
    login.reset();
  }

  useEffect(() => {
    if (isFocused) {
      divisionData.start();
      roleData.start();
      detailAccountData.start();
    } else {
      divisionData.reset();
      roleData.reset();
      detailAccountData.reset();
    }
  }, [isFocused]);
  const handleUploadPicture = picture => {
    cloudStorage.start(picture);
  };
  const handleUpdateData = data => {
    updateAccountData.start(data);
  };

  return React.cloneElement(children, {
    handleUploadPicture,
    cloudStorageData: cloudStorage.data,
    cloudStorageIsLoading: cloudStorage.isLoading,
    detailAccountDataIsLoading: detailAccountData.isLoading,
    detailAccountData: detailAccountData.data,
    roleDataIsLoading: roleData.isLoading ?? true,
    roleData: roleData.data?.map(d => ({value: d.id, label: d.name})),
    divisionDataIsLoading: divisionData.isLoading ?? true,
    divisionData: divisionData.data?.map(d => ({value: d.id, label: d.name})),
    fullName: auth?.userData?.data?.fullName,
    roleName: auth?.userData?.data?.roleName,
    handleUpdateData: handleUpdateData,
    updateAccountDataIsLoading: updateAccountData.isLoading,
    updateAccountDataIsSuccess: updateAccountData.isSuccess,
    updateAccountDataReset: updateAccountData.reset,
    updateAccountData: updateAccountData.data,
    sendChangePasswordLinkIsLoading: sendChangePasswordLink.isLoading,
    sendChangePasswordLinkIsSuccess: sendChangePasswordLink.isSuccess,
    sendChangePasswordLinkReset: sendChangePasswordLink.reset,
    handleSendChangePassword: handleSendChangePassword,
    setUserData: auth.setUserData,
    isFocused: isFocused,
    handleSubmitLogin: handleSubmitLogin,
    isLoginLoading: login.isLoading,
    isLoginError: login.isError,
    isLoginSuccess: login.isSuccess,
    loginData: login.data,
    loginError: login.error,
    loginReset: login.reset,
    visibleModal: visibleModal,
    setVisibleModal: setVisibleModal,
    showModal: showModal,
    hideModal: hideModal,
  });
};

export default AccountSettingsContainer;
