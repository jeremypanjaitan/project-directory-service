import storage from "@react-native-firebase/storage";
import {FIREBASE_UPLOAD_PIC_DIR} from "../shared";
import useFetchCall from "./useFetchCall";

const useCloudStorage = () => {
  const cloudStorage = async file => {
    try {
      const reference = storage().ref(
        FIREBASE_UPLOAD_PIC_DIR + "/" + file.name,
      );
      await reference.putFile(file.uri);
      const url = await reference.getDownloadURL();
      return url;
    } catch (err) {
      throw err;
    }
  };
  return useFetchCall(cloudStorage);
};

export default useCloudStorage;
