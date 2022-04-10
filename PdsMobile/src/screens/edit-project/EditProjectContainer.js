import React, {useEffect} from "react";
import {useIsFocused, useNavigation, useRoute} from "@react-navigation/native";
import {APP_HOME, APP_PROJECT} from "../../config";

const EditProjectContainer = ({
  children,
  useCloudStorage,
  useEditProject,
  useCategoryData,
  useProjectData,
  useAuth,
}) => {
  const categoryData = useCategoryData();
  const cloudStorage = useCloudStorage();
  const cloudStorageProjectVideo = useCloudStorage();
  const cloudStorageProjectPicture = useCloudStorage();
  const editProject = useEditProject();
  const projectData = useProjectData();
  const isFocused = useIsFocused();
  const route = useRoute();
  const navigation = useNavigation();
  useEffect(() => {
    if (isFocused) {
      projectData.start(route.params.projectId);
      categoryData.start();
    } else {
      categoryData.reset();
      projectData.reset();
    }
  }, [isFocused]);

  const handleEditProject = values => {
    if (values.picture === "") {
      alert("Please upload picure");
    } else {
      console.log(values, "in container", route.params.projectId);
      editProject.start({id: route.params.projectId, values: values});
    }
  };
  useEffect(() => {
    if (editProject.isSuccess) {
      console.log("sukses edit")
      navigation.replace(APP_PROJECT, {
        projectId: route.params.projectId,
        backStack: APP_HOME,
      });
    }
  }, [editProject.isSuccess]);
  const handleUploadPicture = picture => {
    cloudStorage.start(picture);
  };
  const handlePictureProject = picture => {
    cloudStorageProjectPicture.start(picture);
  };
  const handleVideoProject = picture => {
    cloudStorageProjectVideo.start(picture);
  };
  return React.cloneElement(children, {
    handleUploadPicture: handleUploadPicture,
    handleEditProject: handleEditProject,
    handlePictureProject: handlePictureProject,
    handleVideoProject: handleVideoProject,
    cloudStorageProjectVideoData: cloudStorageProjectVideo.data,
    cloudStorageProjectVideoIsLoading: cloudStorageProjectVideo.isLoading,
    cloudStorageProjectPictureData: cloudStorageProjectPicture.data,
    cloudStorageProjectPictureIsLoading: cloudStorageProjectPicture.isLoading,
    cloudStorageData: cloudStorage.data,
    cloudStorageIsLoading: cloudStorage.isLoading,
    categoryDataIsLoading: categoryData.isLoading,
    categoryData: categoryData.data?.map(d => ({value: d.id, label: d.name})),
    isFocused: isFocused,
    editProjectIsLoading: editProject.isLoading,
    editProjectIsSuccess: editProject.isSuccess,
    editProjectReset: editProject.reset,
    editProjectData: editProject.data,
    projectData: projectData.data,
    projectIsLoading: projectData.isLoading,
  });
};

export default EditProjectContainer;
