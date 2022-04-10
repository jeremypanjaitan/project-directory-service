import {API_COMMENT, PAGE_SIZE_COMMENT} from "../../config";

const CommentRepo = (axios, useFetchCall) => {
  const useComment = () => {
    const projectData = async dataToPost => {
      try {
        const res = await axios.get(API_COMMENT(dataToPost.id), {
          params: {
            pageNumber: dataToPost.pageNumber,
            pageSize: PAGE_SIZE_COMMENT,
          },
        });
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(projectData);
  };
  const usePostComment = () => {
    const projectData = async dataToPost => {
      try {
        const res = await axios.post(API_COMMENT(dataToPost.id), {
          body: dataToPost.body,
        });
        const data = res.data.data;
        return data;
      } catch (err) {
        throw err;
      }
    };
    return useFetchCall(projectData);
  };
  return {
    useComment,
    usePostComment,
  };
};

export default CommentRepo;
