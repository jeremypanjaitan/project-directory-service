import { API_DISLIKE, API_LIKE, API_VIEWS, axios } from "../../config";
import { useFetchCall } from "../../hooks";
import {
  API_PROJECT,
  API_PROFILE,
  API_CATEGORY,
  API_LIKES,
} from "../../config";
import {} from "../../utils";

export const useProjectData = () => {
  const projectData = async (id) => {
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

export const useDeleteProject = () => {
  const projectData = async (id) => {
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

export const useProfileData = () => {
  const profileData = async (id) => {
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

export const useCategoryData = () => {
  const categoryData = async (id) => {
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

export const useLikesData = () => {
  const likesData = async (id) => {
    try {
      const res = await axios.get(API_LIKES + "/" + id + "/" + API_LIKE);
      const data = res.data.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(likesData);
};

export const useCreateLikes = () => {
  const likesData = async (id) => {
    try {
      const res = await axios.post(API_LIKES + "/" + id + "/" + API_LIKE);
      const data = res.data.description;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(likesData);
};

export const useCreateViews = () => {
  const viewsData = async (id) => {
    try {
      const res = await axios.post(API_VIEWS + "/" + id);
      const data = res.data.description;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(viewsData);
};

export const useCreateDislike = () => {
  const dislikeData = async (id) => {
    try {
      const res = await axios.delete(API_LIKES + "/" + id + "/" + API_DISLIKE);
      const data = res.data.description;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(dislikeData);
};
