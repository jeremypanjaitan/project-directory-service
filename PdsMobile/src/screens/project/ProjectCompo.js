import ProjectContainer from "./ProjectContainer";
import ProjectView from "./ProjectView";
import ProjectRepo from "./ProjectRepo";
import React from "react";
import {useRoute, useNavigation} from "@react-navigation/native";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {useGlobal} from "../../services";

const ProjectCompo = () => {
  const projectRepo = ProjectRepo(axios, useFetchCall);
  return (
    <ProjectContainer
      useRoute={useRoute}
      {...projectRepo}
      useGlobal={useGlobal}>
      <ProjectView useNavigation={useNavigation} />
    </ProjectContainer>
  );
};

export default ProjectCompo;
