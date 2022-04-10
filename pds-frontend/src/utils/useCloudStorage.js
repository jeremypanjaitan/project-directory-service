import { useState } from "react";
import {
  storage,
  STORAGE_DIR_PICTURE,
  STORAGE_DIR_FILE,
} from "../config/firebase";
import { ref, getDownloadURL, uploadBytesResumable } from "firebase/storage";

const useCloudStorage = (
  _uploadBytesResumable = uploadBytesResumable,
  _ref = ref,
  _useState = useState
) => {
  const [progress, setProgress] = _useState(0);
  const upload = async (file, dir) => {
    if (!file) return;
    const sotrageRef = _ref(storage, `${dir}/${file.name}`);
    const uploadTask = _uploadBytesResumable(sotrageRef, file);

    return new Promise((resolve, reject) => {
      uploadTask.on(
        "state_changed",
        (snapshot) => {
          const prog = Math.round(
            (snapshot.bytesTransferred / snapshot.totalBytes) * 100
          );
          setProgress(prog);
        },
        (error) => reject(error),
        () => {
          resolve(getDownloadURL(uploadTask.snapshot.ref));
        }
      );
    });
  };
  const uploadPicture = (picture) => {
    return upload(picture, STORAGE_DIR_PICTURE);
  };
  const uploadFile = (file) => {
    return upload(file, STORAGE_DIR_FILE);
  };
  const resetProgress = () => {
    setProgress(0);
  };
  return { uploadPicture, progress, resetProgress, uploadFile };
};

export default useCloudStorage;
