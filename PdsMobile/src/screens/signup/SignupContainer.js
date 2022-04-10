import React, {useEffect} from "react";

const SignupContainer = ({
  children,
  useDivisionData,
  useRoleData,
  useRegister,
}) => {
  const divisionData = useDivisionData();
  const roleData = useRoleData();
  const register = useRegister();
  const handleSubmit = values => {
    const {reEnterPassword, ...credential} = values;
    register.start(credential);
  };
  useEffect(() => {
    divisionData.start();
    roleData.start();
  }, []);
  return React.cloneElement(children, {
    divisionDataIsError: divisionData.isError,
    divisionDataError: divisionData.error,
    divisionDataIsLoading: divisionData.isLoading,
    divisionData: divisionData.data?.map(d => ({value: d.id, label: d.name})),
    roleDataIsError: roleData.isError,
    roleDataError: roleData.error,
    roleDataIsLoading: roleData.isLoading,
    roleData: roleData.data?.map(d => ({value: d.id, label: d.name})),

    registerIsError: register.isError,
    registerError: register.error,
    registerIsLoading: register.isLoading,
    registerReset: register.reset,
    registerIsSuccess: register.isSuccess,
    registerData: register.data,
    handleSubmit,
  });
};

export default SignupContainer;
