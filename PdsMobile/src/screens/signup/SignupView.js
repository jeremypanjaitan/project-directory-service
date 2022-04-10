import React, {useEffect, useState} from "react";
import {ScrollView, View} from "react-native";
import {TextInput, Button, HelperText} from "react-native-paper";
import {DefaultSpinner, PaperDropdown} from "../../components";
import {en} from "../../shared";
import {useFormik} from "formik";
import * as Yup from "yup";

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

const SignupView = ({
  alert,
  divisionData,
  divisionDataIsError,
  divisionDataError,
  divisionDataIsLoading,
  roleData,
  roleDataIsError,
  roleDataError,
  roleDataIsLoading,
  handleSubmit,
  registerIsLoading,
  registerIsError,
  registerError,
  registerReset,
  goToLogin,
  registerIsSuccess,
}) => {
  const [isSecurePassword, setIsSecurePassword] = useState(true);
  const [isSecureReEnterPassword, setIsSecureReEnterPassword] = useState(true);
  const formik = useFormik({
    initialValues: {
      fullName: "",
      email: "",
      gender: "",
      divisionId: "",
      roleId: "",
      password: "",
    },
    validationSchema: Yup.object({
      fullName: Yup.string().required("This field is required"),
      email: Yup.string()
        .required("This field is required")
        .email("Must be email format"),
      gender: Yup.string().required("This field is required"),
      divisionId: Yup.string().required("This field is required"),
      roleId: Yup.string().required("This field is required"),
      password: Yup.string()
        .required("This field is required")
        .min(6, "Minimum 6 character length"),
      reEnterPassword: Yup.string()
        .required("Re-Enter New Password is required")
        .min(6, "Minimum 6 characters")
        .oneOf([Yup.ref("password"), null], "Passwords must match"),
    }),
    onSubmit: handleSubmit,
  });
  useEffect(() => {
    if (registerIsError) {
      alert(en.error, registerError.response.data.description);
      registerReset();
    }
  }, [registerIsError]);
  useEffect(() => {
    if (registerIsSuccess) {
      alert(en.success, en.successCreateUser);
      goToLogin();
    }
  }, [registerIsSuccess]);
  if (roleDataIsLoading || divisionDataIsLoading) {
    return <DefaultSpinner />;
  }
  if (divisionDataIsError) {
    alert(en.error, divisionDataError.response);
  }
  if (roleDataIsError) {
    alert(en.error, roleDataError.response);
  }
  return (
    <ScrollView style={{flex: 1, padding: 25}}>
      <View style={{marginBottom: 80}}>
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
          visible={Boolean(formik.errors.fullName && formik.touched.fullName)}>
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
          setValue={value => formik.setFieldValue("roleId", value)}
          value={formik.values.roleId}
          inputProps={{
            error: formik.errors.roleId && formik.touched.roleId,
            onBlur: formik.handleBlur("roleId"),
          }}
        />
        <HelperText
          type="error"
          visible={Boolean(formik.errors.roleId && formik.touched.roleId)}>
          {formik.errors.roleId}
        </HelperText>
        <PaperDropdown
          listValue={divisionData}
          label="Division"
          setValue={value => formik.setFieldValue("divisionId", value)}
          value={formik.values.divisionId}
          inputProps={{
            error: formik.errors.divisionId && formik.touched.divisionId,
            onBlur: formik.handleBlur("divisionId"),
          }}
        />
        <HelperText
          type="error"
          visible={Boolean(
            formik.errors.divisionId && formik.touched.divisionId,
          )}>
          {formik.errors.divisionId}
        </HelperText>
        <TextInput
          label="Password"
          mode="outlined"
          style={{height: 50}}
          onChangeText={formik.handleChange("password")}
          secureTextEntry={isSecurePassword}
          right={
            <TextInput.Icon
              name="eye"
              onPress={() => setIsSecurePassword(!isSecurePassword)}
            />
          }
          onBlur={formik.handleBlur("password")}
          value={formik.values.password}
          error={formik.errors.password && formik.touched.password}
        />
        <HelperText
          type="error"
          visible={Boolean(formik.errors.password && formik.touched.password)}>
          {formik.errors.password}
        </HelperText>
        <TextInput
          label="Re-Enter Password"
          mode="outlined"
          style={{height: 50}}
          onChangeText={formik.handleChange("reEnterPassword")}
          onBlur={formik.handleBlur("reEnterPassword")}
          value={formik.values.reEnterPassword}
          error={
            formik.errors.reEnterPassword && formik.touched.reEnterPassword
          }
          secureTextEntry={isSecureReEnterPassword}
          right={
            <TextInput.Icon
              name="eye"
              onPress={() =>
                setIsSecureReEnterPassword(!isSecureReEnterPassword)
              }
            />
          }
        />
        <HelperText
          type="error"
          visible={Boolean(
            formik.errors.reEnterPassword && formik.touched.reEnterPassword,
          )}>
          {formik.errors.reEnterPassword}
        </HelperText>

        <Button
          mode="contained"
          onPress={formik.handleSubmit}
          loading={registerIsLoading}>
          Sign up Now
        </Button>
      </View>
    </ScrollView>
  );
};

export default SignupView;
