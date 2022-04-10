import EditProjectView from "./EditProjectView";
import { useEffect, useState } from "react";
import { useLoading } from "../../context";
import {
  useCategoryData,
  useCategoryDataByID,
  useEditProject,
  useProjectData,
} from "./EditProjectRepo";
import Swal from "sweetalert2";
import { RouteNavigation } from "../../utils/navigation";
import { useParams } from "react-router-dom";

const EditProjectContainer = ({
  children,
  _useCategoryData = useCategoryData,
  _useLoading = useLoading,
  _useEditProject = useEditProject,
  _RouteNavigation = RouteNavigation,
  EditProjectViewTest = EditProjectView,
}) => {
  const categoryData = _useCategoryData();
  const [category, setCategory] = useState();
  const { setIsLoading } = _useLoading();
  const EditProject = useEditProject();
  const navigateTo = _RouteNavigation();
  const projectData = useProjectData();
  const currentCategoryData = useCategoryDataByID();
  const [currentProject, setCurrentProject] = useState({});
  const [currentCategory, setCurrentCategory] = useState("");
  const params = useParams();

  const handleSubmit = async (values) => {
    setIsLoading(true);
    EditProject.start({ id: params.id, values: values });
  };

  if (EditProject.isSuccess) {
    Swal.fire("success edit project!", "", "success");
    navigateTo(-1);
    setIsLoading(false);
    EditProject.reset();
  }
  if (EditProject.isError) {
    Swal.fire("Project Error while edit!", "", "error");
    setIsLoading(false);
  }

  useEffect(() => {
    setIsLoading(true);
    categoryData.start();
    projectData.start(params.id);
  }, []);

  useEffect(() => {
    if (categoryData.isSuccess && projectData.isSuccess) {
      setCategory(categoryData.data);
      setCurrentProject(projectData.data);
      categoryData.reset();
      projectData.reset();
      setIsLoading(false);
    }
    //eslint-disable-next-line
  }, [categoryData.isSuccess, projectData.isSuccess, setIsLoading]);

  return (
    <>
      <EditProjectViewTest
        categories={category || []}
        currentProject={currentProject || {}}
        handleSubmit={handleSubmit}
      >
        {children}
      </EditProjectViewTest>
    </>
  );
};

export default EditProjectContainer;
