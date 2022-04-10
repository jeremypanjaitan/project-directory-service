import React from "react";
import {ActivityIndicator, Text} from "react-native-paper";
import {View} from "react-native";

const DefaultSpinner = ({text = "Loading..."}) => {
  return (
    <View style={{flex: 1, justifyContent: "center"}}>
      <View
        style={{
          alignItems: "center",
          flexDirection: "row",
          justifyContent: "center",
        }}>
        <ActivityIndicator animating={true} style={{marginRight: 10}} />
        <Text>{text}</Text>
      </View>
    </View>
  );
};

export default DefaultSpinner;
