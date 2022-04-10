import React, {useState, useEffect} from "react";
import moment from "moment";

const CommentContainer = ({
  children,
  useRoute,
  useAuth,
  useComment,
  usePostComment,
  useGlobal,
}) => {
  const globalState = useGlobal();
  const commentData = useComment();
  const auth = useAuth();
  const route = useRoute();
  const postComment = usePostComment();
  const [currentPage, setCurrentPage] = useState(1);
  const [comments, setComments] = useState([]);
  const [isEndReachRefresh, setIsEndReachRefresh] = useState(false);
  const [isPullRefresh, setIsPullRefresh] = useState(false);
  const [isEndPage, setIsEndPage] = useState(false);
  useEffect(() => {
    commentData.start({id: route.params.projectId, pageNumber: currentPage});
  }, []);

  const handlePostComment = values => {
    const newComments = [...comments];
    const commentToPost = {
      from: auth.userData.data.fullName,
      body: values.body,
      picture: auth.userData.data.picture,
      createdAt: moment(),
    };
    newComments.unshift(commentToPost);
    setComments(newComments);
    globalState.setTotalComment(globalState.totalComment + 1);
    postComment.start({body: values.body, id: route.params.projectId});
  };
  useEffect(() => {
    if (isPullRefresh) {
      setIsEndPage(false);
      setCurrentPage(1);
      commentData.start({
        pageNumber: 1,
        searchQuery: "",
        id: route.params.projectId,
      });
      commentData.reset();
    }
  }, [isPullRefresh]);
  useEffect(() => {
    if (isPullRefresh && commentData.isSuccess) {
      setComments(commentData?.data?.row);
      setIsPullRefresh(false);
      commentData.reset();
    }
  }, [isPullRefresh, commentData.isSuccess]);
  useEffect(() => {
    if (postComment.isError) {
      const newComments = [...comments];
      newComments.shift();
      setComments(newComments);
      if (globalState.totalComment > 0) {
        globalState.setTotalComment(globalState.totalComment - 1);
      }
    }
  }, [postComment.isError]);
  useEffect(() => {
    if (isEndReachRefresh && !isEndPage) {
      if (commentData?.data?.row?.length === 0) {
        setIsEndPage(true);
      } else {
        setCurrentPage(currentPage + 1);
        commentData.start({
          pageNumber: currentPage + 1,
          searchQuery: "",
          id: route.params.projectId,
        });
        commentData.reset();
      }
    }
  }, [isEndReachRefresh, isEndPage]);
  useEffect(() => {
    if (isEndReachRefresh && commentData.isSuccess) {
      setComments([...comments, ...commentData?.data?.row]);
      setIsEndReachRefresh(false);
      commentData.reset();
    }
  }, [isEndReachRefresh, commentData.isSuccess]);
  useEffect(() => {
    if (commentData.isSuccess && !isEndReachRefresh && !isPullRefresh) {
      setComments(commentData.data.row);
      commentData.reset();
    }
  }, [commentData.isSuccess, isEndReachRefresh, isPullRefresh]);
  console.log(comments.length);
  return React.cloneElement(children, {
    ...route.params,
    comments,
    handlePostComment,
    commentDataIsLoading: commentData.isLoading,
    setIsEndReachRefresh,
    isEndReachRefresh,
    isPullRefresh,
    setIsPullRefresh,
    userProfilePicture: auth.userData.data.picture,
  });
};

export default CommentContainer;
