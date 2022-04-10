import CreateProjectView from "./CreateProjectView";
import { useEffect, useState } from "react";
import { useLoading } from "../../context";
import { useCategoryData, useCreateProject } from "./CreateProjectRepo";
import Swal from "sweetalert2";
import { RouteNavigation } from "../../utils/navigation";

const CreateProjectContainer = ({
  children,
  _useCategoryData = useCategoryData,
  _useLoading = useLoading,
  _useCreateProject = useCreateProject,
  _RouteNavigation = RouteNavigation,
  CreateProjectViewTest = CreateProjectView,
}) => {
  const categoryData = _useCategoryData();
  const [category, setCategory] = useState();
  const { setIsLoading } = _useLoading();
  const createProject = _useCreateProject();
  const navigateTo = _RouteNavigation();

  const handleSubmit = async (values) => {
    if (values.picture === "") {
      Swal.fire("Please upload your project picture", "", "info");
      return;
    }
    setIsLoading(true);
    try {
      createProject.start(values);
    } catch (err) {
      Swal.fire({
        icon: "error",
        text: "Failed to create project!",
      });
    } finally {
      setIsLoading(false);
    }
  };
  if (createProject.isSuccess) {
    Swal.fire("Project Created!", "", "success");
    navigateTo(-1);
    setIsLoading(false);
    createProject.reset();
  }
  if (createProject.isError) {
    setIsLoading(false);
  }

  useEffect(() => {
    categoryData.start();
    //eslint-disable-next-line
  }, []);

  useEffect(() => {
    if (categoryData.isSuccess) {
      setCategory(categoryData.data);
      categoryData.reset();
      setIsLoading(false);
    }
    //eslint-disable-next-line
  }, [categoryData.isSuccess, setIsLoading]);
  return (
    <>
      <CreateProjectViewTest
        categories={category || []}
        handleSubmit={handleSubmit}
      >
        {children}
      </CreateProjectViewTest>
    </>
  );
};

export default CreateProjectContainer;
