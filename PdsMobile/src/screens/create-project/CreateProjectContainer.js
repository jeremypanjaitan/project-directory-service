import React, { useEffect } from "react";
import { useIsFocused, useNavigation } from "@react-navigation/native";
import { APP_HOME } from "../../config";

const CreateProjectContainer = ({
  children,
  useCloudStorage,
  useCreateProject,
  useCategoryData,
  useAuth,
}) => {
  const categoryData = useCategoryData();
  const cloudStorage = useCloudStorage();
  const cloudStorageProjectVideo = useCloudStorage();
  const cloudStorageProjectPicture = useCloudStorage();
  const createProject = useCreateProject();
  const isFocused = useIsFocused();
  const navigation = useNavigation();
  useEffect(() => {
    if (isFocused) {
      categoryData.start();
    } else {
      categoryData.reset();
    }
  }, [isFocused]);

  const handleCreateProject = values => {
    if (values.picture === "") {
      alert("Please upload picure");
    } else {
      console.log(values);
      createProject.start(values);
    }
  };
  useEffect(() => {
    if (createProject.isSuccess) {
      navigation.navigate(APP_HOME);
    }
  }, [createProject.isSuccess]);
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
    handleCreateProject: handleCreateProject,
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
    createProjectIsLoading: createProject.isLoading,
    createProjectIsSuccess: createProject.isSuccess,
    createProjectReset: createProject.reset,
    createProjectData: createProject.data,
  });
};

export default CreateProjectContainer;
