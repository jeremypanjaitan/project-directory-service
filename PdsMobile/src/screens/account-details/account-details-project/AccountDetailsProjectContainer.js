import {useIsFocused} from "@react-navigation/core";
import React, {useEffect, useState} from "react";

const AccountDetailsProjectContainer = ({children, useAllProfileProject}) => {
  const profileProject = useAllProfileProject();

  const [projects, setProjects] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [isEndReachRefresh, setIsEndReachRefresh] = useState(false);
  const [isPullRefresh, setIsPullRefresh] = useState(false);
  const isFocused = useIsFocused();
  const [isEndPage, setIsEndPage] = useState(false);

  useEffect(() => {
    if (isFocused) {
      profileProject.start({
        page: currentPage,
      });
    } else {
      setCurrentPage(1);

      profileProject.reset();
      setProjects([]);
    }
  }, [isFocused]);

  useEffect(() => {
    if (isPullRefresh) {
      setIsEndPage(false);
      setCurrentPage(1);

      profileProject.start({
        page: 1,
      });
      profileProject.reset();
    }
  }, [isPullRefresh]);
  useEffect(() => {
    if (isPullRefresh && profileProject.isSucce) {
      setProjects(profileProject?.data?.row);
      setIsPullRefresh(false);
      profileProject.reset();
    }
  }, [isPullRefresh, profileProject.isSuccess]);

  useEffect(() => {
    if (isEndReachRefresh && !isEndPage) {
      setCurrentPage(currentPage + 1);

      profileProject.start({
        page: currentPage + 1,
      });
      profileProject.reset();
    }
  }, [isEndReachRefresh, isEndPage]);
  useEffect(() => {
    if (isEndReachRefresh && profileProject.isSuccess) {
      if (profileProject?.data?.row?.length === 0) {
        setIsEndPage(true);
      } else {
        setProjects([...projects, ...profileProject?.data?.row]);
        setIsEndReachRefresh(false);
        profileProject.reset();
      }
    }
  }, [isEndReachRefresh, profileProject.isSuccess]);
  useEffect(() => {
    if (
      profileProject?.data?.row &&
      profileProject.isSuccess &&
      !isEndReachRefresh &&
      !isPullRefresh
    ) {
      setProjects([...profileProject?.data?.row]);
      profileProject.reset();
    }
  }, [profileProject.isSuccess, isEndReachRefresh, isPullRefresh]);

  return React.cloneElement(children, {
    projects,
    profileProjectIsLoading: profileProject.isLoading,
    setIsEndReachRefresh,
    isEndReachRefresh,
    isPullRefresh,
    setIsPullRefresh,
  });
};

export default AccountDetailsProjectContainer;
