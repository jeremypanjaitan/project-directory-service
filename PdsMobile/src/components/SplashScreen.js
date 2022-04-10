import React from "react";
import {Image, View} from "react-native";
import {ActivityIndicator} from "react-native-paper";

const SplashScreen = () => {
  return (
    <View style={{flex: 1, justifyContent: "center", alignItems: "center"}}>
      <Image source={require("../assets/ITFUN.png")} />
      <ActivityIndicator animating={true} style={{marginTop: 50}} />
    </View>
  );
};

export default SplashScreen;
