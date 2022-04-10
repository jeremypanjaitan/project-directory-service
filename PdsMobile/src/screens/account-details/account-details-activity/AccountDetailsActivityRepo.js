import {API_PROFILE_ACTIVITY, PAGE_SIZE_ACTIVITY} from "../../../config";

const AccountDetailsActivityRepo = (axios, useFetchCall) => {
  const useAllProfileActivity = () => {
    const allProfileActivity = async ({page}) => {
      try {
        const res = await axios.get(API_PROFILE_ACTIVITY, {
          params: {
            pageNumber: page,
            pageSize: PAGE_SIZE_ACTIVITY,
          },
        });
        const data = res.data.data;
        console.log("data profile activity", data);
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(allProfileActivity);
  };

  return {
    useAllProfileActivity,
  };
};

export default AccountDetailsActivityRepo;
