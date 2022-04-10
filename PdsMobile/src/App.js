import React from "react";
import {DefaultTheme, Provider as PaperProvider} from "react-native-paper";
import {ScreenNavigator} from "./navigator";
import {AuthProvider, GlobalProvider} from "./services";
import {NavigationContainer} from "@react-navigation/native";
import {navigationRef} from "./navigator";
import {LogBox} from "react-native";
import {AxiosRefreshToken} from "./utils";

LogBox.ignoreLogs([
  "EventEmitter.removeListener",
  "Warning: Can't perform a React state update on an unmounted component. This is a no-op, but it indicates a memory leak in your application. To fix, cancel all subscriptions and asynchronous tasks in a useEffect cleanup function.",
  "No task registered for key ReactNativeFirebaseMessagingHeadlessTask",
  'No background message handler has been set. Set a handler via the "setBackgroundMessageHandler" method.',
]);

const theme = {
  ...DefaultTheme,
  roundness: 2,
  colors: {
    ...DefaultTheme.colors,
    primary: "#547AF0",
    editor: "#dce3fa",
    background: "white",
    headingText: "black",
  },
};
const App = () => {
  return (
    <PaperProvider theme={theme}>
      <AuthProvider>
        <GlobalProvider>
          <NavigationContainer ref={navigationRef}>
            <AxiosRefreshToken>
              <ScreenNavigator />
            </AxiosRefreshToken>
          </NavigationContainer>
        </GlobalProvider>
      </AuthProvider>
    </PaperProvider>
  );
};

export default App;
