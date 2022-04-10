import CreateProjectView from "./CreateProjectView";
import CreateProjectContainer from "./CreateProjectContainer";
import CreateProjectRepo from "./CreateProjectRepo";
import {alert, confirmAlert, useCloudStorage} from "../../utils";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {useAuth} from "../../services";

import React from "react";

const CreateProjectCompo = props => {
  const createProjectRepo = CreateProjectRepo(axios, useFetchCall);

  return (
    <CreateProjectContainer
      useCloudStorage={useCloudStorage}
      {...props}
      {...createProjectRepo}>
      <CreateProjectView alert={alert} confirmAlert={confirmAlert} />
    </CreateProjectContainer>
  );
};

export default CreateProjectCompo;
