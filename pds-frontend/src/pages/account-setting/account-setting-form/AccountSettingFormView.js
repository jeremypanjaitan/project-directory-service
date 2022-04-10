import {
  Avatar,
  Badge,
  Button,
  Container,
  Grid,
  IconButton,
  InputAdornment,
  MenuItem,
  Modal,
  TextField,
  Typography,
} from "@mui/material";
import { Box } from "@mui/system";
import { useFormik } from "formik";
import * as Yup from "yup";
import * as React from "react";
import { colors } from "../../../constants";
import { useAuth } from "../../../services/auth";
import { useCloudStorage } from "../../../utils";
import CircularProgress from "@mui/material/CircularProgress";
import defaultPicture from "../../../static/defaultImage.png";
import Swal from "sweetalert2";
import AddIcon from "@mui/icons-material/Add";
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

const AccountSettingFormView = ({
  handleSubmitOldPassword,
  handleSubmitPassword,
  handleSubmitData,
  roles,
  divisions,
  defaultData,
  currentRole,
  currentDivision,
  passwordVisible,
  setPasswordVisible,
  isError,
  loginReset,
}) => {
  const [showPassword, setShowPassword] = React.useState(false);
  const handleClickShowPassword = () => setShowPassword(!showPassword);
  const handleMouseDownPassword = () => setShowPassword(!showPassword);
  const { uploadPicture, resetProgress, progress } = useCloudStorage();
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
    onSubmit: handleSubmitData,
  });
  const formikChangePassword = useFormik({
    initialValues: {
      password: "",
    },
    validationSchema: Yup.object({
      password: Yup.string()
        .min(6, "6 min character length")
        .required("Password is required"),
    }),
    onSubmit: handleSubmitOldPassword,
  });

  if (isError) {
    Swal.fire({
      title: "Invalid password!",
      confirmButtonText: "Ok",
    }).then(
      (result) => {
        if (result.isConfirmed) {
          loginReset();
          setPasswordVisible(true);
        }
      },
      "",
      "warning"
    );
  }

  React.useEffect(() => {
    if (defaultData) {
      formik.setFieldValue("fullName", defaultData.data.fullName);
      formik.setFieldValue("email", defaultData.data.email);
      formik.setFieldValue("gender", defaultData.data.gender);
      formik.setFieldValue("division", defaultData.data.divisionId);
      formik.setFieldValue("role", defaultData.data.roleId);
      formik.setFieldValue("biography", defaultData.data.biography);
      formik.setFieldValue("picture", defaultData.data.picture);
    }
    //eslint-disable-next-line
  }, [defaultData]);

  const [openModal, setOpenModal] = React.useState(false);
  const handleOpenModal = () => setOpenModal(true);
  const handleCloseModal = () => setOpenModal(false);

  const auth = useAuth();

  const handleUploadPicture = async (picture) => {
    if (picture.size / (1024 * 1024) > 10) {
      setOpenModal(false);
      Swal.fire({
        title: "Unable to choose file",
        text: "Maximum file size 10MB, please choose other image",
        icon: "warning",
        showCancelButton: true,
        cancelButtonColor: "#d33",
      });
    } else {
      const url = await uploadPicture(picture);
      formik.setFieldValue("picture", url);
      resetProgress();
    }
  };

  return (
    <>
      <Container
        sx={{
          height: "100vh",
          padding: "50px",
          display: "flex",
          flexDirection: "column",
          gap: "30px",
        }}
        maxWidth="lg"
      >
        <h1 style={{ fontFamily: "montserrat", marginLeft: "30px" }}>
          Account Settings
        </h1>
        <Grid
          container
          sx={{
            display: "flex",
            justifyContent: "center",
          }}
        >
          <Grid item>
            <Container
              sx={{ display: "flex", flexDirection: "column", gap: "10px" }}
            >
              <Box
                sx={{
                  padding: "15px",
                  width: "150px",
                  height: "600 px",
                  background: `linear-gradient(229.93deg, ${colors.SOFT_BLUE} -0.7%, rgba(102, 133, 227, 0) 170%)`,
                  borderRadius: "20px",
                  boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",

                  display: "flex",
                  justifyContent: "center",
                  alignItems: "center",
                }}
              >
                <Box
                  sx={{
                    padding: "15px",
                    display: "flex",
                    flexDirection: "column",
                    gap: "30px",
                  }}
                >
                  <Box
                    sx={{
                      display: "flex",
                      justifyContent: "center",
                    }}
                  >
                    <Badge
                      badgeContent={<AddIcon fontSize="small" />}
                      color="primary"
                      overlap="circular"
                      anchorOrigin={{
                        vertical: "bottom",
                        horizontal: "right",
                      }}
                      onClick={handleOpenModal}
                    >
                      <Avatar
                        alt="Remy Sharp"
                        src={formik.values.picture}
                        sx={{
                          width: "70px",
                          height: "70px",
                          boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                        }}
                      />
                    </Badge>
                  </Box>
                  <Box>
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "center",
                      }}
                    >
                      <Typography
                        noWrap
                        variant="h6"
                        sx={{
                          color: "white",
                          fontSize: "15px",
                          fontWeight: 600,
                        }}
                      >
                        {auth?.user?.fullName}
                      </Typography>
                    </Box>
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "center",
                      }}
                    >
                      <Typography
                        noWrap
                        variant="h8"
                        sx={{
                          color: "white",
                          fontSize: "14px",
                          fontWeight: 400,
                        }}
                      >
                        {auth?.user?.roleName}
                      </Typography>
                    </Box>
                  </Box>
                </Box>
              </Box>

              <Box>
                {" "}
                <Button
                  fullWidth
                  sx={{
                    boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
                    borderRadius: "20px",
                    backgroundColor: colors.SOFT_BLUE,
                    textTransform: "none",
                    fontSize: "13px",
                    fontWeight: 500,
                  }}
                  variant="contained"
                  onClick={() => {
                    setPasswordVisible(true);
                  }}
                >
                  Change Password
                </Button>
              </Box>

              <Modal
                open={passwordVisible}
                onClose={() => setPasswordVisible(false)}
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
                    Input your old password
                  </Typography>

                  <TextField
                    margin="normal"
                    fullWidth
                    id="password"
                    name="password"
                    label="Old password"
                    value={formikChangePassword.values.password || ""}
                    onChange={formikChangePassword.handleChange}
                    onBlur={formikChangePassword.handleBlur}
                    error={formikChangePassword.errors.password}
                    helperText={
                      formikChangePassword.touched.password &&
                      formikChangePassword.errors.password
                    }
                    type={showPassword ? "text" : "password"}
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
                    autoComplete="current-password"
                    sx={{
                      width: "380px",
                      borderRadius: "5px",
                      background: "#FFFFFF",
                      mt: "30px",
                      ml: "8px",
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
                        // handleSubmitPassword(auth?.user?.email);

                        formikChangePassword.handleSubmit();
                        setPasswordVisible(false);
                      }}
                    >
                      Submit
                    </Button>
                  </Box>
                </Box>
              </Modal>

              <Modal open={openModal} onClose={handleCloseModal}>
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
                    width: 500,
                    height: 360,
                    background:
                      "linear-gradient(227.51deg, #5B8CEC 11.78%, #77D3DC 81.58%, rgba(107, 181, 227, 0) 113.09%)",
                    filter: "drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25))",
                    borderRadius: "25px",
                  }}
                >
                  <p
                    align="center"
                    style={{
                      color: "white",
                      fontWeight: "600",
                      fontSize: "20px",
                      fontFamily: "Montserrat",
                    }}
                  >
                    Change your profile picture
                  </p>
                  <Box
                    sx={{
                      paddingTop: "20px",
                      paddingButtom: "20px",
                      margin: "10px",
                      width: "300px",
                      height: "227px",
                      background:
                        "linear-gradient(196.11deg, #FFFFFF -100.17%, rgba(255, 255, 255, 0) 171.14%)",
                      boxShadow: "0px 0px 6px 5px rgba(0, 0, 0, 0.1)",
                      ml: "auto",
                      mr: "auto",
                      mt: "25px",
                      borderRadius: "20px",
                    }}
                  >
                    <Box style={{ display: "flex", justifyContent: "center" }}>
                      {progress > 1 ? (
                        <Box
                          sx={{
                            width: "120px",
                            height: "120px",
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center",
                          }}
                        >
                          <CircularProgress
                            variant="determinate"
                            value={progress}
                          />
                        </Box>
                      ) : (
                        <img
                          src={formik.values.picture || defaultPicture}
                          alt="preview"
                          style={{
                            width: "120px",
                            height: "120px",
                          }}
                        />
                      )}
                    </Box>

                    <Box sx={{ padding: "40px" }}>
                      <Button
                        sx={{
                          paddingButtom: "20px",
                          display: "flex",
                          boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
                          borderRadius: "20px",
                          margin: 0,
                          padding: "10 px",
                          backgroundColor: colors.SOFT_BLUE,
                          textTransform: "none",
                        }}
                        variant="contained"
                        component="label"
                      >
                        <Typography
                          sx={{
                            fontSize: "13px",
                            fontWeight: 400,
                          }}
                        >
                          Browse File
                        </Typography>
                        <input
                          type="file"
                          name="picture"
                          accept="image/*"
                          style={{
                            display: "none",
                            width: "12 px",
                            height: "38 px",
                          }}
                          onChange={(e) => {
                            handleUploadPicture(e.currentTarget.files[0]);
                          }}
                        />
                      </Button>
                    </Box>
                  </Box>
                </Box>
              </Modal>
            </Container>
          </Grid>
          <Grid item>
            <Container>
              <Box
                sx={{
                  padding: "15px",
                  width: "700px",
                  height: "655 px",
                  borderRadius: "20px",
                  mt: "138 px",

                  boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                }}
              >
                <form onSubmit={formik.handleSubmit}>
                  <Box sx={{ display: "flex", gap: "20px" }}>
                    <Box sx={{ width: "50%" }}>
                      <TextField
                        fullWidth
                        id="fullName"
                        margin="normal"
                        name="fullName"
                        label="Full Name"
                        value={formik.values.fullName}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={
                          formik.errors.fullName && formik.touched.fullName
                        }
                        helperText={
                          formik.errors.fullName &&
                          formik.touched.fullName &&
                          formik.errors.fullName
                        }
                      />
                    </Box>
                    <Box sx={{ width: "50%" }}>
                      <TextField
                        fullWidth
                        id="gender"
                        select
                        margin="normal"
                        name="gender"
                        label="Gender"
                        defaultValue={formik.values.gender || ""}
                        value={formik.values.gender || ""}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
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
                    </Box>
                  </Box>
                  <Box sx={{ display: "flex", gap: "20px" }}>
                    <Box sx={{ width: "50%" }}>
                      <TextField
                        fullWidth
                        id="email"
                        name="email"
                        margin="normal"
                        label="Email"
                        inputProps={{ readOnly: true }}
                        value={formik.values.email}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.errors.email && formik.touched.email}
                        helperText={
                          formik.errors.email &&
                          formik.touched.email &&
                          formik.errors.email
                        }
                        disabled
                      />
                    </Box>
                    <Box sx={{ width: "50%" }}>
                      <TextField
                        fullWidth
                        id="division"
                        name="division"
                        select
                        margin="normal"
                        label="Division"
                        defaultValue={currentDivision || ""}
                        value={formik.values.division}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={
                          formik.errors.division && formik.touched.division
                        }
                        helperText={
                          formik.errors.division &&
                          formik.touched.division &&
                          formik.errors.division
                        }
                      >
                        {divisions.map((option, key) => (
                          <MenuItem key={key} value={option.id}>
                            {option.name}
                          </MenuItem>
                        ))}
                      </TextField>
                    </Box>
                  </Box>
                  <Box sx={{ display: "flex", gap: "20px" }}>
                    <Box
                      sx={{
                        width: "50%",
                      }}
                    >
                      <TextField
                        fullWidth
                        id="role"
                        name="role"
                        select
                        margin="normal"
                        label="Role"
                        defaultValue={currentRole || ""}
                        value={formik.values.role}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.errors.role && formik.touched.role}
                        helperText={
                          formik.errors.role &&
                          formik.touched.role &&
                          formik.errors.role
                        }
                      >
                        {roles.map((option, key) => (
                          <MenuItem key={key} value={option.id}>
                            {option.name}
                          </MenuItem>
                        ))}
                      </TextField>
                    </Box>
                    <Box
                      sx={{
                        width: "50%",
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                      }}
                    ></Box>
                  </Box>
                  <Box sx={{ display: "flex", gap: "20px" }}>
                    <Box
                      sx={{
                        width: "49%",
                      }}
                    >
                      <TextField
                        fullWidth
                        id="biography"
                        name="biography"
                        label="Bio"
                        margin="normal"
                        value={formik.values.biography}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        multiline
                        rows={4}
                        maxRows={4}
                        error={
                          formik.errors.biography && formik.touched.biography
                        }
                        helperText={
                          formik.errors.biography &&
                          formik.touched.biography &&
                          formik.errors.biography
                        }
                      />
                    </Box>
                  </Box>
                  <Box sx={{ display: "flex", gap: "20px", marginTop: "15px" }}>
                    <Box
                      sx={{
                        width: "50%",
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                      }}
                    >
                      <Button
                        fullWidth
                        sx={{
                          boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
                          borderRadius: "20px",
                          backgroundColor: colors.SOFT_BLUE,
                          textTransform: "none",
                          fontSize: "13px",
                          fontWeight: 500,
                        }}
                        variant="contained"
                        type="submit"
                      >
                        Save Details
                      </Button>
                    </Box>
                  </Box>
                </form>
              </Box>
            </Container>
          </Grid>
        </Grid>
      </Container>
    </>
  );
};

export default AccountSettingFormView;
