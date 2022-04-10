import React from "react";
import MainLayoutView from "./MainLayoutView";
import { useLogout } from "../../../services";
import { useLoading } from "../../../context";
import { useAuth } from "../../../services";
import Swal from "sweetalert2";

const MainLayoutContainer = ({
  children,
  _useLoading = useLoading,
  _useAuth = useAuth,
  _useLogout = useLogout,
  MainLayoutViewTest = MainLayoutView,
}) => {
  const logout = _useLogout();
  const { setUser } = _useAuth();
  const { setIsLoading, isLoading } = _useLoading();
  const drawerWidth = 320;
  const menu = {
    home: "Home",
    accountSettings: "Account Settings",
    logout: "Logout",
  };
  const handleLogout = () => {
    Swal.fire({
      title: "Are you sure want to logout?",
      showCancelButton: true,
      confirmButtonText: "OK",
    }).then((result) => {
      if (result.isConfirmed) {
        setIsLoading(true);
        logout.start();
      }
    });
  };

  if (logout.isSuccess) {
    setIsLoading(false);
    setUser(null);
  }
  return (
    <>
      <MainLayoutViewTest
        drawerWidth={drawerWidth}
        menu={menu}
        handleLogout={handleLogout}
        isLoading={isLoading}
      >
        {children}
      </MainLayoutViewTest>
    </>
  );
};

export default MainLayoutContainer;
