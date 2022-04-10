import React, { useState } from "react";
import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Grid from "@mui/material/Grid";
import SearchIcon from "@mui/icons-material/Search";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import Pagination from "@mui/material/Pagination";
import { styled, alpha } from "@mui/material/styles";
import Toolbar from "@mui/material/Toolbar";
import InputBase from "@mui/material/InputBase";
import Box from "@mui/material/Box";
import { ButtonBase, CardActions } from "@mui/material";
import ThumbUpOutlinedIcon from "@mui/icons-material/ThumbUpOutlined";
import ChatBubbleOutlineOutlinedIcon from "@mui/icons-material/ChatBubbleOutlineOutlined";
import VisibilityOutlinedIcon from "@mui/icons-material/VisibilityOutlined";
import MenuIcon from "@mui/icons-material/Menu";
import CategoryIcon from "@mui/icons-material/Category";
import ThumbUpTwoToneIcon from "@mui/icons-material/ThumbUpTwoTone";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import KeyboardArrowDownIcon from "@mui/icons-material/KeyboardArrowDown";
import { ASCENDING, colors, DESCENDING } from "../../constants";
import { useNavigate } from "react-router-dom";
import { SHOW } from "../../router";
import AddIcon from "@mui/icons-material/Add";
import List from "@mui/material/List";
import ListItemButton from "@mui/material/ListItemButton";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import Collapse from "@mui/material/Collapse";
import ExpandLess from "@mui/icons-material/ExpandLess";
import ExpandMore from "@mui/icons-material/ExpandMore";
import VisibilityTwoToneIcon from "@mui/icons-material/VisibilityTwoTone";
import { useLoading } from "../../context";

const Search = styled("div")(({ theme }) => ({
  position: "relative",
  borderRadius: theme.shape.borderRadius,
  backgroundColor: alpha(theme.palette.common.white, 0.15),
  "&:hover": {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  marginLeft: 0,
  width: "100%",
  [theme.breakpoints.up("sm")]: {
    marginLeft: theme.spacing(1),
    width: "auto",
  },
}));

const SearchIconWrapper = styled("div")(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: "100%",
  position: "absolute",
  pointerEvents: "none",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
  color: "inherit",
  boxShadow: "0px 0px 2px 2px rgba(0, 0, 0, 0.1)",
  borderRadius: "10px",
  "& .MuiInputBase-input": {
    padding: theme.spacing(1, 1, 1, 0),
    paddingLeft: `calc(1em + ${theme.spacing(4)})`,
    width: "30ch",
  },
}));
const HomeView = ({
  data,
  handleNavigateAddPage,
  currentPage,
  setCurrentPage,
  totalPage,
  handleSearch,
  handleLike,
  handleCategoryId,
  categories,
}) => {
  const navigate = useNavigate();
  const [inputSearchValue, setInputSearchValue] = useState("");
  const handleKeypress = (e) => {
    if (e.keyCode === 13) {
      handleSearch(inputSearchValue);
    }
  };

  const { setIsLoading } = useLoading();
  const [header, setHeader] = React.useState("");
  const [subHeader, setSubHeader] = React.useState("");

  const [anchorEl, setAnchorEl] = React.useState(null);
  const open = Boolean(anchorEl);
  const handleClick = (event) => {
    setAnchorEl(event.currentTarget);
  };
  const handleClose = () => {
    setAnchorEl(null);
  };

  const [openItemsCategories, setOpenItemsCategories] = React.useState(false);

  const handleClickItemsCategories = () => {
    setOpenItemsCategories(!openItemsCategories);
  };

  const [openItemsLikes, setOpenItemsLikes] = React.useState(false);

  const handleClickItemsLikes = () => {
    setOpenItemsLikes(!openItemsLikes);
  };

  console.log("data", data);
  console.log("categories", categories);
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
        maxWidth="xl"
      >
        <Toolbar
          sx={{ display: "flex", justifyContent: "space-between", padding: 0 }}
        >
          <Button
            sx={{
              display: "flex",
              boxShadow: "1px -1px 15px -6px rgba(0,0,0,0.83)",
              background: `linear-gradient(234.28deg, ${colors.BLUEBERRY} -70.98%, ${colors.VERY_SOFT_CYAN} 154.01%)`,
              borderRadius: "20px",
              margin: 0,
              width: "216 px",
              height: "52 px",
              mt: "41 px",
              ml: "488 px",
              padding: "11 px",
            }}
            onClick={handleNavigateAddPage}
            variant="contained"
            style={{
              textTransform: "none",
              fontFamily: "montserrat",
              fontSize: "15px",
              fontWeight: "700",
            }}
          >
            <AddIcon />
            Create Project
          </Button>
          <Toolbar>
            <Search>
              <SearchIconWrapper>
                <SearchIcon />
              </SearchIconWrapper>
              <StyledInputBase
                placeholder="Search by title…"
                inputProps={{ "aria-label": "search" }}
                onChange={(e) => setInputSearchValue(e.target.value)}
                onKeyDown={(e) => {
                  handleKeypress(e);
                }}
              />
            </Search>
            <ButtonBase>
              <ButtonBase
                id="demo-customized-button"
                aria-controls={open ? "demo-customized-menu" : undefined}
                aria-haspopup="true"
                aria-expanded={open ? "true" : undefined}
                variant="contained"
                disableElevation
                onClick={handleClick}
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
                open={open}
                onClose={handleClose}
              >
                <List
                  sx={{
                    width: "100%",
                    maxWidth: 360,
                    bgcolor: "background.paper",
                  }}
                >
                  <ListItemButton
                    onClick={() => {
                      setHeader("");
                      setSubHeader("");
                      handleSearch(null);
                      handleCategoryId(null);
                      handleLike(null);
                      handleClose();
                      setIsLoading(true);
                    }}
                    disableRipple
                  >
                    <ListItemIcon>
                      <VisibilityTwoToneIcon />
                      <Typography>All</Typography>
                    </ListItemIcon>
                  </ListItemButton>
                  <ListItemButton
                    onClick={handleClickItemsCategories}
                    disableRipple
                  >
                    <ListItemIcon>
                      <CategoryIcon />
                      <Typography>Category</Typography>
                    </ListItemIcon>
                    <ListItemText />
                    {openItemsCategories ? <ExpandLess /> : <ExpandMore />}
                  </ListItemButton>
                  <Collapse
                    in={openItemsCategories}
                    timeout="auto"
                    unmountOnExit
                  >
                    <List component="div" disablePadding>
                      <ListItemButton sx={{ pl: 4 }}>
                        <Typography>
                          {categories.map((option, key) => (
                            <MenuItem
                              key={key}
                              value={option.id}
                              onClick={() => {
                                handleCategoryId(option.id);
                                handleClose();
                                setHeader(option.name);
                                setIsLoading(true);
                              }}
                            >
                              {option.name}
                            </MenuItem>
                          ))}
                        </Typography>
                      </ListItemButton>
                    </List>
                  </Collapse>
                </List>

                <List
                  sx={{
                    width: "100%",
                    maxWidth: 360,
                    bgcolor: "background.paper",
                  }}
                  component="nav"
                  aria-labelledby="nested-list-subheader"
                >
                  <ListItemButton onClick={handleClickItemsLikes} disableRipple>
                    <ListItemIcon>
                      <ThumbUpTwoToneIcon />
                      <Typography>Like</Typography>
                    </ListItemIcon>
                    <ListItemText />
                    {openItemsLikes ? <ExpandLess /> : <ExpandMore />}
                  </ListItemButton>
                  <Collapse in={openItemsLikes} timeout="auto" unmountOnExit>
                    <List component="div" disablePadding>
                      <ListItemButton sx={{ pl: 4 }}>
                        <Typography sx={{ ml: "10px" }}>
                          <MenuItem
                            onClick={() => {
                              setSubHeader("Most Likes");
                              handleLike(DESCENDING);
                              handleClose();
                              setIsLoading(true);
                            }}
                          >
                            Most likes
                          </MenuItem>
                          <MenuItem
                            onClick={() => {
                              setSubHeader("Least Likes");
                              handleLike(ASCENDING);
                              handleClose();
                              setIsLoading(true);
                            }}
                          >
                            Least likes
                          </MenuItem>
                        </Typography>
                      </ListItemButton>
                    </List>
                  </Collapse>
                </List>
              </StyledMenu>
            </ButtonBase>
          </Toolbar>
        </Toolbar>

        {data.length === 0 ? (
          <>
            {header === "" ? (
              <h2
                style={{
                  fontFamily: "montserrat",
                  textAlign: "center",
                }}
              >
                All Projects
                <h6
                  style={{
                    fontFamily: "montserrat",
                    textAlign: "center",
                  }}
                >
                  {subHeader}
                </h6>
              </h2>
            ) : (
              <h2
                style={{
                  fontFamily: "montserrat",
                  textAlign: "center",
                }}
              >
                Category: {header}
                <h6
                  style={{
                    fontFamily: "montserrat",
                    textAlign: "center",
                  }}
                >
                  {subHeader}
                </h6>
              </h2>
            )}
            <h1
              style={{
                fontFamily: "montserrat",
                margin: "auto",
              }}
            >
              No Data ...
            </h1>
          </>
        ) : (
          <>
            {header === "" ? (
              <h2
                style={{
                  fontFamily: "montserrat",
                  textAlign: "center",
                }}
              >
                All Projects
                <h6
                  style={{
                    fontFamily: "montserrat",
                    textAlign: "center",
                  }}
                >
                  {subHeader}
                </h6>
              </h2>
            ) : (
              <h2
                style={{
                  fontFamily: "montserrat",
                  textAlign: "center",
                }}
              >
                Category: {header}
                <h6
                  style={{
                    fontFamily: "montserrat",
                    textAlign: "center",
                  }}
                >
                  {subHeader}
                </h6>
              </h2>
            )}
            <Grid container spacing={5}>
              {data?.map((d) => (
                <Grid item xs={12} sm={6} md={4}>
                  <ButtonBase
                    onClick={() => {
                      setIsLoading(true);
                      navigate(SHOW + "/" + d.ID);
                    }}
                  >
                    <Card
                      sx={{
                        display: "flex",
                        flexDirection: "column",
                        width: "300px",
                      }}
                    >
                      <CardMedia
                        component="img"
                        image={d.picture}
                        sx={{ height: "300px" }}
                      />
                      <CardContent>
                        <Typography
                          gutterBottom
                          variant="h5"
                          component="h5"
                          noWrap
                        >
                          {d.title}
                        </Typography>
                        <Typography noWrap>{d.description}</Typography>
                      </CardContent>
                      <CardActions>
                        <Grid container>
                          <Grid item xs={1} sx={{ ml: "20px" }}>
                            <Typography>{d.totalLikes}</Typography>
                          </Grid>
                          <Grid item>
                            <ThumbUpOutlinedIcon />
                          </Grid>
                          <Grid item xs={1} sx={{ ml: "40px" }}>
                            <Typography>{d.totalComments}</Typography>
                          </Grid>
                          <Grid item>
                            <ChatBubbleOutlineOutlinedIcon />
                          </Grid>
                          <Grid item xs={1} sx={{ ml: "50px" }}>
                            <Typography>{d.totalViews}</Typography>
                          </Grid>
                          <Grid item>
                            <VisibilityOutlinedIcon />
                          </Grid>
                        </Grid>
                      </CardActions>
                    </Card>
                  </ButtonBase>
                </Grid>
              ))}
            </Grid>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
              }}
            >
              <Pagination
                variant="outlined"
                shape="rounded"
                count={totalPage}
                page={currentPage}
                color="secondary"
                sx={{ marginBottom: "40px" }}
                onChange={(e, page) => {
                  window.scrollTo({ top: 0, left: 0, behavior: "smooth" });
                  setCurrentPage(page);
                }}
              />
            </Box>
          </>
        )}
      </Container>
    </>
  );
};

export default HomeView;
