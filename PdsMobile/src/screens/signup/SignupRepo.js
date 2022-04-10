import {API_ROLE, API_DIVISION, API_REGISTER} from "../../config";

const SignupRepo = (axios, useFetchCall) => {
  const useRoleData = () => {
    const roleData = async () => {
      try {
        const res = await axios.get(API_ROLE);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(roleData);
  };

  const useDivisionData = () => {
    const divisionData = async () => {
      try {
        const res = await axios.get(API_DIVISION);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(divisionData);
  };

  const useRegister = () => {
    const addRegister = async dataToPost => {
      try {
        const res = await axios.post(API_REGISTER, dataToPost);
        const data = res.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(addRegister);
  };
  return {useRegister, useDivisionData, useRoleData};
};

export default SignupRepo;
