import {API_CATEGORY, API_PROJECT} from "../../config";

const EditProjectRepo = (axios, useFetchCall) => {
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

  const useEditProject = () => {
    const addProject = async ({id, values}) => {
      try {
        const url = API_PROJECT + "/" + id;
        const res = await axios.put(url, {
          title: values.title,
          picture: values.picture,
          story: values.story,
          description: values.description,
          categoryId: values.categoryId,
        });
        const data = res.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(addProject);
  };
  const useProjectData = () => {
    const projectData = async id => {
      try {
        const res = await axios.get(API_PROJECT + "/" + id);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(projectData);
  };
  return {
    useEditProject,
    useCategoryData,
    useProjectData,
  };
};

export default EditProjectRepo;
