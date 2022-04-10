import { useFormik } from "formik";
import * as Yup from "yup";
import * as React from "react";
import LoadingOverlay from "react-loading-overlay-ts";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Link from "@mui/material/Link";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { useNavigate } from "react-router-dom";
import { SIGN_UP } from "../../router";
import backgroundImage from "../../static/project.svg";
import { colors } from "../../constants";
import Swal from "sweetalert2";
import { IconButton, InputAdornment, Modal } from "@mui/material";
import { Visibility, VisibilityOff } from "@mui/icons-material";

const LoginPageView = ({
  isLoading,
  isError,
  loginReset,
  handleSubmitLogin,
  handleSubmitForgetPassword,
  loginError,
  InputEmailVisible,
  setInputEmailVisible,
  postForgetPasswordError,
  postForgetPasswordReset,
  postPasswordData,
}) => {
  const navigate = useNavigate();
  const [showPassword, setShowPassword] = React.useState(false);
  const handleClickShowPassword = () => setShowPassword(!showPassword);
  const handleMouseDownPassword = () => setShowPassword(!showPassword);

  const formikLogin = useFormik({
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
    onSubmit: handleSubmitLogin,
  });

  if (isError) {
    Swal.fire(loginError.response.data.description, "", "warning");
    loginReset();
  }

  if (postForgetPasswordError) {
    postForgetPasswordReset();
    Swal.fire({
      title: postPasswordData.error.response.data.description,
      confirmButtonText: "Ok",
    }).then(
      (result) => {
        if (result.isConfirmed) {
          setInputEmailVisible(true);
        }
      },
      "",
      "warning"
    );
  }

  const FormikForgetPassword = useFormik({
    initialValues: {
      email: "",
    },
    validationSchema: Yup.object({
      email: Yup.string()
        .email("Invalid email format")
        .required("Email is required"),
    }),
    onSubmit: (values) => {
      handleSubmitForgetPassword(values);
    },
  });

  return (
    <LoadingOverlay active={isLoading}>
      <Grid container component="main" sx={{ height: "100vh" }}>
        <CssBaseline />
        <Grid
          item
          sm={4}
          md={7}
          sx={{
            backgroundImage: `url(${backgroundImage})`,
            backgroundRepeat: "no-repeat",
            backgroundSize: "cover",
          }}
        />

        <Grid
          item
          xs={12}
          sm={8}
          md={5}
          mt={0.25}
          sx={{ height: "99vh" }}
          component={Paper}
          elevation={6}
          boxShadow={20}
        >
          <Box
            sx={{
              my: 20,
              mx: 4,
              display: "flex",
              flexDirection: "column",
              alignItems: "center",
            }}
          >
            <Typography component="h1" variant="h5" style={{ fontWeight: 600 }}>
              Sign in
            </Typography>
            <form onSubmit={formikLogin.handleSubmit} sx={{ mt: 1 }}>
              <TextField
                margin="normal"
                fullWidth
                id="email"
                name="email"
                label="Email Address"
                value={formikLogin.values.email || ""}
                onChange={formikLogin.handleChange}
                onBlur={formikLogin.handleBlur}
                error={formikLogin.errors.email}
                helperText={
                  formikLogin.touched.email && formikLogin.errors.email
                }
              />
              <TextField
                margin="normal"
                fullWidth
                id="password"
                name="password"
                label="Password"
                value={formikLogin.values.password || ""}
                onChange={formikLogin.handleChange}
                onBlur={formikLogin.handleBlur}
                error={formikLogin.errors.password}
                InputProps={{
                  // <-- This is where the toggle button is added.
                  endAdornment: (
                    <InputAdornment position="end">
                      <IconButton
                        aria-label="toggle password visibility"
                        onClick={handleClickShowPassword}
                        onMouseDown={handleMouseDownPassword}
                      >
                        {showPassword ? <Visibility /> : <VisibilityOff />}
                      </IconButton>
                    </InputAdornment>
                  ),
                }}
                helperText={
                  formikLogin.touched.password && formikLogin.errors.password
                }
                type={showPassword ? "text" : "password"}
                autoComplete="current-password"
              />

              <br></br>
              <Button
                fullWidth
                variant="contained"
                sx={{ mt: 3, mb: 2 }}
                style={{ backgroundColor: "#5D6CE6" }}
                type="submit"
                disabled={!(formikLogin.isValid && formikLogin.dirty)}
              >
                Sign In
              </Button>
              <Grid container>
                <Grid item xs>
                  <Link
                    href="#"
                    variant="body2"
                    onClick={() => {
                      setInputEmailVisible(true);
                    }}
                  >
                    {" Forget password?"}
                  </Link>
                  <Modal
                    open={InputEmailVisible}
                    onClose={() => setInputEmailVisible(false)}
                  >
                    <Box
                      sx={{
                        outline: "none",
                        gap: "20px",
                        padding: "10px",
                        margin: "10px",
                        paddingBottom: "10px",
                        flexGrow: 1,
                        position: "absolute",
                        top: "50%",
                        left: "50%",
                        transform: "translate(-50%, -50%)",
                        width: 400,
                        height: 250,
                        background:
                          "linear-gradient(227.51deg, #5B8CEC 11.78%, #77D3DC 81.58%, rgba(107, 181, 227, 0) 113.09%)",
                        filter: "drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25))",
                        borderRadius: "25px",
                        flexWrap: "wrap",
                      }}
                    >
                      <Typography
                        align="center"
                        style={{
                          marginTop: "20px",
                          color: "white",
                          fontWeight: "600",
                          fontSize: "20px",
                          fontFamily: "Montserrat",
                        }}
                      >
                        Input your email
                      </Typography>
                      <TextField
                        margin="normal"
                        fullWidth
                        id="email"
                        name="email"
                        label="Email Address"
                        value={FormikForgetPassword.values.email || ""}
                        onChange={FormikForgetPassword.handleChange}
                        onBlur={FormikForgetPassword.handleBlur}
                        error={FormikForgetPassword.errors.email}
                        helperText={
                          FormikForgetPassword.touched.email &&
                          FormikForgetPassword.errors.email
                        }
                        sx={{
                          mt: "30px",
                          borderRadius: "5px",
                          background: "#FFFFFF",
                        }}
                      />
                      <Box
                        sx={{
                          ml: "270px",
                          mt: "20px",
                        }}
                      >
                        <Button
                          sx={{
                            display: "flex",
                            boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
                            borderRadius: "15px",
                            margin: 0,
                            width: "270 px",
                            height: "34 px",
                            mt: "236 px",

                            padding: "10 px",
                            background: `linear-gradient(229.93deg, ${colors.SOFT_BLUE} -0.7%, rgba(102, 133, 227, 0) 170%)`,
                          }}
                          style={{
                            textTransform: "none",
                            fontSize: "15px",
                            fontWeight: 500,
                          }}
                          variant="contained"
                          type="submit"
                          onClick={() => {
                            FormikForgetPassword.handleSubmit();
                            setInputEmailVisible(false);
                          }}
                        >
                          Submit
                        </Button>
                      </Box>
                    </Box>
                  </Modal>
                </Grid>
                <Grid item>
                  <Link
                    href="#"
                    variant="body2"
                    onClick={(e) => {
                      e.preventDefault();
                      navigate(SIGN_UP);
                    }}
                  >
                    {"Don't have an account? Sign Up"}
                  </Link>
                </Grid>
              </Grid>
            </form>
          </Box>
        </Grid>
      </Grid>
    </LoadingOverlay>
  );
};

export default LoginPageView;
