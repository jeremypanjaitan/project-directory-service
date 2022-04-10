import {Alert} from "react-native";

export const alert = (title, message) => {
  Alert.alert(title, message);
};

export const confirmAlert = (title, message, okHandler) => {
  Alert.alert(title, message, [
    {
      text: "Cancel",
      style: "cancel",
    },
    {
      text: "OK",
      onPress: okHandler,
    },
  ]);
};
