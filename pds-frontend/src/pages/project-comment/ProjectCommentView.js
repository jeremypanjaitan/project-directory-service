import { Button, ButtonBase, Container, Modal, TextField } from "@mui/material";
import Typography from "@mui/material/Typography";
import Box from "@mui/material/Box";
import Avatar from "@mui/material/Avatar";
import Grid from "@mui/material/Grid";
import moment from "moment";
import React, { useState } from "react";
import ArrowBackIosNewIcon from "@mui/icons-material/ArrowBackIosNew";
import { useFormik } from "formik";
import * as Yup from "yup";
import List from "@mui/material/List";
import { useNavigate } from "react-router-dom";
import { colors } from "../../constants";

const ProjectCommentView = ({
  category,
  handleLoadMore,
  profile,
  totalPage,
  currentPage,
  comments,
  project,
  handleCreateComment,
}) => {
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const formik = useFormik({
    initialValues: {
      body: "",
    },
    validationSchema: Yup.object({
      body: Yup.string()
        .required("This field is required")
        .min(2, "Minimum 2 character length")
        .max(300, "Maximum 300 character"),
    }),
    onSubmit: (values) => {
      handleCreateComment(values.body);
      formik.values.body = "";
    },
  });

  const navigate = useNavigate();
  return (
    <Container
      sx={{
        height: "100vh",
        padding: "50px",
        display: "flex",
        gap: "30px",
        flexDirection: "column",
      }}
    >
      <Box container>
        <Grid container xs={12} direction="column" spacing={2}>
          <Grid container spacing={2}>
            <Grid item xs={1}>
              <Button
                data-testid="comment"
                startIcon={<ArrowBackIosNewIcon sx={{ color: "black" }} />}
                style={{
                  marginRight: "10px",
                  maxWidth: "80px",
                  maxHeight: "30px",
                  minWidth: "80px",
                  minHeight: "40px",
                }}
                onClick={() => navigate(-1)}
              />
            </Grid>
            <Grid item xs>
              <Typography
                style={{
                  fontFamily: "montserrat",
                  fontWeight: "bold",
                  fontSize: "30px",
                }}
              >
                {project.title || ""}
              </Typography>
            </Grid>
          </Grid>
          <Grid item xs={3}>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                gap: "20px",
                padding: "20px",
                background:
                  "linear-gradient(188.09deg, #6685E3 -28.83%, rgba(102, 133, 227, 0) 176.25%)",
                backgroundBlendMode: "darken",
                filter: "drop-shadow(0px 4px 12px rgba(0, 0, 0, 0.14))",
                borderRadius: "20px",
                width: "400px",
                boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
              }}
              data-testid="profile-box"
              onClick={handleOpen}
            >
              <Avatar
                alt="picture of profile"
                src={profile.picture || ""}
                sx={{ width: 75, height: 75 }}
              />
              <Box>
                <Typography variant="h6" sx={{ color: "white" }}>
                  {project.title}
                </Typography>
                <Typography variant="h6" sx={{ color: "white" }}>
                  {profile.fullName || ""}
                </Typography>
                <Typography variant="h8" sx={{ color: "white" }}>
                  {category.name || ""}
                </Typography>
              </Box>
            </Box>
          </Grid>
        </Grid>

        <Box sx={{ justifyContent: "space-between" }} marginTop={2}>
          <List
            sx={{
              width: "80%",
              bgcolor: "background.paper",
              overflow: "auto",
              maxHeight: "600px",
            }}
          >
            {comments.length > 0 &&
              comments.map((value, key) => (
                <Box
                  key={key}
                  sx={{
                    boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                    marginBottom: 3,
                    width: "95%",
                    wordWrap: "break-word",
                  }}
                >
                  <Grid
                    container
                    xs={12}
                    orientation="row"
                    sx={{ margin: 1, padding: 1 }}
                  >
                    <Grid item xs={1}>
                      <Avatar
                        alt="picture of profile"
                        src={value.picture || ""}
                        sx={{ width: 75, height: 75 }}
                      />
                    </Grid>
                    <Grid item xs={6}>
                      <Typography variant="h5" padding={3}>
                        {" "}
                        {value.from}{" "}
                      </Typography>
                    </Grid>
                  </Grid>
                  <Typography variant="h6" sx={{ margin: 1, padding: 1 }}>
                    {" "}
                    {value.body}{" "}
                  </Typography>
                  <Typography variant="h7" sx={{ margin: 1, padding: 1 }}>
                    {" "}
                    {moment.utc(value.createdAt).fromNow()}{" "}
                  </Typography>
                </Box>
              ))}
            {totalPage !== currentPage && totalPage !== 0 && (
              <ButtonBase
                className="btn-load-more"
                onClick={() => handleLoadMore(currentPage + 1)}
              >
                <Typography sx={{ mt: "10px", ml: "160px" }}>
                  Load more...
                </Typography>
              </ButtonBase>
            )}
          </List>

          <TextField
            sx={{ width: "80%" }}
            fullWidth
            multiline
            rows={4}
            maxRows={4}
            margin="normal"
            name="body"
            label="Write your comment about this project"
            value={formik.values.body}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.errors.body && formik.touched.body}
            helperText={
              formik.errors.body && formik.touched.body && formik.errors.body
            }
          />
          <Button
            onClick={formik.handleSubmit}
            fullWidth
            sx={{
              boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
              borderRadius: "20px",
              backgroundColor: colors.BLUEBERRY,
              textTransform: "none",
              fontSize: "13px",
              fontWeight: 500,
              width: "80%",
              marginTop: "5px",
              marginBottom: "5px",
              alignSelf: "center",
            }}
            variant="contained"
          >
            Comment Project
          </Button>
        </Box>
        <Modal
          open={open}
          onClose={handleClose}
          aria-labelledby="modal-modal-title"
          aria-describedby="modal-modal-description"
        >
          <Box
            sx={{
              outline: "none",
              gap: "20px",
              flexGrow: 1,
              position: "absolute",
              top: "50%",
              left: "50%",
              transform: "translate(-50%, -50%)",
              width: 500,
              height: 320,
              background:
                "linear-gradient(227.51deg, #5B8CEC 11.78%, #77D3DC 81.58%, rgba(107, 181, 227, 0) 113.09%)",
              filter: "drop-shadow(0px 4px 4px rgba(0, 0, 0, 0.25))",
              borderRadius: "25px",
              padding: "30px",
            }}
          >
            <Grid
              container
              sx={{ display: "flex", flexDirection: "column", gap: "30px" }}
            >
              <Box sx={{ display: "flex", alignItems: "center", gap: "35px" }}>
                <Grid item>
                  <Avatar
                    alt="picture of profile"
                    src={project.picture || ""}
                  />
                </Grid>
                <Grid item>
                  <Typography
                    variant="h4"
                    sx={{ color: "white", marginBottom: "10px" }}
                  >
                    {project.title || ""}
                  </Typography>
                  <Typography variant="h6" sx={{ color: "white" }}>
                    {profile.fullName || ""}
                  </Typography>
                  <Typography variant="h8" sx={{ color: "white" }}>
                    {category.name || ""}
                  </Typography>
                </Grid>
              </Box>
              <Grid item>
                <Box
                  sx={{
                    height: "120px",
                    background:
                      "linear-gradient(196.11deg, #FFFFFF -15.17%, rgba(255, 255, 255, 0) 171.14%)",
                    boxShadow: "0px 0px 6px 5px rgba(0, 0, 0, 0.1)",
                    borderRadius: "25px",
                    padding: "20px",
                  }}
                >
                  {project.description || ""}
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Modal>
      </Box>
    </Container>
  );
};

export default ProjectCommentView;
