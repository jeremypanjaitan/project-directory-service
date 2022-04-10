import React, {useEffect, useState} from "react";
import {useIsFocused, useNavigation} from "@react-navigation/native";
import {Alert} from "react-native";

const HomeContainer = ({children, useAllProjectData, useAllCategoryData}) => {
  const allProjectData = useAllProjectData();
  const allCategoriesData = useAllCategoryData();
  const [searchByTitle, setSearchByTitle] = useState("");
  const [sortByLikes, setSortByLikes] = useState("");
  const [categoryId, setCategoryId] = useState();
  const isFocused = useIsFocused();
  const [isPullRefresh, setIsPullRefresh] = useState(false);
  const [isEndReachRefresh, setIsEndReachRefresh] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const [projects, setProjects] = useState([]);
  const [isEndPage, setIsEndPage] = useState(false);
  const handleSearch = value => {
    setSearchByTitle(value);
  };

  const handleCategoryId = value => {
    setCategoryId(value);
  };

  const handleLike = value => {
    setSortByLikes(value);
  };
  const handleStartSearch = ({searchByTitle, sortByLikes, categoryId}) => {
    allProjectData.start({
      page: 1,
      searchByTitle,
      sortByLikes,
      categoryId,
    });
  };
  useEffect(() => {
    if (isFocused) {
      console.log("start a;; project data")
      allProjectData.start({
        page: currentPage,
        searchByTitle,
        sortByLikes,
        categoryId,
      });
      allCategoriesData.start();
    } else {
      setCurrentPage(1);
      allProjectData.reset();
      allCategoriesData.reset();
      setProjects([]);
    }
  }, [isFocused]);
  useEffect(() => {
    if (isPullRefresh) {
      setIsEndPage(false);
      setCurrentPage(1);
      allProjectData.start({
        page: 1,
        searchByTitle,
        sortByLikes,
        categoryId,
      });
    }
  }, [isPullRefresh]);
  useEffect(() => {
    if (isPullRefresh && allProjectData.isSuccess) {
      console.log("Set the project data 66")
      setProjects(allProjectData?.data?.row);
      setIsPullRefresh(false);
      allProjectData.reset();
    }
  }, [isPullRefresh, allProjectData.isSuccess]);
  useEffect(() => {
    if (isEndReachRefresh && !isEndPage) {
      setCurrentPage(currentPage + 1);
      allProjectData.start({
        page: currentPage + 1,
        searchByTitle,
        sortByLikes,
        categoryId,
      });
    }
  }, [isEndReachRefresh, isEndPage]);
  useEffect(() => {
    if (isEndReachRefresh && allProjectData.isSuccess) {
      if (allProjectData?.data?.row?.length === 0) {
        setIsEndPage(true);
      } else {
        console.log("Set the project data, 87")
        setProjects([...projects, ...allProjectData?.data?.row]);
        setIsEndReachRefresh(false);
        allProjectData.reset();
      }
    }
  }, [isEndReachRefresh, allProjectData.isSuccess]);
  useEffect(() => {
    if (
      !isPullRefresh &&
      !isEndReachRefresh &&
      allProjectData.isSuccess &&
      allProjectData?.data?.row
    ) {
      console.log("Set the project data, 100", allProjectData.data.row)
      setProjects([...allProjectData?.data?.row]);
      allProjectData.reset();
    }
  }, [allProjectData.isSuccess, isPullRefresh, isEndReachRefresh]);
  return React.cloneElement(children, {
    allProjectData: projects,
    allProjectDataIsLoading: allProjectData.isLoading,
    isPullRefresh,
    setCurrentPage,
    setIsPullRefresh,
    setIsEndReachRefresh,
    isEndReachRefresh,
    handleSearch,
    handleCategoryId,
    handleLike,
    allCategoriesData: allCategoriesData.data?.map(d => ({
      value: d.id,
      label: d.name,
    })),
    handleStartSearch,
  });
};

export default HomeContainer;
