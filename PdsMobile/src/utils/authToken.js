import AsyncStorage from "@react-native-async-storage/async-storage";
import {REFRESH_TOKEN, TOKEN} from "../shared";

export const saveAuthToken = async token => {
  try {
    await AsyncStorage.setItem(TOKEN, token);
  } catch (e) {
    console.log(e);
  }
};

export const getAuthToken = async () => {
  try {
    const token = await AsyncStorage.getItem(TOKEN);
    return token;
  } catch (e) {
    console.log(e);
  }
};

export const deleteAuthToken = async () => {
  try {
    const token = await AsyncStorage.removeItem(TOKEN, token);
    return token;
  } catch (e) {
    console.log(e);
  }
};

export const saveRefreshToken = async token => {
  try {
    await AsyncStorage.setItem(REFRESH_TOKEN, token);
  } catch (e) {
    console.log(e);
  }
};

export const getRefreshToken = async () => {
  try {
    const token = await AsyncStorage.getItem(REFRESH_TOKEN);
    return token;
  } catch (e) {
    console.log(e);
  }
};

export const deleteRefreshToken = async () => {
  try {
    const token = await AsyncStorage.removeItem(REFRESH_TOKEN, token);
    return token;
  } catch (e) {
    console.log(e);
  }
};
