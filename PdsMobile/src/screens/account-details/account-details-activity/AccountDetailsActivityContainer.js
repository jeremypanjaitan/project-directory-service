import {useIsFocused} from "@react-navigation/core";
import React, {useEffect, useState} from "react";

const AccountDetailsActivityContainer = ({children, useAllProfileActivity}) => {
  const profileActivity = useAllProfileActivity();
  const [activities, setActivities] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [isEndReachRefresh, setIsEndReachRefresh] = useState(false);
  const [isPullRefresh, setIsPullRefresh] = useState(false);
  const isFocused = useIsFocused();
  const [isEndPage, setIsEndPage] = useState(false);

  useEffect(() => {
    if (isFocused) {
      profileActivity.start({
        page: currentPage,
      });
    } else {
      setCurrentPage(1);
      profileActivity.reset();
      setActivities([]);
    }
  }, [isFocused]);

  useEffect(() => {
    if (isPullRefresh) {
      setIsEndPage(false);
      setCurrentPage(1);
      profileActivity.start({
        page: 1,
      });
      profileActivity.reset();
    }
  }, [isPullRefresh]);
  useEffect(() => {
    if (isPullRefresh && profileActivity.isSuccess) {
      setActivities(profileActivity?.data?.row);
      setIsPullRefresh(false);
      profileActivity.reset();
    }
  }, [isPullRefresh, profileActivity.isSuccesss]);

  useEffect(() => {
    if (isEndReachRefresh && !isEndPage) {
      if (profileActivity?.data?.row?.length === 0) {
        setIsEndPage(true);
      } else {
        setCurrentPage(currentPage + 1);
        profileActivity.start({
          page: currentPage + 1,
        });
        profileActivity.reset();
      }
    }
  }, [isEndReachRefresh]);
  useEffect(() => {
    if (isEndReachRefresh && profileActivity.isSuccess) {
      setActivities([...activities, ...profileActivity?.data?.row]);
      setIsEndReachRefresh(false);
      profileActivity.reset();
    }
  }, [isEndReachRefresh, profileActivity.isSuccess]);
  useEffect(() => {
    if (
      profileActivity?.data?.row &&
      profileActivity.isSuccess &&
      !isEndReachRefresh &&
      !isPullRefresh
    ) {
      setActivities([...profileActivity?.data?.row]);
      profileActivity.reset();
    }
  }, [profileActivity.isSuccess, isEndReachRefresh, isPullRefresh]);

  return React.cloneElement(children, {
    activities,
    profileActivityIsLoading: profileActivity.isLoading,

    setIsEndReachRefresh,
    isEndReachRefresh,
    isPullRefresh,
    setIsPullRefresh,
  });
};

export default AccountDetailsActivityContainer;
