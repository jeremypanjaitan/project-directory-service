import {API_PROFILE_PASSWORD} from "../../config";
const ForgetPasswordRepo = (axios, useFetchCall) => {
  const useSendForgetPasswordLink = () => {
    const addRegister = async dataToPost => {
      try {
        const res = await axios.post(API_PROFILE_PASSWORD, dataToPost);
        const data = res.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(addRegister);
  };
  return {useSendForgetPasswordLink};
};

export default ForgetPasswordRepo;
