import {API_CATEGORY, API_PROJECT} from "../../config";

const CreateProjectRepo = (axios, useFetchCall) => {
  const useCategoryData = () => {
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

  const useCreateProject = () => {
    const addProject = async (dataToPost) => {
      try {
        const res = await axios.post(API_PROJECT, dataToPost);
        const data = res.data;
        console.log(data)
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(addProject);
  };
  return {
    useCreateProject,
    useCategoryData,
  };
};

export default CreateProjectRepo;
