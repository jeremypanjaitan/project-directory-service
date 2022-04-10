import EditProjectView from "./EditProjectView";
import EditProjectContainer from "./EditProjectContainer";
import EditProjectRepo from "./EditProjectRepo";
import {alert, confirmAlert, useCloudStorage} from "../../utils";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {useAuth} from "../../services";

import React from "react";

const EditProjectCompo = props => {
  const editProjectRepo = EditProjectRepo(axios, useFetchCall);

  return (
    <EditProjectContainer
      useCloudStorage={useCloudStorage}
      {...props}
      {...editProjectRepo}>
      <EditProjectView alert={alert} confirmAlert={confirmAlert} />
    </EditProjectContainer>
  );
};

export default EditProjectCompo;
