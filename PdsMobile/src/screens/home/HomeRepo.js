import {API_CATEGORY, API_PROJECT, PAGE_SIZE} from "../../config";
const HomeRepo = (axios, useFetchCall) => {
  const useAllProjectData = () => {
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
            pageSize: PAGE_SIZE,
            searchByTitle,
            sortByLikes,
            categoryId,
          },
        });
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(allProjectData);
  };
  const useAllCategoryData = () => {
    const allCategoryData = async () => {
      try {
        const res = await axios.get(API_CATEGORY);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(allCategoryData);
  };

  return {
    useAllProjectData,
    useAllCategoryData,
  };
};

export default HomeRepo;
