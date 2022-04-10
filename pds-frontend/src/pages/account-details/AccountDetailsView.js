import React from "react";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Grid from "@mui/material/Grid";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { ButtonBase, List } from "@mui/material";
import { Box } from "@mui/system";
import VisibilityOutlinedIcon from "@mui/icons-material/VisibilityOutlined";
import ChatBubbleOutlineOutlinedIcon from "@mui/icons-material/ChatBubbleOutlineOutlined";
import ThumbUpOutlinedIcon from "@mui/icons-material/ThumbUpOutlined";
import ThumbDownOutlinedIcon from "@mui/icons-material/ThumbDownOutlined";
import moment from "moment";
import { SHOW } from "../../router";
import { useLoading } from "../../context";
import { useNavigate } from "react-router";
import { useAuth } from "../../services";

const AccountDetailsView = ({
  totalPageActivity,
  activities,
  currentPageActivity,
  setCurrentPageActivity,
  handleLoadMoreActivity,
  totalPageProject,
  projects,
  currentPageProject,
  setCurrentPageProject,
  handleLoadMoreProject,
}) => {
  const navigate = useNavigate();
  const { setIsLoading } = useLoading();
  const auth = useAuth();

  return (
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
        Hello {auth?.user?.fullName}!
      </h1>
      <Grid items container>
        <Grid xs>
          <Box
            sx={{
              padding: "15px",
              width: "440px",
              height: "500 px",
              borderRadius: "25px",
              mt: "209 px",
              ml: "554 px",
              boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
            }}
          >
            <h3 style={{ fontFamily: "montserrat", marginLeft: "20px" }}>
              Your Activity
            </h3>
            <List
              sx={{
                width: "440px",
                maxWidth: "440px",
                bgcolor: "background.paper",
                position: "relative",
                overflow: "auto",
                maxHeight: "450px",
              }}
            >
              {activities.length > 0 ? (
                activities.map((value, key) => (
                  <Box
                    key={key}
                    sx={{
                      padding: "5px",
                      width: "350px",
                      height: "100 px",
                      borderRadius: "25px",
                      mt: "10px",
                      ml: "20px",
                      boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                    }}
                  >
                    <Grid
                      container
                      orientation="row"
                      sx={{ margin: 1, padding: 1 }}
                    >
                      <Grid container>
                        <Grid item xs={1}>
                          {value.Type === "LIKE" ? (
                            <ThumbUpOutlinedIcon />
                          ) : value.Type === "DISLIKE" ? (
                            <ThumbDownOutlinedIcon />
                          ) : (
                            <ChatBubbleOutlineOutlinedIcon />
                          )}
                        </Grid>
                        <Grid item xs>
                          <Typography>{value.Header}</Typography>
                          {value.Body === null ? null : (
                            <Typography sx={{ fontStyle: "italic" }}>
                              "{value.Body}"
                            </Typography>
                          )}
                        </Grid>
                      </Grid>
                    </Grid>
                    <Typography
                      sx={{
                        textAlign: "center",
                        fontSize: "15px",
                      }}
                    >
                      {moment.utc(value.CreatedAt).fromNow()}
                    </Typography>
                  </Box>
                ))
              ) : (
                <Typography> No Data</Typography>
              )}
              {totalPageActivity !== currentPageActivity && (
                <ButtonBase
                  className="btn-load-more"
                  onClick={() =>
                    handleLoadMoreActivity(currentPageActivity + 1)
                  }
                >
                  <Typography sx={{ mt: "10px", ml: "160px" }}>
                    Load more...
                  </Typography>
                </ButtonBase>
              )}
            </List>
          </Box>
        </Grid>
        <Grid xs>
          <Box
            sx={{
              padding: "15px",
              width: "440px",
              height: "500 px",
              borderRadius: "25px",
              mt: "209 px",
              ml: "554 px",
              boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
            }}
          >
            <h3 style={{ fontFamily: "montserrat", marginLeft: "20px" }}>
              Your Project
            </h3>
            <List
              sx={{
                width: "440px",
                maxWidth: "440px",
                bgcolor: "background.paper",
                position: "relative",
                overflow: "auto",
                maxHeight: "450px",
              }}
            >
              {projects.length > 0 ? (
                projects.map((value, key) => (
                  <ButtonBase
                    onClick={() => {
                      setIsLoading(true);
                      navigate(SHOW + "/" + value.ID);
                    }}
                  >
                    <Card
                      sx={{
                        display: "flex",
                        flexDirection: "column",
                        width: "300px",
                        height: "200px",
                        mt: "20px",
                        ml: "50px",
                      }}
                    >
                      <CardMedia
                        component="img"
                        image={value.picture}
                        sx={{ height: "100px" }}
                      />
                      <CardContent>
                        <Typography gutterBottom noWrap>
                          {value.title}
                        </Typography>

                        <Grid container>
                          <Grid item xs={1} sx={{ ml: "20px" }}>
                            <Typography>{value.totalLikes}</Typography>
                          </Grid>
                          <Grid item>
                            <ThumbUpOutlinedIcon />
                          </Grid>
                          <Grid item xs={1} sx={{ ml: "40px" }}>
                            <Typography>{value.totalComments}</Typography>
                          </Grid>
                          <Grid item>
                            <ChatBubbleOutlineOutlinedIcon />
                          </Grid>
                          <Grid item xs={1} sx={{ ml: "50px" }}>
                            <Typography>{value.totalViews}</Typography>
                          </Grid>
                          <Grid item>
                            <VisibilityOutlinedIcon />
                          </Grid>
                        </Grid>
                      </CardContent>
                    </Card>
                  </ButtonBase>
                ))
              ) : (
                <Typography> No Data</Typography>
              )}
              {totalPageProject !== currentPageProject && (
                <ButtonBase
                  className="btn-load-more"
                  onClick={() => handleLoadMoreProject(currentPageProject + 1)}
                >
                  <Typography sx={{ mt: "10px", ml: "160px" }}>
                    Load more...
                  </Typography>
                </ButtonBase>
              )}
            </List>
          </Box>
        </Grid>
      </Grid>
    </Container>
  );
};

export default AccountDetailsView;
