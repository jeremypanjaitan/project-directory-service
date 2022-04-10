import {
  API_CATEGORY,
  API_DISLIKE,
  API_LIKE,
  API_PROFILE,
  API_PROJECT,
  API_VIEWS,
} from "../../config";

const ProjectRepo = (axios, useFetchCall) => {
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
  const useDeleteProject = () => {
    const projectData = async id => {
      try {
        const res = await axios.delete(API_PROJECT + "/" + id);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(projectData);
  };
  const useProjectLike = () => {
    const projectLike = async id => {
      try {
        const res = await axios.get(API_LIKE(id));
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(projectLike);
  };
  const useLike = () => {
    const like = async id => {
      try {
        const res = await axios.post(API_LIKE(id));
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(like);
  };
  const useDislike = () => {
    const like = async id => {
      try {
        const res = await axios.delete(API_DISLIKE(id));
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(like);
  };
  const useProfileData = () => {
    const profileData = async id => {
      try {
        const res = await axios.get(API_PROFILE + "/" + id);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(profileData);
  };
  const useCategoryData = () => {
    const categoryData = async id => {
      try {
        const res = await axios.get(API_CATEGORY + "/" + id);
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(categoryData);
  };
  const useView = () => {
    const view = async id => {
      try {
        const res = await axios.post(API_VIEWS(id));
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(view);
  };
  const useTotalView = () => {
    const totalView = async id => {
      try {
        const res = await axios.get(API_VIEWS(id));
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(totalView);
  };
  return {
    useProjectData,
    useProfileData,
    useCategoryData,
    useProjectLike,
    useLike,
    useDislike,
    useView,
    useTotalView,
    useDeleteProject,
  };
};

export default ProjectRepo;
