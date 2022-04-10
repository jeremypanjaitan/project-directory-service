import React, {useEffect} from "react";
import {
  Modal,
  Portal,
  TextInput,
  HelperText,
  Text,
  Button,
} from "react-native-paper";
import {useFormik} from "formik";
import * as Yup from "yup";
import {en} from "../../shared";

const ForgetPasswordView = ({
  visible,
  setVisible,
  handleSubmit,
  sendForgetPasswordLinkIsSuccess,
  sendForgetPasswordLinkIsError,
  sendForgetPasswordLinkError,
  sendForgetPasswordLinkData,
  sendForgetPasswordLinkIsLoading,
  sendForgetPasswordLinkReset,
  alert,
}) => {
  const formik = useFormik({
    initialValues: {
      email: "",
    },
    validationSchema: Yup.object({
      email: Yup.string()
        .email("Invalid email format")
        .required("Email is required"),
    }),
    onSubmit: handleSubmit,
  });
  const hideModal = () => setVisible(false);
  const containerStyle = {backgroundColor: "white", padding: 20};
  useEffect(() => {
    if (sendForgetPasswordLinkError) {
      alert(en.error, sendForgetPasswordLinkError.response.data.description);
      sendForgetPasswordLinkReset();
    }
  }, [sendForgetPasswordLinkError]);

  useEffect(() => {
    if (sendForgetPasswordLinkIsSuccess) {
      alert(en.success, sendForgetPasswordLinkData.data.message);
      sendForgetPasswordLinkReset();
      formik.resetForm();
      setVisible(false);
    }
  }, [sendForgetPasswordLinkIsSuccess]);

  return (
    <Portal>
      <Modal
        style={{padding: 20}}
        visible={visible}
        onDismiss={!sendForgetPasswordLinkIsLoading && hideModal}
        contentContainerStyle={containerStyle}>
        <Text>Link to change password will be sent to your email</Text>
        <TextInput
          style={{height: 40, marginTop: 10}}
          label="Email"
          mode="outlined"
          onChangeText={formik.handleChange("email")}
          onBlur={formik.handleBlur("email")}
          value={formik.values.email}
          error={formik.errors.email && formik.touched.email}
        />
        <HelperText
          type="error"
          visible={Boolean(formik.errors.email && formik.touched.email)}>
          {formik.errors.email}
        </HelperText>
        <Button
          mode="contained"
          onPress={formik.handleSubmit}
          loading={sendForgetPasswordLinkIsLoading}>
          Send
        </Button>
      </Modal>
    </Portal>
  );
};

export default ForgetPasswordView;
