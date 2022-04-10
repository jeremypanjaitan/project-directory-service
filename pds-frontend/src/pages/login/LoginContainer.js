import { useState } from "react";
import Swal from "sweetalert2";
import { useLoading } from "../../context";
import { useAuth, useLogin } from "../../services/auth";
import { usePostForgetPasswordData } from "./LoginRepo";
import LoginPageView from "./LoginView";

const LoginContainer = () => {
  const login = useLogin();
  const auth = useAuth();
  const postPasswordData = usePostForgetPasswordData();
  const [isModalInputEmailVisible, setIsModalInputEmailVisible] =
    useState(false);
  const { setIsLoading } = useLoading();

  const handleSubmitLogin = (values) => {
    console.log("values login", values);
    login.start(values);
  };

  if (login.isSuccess) {
    auth.setUser(login.data.data.userData);
  }

  const handleSubmitForgetPassword = (values) => {
    console.log("values forget", values);
    Swal.fire({
      title: "Are you sure want to set your password?",
      showCancelButton: true,
      confirmButtonText: "Yes",
    }).then((result) => {
      if (result.isConfirmed) {
        setIsLoading(true);
        postPasswordData.start(values);
        setIsModalInputEmailVisible(false);
      }
    });
  };

  if (postPasswordData.isSuccess) {
    Swal.fire("Please check your email!", "", "success");
    setIsLoading(false);
    postPasswordData.reset();
    setIsModalInputEmailVisible(false);
  }
  // if (postPasswordData.isError) {
  //   Swal.fire(postPasswordData.error.response.data.description, "", "warning");
  //   postPasswordData.reset();
  //   setIsModalInputEmailVisible(true);
  //   setIsLoading(false);
  // }

  return (
    <>
      <LoginPageView
        handleSubmitLogin={handleSubmitLogin}
        handleSubmitForgetPassword={handleSubmitForgetPassword}
        isLoading={login.isLoading}
        isError={login.isError}
        loginReset={login.reset}
        loginError={login.error}
        postPasswordData={postPasswordData}
        postForgetPasswordError={postPasswordData.error}
        postForgetPasswordReset={postPasswordData.reset}
        InputEmailVisible={isModalInputEmailVisible}
        setInputEmailVisible={setIsModalInputEmailVisible}
      />
    </>
  );
};

export default LoginContainer;
