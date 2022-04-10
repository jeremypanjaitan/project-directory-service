import {API_AUTH} from "../../config";
import {saveAuthToken, saveRefreshToken} from "../../utils";
const LoginRepo = (axios, useFetchCall) => {
  const useLogin = () => {
    const login = async credential => {
      try {
        const res = await axios.post(API_AUTH, credential);
        const data = res.data.data;
        console.log("credential", credential);
        saveAuthToken(data.tokenData.accessToken);
        saveRefreshToken(data.tokenData.refreshToken);
        return data;
      } catch (err) {
        console.log(err);
        throw err;
      }
    };
    return useFetchCall(login);
  };
  return {
    useLogin,
  };
};

export default LoginRepo;
