import React, {useState, useEffect} from "react";
import {
  TextInput,
  Button,
  HelperText,
  Text,
  withTheme,
} from "react-native-paper";
import {View, Image, TouchableOpacity} from "react-native";
import {useFormik} from "formik";
import {en} from "../../shared";
import {ForgetPassword} from "../../components";
import * as Yup from "yup";

const LoginView = ({
  handleSubmit,
  isLoginLoading,
  isLoginError,
  alert,
  loginError,
  loginReset,
  theme,
  goToRegister,
}) => {
  const {colors} = theme;
  const [isSecureText, setIsSecureText] = useState(true);
  const [forgetPasswordIsVisible, setForgetPasswordIsVisible] = useState(false);
  const formik = useFormik({
    initialValues: {
      email: "",
      password: "",
    },
    validationSchema: Yup.object({
      email: Yup.string()
        .email("Invalid email format")
        .required("Email is required"),
      password: Yup.string()
        .min(6, "6 min character length")
        .required("Password is required"),
    }),
    onSubmit: handleSubmit,
  });
  useEffect(() => {
    if (isLoginError) {
      alert(en.warning, loginError.response.data.description);
      loginReset();
    }
  }, [isLoginError]);
  return (
    <>
      <ForgetPassword
        visible={forgetPasswordIsVisible}
        setVisible={setForgetPasswordIsVisible}
      />
      <View style={{flex: 1, padding: 25}}>
        <View style={{alignItems: "center", marginTop: 50, marginBottom: 15}}>
          <Image source={require("../../assets/Logo.png")} />
        </View>
        <View style={{marginBottom: 15}}>
          <TextInput
            style={{height: 50}}
            label="Email"
            onChangeText={formik.handleChange("email")}
            onBlur={formik.handleBlur("email")}
            value={formik.values.email}
            mode="outlined"
            error={formik.errors.email && formik.touched.email}
          />
          <HelperText
            type="error"
            visible={Boolean(formik.errors.email && formik.touched.email)}>
            {formik.errors.email}
          </HelperText>
          <TextInput
            style={{height: 50}}
            label="Password"
            onChangeText={formik.handleChange("password")}
            onBlur={formik.handleBlur("password")}
            value={formik.values.password}
            secureTextEntry={isSecureText}
            mode="outlined"
            error={Boolean(formik.errors.password && formik.touched.password)}
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
              formik.errors.password && formik.touched.password,
            )}>
            {formik.errors.password}
          </HelperText>
        </View>
        <View>
          <Button
            mode="contained"
            onPress={formik.handleSubmit}
            loading={isLoginLoading}>
            Login
          </Button>
        </View>
        <View style={{alignItems: "center", marginTop: 20}}>
          <TouchableOpacity onPress={() => setForgetPasswordIsVisible(true)}>
            <Text style={{color: colors.primary}}>Forget password ?</Text>
          </TouchableOpacity>
        </View>
        <View
          style={{
            alignItems: "center",
            marginTop: 20,
            flexDirection: "row",
            justifyContent: "center",
          }}>
          <Text>Don't have an account ?</Text>
          <TouchableOpacity onPress={() => goToRegister()}>
            <Text style={{color: colors.primary, marginLeft: 5}}>Sign Up</Text>
          </TouchableOpacity>
        </View>
      </View>
    </>
  );
};

export default withTheme(LoginView);
