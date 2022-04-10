import React from "react";
import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List";
import Typography from "@mui/material/Typography";
import ListItem from "@mui/material/ListItem";
import ListItemIcon from "@mui/material/ListItemIcon";
import ListItemText from "@mui/material/ListItemText";
import HomeIcon from "@mui/icons-material/Home";
import SettingsIcon from "@mui/icons-material/Settings";
import LogoutIcon from "@mui/icons-material/Logout";
import Avatar from "@mui/material/Avatar";
import Divider from "@mui/material/Divider";
import { ButtonBase } from "@mui/material";
import LoadingOverlay from "react-loading-overlay-ts";
import { useLocation, useNavigate } from "react-router-dom";
import { MAIN, ACCOUNT_SETTINGS, ACCOUNT_DETAILS } from "../../../router";
import { useAuth } from "../../../services/auth";

import { colors } from "../../../constants";

const MainLayoutView = ({
  menu,
  drawerWidth,
  handleLogout,
  children,
  isLoading,
}) => {
  const location = useLocation();
  const navigate = useNavigate();

  const auth = useAuth();
  return (
    <Box sx={{ display: "flex" }}>
      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          "& .MuiDrawer-paper": {
            width: drawerWidth,
            boxSizing: "border-box",
            background: `linear-gradient(200.56deg, ${colors.BLUEBERRY} 28.59%, ${colors.MIDDLE_BLUE}  106.16%)`,
            boxShadow: "1px -1px 21px -6px rgba(0,0,0,0.83)",
          },
        }}
        variant="permanent"
        anchor="left"
      >
        {location.pathname !== ACCOUNT_SETTINGS &&
        location.pathname !== ACCOUNT_DETAILS ? (
          <ButtonBase onClick={() => navigate(ACCOUNT_DETAILS)}>
            <Box
              sx={{
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
                gap: "20px",
                backgroundColor:
                  "linear-gradient(229.93deg, #6685E3 -0.7%, rgba(102, 133, 227, 0) 115.3%)",
                borderRadius: "20px",
                margin: "15px",
              }}
            >
              <Avatar
                alt="Remy Sharp"
                src={auth?.user?.picture}
                sx={{ width: "55px", height: "55px" }}
              />
              <Box sx={{ width: "60%" }}>
                <Typography
                  noWrap
                  variant="h6"
                  sx={{
                    color: "white",
                    textAlign: "center",
                    fontFamily: "montserrat",
                    fontWeight: "600",
                    fontSize: "18px",
                  }}
                >
                  {auth?.user?.fullName}
                </Typography>
                <Typography
                  noWrap
                  variant="h8"
                  component="h5"
                  sx={{
                    color: "white",
                    textAlign: "center",
                    fontFamily: "montserrat",
                    fontWeight: "600",
                  }}
                >
                  {auth?.user?.roleName}
                </Typography>
              </Box>
            </Box>
          </ButtonBase>
        ) : null}

        <Divider />
        <List>
          <ListItem
            button
            key={menu.home}
            disabled={location.pathname === MAIN}
            onClick={() => navigate(MAIN)}
          >
            <ListItemIcon>
              <HomeIcon sx={{ color: "white" }} />
            </ListItemIcon>
            <ListItemText primary={menu.home} sx={{ color: "white" }} />
          </ListItem>

          <ListItem
            button
            key={menu.accountSettings}
            disabled={location.pathname === ACCOUNT_SETTINGS}
            onClick={() => navigate(ACCOUNT_SETTINGS)}
          >
            <ListItemIcon>
              <SettingsIcon sx={{ color: "white" }} />
            </ListItemIcon>
            <ListItemText
              primary={menu.accountSettings}
              sx={{ color: "white" }}
            />
          </ListItem>

          <ListItem button key={menu.logout} onClick={handleLogout}>
            <ListItemIcon>
              <LogoutIcon sx={{ color: "white" }} />
            </ListItemIcon>
            <ListItemText primary={menu.logout} sx={{ color: "white" }} />
          </ListItem>
        </List>
      </Drawer>
      <Box
        component="main"
        sx={{
          flexGrow: 1,
          bgcolor: "background.default",
          p: 3,
          height: "100vh",
          padding: 0,
        }}
      >
        <LoadingOverlay active={isLoading}>{children}</LoadingOverlay>
      </Box>
    </Box>
  );
};
export default MainLayoutView;
