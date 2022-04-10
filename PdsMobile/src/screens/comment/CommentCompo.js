import React from "react";
import CommentContainer from "./CommentContainer";
import CommentView from "./CommentView";
import CommentRepo from "./CommentRepo";
import {useRoute} from "@react-navigation/native";
import {alert} from "../../utils";
import {useAuth} from "../../services";
import {axios} from "../../config";
import {useFetchCall} from "../../utils";
import {useGlobal} from "../../services";

const CommentCompo = () => {
  const commentRepo = CommentRepo(axios, useFetchCall);
  return (
    <CommentContainer
      useRoute={useRoute}
      useAuth={useAuth}
      {...commentRepo}
      useGlobal={useGlobal}>
      <CommentView alert={alert} />
    </CommentContainer>
  );
};

export default CommentCompo;
