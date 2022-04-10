import { API_PROFILE_ACTIVITY, API_PROFILE_PROJECT, axios } from "../../config";
import { useFetchCall } from "../../hooks";
import { PAGE_SIZE_ACTIVITY, PAGE_SIZE_PROJECT } from "../../constants";

export const useProfileProject = () => {
  const profileProject = async ({ page }) => {
    try {
      const res = await axios.get(API_PROFILE_PROJECT, {
        params: { pageNumber: page, pageSize: PAGE_SIZE_PROJECT },
      });
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(profileProject);
};

export const useProfileActivity = () => {
  const profileActivity = async ({ page }) => {
    try {
      const res = await axios.get(API_PROFILE_ACTIVITY, {
        params: { pageNumber: page, pageSize: PAGE_SIZE_ACTIVITY },
      });
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(profileActivity);
};
