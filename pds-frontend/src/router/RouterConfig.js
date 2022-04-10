import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import CustomRouter from "./CustomRouter";
import customHistory from "./customHistory";
import {
  CreateProjectPage,
  LoginPage,
  SignupPage,
  ProjectPage,
  AccountSettingUpdatePage,
  AccountDetailsPage,
  HomePage,
  ProjectCommentPage,
  EditProjectPage,
} from "../pages";

import {
  CREATE,
  LOGIN,
  SIGN_UP,
  SHOW_ID,
  ACCOUNT_SETTINGS,
  MAIN,
  DEFAULT,
  COMMENT,
  EDIT,
} from "../router";

import { AlreadyAuth, RequireAuth } from "../services";
import { SignupLayout, MainLayout } from "../components";
import { ACCOUNT_DETAILS } from "./routesName";

const RouterConfig = () => {
  return (
    <CustomRouter
      history={customHistory}
      basename={process.env.REACT_APP_BASE_URL}
    >
      <Routes>
        <Route
          path={LOGIN}
          element={
            <AlreadyAuth>
              <LoginPage />
            </AlreadyAuth>
          }
        />
        <Route
          path={ACCOUNT_SETTINGS}
          element={
            <RequireAuth>
              <MainLayout>
                <AccountSettingUpdatePage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route
          path={ACCOUNT_DETAILS}
          element={
            <RequireAuth>
              <MainLayout>
                <AccountDetailsPage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route
          path={SIGN_UP}
          element={
            <AlreadyAuth>
              <SignupLayout>
                <SignupPage />
              </SignupLayout>
            </AlreadyAuth>
          }
        />
        <Route
          path={CREATE}
          element={
            <RequireAuth>
              <MainLayout>
                <CreateProjectPage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route
          path={MAIN}
          element={
            <RequireAuth>
              <MainLayout>
                <HomePage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route
          path={SHOW_ID}
          element={
            <RequireAuth>
              <MainLayout>
                <ProjectPage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route
          path={COMMENT}
          element={
            <RequireAuth>
              <MainLayout>
                <ProjectCommentPage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route
          path={EDIT}
          element={
            <RequireAuth>
              <MainLayout>
                <EditProjectPage />
              </MainLayout>
            </RequireAuth>
          }
        />
        <Route path={DEFAULT} element={<Navigate to={MAIN} />} />
        <Route path={"*"} element={<Navigate to={MAIN} />} />
      </Routes>
    </CustomRouter>
  );
};

export default RouterConfig;
