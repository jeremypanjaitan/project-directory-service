import React from "react";
import {createNativeStackNavigator} from "@react-navigation/native-stack";
import {APP_LOGIN, APP_SIGNUP} from "../config";
import {LoginScreen, SignupScreen} from "../screens";

const Stack = createNativeStackNavigator();

const NoAuthScreenNavigator = () => {
  return (
    <Stack.Navigator>
      <Stack.Screen
        name={APP_LOGIN}
        component={LoginScreen}
        options={{headerShown: false}}
      />
      <Stack.Screen name={APP_SIGNUP} component={SignupScreen} />
    </Stack.Navigator>
  );
};
export default NoAuthScreenNavigator;
