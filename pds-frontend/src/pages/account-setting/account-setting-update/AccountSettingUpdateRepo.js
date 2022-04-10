import { API_PROFILE, API_PROFILE_PASSWORD, axios } from "../../../config";
import { useFetchCall } from "../../../hooks";

export const useDetailAccountData = () => {
  const detailAccountData = async () => {
    try {
      const res = await axios.get(API_PROFILE);
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(detailAccountData);
};

export const useUpdateAccountData = () => {
  const detailAccountData = async (dataToPut) => {
    try {
      const res = await axios.put(API_PROFILE, dataToPut);
      const data = res.data;
      return data;
    } catch (err) {
      console.log(err.response);
      throw err;
    }
  };
  return useFetchCall(detailAccountData);
};

export const usePostChangePasswordData = () => {
  const postPasswordData = async (email) => {
    try {
      const res = await axios.post(API_PROFILE_PASSWORD, email);
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(postPasswordData);
};
