import { useNavigate } from "react-router-dom";
import {
  useUpdateAccountData,
  usePostChangePasswordData,
} from "./AccountSettingUpdateRepo";
import { useLoading } from "../../../context";
import { useEffect, useState } from "react";
import AccountSettingUpdateView from "./AccountSettingUpdateView";
import { useCheckAuth, useAuth, useLogin } from "../../../services";
import Swal from "sweetalert2";

const AccountSettingUpdateContainer = () => {
  const { setUser } = useAuth();
  const auth = useAuth();
  const detailAccountData = useCheckAuth();
  const updateAccountData = useUpdateAccountData();
  const changePasswordData = usePostChangePasswordData();
  const navigate = useNavigate();
  const { setIsLoading } = useLoading();
  const [isChangePasswordVisible, setIsChangePasswordVisible] = useState(false);
  const login = useLogin();

  const handleSubmitChangePassword = () => {
    changePasswordData.start({ email: auth?.user?.email });
  };

  const handleSubmitOldPassword = (values) => {
    console.log({ email: auth?.user?.email, password: values.password });

    Swal.fire({
      title: "Are you sure want to change your password?",
      showCancelButton: true,
      confirmButtonText: "Yes",
    }).then((result) => {
      if (result.isConfirmed) {
        login.start({ email: auth?.user?.email, password: values.password });
      }
    });
  };

  if (login.isSuccess) {
    setIsChangePasswordVisible(false);
    setIsLoading(true);
    changePasswordData.start({ email: auth?.user?.email });
    login.reset();
  }

  const handleSubmitData = async (values) => {
    Swal.fire({
      title: "Are you sure want to update this data?",
      showCancelButton: true,
      confirmButtonText: "Update",
    }).then((result) => {
      if (result.isConfirmed) {
        setIsLoading(true);
        try {
          updateAccountData.start(values);
        } catch (err) {
          Swal.fire("Error to save data", "", "warning");
        } finally {
          setIsLoading(false);
        }
      } else if (result.isDenied) {
        Swal.fire("Changes are not saved", "", "info");
      }
    });
  };
  useEffect(() => {
    setIsLoading(true);
    detailAccountData.start();
    //eslint-disable-next-line
  }, []);

  if (updateAccountData.isSuccess) {
    setIsLoading(false);
    Swal.fire("Success to update data!", "", "success");
    setUser(updateAccountData.data.data);
    updateAccountData.reset();
    changePasswordData.reset();
    navigate(-1);
  }

  if (changePasswordData.isSuccess) {
    setIsLoading(true);
    Swal.fire("Please check your email!", "", "success");
    setIsLoading(false);
    changePasswordData.reset();
    setIsChangePasswordVisible(false);
  }
  if (changePasswordData.isError) {
    changePasswordData.reset();
    Swal.fire(
      changePasswordData.error.response.data.description,
      "",
      "warning"
    );
    setIsLoading(false);
  }

  return (
    <>
      <AccountSettingUpdateView
        defaultData={detailAccountData.data}
        handleSubmitData={handleSubmitData}
        handleSubmitPassword={handleSubmitChangePassword}
        handleSubmitOldPassword={handleSubmitOldPassword}
        passwordVisible={isChangePasswordVisible}
        setPasswordVisible={setIsChangePasswordVisible}
        isError={login.isError}
        isSuccess={login.isSuccess}
        loginReset={login.reset}
      />
    </>
  );
};

export default AccountSettingUpdateContainer;
