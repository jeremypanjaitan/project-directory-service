import React from "react";
import Box from "@mui/material/Box";
import LoadingOverlay from "react-loading-overlay-ts";
import { useLoading } from "../../../context";

const Signup = ({ children, _useLoading = useLoading }) => {
  const { isLoading } = _useLoading();
  return (
    <LoadingOverlay active={isLoading}>
      <Box sx={{ height: "100vh" }}>{children}</Box>
    </LoadingOverlay>
  );
};
export default Signup;
