import { API_DIVISION, API_ROLE, axios, API_REGISTER } from "../../config";
import { useFetchCall } from "../../hooks";

export const useRoleData = () => {
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

export const useDivisionData = () => {
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

export const useRegister = () => {
  const addRegister = async (dataToPost) => {
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

