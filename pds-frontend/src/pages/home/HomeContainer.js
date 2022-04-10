import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import Swal from "sweetalert2";
import { useLoading } from "../../context";
import { CREATE } from "../../router";
import { useAllProjectData, useCategoryData } from "./HomeRepo";
import HomeView from "./HomeView";

const HomeContainer = () => {
  const [searchByTitle, setSearchByTitle] = useState("");
  const [sortByLikes, setSortByLikes] = useState();
  const [categoryId, setCategoryId] = useState();
  const [currentPage, setCurrentPage] = useState(1);
  const categoriesData = useCategoryData();
  const [categories, setCategories] = useState();
  const projectAllData = useAllProjectData();
  const [projects, setProjects] = useState();
  const { setIsLoading } = useLoading();

  const navigate = useNavigate();
  const handleSearch = (value) => {
    setSearchByTitle(value);
  };

  const handleCategoryId = (value) => {
    setCategoryId(value);
  };

  const handleLike = (value) => {
    setSortByLikes(value);
  };
  const handleNavigateAddProject = () => {
    navigate(CREATE);
  };

  useEffect(() => {
    setIsLoading(true);
    setProjects([]);
    setCategories([]);

    projectAllData.start({
      page: currentPage,
      searchByTitle,
      sortByLikes,
      categoryId,
    });

    categoriesData.start();

    //eslint-disable-next-line
  }, [currentPage, searchByTitle, sortByLikes, categoryId]);

  useEffect(() => {
    if (projectAllData.isSuccess && categoriesData.isSuccess) {
      setProjects(projectAllData.data.data);
      console.log("project", projectAllData.data);
      setCategories(categoriesData.data);
      projectAllData.reset();
      setIsLoading(false);
    }

    //eslint-disable-next-line
  }, [projectAllData.isSuccess, categoriesData.isSuccess, setIsLoading]);

  if (projectAllData.isError) {
    Swal.fire("Error get all projects data", "", "warning");
    setIsLoading(false);
  }
  return (
    <>
      <HomeView
        data={projects?.row || []}
        totalPage={projects?.totalPage}
        handleNavigateAddPage={handleNavigateAddProject}
        currentPage={currentPage}
        setCurrentPage={setCurrentPage}
        handleSearch={handleSearch}
        handleCategoryId={handleCategoryId}
        handleLike={handleLike}
        categories={categories || []}
      />
    </>
  );
};

export default HomeContainer;
