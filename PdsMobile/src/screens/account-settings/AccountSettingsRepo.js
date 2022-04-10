import {
  API_ROLE,
  API_DIVISION,
  API_PROFILE,
  API_PROFILE_PASSWORD,
} from "../../config";

const AccountSettingsRepo = (axios, useFetchCall) => {
  const useUpdateAccountData = () => {
    const detailAccountData = async dataToPut => {
      try {
        const res = await axios.put(API_PROFILE, dataToPut);
        const data = res.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(detailAccountData);
  };
  const useSendChangePasswordLink = () => {
    const detailAccountData = async email => {
      try {
        const res = await axios.post(API_PROFILE_PASSWORD, {email});
        const data = res.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(detailAccountData);
  };
  const useDetailAccountData = () => {
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
  return {
    useRoleData,
    useDivisionData,
    useDetailAccountData,
    useUpdateAccountData,
    useSendChangePasswordLink,
  };
};

export default AccountSettingsRepo;
