import React, {useEffect, useState} from "react";
import {useNavigation} from "@react-navigation/native";
import { Alert } from "react-native";
import { confirmAlert } from "../../utils";
import { en } from "../../shared";
import {navigationRef} from "../../navigator";
import {
  ACCOUNT_SETTINGS,
  APP_ACCOUNT_DETAILS_PROJECT,
  APP_HOME,
  APP_HOME_SCREEN, APP_MIDDLE,
  APP_PROJECT
} from "../../config";

const ProjectContainer = ({
  children,
  useRoute,
  useProjectData,
  useProfileData,
  useCategoryData,
  useProjectLike,
  useLike,
  useDislike,
  useView,
  useTotalView,
  useGlobal,
  useDeleteProject,
}) => {
  const projectData = useProjectData();
  const profileData = useProfileData();
  const categoryData = useCategoryData();
  const projectLike = useProjectLike();
  const deleteProject = useDeleteProject();
  const dislike = useDislike();
  const like = useLike();
  const view = useView();
  const totalView = useTotalView();
  const globalState = useGlobal();
  const [isUserLike, setIsUserLike] = useState(false);
  const [totalLikes, setTotalLikes] = useState(0);
  const route = useRoute();
  const navigation = useNavigation();
  const handleDeleteProject = () => {
    confirmAlert(en.alertDelete,en.confirmDeleteProject, () => deleteProject.start(route.params.projectId));
  };
  useEffect(() => {
    if (deleteProject.isSuccess) {
      navigationRef.navigate(APP_HOME_SCREEN, {screen: APP_MIDDLE,
        params: {backStack: route.params.backStack}});
      }
  }, [deleteProject.isSuccess]);
  const likeStart = () => {
    setIsUserLike(!isUserLike);
    setTotalLikes(totalLikes + 1);
    like.start(route.params.projectId);
  };
  const dislikeStart = () => {
    setIsUserLike(!isUserLike);
    setTotalLikes(totalLikes - 1);
    dislike.start(route.params.projectId);
  };
  useEffect(() => {
    view.start(route.params.projectId);
  }, []);
  useEffect(() => {
    if (view.isSuccess) {
      projectData.start(route.params.projectId);
      projectLike.start(route.params.projectId);
      totalView.start(route.params.projectId);
    }
  }, [view.isSuccess]);
  useEffect(() => {
    if (projectData.isSuccess) {
      profileData.start(projectData.data.userId);
      categoryData.start(projectData.data.categoryId);
      globalState.setTotalComment(projectData?.data?.totalComments);
    }
  }, [projectData.isSuccess]);
  useEffect(() => {
    if (projectLike.isSuccess) {
      setIsUserLike(projectLike.data.isUserLike);
      setTotalLikes(projectLike.data.totalLikes);
    }
  }, [projectLike.isSuccess]);
  useEffect(() => {
    if (like.isError) {
      setIsUserLike(!isUserLike);
      setTotalLikes(totalLikes - 1);
    }
  }, [like.isError]);
  useEffect(() => {
    if (dislike.isError) {
      setIsUserLike(!isUserLike);
      setTotalLikes(totalLikes + 1);
    }
  }, [dislike.isError]);
  return React.cloneElement(children, {
    profileDataIsLoading: profileData.isLoading,
    projectDataIsLoading: projectData.isLoading,
    categoryDataIsLoading: categoryData.isLoading,
    projectLikeIsLoading: projectLike.isLoading,
    projectLikeData: projectLike.data,
    projectData: projectData.data,
    profileData: profileData.data,
    totalLikes,
    isUserLike,
    categoryData: categoryData.data,
    likeStart,
    dislikeStart,
    handleDeleteProject: handleDeleteProject,
    deleteProjectIsLoading: deleteProject.isLoading,
    deleteProjectData: deleteProject.data,
    viewIsLoading: view.isLoading,
    viewData: view.data,
    totalViewIsLoading: totalView.isLoading,
    totalViewData: totalView.data,
    totalComments: globalState.totalComment,
    projectId: route.params.projectId,
    backStack: route.params.backStack,
  });
};

export default ProjectContainer;
