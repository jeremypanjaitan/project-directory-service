import * as RootNavigation from "./RootNavigation";
import {
  APP_ACCOUNT_DETAILS,
  APP_ACCOUNT_DETAILS_ACTIVITY,
  APP_ACCOUNT_DETAILS_PROJECT,
  APP_LOGIN,
  APP_SIGNUP,
  APP_PROJECT,
} from "../config";

export const goToRegister = () => {
  RootNavigation.navigateStack(APP_SIGNUP, null, false);
};

export const goToLogin = () => {
  RootNavigation.navigateStack(APP_LOGIN, null, true);
};

export const goToProject = params => {
  RootNavigation.navigateStack(APP_PROJECT, params, true);
};

export const goToAccountDetails = () => {
  RootNavigation.navigateStack(APP_ACCOUNT_DETAILS, null, true);
};

export const goToAccountDetailsActivity = () => {
  RootNavigation.navigateStack(APP_ACCOUNT_DETAILS_ACTIVITY, null, true);
};

export const goToAccountDetailsProject = () => {
  RootNavigation.navigateStack(APP_ACCOUNT_DETAILS_PROJECT, null, true);
};
