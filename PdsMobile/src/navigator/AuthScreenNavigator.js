import React from "react";
import {withTheme} from "react-native-paper";
import {createMaterialBottomTabNavigator} from "@react-navigation/material-bottom-tabs";
import {createNativeStackNavigator} from "@react-navigation/native-stack";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";
import {
  AccountSettingsScreen,
  CreateProjectScreen,
  EditProjectScreen,
  HomeScreen, MiddleScreen,
  ProjectScreen,
} from "../screens";
import {
  ACCOUNT_SETTINGS,
  APP_ACCOUNT_DETAILS,
  APP_ACCOUNT_DETAILS_ACTIVITY,
  APP_ACCOUNT_DETAILS_PROJECT,
  APP_ACCOUNT_DETAILS_SCREEN,
  APP_COMMENT,
  APP_HOME,
  APP_HOME_SCREEN, APP_MIDDLE,
  APP_PROJECT,
  CREATE_PROJECT,
  EDIT_PROJECT,
} from "../config";
import {CommentScreen} from "../screens/comment";

import {AccountDetailsProjectScreen} from "../screens/account-details/account-details-project";
import {AccountDetailsActivityScreen} from "../screens/account-details/account-details-activity";
import AccountDetailsScreen from "../screens/account-details/account-details-screen-navigator/AccountDetailsScreen";

const Tab = createMaterialBottomTabNavigator();
const Stack = createNativeStackNavigator();

const AuthScreenNavigator = ({theme}) => {
  const {colors} = theme;
  return (
    <Tab.Navigator barStyle={{backgroundColor: colors.primary}}>
      <Tab.Screen
        name={APP_HOME_SCREEN}
        component={ProjectScreenNavigator}
        options={{
          tabBarIcon: ({color}) => {
            return (
              <MaterialCommunityIcons name="home" color={color} size={26} />
            );
          },
        }}
      />
      <Tab.Screen
        name={CREATE_PROJECT}
        component={CreateProjectScreen}
        options={{
          tabBarIcon: ({color}) => {
            return (
              <MaterialCommunityIcons name="plus" color={color} size={26} />
            );
          },
        }}
      />
      <Tab.Screen
        name={ACCOUNT_SETTINGS}
        component={AccountSettingsScreen}
        options={{
          tabBarIcon: ({color}) => {
            return (
              <MaterialCommunityIcons
                name="cog-outline"
                color={color}
                size={26}
              />
            );
          },
        }}
      />
      <Tab.Screen
        name={APP_ACCOUNT_DETAILS_SCREEN}
        component={AccountDetailsNavigator}
        options={{
          tabBarIcon: ({color}) => {
            return (
              <MaterialCommunityIcons name="account" color={color} size={26} />
            );
          },
        }}
      />
    </Tab.Navigator>
  );
};

const ProjectScreenNavigator = () => {
  return (
    <Stack.Navigator screenOptions={{headerShown: false}}>
      <Stack.Screen name={APP_HOME} component={HomeScreen} />
      <Stack.Screen name={APP_MIDDLE} component={MiddleScreen} />
      <Stack.Screen name={APP_PROJECT} component={ProjectScreen} />
      <Stack.Screen name={APP_COMMENT} component={CommentScreen} />
      <Stack.Screen name={EDIT_PROJECT} component={EditProjectScreen} />
    </Stack.Navigator>
  );
};

const AccountDetailsNavigator = () => {
  return (
    <Stack.Navigator screenOptions={{headerShown: false}}>
      <Stack.Screen
        name={APP_ACCOUNT_DETAILS}
        component={AccountDetailsScreen}
      />
      <Stack.Screen
        name={APP_ACCOUNT_DETAILS_PROJECT}
        component={AccountDetailsProjectScreen}
      />
      <Stack.Screen
        name={APP_ACCOUNT_DETAILS_ACTIVITY}
        component={AccountDetailsActivityScreen}
      />
      <Stack.Screen name={APP_PROJECT} component={ProjectScreen} />
      <Stack.Screen name={APP_COMMENT} component={CommentScreen} />
      <Stack.Screen name={EDIT_PROJECT} component={EditProjectScreen} />
    </Stack.Navigator>
  );
};
export default withTheme(AuthScreenNavigator);
