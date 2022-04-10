import React from "react";
import HomeContainer from "./HomeContainer";
import HomeView from "./HomeView";
import HomeRepo from "./HomeRepo";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";

const HomeCompo = () => {
  const homeRepo = HomeRepo(axios, useFetchCall);
  return (
    <HomeContainer {...homeRepo}>
      <HomeView />
    </HomeContainer>
  );
};

export default HomeCompo;
