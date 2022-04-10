import {API_PROFILE_PROJECT, PAGE_SIZE_PROJECT} from "../../../config";

const AccountDetailsProjectRepo = (axios, useFetchCall) => {
  const useAllProfileProject = () => {
    const allProfileProject = async ({page}) => {
      try {
        const res = await axios.get(API_PROFILE_PROJECT, {
          params: {
            pageNumber: page,
            pageSize: PAGE_SIZE_PROJECT,
          },
        });
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(allProfileProject);
  };

  return {
    useAllProfileProject,
  };
};

export default AccountDetailsProjectRepo;
