import {
  Button,
  Container,
  MenuItem,
  Modal,
  TextField,
  Toolbar,
} from "@mui/material";
import Box from "@mui/material/Box";
import { useFormik } from "formik";
import * as Yup from "yup";
import React, { useEffect, useRef } from "react";
import { Editor } from "@tinymce/tinymce-react";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import { colors } from "../../constants";
import defaultPicture from "../../static/defaultImage.png";
import { useCloudStorage } from "../../utils";
import CircularProgress from "@mui/material/CircularProgress";
import Swal from "sweetalert2";
import { useEditProject } from "./EditProjectRepo";
import { useLoading } from "../../context";

const EditProjectView = ({ categories, handleSubmit, currentProject }) => {
  const { uploadPicture, resetProgress, progress } = useCloudStorage();
  const editorRef = useRef(null);
  const { setIsLoading } = useLoading();

  useEffect(() => {
    if (currentProject) {
      formik.setFieldValue("title", currentProject.title);
      formik.setFieldValue("picture", currentProject.picture);
      formik.setFieldValue("categoryId", currentProject.categoryId);
      formik.setFieldValue("description", currentProject.description);
      formik.values.story = currentProject.story;
    }
  }, [currentProject]);

  const formik = useFormik({
    initialValues: {
      title: "",
      picture: "",
      description: "",
      story: "<p>Write story about you project here...</p>",
      categoryId: "",
    },
    validationSchema: Yup.object({
      title: Yup.string()
        .required("This field is required")
        .min(6, "Minimum 6 character length")
        .max(150, "Maximum 150 character"),
      picture: Yup.string(),
      description: Yup.string()
        .required("This field is required")
        .min(10, "Minimum 10 character length")
        .max(300, "Maximum 300 character"),
      story: Yup.string()
        .required("This field is required")
        .min(10, "Minimum 10 character length"),
      categoryId: Yup.number().required("This field is required"),
    }),
    onSubmit: (values) => {
      handleSubmit(values);
    },
  });

  const [openModal, setOpenModal] = React.useState(false);
  const handleOpenModal = () => setOpenModal(true);
  const handleCloseModal = () => setOpenModal(false);

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
        maxWidth="xl"
      >
        <Typography variant="h4"> Edit your project </Typography>
        <form onSubmit={formik.handleSubmit}>
          <Grid container>
            <Grid container columnSpacing={2} direction="row">
              <Grid item xs={8} direction="column">
                <Grid item xs={4}>
                  <TextField
                    fullWidth
                    margin="normal"
                    name="title"
                    label="Project Title"
                    value={formik.values.title || ""}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    error={formik.errors.title && formik.touched.title}
                    helperText={
                      formik.errors.title &&
                      formik.touched.title &&
                      formik.errors.title
                    }
                  />
                </Grid>
                <Grid item xs={4}>
                  <TextField
                    fullWidth
                    select
                    inputProps={{ "data-testid": "category-select" }}
                    margin="normal"
                    name="categoryId"
                    label="Category"
                    value={formik.values.categoryId}
                    onChange={formik.handleChange}
                    error={
                      formik.errors.categoryId && formik.touched.categoryId
                    }
                    helperText={
                      formik.errors.categoryId &&
                      formik.touched.categoryId &&
                      formik.errors.categoryId
                    }
                  >
                    {categories.map((option) => (
                      <MenuItem key={option.id} value={option.id}>
                        {option.name}
                      </MenuItem>
                    ))}
                  </TextField>
                </Grid>
                <Grid item xs={4}>
                  <TextField
                    fullWidth
                    multiline
                    maxRows={3}
                    margin="normal"
                    name="description"
                    label="Short Description"
                    value={formik.values.description || ""}
                    onChange={formik.handleChange}
                    onBlur={formik.handleBlur}
                    error={
                      formik.errors.description && formik.touched.description
                    }
                    helperText={
                      formik.errors.description &&
                      formik.touched.description &&
                      formik.errors.description
                    }
                  />
                </Grid>
              </Grid>
              <Grid item xs={3} direction="column">
                <Grid item xs={6}>
                  <img
                    src={formik.values.picture || defaultPicture}
                    alt="preview"
                    style={{
                      width: "150px",
                      height: "150px",
                    }}
                  />
                </Grid>
                <Grid item xs={6}>
                  <Box>
                    <Box>
                      <Button
                        sx={{
                          display: "flex",
                          boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
                          borderRadius: "15px",
                          margin: "25px",
                          width: "200 px",
                          height: "36 px",
                          mt: "429 px",
                          ml: "500 px",
                          padding: "6 px",
                          background: `linear-gradient(229.93deg, ${colors.SOFT_BLUE} -0.7%, rgba(102, 133, 227, 0) 300%)`,
                          textTransform: "none",
                          fontSize: "12px",
                          fontWeight: 400,
                        }}
                        onClick={handleOpenModal}
                        variant="contained"
                      >
                        Upload Picture
                      </Button>
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
                            filter:
                              "drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25))",
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
                            Upload your project cover image
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
                              ml: "95px",
                              mt: "25px",
                              borderRadius: "20px",
                            }}
                          >
                            <Box
                              style={{
                                display: "flex",
                                justifyContent: "center",
                              }}
                            >
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
                                  boxShadow:
                                    "1px -1px 15px -6px rgba(0,0,0,0.83)",
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
                                    handleUploadPicture(
                                      e.currentTarget.files[0]
                                    );
                                  }}
                                />
                              </Button>
                            </Box>
                          </Box>
                        </Box>
                      </Modal>
                    </Box>
                  </Box>
                </Grid>
              </Grid>
            </Grid>
          </Grid>
          <Grid container columnSpacing={2} direction="row"></Grid>

          <Editor
            apiKey={process.env.REACT_APP_TINY_MCE_API_KEY}
            id="tiny-mce"
            onInit={(editor) => (editorRef.current = editor)}
            initialValue={currentProject.story}
            onChange={(e) =>
              formik.setFieldValue("story", e.target.getContent())
            }
            onBlur={formik.handleBlur}
            init={{
              height: 500,
              plugins: [
                "advlist autolink lists link image charmap print preview anchor",
                "searchreplace visualblocks code fullscreen",
                "insertdatetime media table paste code help wordcount",
              ],

              toolbar:
                "undo redo | formatselect image| " +
                "bold italic backcolor | alignleft aligncenter " +
                "alignright alignjustify | bullist numlist outdent indent | " +
                "removeformat | help",
              images_upload_url: "postAcceptor.php",
              images_upload_handler: function (blobInfo, success, failure) {
                if (blobInfo.blob().size / (1024 * 1024) > 10) {
                  failure("Maximum 10 MB file size");
                } else {
                  uploadPicture(blobInfo.blob())
                    .then((url) => success(url))
                    .catch((error) => failure(error));
                }
              },
              content_style:
                "body { font-family:Helvetica,Arial,sans-serif; font-size:14px }",
            }}
          />
          <Button
            className="SignupButton"
            variant="contained"
            width="900px"
            fullWidth
            margin="normal"
            type="submit"
            sx={{
              marginTop: "20px",
              boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
              borderRadius: "15px",
              width: "200 px",
              height: "36 px",
              padding: "6 px",
              background: `linear-gradient(229.93deg, ${colors.SOFT_BLUE} -0.7%, rgba(102, 133, 227, 0) 300%)`,
              textTransform: "none",
              fontSize: "12px",
              fontWeight: 400,
            }}
            onClick={() => handleSubmit(formik.values)}
          >
            {" "}
            Edit Project{" "}
          </Button>
        </form>

        <Toolbar />
      </Container>
    </>
  );
};

export default EditProjectView;
