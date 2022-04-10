import { API_CATEGORY, API_PROJECT, axios } from "../../config";
import { useFetchCall } from "../../hooks";

export const useCategoryData = () => {
  const categoryData = async () => {
    try {
      const res = await axios.get(API_CATEGORY);
      const data = res.data.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(categoryData);
};

export const useCreateProject = () => {
  const addProject = async (dataToPost) => {
    try {
      const res = await axios.post(API_PROJECT, dataToPost);
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(addProject);
};
