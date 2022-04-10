import React, { useEffect, useState } from "react";
import SignupView from "./SignupView";
import { useDivisionData, useRoleData, useRegister } from "./SignupRepo";
import { useLoading } from "../../context";
import { RouteNavigation } from "../../utils/navigation";
import Swal from "sweetalert2";

const SignupContainer = ({
  children,
  _useRoleData = useRoleData,
  _useDivisionData = useDivisionData,
  _useLoading = useLoading,
  _useRegister = useRegister,
  _RouteNavigation = RouteNavigation,
  SignupViewTest = SignupView,
}) => {
  const roleData = _useRoleData();
  const divisionData = _useDivisionData();
  const [division, setDivision] = useState();
  const [role, setRole] = useState();
  const { setIsLoading } = _useLoading();
  const register = _useRegister();
  const navigateTo = _RouteNavigation();

  const handleSubmit = (values) => {
    Swal.fire({
      title: "Do you want to create this user?",
      showCancelButton: true,
      confirmButtonText: "Create user",
    }).then((result) => {
      /* Read more about isConfirmed, isDenied below */
      if (result.isConfirmed) {
        setIsLoading(true);
        register.start(values);
      }
    });
  };
  if (register.isSuccess) {
    Swal.fire("User Created, please check your email!", "", "success");
    navigateTo(-1);
    setIsLoading(false);
    register.reset();
  }
  if (register.isError) {
    register.reset();
    Swal.fire(register.error.response.data.description, "", "warning");
    setIsLoading(false);
  }

  useEffect(() => {
    // setIsLoading(true);
    roleData.start();
    divisionData.start();
    //eslint-disable-next-line
  }, []);
  useEffect(() => {
    if (roleData.isSuccess && divisionData.isSuccess) {
      setRole(roleData.data);
      setDivision(divisionData.data);
      roleData.reset();
      divisionData.reset();
      setIsLoading(false);
    }
    //eslint-disable-next-line
  }, [roleData.isSuccess, divisionData.isSuccess, setIsLoading]);

  return (
    <>
      <SignupViewTest
        roles={role || []}
        divisions={division || []}
        handleSubmit={handleSubmit}
        navigation={navigateTo}
      >
        {children}
      </SignupViewTest>
    </>
  );
};

export default SignupContainer;
