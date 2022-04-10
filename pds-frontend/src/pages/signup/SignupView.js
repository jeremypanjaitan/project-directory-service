import React, { useState } from "react";
import { useFormik } from "formik";
import * as Yup from "yup";
import {
  Box,
  Button,
  Container,
  IconButton,
  InputAdornment,
  MenuItem,
  TextField,
  Typography,
} from "@mui/material";
import "./Signup.module.css";
import ArrowBackIcon from "@mui/icons-material/ArrowBack";
import { LOGIN } from "../../router";
import { Visibility, VisibilityOff } from "@mui/icons-material";
const gender = [
  {
    index: 1,
    value: "M",
    label: "Male",
  },
  {
    index: 2,
    value: "F",
    label: "Female",
  },
];

const SignupView = ({ roles, divisions, handleSubmit, navigation }) => {
  const [showReenterPassword, setShowReenterPassword] = useState(false);
  const handleClickShowReEnterPassword = () =>
    setShowReenterPassword(!showReenterPassword);
  const handleMouseDownReEnterPassword = () =>
    setShowReenterPassword(!showReenterPassword);
  const [showPassword, setShowPassword] = useState(false);
  const handleClickShowPassword = () => setShowPassword(!showPassword);
  const handleMouseDownPassword = () => setShowPassword(!showPassword);
  const navigateTo = navigation;
  const formik = useFormik({
    initialValues: {
      fullName: "",
      email: "",
      gender: "",
      divisionId: "",
      roleId: "",
      password: "",
      reEnterPassword: "",
    },
    validationSchema: Yup.object({
      fullName: Yup.string()
        .required("This field is required")
        .max(30, "Maximum 30 characters"),
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
        .required("This field is required")
        .min(6, "Minimum 6 character length")
        .oneOf([Yup.ref("password"), null], "Passwords must match"),
    }),
    onSubmit: (values) => {
      handleSubmit({
        fullName: values.fullName,
        email: values.email,
        gender: values.gender,
        divisionId: values.divisionId,
        roleId: values.roleId,
        password: values.password,
      });
    },
  });

  return (
    <>
      <Box
        component="main"
        sx={{
          alignItems: "center",
          display: "flex",
          flexGrow: 1,
          minHeight: "100%",
        }}
      >
        <Container maxWidth="sm">
          <Button
            component="a"
            startIcon={<ArrowBackIcon fontSize="small" />}
            onClick={(e) => {
              navigateTo(LOGIN);
            }}
          >
            Login Page
          </Button>
          <form onSubmit={formik.handleSubmit}>
            <Box sx={{ my: 3 }}>
              <Typography color="textPrimary" variant="h4">
                Create a new account
              </Typography>
            </Box>
            <TextField
              fullWidth
              className="SignupForm"
              data-testid="SignupForm-fullName"
              id="fullName"
              margin="normal"
              name="fullName"
              label="Full Name"
              value={formik.values.fullName}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.errors.fullName && formik.touched.fullName}
              helperText={
                formik.errors.fullName &&
                formik.touched.fullName &&
                formik.errors.fullName
              }
            />
            <TextField
              fullWidth
              data-testid="SignupForm-email"
              className="SignupForm"
              id="email"
              name="email"
              margin="normal"
              label="Email"
              value={formik.values.email}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.errors.email && formik.touched.email}
              helperText={
                formik.errors.email &&
                formik.touched.email &&
                formik.errors.email
              }
            />
            <TextField
              fullWidth
              className="SignupForm"
              id="gender"
              select
              margin="normal"
              name="gender"
              inputProps={{ "data-testid": "SignupForm-gender" }}
              label="Gender"
              value={formik.values.gender}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={formik.errors.gender && formik.touched.gender}
              helperText={
                formik.errors.gender &&
                formik.touched.gender &&
                formik.errors.gender
              }
            >
              {gender.map((option) => (
                <MenuItem key={option.index} value={option.value}>
                  {option.label}
                </MenuItem>
              ))}
            </TextField>
            <TextField
              fullWidth
              className="SignupForm"
              id="divisionId"
              name="divisionId"
              select
              margin="normal"
              label="Division"
              inputProps={{ "data-testid": "SignupForm-division" }}
              value={formik.values.divisionId}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={formik.errors.divisionId && formik.touched.divisionId}
              helperText={
                formik.errors.divisionId &&
                formik.touched.divisionId &&
                formik.errors.divisionId
              }
            >
              {divisions.map((option, key) => (
                <MenuItem key={key} value={option.id}>
                  {option.name}
                </MenuItem>
              ))}
            </TextField>
            <TextField
              fullWidth
              className="SignupForm"
              id="roleId"
              name="roleId"
              select
              margin="normal"
              label="Role"
              inputProps={{ "data-testid": "SignupForm-role" }}
              value={formik.values.roleId}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={formik.errors.roleId && formik.touched.roleId}
              helperText={
                formik.errors.roleId &&
                formik.touched.roleId &&
                formik.errors.roleId
              }
            >
              {roles.map((option, key) => (
                <MenuItem key={key} value={option.id}>
                  {option.name}
                </MenuItem>
              ))}
            </TextField>
            <TextField
              fullWidth
              inputProps={{ "data-testid": "SignupForm-password" }}
              className="SignupForm"
              id="password"
              name="password"
              label="Password"
              margin="normal"
              type={showPassword ? "text" : "password"}
              value={formik.values.password}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={formik.errors.password && formik.touched.password}
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
                formik.errors.password &&
                formik.touched.password &&
                formik.errors.password
              }
            />
            <TextField
              fullWidth
              inputProps={{ "data-testid": "SignupForm-password" }}
              className="SignupForm"
              id="reEnterPassword"
              name="reEnterPassword"
              label="Confirm Password"
              margin="normal"
              type={showReenterPassword ? "text" : "password"}
              value={formik.values.reEnterPassword}
              onBlur={formik.handleBlur}
              onChange={formik.handleChange}
              error={
                formik.errors.reEnterPassword && formik.touched.reEnterPassword
              }
              InputProps={{
                // <-- This is where the toggle button is added.
                endAdornment: (
                  <InputAdornment position="end">
                    <IconButton
                      aria-label="toggle password visibility"
                      onClick={handleClickShowReEnterPassword}
                      onMouseDown={handleMouseDownReEnterPassword}
                    >
                      {showReenterPassword ? <Visibility /> : <VisibilityOff />}
                    </IconButton>
                  </InputAdornment>
                ),
              }}
              helperText={
                formik.errors.reEnterPassword &&
                formik.touched.reEnterPassword &&
                formik.errors.reEnterPassword
              }
            />
            <Button
              className="SignupButton"
              variant="contained"
              width="900px"
              fullWidth
              margin="normal"
              onClick={formik.handleSubmit}
            >
              {" "}
              Sign Up{" "}
            </Button>
          </form>
        </Container>
      </Box>
    </>
  );
};

export default SignupView;
