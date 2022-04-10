import React, {useEffect} from "react";
import {View} from "react-native";
import {navigationRef} from "../../navigator";
import {APP_HOME, APP_HOME_SCREEN} from "../../config";

const MiddleScreenContainer = () => {
  useEffect(() => {
    navigationRef.navigate(APP_HOME_SCREEN, {
      screen: APP_HOME});
  }, []);
  return (
    <View />
  );
};

export default MiddleScreenContainer;
