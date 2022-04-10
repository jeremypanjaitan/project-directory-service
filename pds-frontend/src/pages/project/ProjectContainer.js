import ProjectView from "./ProjectView";
import { useEffect, useState } from "react";
import { useLoading } from "../../context";
import {
  useProfileData,
  useProjectData,
  useCategoryData,
  useLikesData,
  useCreateLikes,
  useCreateDislike,
  useCreateViews,
  useDeleteProject,
} from "./ProjectRepo";
import { useNavigate, useParams } from "react-router-dom";
import parse from "html-react-parser";
import Swal from "sweetalert2";

const ProjectContainer = () => {
  const params = useParams();
  const projectData = useProjectData();
  const profileData = useProfileData();
  const categoryData = useCategoryData();
  const likesData = useLikesData();
  const createLikes = useCreateLikes();
  const createDislike = useCreateDislike();
  const createViews = useCreateViews();
  const deleteProject = useDeleteProject();
  const [project, setProject] = useState();
  const [profile, setProfile] = useState();
  const [category, setCategory] = useState();
  const [likes, setLikes] = useState();
  const { setIsLoading } = useLoading();
  const navigate = useNavigate();
  const handleLike = async () => {
    setIsLoading(true);
    createLikes.start(params.id);
  };

  if (createLikes.isSuccess) {
    likesData.start(params.id);
    setIsLoading(false);
    createLikes.reset();
  }
  if (createLikes.isError) {
    setIsLoading(false);
  }
  const handleDelete = async () => {
    Swal.fire({
      title: "Are you sure delete this project?",
      text: "You won't be able to revert this!",
      icon: "warning",
      showCancelButton: true,
      confirmButtonColor: "#228228",
      cancelButtonColor: "#d33",
      confirmButtonText: "Yes, delete it!",
    }).then((result) => {
      if (result.isConfirmed) {
        setIsLoading(true);
        deleteProject.start(params.id);
        projectData.start(params.id);
      }
    });
  };

  if (deleteProject.isSuccess) {
    Swal.fire("success delete project!", "", "success");
    setIsLoading(false);
    navigate(-1);
    deleteProject.reset();
  }
  if (createLikes.isError) {
    Swal.fire("error while delete project!", "", "error");
    setIsLoading(false);
  }

  const handleDislike = async () => {
    setIsLoading(true);
    createDislike.start(params.id);
  };

  if (createDislike.isSuccess) {
    likesData.start(params.id);
    setIsLoading(false);
    createDislike.reset();
  }
  if (createDislike.isError) {
    createDislike(false);
  }

  useEffect(() => {
    createViews.start(params.id);
    //eslint-disable-next-line
  }, []);

  useEffect(() => {
    if (createViews.isSuccess) {
      projectData.start(params.id);
      likesData.start(params.id);
    }
  }, [createViews.isSuccess]);

  useEffect(() => {
    if (likesData.isSuccess) {
      setLikes(likesData.data);
      likesData.reset();
      setIsLoading(false);
    }
  }, [likesData.isSuccess]);
  useEffect(() => {
    if (projectData.isSuccess) {
      profileData.start(projectData.data.userId);
      categoryData.start(projectData.data.categoryId);
      projectData.data.story = parse(projectData.data.story);
      setProject(projectData.data);
      projectData.reset();
      setIsLoading(false);
    }

    //eslint-disable-next-line
  }, [projectData.isSuccess, setIsLoading]);

  if (profileData.isSuccess && categoryData.isSuccess) {
    setProfile(profileData.data);
    setCategory(categoryData.data);
    profileData.reset();
    categoryData.reset();
    setIsLoading(false);
  }

  return (
    <>
      <ProjectView
        project={project || ""}
        likes={likes || {}}
        handleLike={handleLike}
        handleDislike={handleDislike}
        category={category || ""}
        profile={profile || ""}
        handleDelete={handleDelete}
      />
    </>
  );
};

export default ProjectContainer;
