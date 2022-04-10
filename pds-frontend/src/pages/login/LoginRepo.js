import { API_PROFILE_PASSWORD, axios } from "../../../src/config";
import { useFetchCall } from "../../../src/hooks";
export const usePostForgetPasswordData = () => {
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
