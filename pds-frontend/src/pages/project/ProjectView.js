import {
  Button,
  Modal,
  styled,
  Container,
  Toolbar,
  Tooltip,
  ButtonBase,
  MenuItem,
  Menu,
  Badge,
} from "@mui/material";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import React, { useEffect, useState } from "react";
import Typography from "@mui/material/Typography";
import Avatar from "@mui/material/Avatar";
import Grid from "@mui/material/Grid";
import ThumbUpIcon from "@mui/icons-material/ThumbUp";
import moment from "moment";
import { formatDate } from "../../utils";
import Comment from "@mui/icons-material/Comment";
import Visibility from "@mui/icons-material/Visibility";
import MenuIcon from "@mui/icons-material/Menu";
import DeleteIcon from "@mui/icons-material/Delete";
import EditIcon from "@mui/icons-material/Edit";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import { alpha } from "@mui/material/styles";
import { useNavigate, useParams } from "react-router-dom";
import ArrowBackIosNewIcon from "@mui/icons-material/ArrowBackIosNew";

const ProjectView = ({
  project,
  category,
  profile,
  handleLike,
  handleDislike,
  likes,
  handleDelete,
}) => {
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const navigate = useNavigate();
  const [totalLikes, setTotalLikes] = useState(0);
  const [totalViews, setTotalViews] = useState(0);
  const [totalComments, setTotalComments] = useState(0);

  useEffect(() => {
    setTotalLikes(project.totalLikes);
    setTotalViews(project.totalViews);
    setTotalComments(project.totalComments);
  }, [project]);

  const handleViewLike = () => {
    handleLike();
    setTotalLikes(totalLikes + 1);
  };

  const handleViewDislike = () => {
    handleDislike();
    setTotalLikes(totalLikes - 1);
  };

  const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === "dark" ? "#1A2027" : "#fff",
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: "center",
  }));
  const params = useParams();
  const [anchorEl, setAnchorEl] = React.useState(null);
  const openAnchor = Boolean(anchorEl);
  const handleClickAnchor = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleCloseAnchor = () => {
    setAnchorEl(null);
  };

  const handleDeleteProject = () => {
    handleDelete();
    handleCloseAnchor();
  };

  const StyledMenu = styled((props) => (
    <Menu
      elevation={0}
      anchorOrigin={{
        vertical: "bottom",
        horizontal: "right",
      }}
      transformOrigin={{
        vertical: "top",
        horizontal: "right",
      }}
      {...props}
    />
  ))(({ theme }) => ({
    "& .MuiPaper-root": {
      borderRadius: 6,
      marginTop: theme.spacing(1),
      minWidth: 180,
      color:
        theme.palette.mode === "light"
          ? "rgb(55, 65, 81)"
          : theme.palette.grey[300],
      boxShadow:
        "rgb(255, 255, 255) 0px 0px 0px 0px, rgba(0, 0, 0, 0.05) 0px 0px 0px 1px, rgba(0, 0, 0, 0.1) 0px 10px 15px -3px, rgba(0, 0, 0, 0.05) 0px 4px 6px -2px",
      "& .MuiMenu-list": {
        padding: "4px 0",
      },
      "& .MuiMenuItem-root": {
        "& .MuiSvgIcon-root": {
          fontSize: 18,
          color: theme.palette.text.secondary,
          marginRight: theme.spacing(1.5),
        },
        "&:active": {
          backgroundColor: alpha(
            theme.palette.primary.main,
            theme.palette.action.selectedOpacity
          ),
        },
      },
    },
  }));

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
              Creator
            </Typography>
            <Typography variant="h6" sx={{ color: "white" }}>
              {profile.fullName || ""}
            </Typography>
            <Typography variant="h8" sx={{ color: "white" }}>
              {profile.roleName || ""}
            </Typography>
          </Box>
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
                    src={profile.picture || ""}
                    sx={{ width: 125, height: 125 }}
                  />
                </Grid>
                <Grid item>
                  <Typography
                    variant="h4"
                    sx={{ color: "white", marginBottom: "10px" }}
                  >
                    {profile.fullName || ""}
                  </Typography>
                  <Typography variant="h6" sx={{ color: "white" }}>
                    {profile.divisionName || ""}
                  </Typography>
                  <Typography variant="h8" sx={{ color: "white" }}>
                    {profile.roleName || ""}
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
                  {profile.biography || ""}
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Modal>

        <Box
          container
          sx={{
            display: "flex",
            justifyContent: "space-between",
            alignItems: "center",
          }}
        >
          <div>
            <Badge badgeContent={totalLikes} overlap="circular" color="success">
              <Button
                data-testid="thumb-up"
                startIcon={
                  <ThumbUpIcon
                    style={{
                      color: likes.isUserLike === true ? "white" : "#0e80ce",
                    }}
                  />
                }
                style={{
                  maxWidth: "80px",
                  marginLeft: "20px",
                  maxHeight: "30px",
                  minWidth: "80px",
                  minHeight: "40px",
                  border: "1px solid rgba(0, 0, 0, 0.2)",
                  borderRadius: "20px",
                  boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                  backgroundColor:
                    likes.isUserLike === true ? "#0e80ce" : "white",
                }}
                onClick={
                  likes.isUserLike === true
                    ? () => handleViewDislike()
                    : () => handleViewLike()
                }
              />
            </Badge>
            <Badge
              badgeContent={totalComments}
              overlap="circular"
              color="success"
            >
              <Button
                onClick={() => {
                  navigate("/project" + "/" + params.id + "/comment");
                }}
                data-testid="comment"
                startIcon={<Comment style={{ color: "#0e80ce" }} />}
                style={{
                  marginLeft: "20px",
                  maxWidth: "80px",
                  maxHeight: "30px",
                  minWidth: "80px",
                  minHeight: "40px",
                  border: "1px solid rgba(0, 0, 0, 0.2)",
                  borderRadius: "20px",
                  boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                }}
              />
            </Badge>
            <Badge badgeContent={totalViews} overlap="circular" color="success">
              <Button
                data-testid="visibility"
                startIcon={<Visibility style={{ color: "#0e80ce" }} />}
                style={{
                  marginLeft: "20px",
                  maxWidth: "80px",
                  maxHeight: "30px",
                  minWidth: "80px",
                  minHeight: "40px",
                  border: "1px solid rgba(0, 0, 0, 0.2)",
                  borderRadius: "20px",
                  boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                }}
              />
            </Badge>
          </div>
          {project.canEdit === true && project.canDelete === true ? (
            <div>
              <ButtonBase>
                <div>
                  <ButtonBase
                    id="demo-customized-button"
                    aria-controls={
                      openAnchor ? "demo-customized-menu" : undefined
                    }
                    aria-haspopup="true"
                    aria-expanded={open ? "true" : undefined}
                    variant="contained"
                    disableElevation
                    onClick={handleClickAnchor}
                    endIcon={<KeyboardArrowDownIcon />}
                    sx={{ ml: "10px" }}
                  >
                    <MenuIcon />
                  </ButtonBase>
                  <StyledMenu
                    id="demo-customized-menu"
                    MenuListProps={{
                      "aria-labelledby": "demo-customized-button",
                    }}
                    anchorEl={anchorEl}
                    open={openAnchor}
                    onClose={handleCloseAnchor}
                  >
                    <MenuItem
                      onClick={() => {
                        navigate("/project" + "/" + params.id + "/edit");
                      }}
                      disableRipple
                    >
                      <EditIcon />
                      Edit
                    </MenuItem>
                    <MenuItem onClick={handleDeleteProject} disableRipple>
                      <DeleteIcon />
                      Delete
                    </MenuItem>
                  </StyledMenu>
                </div>
              </ButtonBase>
            </div>
          ) : (
            <div></div>
          )}
        </Box>

        <Grid container>
          <Box
            sx={{
              width: "100%",
              borderRadius: "20px",
              boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
              padding: "35px",
            }}
          >
            <Grid container spacing={3}>
              <Grid item xs={12}>
                <Box
                  container
                  sx={{
                    display: "flex",
                    justifyContent: "space-between",
                    alignItems: "center",
                  }}
                >
                  <div>
                    <Item
                      sx={{
                        padding: "10px",
                        margin: "10px",
                        width: "200px",
                      }}
                    >
                      {category.name || ""}
                    </Item>
                  </div>
                  <div>
                    <Tooltip
                      title={formatDate(project.createdAt)}
                      placement="top"
                    >
                      <Typography
                        data-testid="date-pub"
                        variant="h6"
                        sx={{
                          textAlign: "center",
                          backgroundColor: "white",
                          border: "1px solid rgba(0, 0, 0, 0.2)",
                          borderRadius: "30px",
                          paddingLeft: "30px",
                          paddingRight: "30px",
                          paddingTop: "5px",
                          paddingBottom: "5px",
                          boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
                          width: "200px",
                        }}
                      >
                        Published, {moment.utc(project.createdAt).fromNow()}
                      </Typography>
                    </Tooltip>
                  </div>
                </Box>
              </Grid>
              <Grid
                item
                xs={12}
                sx={{
                  display: "flex",
                  justifyContent: "center",
                }}
              >
                <Box
                  sx={{
                    padding: "10px",
                    margin: "10px",
                    width: "900px",
                    overflow: "hidden",
                  }}
                >
                  {project.story}
                </Box>
              </Grid>
            </Grid>
          </Box>
        </Grid>

        <Toolbar />
      </Container>
    </>
  );
};

export default ProjectView;
