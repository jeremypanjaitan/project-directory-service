import {
  API_FCM_TOKEN,
  APP_AUTH_SCREEN,
  APP_HOME_SCREEN,
  APP_PROJECT,
  axios,
  USERS_COLLECTION,
} from "../config";
import React, {useEffect} from "react";
import {useAuth} from "../services";
import {API_PROFILE, API_LOGOUT} from "../config";
import {confirmAlert, useFetchCall} from "../utils";
import NoAuthScreenNavigator from "./NoAuthScreenNavigator";
import AuthScreenNavigator from "./AuthScreenNavigator";
import {SplashScreen} from "../components";
import {createNativeStackNavigator} from "@react-navigation/native-stack";
import {Image, TouchableOpacity, View} from "react-native";
import {Button} from "react-native-paper";
import {deleteAuthToken, deleteRefreshToken} from "../utils";
import {en, FCM_NOTIFICATION_CHANNEL} from "../shared";
import messaging from "@react-native-firebase/messaging";
import PushNotification from "react-native-push-notification";
import firestore from "@react-native-firebase/firestore";
import {useNetInfo} from "@react-native-community/netinfo";
import {NoConnectionScreen} from "../screens";
import {navigationRef} from ".";

const saveTokenToDatabase = async (token, uid) => {
  await firestore()
    .collection(USERS_COLLECTION)
    .doc(uid)
    .set(
      {
        fcmTokens: firestore.FieldValue.arrayUnion(token),
      },
      {merge: true},
    );
};

const Stack = createNativeStackNavigator();
const useUserData = () => {
  const userData = async () => {
    try {
      const res = await axios.get(API_PROFILE);
      const data = res.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(userData);
};
PushNotification.createChannel({
  channelId: FCM_NOTIFICATION_CHANNEL, // (required)
  channelName: FCM_NOTIFICATION_CHANNEL, // (required)
});
const useLogout = () => {
  const logout = async () => {
    try {
      const res = await axios.post(API_LOGOUT);
      const data = res.data.data;
      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(logout);
};

const useDeleteFcmToken = () => {
  const deleteFcmToken = async auth => {
    try {
      const res = await axios.delete(API_FCM_TOKEN);
      const data = res.data.data;

      auth.setUserData();
      await deleteAuthToken();
      await deleteRefreshToken();

      return data;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(deleteFcmToken);
};
PushNotification.configure({
  permissions: {
    alert: true,
    badge: true,
    sound: true,
  },
  popInitialNotification: true,
  requestPermissions: true,
  onNotification: notification => {
    if (notification.userInteraction) {
      navigationRef.navigate(APP_HOME_SCREEN, {
        screen: APP_PROJECT,
        params: {projectId: notification.data.projectId},
      });
    }
    if (notification.foreground) {
      PushNotification.localNotification({
        channelId: FCM_NOTIFICATION_CHANNEL,
        title: notification.data.title,
        message: notification.data.body,
        smallIcon: "ic_notification",
        largeIconUrl: notification.data.largeIconUrl,
        bigPictureUrl: notification.data.projectPicture,
        data: {projectId: notification.data.projectId},
      });
    } else {
      navigationRef.navigate(APP_HOME_SCREEN, {
        screen: APP_PROJECT,
        params: {projectId: notification.data.projectId},
      });
    }
  },
});

const ScreenNavigator = () => {
  const auth = useAuth();
  const userData = useUserData();
  const logout = useLogout();
  const deleteFcmToken = useDeleteFcmToken();
  const netInfo = useNetInfo();
  useEffect(() => {
    userData.start();
  }, []);
  useEffect(() => {
    if (userData.isSuccess) {
      auth.setUserData(userData.data);
      userData.reset();
    }
  }, [userData.isSuccess]);
  useEffect(() => {
    if (auth.userData) {
      const effectFunc = async () => {
        messaging()
          .getToken()
          .then(token => {
            return saveTokenToDatabase(token, auth.userData.data.firebaseUid);
          });
        messaging().setBackgroundMessageHandler(async notification => {
          PushNotification.localNotification({
            channelId: FCM_NOTIFICATION_CHANNEL,
            title: notification.data.title,
            message: notification.data.body,
            smallIcon: "ic_notification",
            largeIconUrl: notification.data.largeIconUrl,
            bigPictureUrl: notification.data.projectPicture,
            data: {projectId: notification.data.projectId},
          });
        });
        return messaging().onTokenRefresh(token => {
          saveTokenToDatabase(token, auth.userData.data.firebaseUid);
        });
      };
      effectFunc();
    }
  }, [auth.userData]);

  if (userData.isLoading || logout.isLoading || deleteFcmToken.isLoading) {
    return <SplashScreen />;
  }

  return netInfo.isConnected ? (
    auth.userData ? (
      <Stack.Navigator>
        <Stack.Screen
          component={AuthScreenNavigator}
          name={APP_AUTH_SCREEN}
          options={{
            title: "",
            headerLeft: () => (
              <View style={{width: 100}}>
                <Image
                  source={require("../assets/ITFUN.png")}
                  style={{
                    flex: 1,
                    width: null,
                    height: null,
                    resizeMode: "contain",
                  }}
                />
              </View>
            ),
            headerRight: () => (
              <TouchableOpacity
                onPress={() => {
                  confirmAlert(en.warning, en.confirmLogout, () =>
                    deleteFcmToken.start(auth),
                  );
                }}>
                <Button icon="logout" compact style={{borderRadius: 20}} />
              </TouchableOpacity>
            ),
          }}
        />
      </Stack.Navigator>
    ) : (
      <NoAuthScreenNavigator />
    )
  ) : (
    <NoConnectionScreen />
  );
};
export default ScreenNavigator;
