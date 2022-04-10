import React from "react";

const ForgetPasswordContainer = ({children, useSendForgetPasswordLink}) => {
  const sendForgetPasswordLink = useSendForgetPasswordLink();
  const handleSubmit = values => {
    sendForgetPasswordLink.start(values);
  };
  return React.cloneElement(children, {
    handleSubmit,
    sendForgetPasswordLinkIsSuccess: sendForgetPasswordLink.isSuccess,
    sendForgetPasswordLinkIsError: sendForgetPasswordLink.isError,
    sendForgetPasswordLinkError: sendForgetPasswordLink.error,
    sendForgetPasswordLinkData: sendForgetPasswordLink.data,
    sendForgetPasswordLinkIsLoading: sendForgetPasswordLink.isLoading,
    sendForgetPasswordLinkReset: sendForgetPasswordLink.reset,
  });
};

export default ForgetPasswordContainer;
