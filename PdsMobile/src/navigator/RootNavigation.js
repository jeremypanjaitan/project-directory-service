import {
  createNavigationContainerRef,
  StackActions,
} from "@react-navigation/native";

export const navigationRef = createNavigationContainerRef();

export const navigateStack = (name, params, isReplace = false) => {
  if (navigationRef.isReady()) {
    if (isReplace) {
      navigationRef.current.dispatch(StackActions.replace(name, params));
    } else {
      navigationRef.navigate(name, params);
    }
  }
};
