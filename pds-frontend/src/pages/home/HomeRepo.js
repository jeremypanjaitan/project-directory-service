import { API_CATEGORY, API_PROJECT, axios } from "../../config";
import { useFetchCall } from "../../hooks";
import { PAGE_SIZE_HOME } from "../../constants";

export const useAllProjectData = () => {
  const allProjectData = async ({
    page,
    searchByTitle,
    sortByLikes,
    categoryId,
  }) => {
    try {
      const res = await axios.get(API_PROJECT, {
        params: {
          pageNumber: page,
          pageSize: PAGE_SIZE_HOME,
          searchByTitle,
          sortByLikes,
          categoryId,
        },
      });
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(allProjectData);
};

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
