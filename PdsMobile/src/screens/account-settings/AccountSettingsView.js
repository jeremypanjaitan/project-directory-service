import React, {useEffect, useState} from "react";
import {ScrollView, StyleSheet, TouchableOpacity, View} from "react-native";
import {
  withTheme,
  Text,
  ActivityIndicator,
  Avatar,
  Button,
  TextInput,
  HelperText,
  Modal,
} from "react-native-paper";
import {launchCamera, launchImageLibrary} from "react-native-image-picker";
import {useFormik} from "formik";
import {DefaultSpinner} from "../../components";
import {PaperDropdown} from "../../components";
import * as Yup from "yup";
import {en} from "../../shared";

const gender = [
  {
    value: "M",
    label: "Male",
  },
  {
    value: "F",
    label: "Female",
  },
];
const AccountSettingsView = ({
  theme,
  handleUploadPicture,
  cloudStorageData,
  cloudStorageIsLoading,
  divisionDataIsLoading,
  divisionData,
  roleDataIsLoading,
  roleData,
  detailAccountDataIsLoading,
  detailAccountData,
  fullName,
  roleName,
  handleUpdateData,
  updateAccountDataIsLoading,
  updateAccountDataIsSuccess,
  updateAccountDataReset,
  alert,
  confirmAlert,
  sendChangePasswordLinkIsLoading,
  sendChangePasswordLinkIsSuccess,
  sendChangePasswordLinkReset,
  handleSendChangePassword,
  updateAccountData,
  setUserData,
  isFocused,
  handleSubmitLogin,
  isLoginLoading,
  isLoginError,
  isLoginSuccess,
  loginData,
  loginError,
  loginReset,
  setVisibleModal,
  visibleModal,
  showModal,
  hideModal,
}) => {
  const {colors} = theme;

  const [isSecureText, setIsSecureText] = useState(true);
  const formikChangePassword = useFormik({
    initialValues: {
      password: "",
    },
    validationSchema: Yup.object({
      password: Yup.string()
        .min(6, "6 min character length")
        .required("Password is required"),
    }),
    onSubmit: handleSubmitLogin,
  });

  if (isLoginError) {
    setVisibleModal(true);
    alert(en.warning, loginError.response.data.description);
    loginReset();
  }

  const formik = useFormik({
    initialValues: {
      fullName: "",
      email: "",
      gender: "",
      division: "",
      role: "",
      biography: "",
      picture: "",
    },
    validationSchema: Yup.object({
      fullName: Yup.string()
        .required("Name is required")
        .max(30, "Maximum 30 characters"),
      email: Yup.string()
        .email("Invalid email format")
        .required("Email is required"),
      gender: Yup.string().required("Gender is required"),
      division: Yup.string().required("Division is required"),
      role: Yup.string().required("Role is required"),
      picture: Yup.mixed(),
      biography: Yup.string().max(150, "Maximum 150 characters"),
    }),
    onSubmit: values => {
      handleUpdateData(values);
    },
  });
  useEffect(() => {
    if (!isFocused) {
      formik.resetForm();
    }
  }, [isFocused]);
  useEffect(() => {
    if (cloudStorageData) {
      formik.setFieldValue("picture", cloudStorageData);
    }
  }, [cloudStorageData]);
  useEffect(() => {
    if (updateAccountDataIsSuccess) {
      alert(en.success, en.successUpdateUserData);
      setUserData(updateAccountData);
      updateAccountDataReset();
    }
  }, [updateAccountDataIsSuccess]);
  useEffect(() => {
    if (sendChangePasswordLinkIsSuccess) {
      alert(en.success, en.successSendPasswordLink);
      sendChangePasswordLinkReset();
    }
  }, [sendChangePasswordLinkIsSuccess]);
  useEffect(() => {
    if (detailAccountData) {
      formik.setFieldValue("fullName", detailAccountData.data.fullName);
      formik.setFieldValue("email", detailAccountData.data.email);
      formik.setFieldValue("gender", detailAccountData.data.gender);
      formik.setFieldValue("division", detailAccountData.data.divisionId);
      formik.setFieldValue("role", detailAccountData.data.roleId);
      formik.setFieldValue("biography", detailAccountData.data.biography);
      formik.setFieldValue("picture", detailAccountData.data.picture);
    }
  }, [detailAccountData]);
  if (
    roleDataIsLoading ||
    divisionDataIsLoading ||
    detailAccountDataIsLoading
  ) {
    return <DefaultSpinner />;
  }

  return (
    <>
      <ScrollView style={{flex: 1}}>
        <View style={{alignItems: "center", padding: 10}}>
          <View
            style={{
              backgroundColor: colors.primary,
              width: 180,
              height: 200,
              justifyContent: "center",
              alignItems: "center",
              borderRadius: 20,
              marginTop: 20,
            }}>
            {cloudStorageIsLoading ? (
              <View style={{marginBottom: 56}}>
                <ActivityIndicator animating={true} color="white" />
              </View>
            ) : (
              <Avatar.Image
                source={{
                  uri:
                    formik.values.picture ||
                    "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_960_720.png",
                }}
                size={80}
              />
            )}
            <View style={{alignItems: "center", marginTop: 15}}>
              <Text style={{fontSize: 20, color: "white"}}>{fullName}</Text>
              <Text style={{fontSize: 15, color: "white"}}>{roleName}</Text>
            </View>
          </View>
          <View
            style={{
              flexDirection: "row",
              marginTop: 10,
            }}>
            <TouchableOpacity
              onPress={async () =>
                await launchImageLibrary(null, async res => {
                  if (!res.didCancel) {
                    if (res.assets[0].fileSize > 10 * 1024 * 1024) {
                      alert(en.error, "Maximum file size 10 MB");
                    } else {
                      handleUploadPicture({
                        name: res.assets[0].fileName,
                        uri: res.assets[0].uri,
                      });
                    }
                  }
                })
              }>
              <Button
                icon="upload"
                mode="outlined"
                compact
                style={{borderRadius: 20}}
              />
            </TouchableOpacity>
            <TouchableOpacity
              onPress={async () =>
                await launchCamera(null, async res => {
                  if (!res.didCancel) {
                    if (res.assets[0].fileSize > 10 * 1024 * 1024) {
                      alert(en.error, "Maximum file size 10 MB");
                    } else {
                      handleUploadPicture({
                        name: res.assets[0].fileName,
                        uri: res.assets[0].uri,
                      });
                    }
                  }
                })
              }>
              <Button
                icon="camera"
                mode="contained"
                compact
                style={{borderRadius: 20, marginLeft: 10}}
              />
            </TouchableOpacity>
            <TouchableOpacity>
              <Button
                icon="lock"
                mode="outlined"
                compact
                style={{borderRadius: 20, marginLeft: 10}}
                loading={sendChangePasswordLinkIsLoading}
                onPress={showModal}
                // onPress={() => {
                //   confirmAlert(en.warning, en.confirmChangePassword, () =>
                //     handleSendChangePassword(),
                //   );
                // }}
              />
            </TouchableOpacity>
          </View>
        </View>
        <View style={{flex: 1, padding: 25}}>
          <TextInput
            label="Full Name"
            mode="outlined"
            style={{height: 50}}
            onChangeText={formik.handleChange("fullName")}
            onBlur={formik.handleBlur("fullName")}
            value={formik.values.fullName}
            error={formik.errors.fullName && formik.touched.fullName}
          />
          <HelperText
            type="error"
            visible={Boolean(
              formik.errors.fullName && formik.touched.fullName,
            )}>
            {formik.errors.fullName}
          </HelperText>
          <TextInput
            label="Email Address"
            mode="outlined"
            style={{height: 50}}
            onChangeText={formik.handleChange("email")}
            onBlur={formik.handleBlur("email")}
            value={formik.values.email}
            error={formik.errors.email && formik.touched.email}
            disabled
          />
          <HelperText
            type="error"
            visible={Boolean(formik.errors.email && formik.touched.email)}>
            {formik.errors.email}
          </HelperText>
          <PaperDropdown
            listValue={gender}
            label="Gender"
            setValue={value => formik.setFieldValue("gender", value)}
            value={formik.values.gender}
            inputProps={{
              error: formik.errors.gender && formik.touched.gender,
              onBlur: formik.handleBlur("gender"),
            }}
          />
          <HelperText
            type="error"
            visible={Boolean(formik.errors.gender && formik.touched.gender)}>
            {formik.errors.gender}
          </HelperText>
          <PaperDropdown
            listValue={roleData}
            label="Role"
            setValue={value => formik.setFieldValue("role", value)}
            value={formik.values.role}
            inputProps={{
              error: formik.errors.role && formik.touched.role,
              onBlur: formik.handleBlur("role"),
            }}
          />
          <HelperText
            type="error"
            visible={Boolean(formik.errors.role && formik.touched.role)}>
            {formik.errors.role}
          </HelperText>
          <PaperDropdown
            listValue={divisionData}
            label="Division"
            setValue={value => formik.setFieldValue("division", value)}
            value={formik.values.division}
            inputProps={{
              error: formik.errors.division && formik.touched.division,
              onBlur: formik.handleBlur("division"),
            }}
          />
          <HelperText
            type="error"
            visible={Boolean(
              formik.errors.division && formik.touched.division,
            )}>
            {formik.errors.division}
          </HelperText>
          <TextInput
            label="Biography"
            mode="outlined"
            multiline
            numberOfLines={5}
            style={{height: 100}}
            onChangeText={formik.handleChange("biography")}
            onBlur={formik.handleBlur("biography")}
            value={formik.values.biography}
            error={formik.errors.biography && formik.touched.biography}
          />
          <HelperText
            type="error"
            visible={Boolean(
              formik.errors.biography && formik.touched.biography,
            )}>
            {formik.errors.biography}
          </HelperText>

          <Button
            mode="contained"
            onPress={formik.handleSubmit}
            loading={updateAccountDataIsLoading}>
            Save
          </Button>
        </View>
      </ScrollView>

      <Modal
        visible={visibleModal}
        onDismiss={hideModal}
        contentContainerStyle={styles.modal}>
        <View style={{padding: 10, marginTop: 20}}>
          <Text
            style={{
              alignSelf: "center",
              fontWeight: "500",
              fontSize: 18,
              marginBottom: 10,
            }}>
            Input your old password
          </Text>
          <TextInput
            style={{height: 50}}
            label="Old password"
            onChangeText={formikChangePassword.handleChange("password")}
            onBlur={formikChangePassword.handleBlur("password")}
            value={formikChangePassword.values.password}
            secureTextEntry={isSecureText}
            mode="outlined"
            error={Boolean(
              formikChangePassword.errors.password &&
                formikChangePassword.touched.password,
            )}
            right={
              <TextInput.Icon
                name="eye"
                onPress={() => setIsSecureText(!isSecureText)}
              />
            }
          />
          <HelperText
            type="error"
            visible={Boolean(
              formikChangePassword.errors.password &&
                formikChangePassword.touched.password,
            )}>
            {formikChangePassword.errors.password}
          </HelperText>
          <Button
            style={{
              alignItems: "center",
              marginTop: 5,
              flexDirection: "row",
              justifyContent: "center",
            }}
            mode="contained"
            onPress={formikChangePassword.handleSubmit}>
            Submit
          </Button>
        </View>
      </Modal>
    </>
  );
};

const styles = StyleSheet.create({
  modal: {
    backgroundColor: "white",
    padding: 20,
    width: 300,
    alignSelf: "center",
  },
});

export default withTheme(AccountSettingsView);
