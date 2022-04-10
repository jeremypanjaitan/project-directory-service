import {
  API_CATEGORY,
  API_COMMENT,
  API_PROFILE,
  API_PROJECT,
  axios,
} from "../../config";
import { useFetchCall } from "../../hooks";
import { COMMENT_SIZE } from "../../constants";

export const useCreateComment = () => {
  const commentData = async (data) => {
    try {
      const url = API_COMMENT + "/" + data.id;
      const res = await axios.post(url, { body: data.dataToPost });
      return res.data.description;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(commentData);
};

export const useCommentData = () => {
  const commentData = async ({ id, page }) => {
    try {
      const url = API_COMMENT + "/" + id;
      const res = await axios.get(url, {
        params: { pageNumber: page, pageSize: COMMENT_SIZE },
      });
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(commentData);
};

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
