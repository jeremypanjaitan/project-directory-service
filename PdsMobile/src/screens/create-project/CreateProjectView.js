import React, {useEffect, useState} from "react";
import {Image, Modal, ScrollView, TouchableOpacity, View} from "react-native";
import {
  withTheme,
  Text,
  ActivityIndicator,
  Button,
  TextInput,
  HelperText,
  Divider,
} from "react-native-paper";
import {launchImageLibrary} from "react-native-image-picker";
import {useFormik} from "formik";
import {DefaultSpinner} from "../../components";
import {PaperDropdown} from "../../components";
import * as Yup from "yup";
import {en} from "../../shared";
import {actions, RichEditor, RichToolbar} from "react-native-pell-rich-editor";
import ImageView from "react-native-image-viewing";

const CreateProjectView = ({
  alert,
  theme,
  handleUploadPicture,
  handleCreateProject,
  handlePictureProject,
  handleVideoProject,
  cloudStorageProjectVideoData,
  cloudStorageProjectVideoIsLoading,
  cloudStorageProjectPictureData,
  cloudStorageProjectPictureIsLoading,
  cloudStorageData,
  cloudStorageIsLoading,
  categoryDataIsLoading,
  categoryData,
  isFocused,
  createProjectIsLoading,
  createProjectIsSuccess,
  createProjectReset,
  createProjectData,
}) => {
  const [isUploading, setIsUploading] = useState(false);
  const [isImageViewVisible, setIsImageViewVisible] = useState(false);

  const {colors} = theme;
  const formik = useFormik({
    initialValues: {
      title: "",
      picture: "",
      description: "",
      story: "<p>Write story about you project here...</p>",
      categoryId: "",
    },
    validationSchema: Yup.object({
      title: Yup.string()
        .required("This field is required")
        .min(6, "Minimum 6 character length")
        .max(150, "Maximum 150 character"),
      picture: Yup.string(),
      description: Yup.string()
        .required("This field is required")
        .min(10, "Minimum 10 character length")
        .max(300, "Maximum 300 character"),
      story: Yup.string()
        .required("This field is required")
        .min(10, "Minimum 10 character length"),
      categoryId: Yup.number().required("This field is required"),
    }),
    onSubmit: values => {
      handleCreateProject(values);
    },
  });
  useEffect(() => {
    if (!isFocused) {
      formik.resetForm();
    }
  }, [isFocused]);
  useEffect(() => {
    if (cloudStorageData) {
      formik.setFieldValue("picture", cloudStorageData);
    }
  }, [cloudStorageData]);
  useEffect(() => {
    if (cloudStorageProjectPictureData) {
      richText.current?.insertImage(cloudStorageProjectPictureData);
      setIsUploading(false);
    }
  }, [cloudStorageProjectPictureData]);
  useEffect(() => {
    if (cloudStorageProjectVideoData) {
      console.log("masuk upload video")
      richText.current?.insertVideo(cloudStorageProjectVideoData);
      setIsUploading(false);
    }
  }, [cloudStorageProjectVideoData]);
  useEffect(() => {
    if (createProjectIsSuccess) {
      alert(en.success, en.successCreateProject);
      createProjectReset();
    }
  }, [createProjectIsSuccess]);

  const handleEditor = text => {
    formik.values.story = text;
  };

  const richText = React.useRef();

  function onPressAddImage() {
    LaunchGaleryPicture();
  }
  console.log(isUploading, "Hasil loading");

  const LaunchGaleryPicture = async () => {
    await launchImageLibrary(null, async res => {
      if (!res.didCancel) {
        if (res.assets[0].fileSize > (10 * 1024 * 1024)){
          alert(en.error, "Maximum file size 10 MB");
        } else {
          handlePictureProject({
            name: res.assets[0].fileName,
            uri: res.assets[0].uri,
          });
          setIsUploading(true);
        }
      }
    });
  };

  const LaunchGaleryVideo = async () => {
    await launchImageLibrary({mediaType: 'video'}, async res => {
      if (!res.didCancel) {
        if (res.assets[0].fileSize > 20 * 1024 * 1024) {
          alert(en.error, "Maximum file size 20 MB");
        } else {
          handleVideoProject({
            name: res.assets[0].fileName,
            uri: res.assets[0].uri,
          });
          setIsUploading(true);
        }
      }
    });
  };
  function insertVideo() {
    LaunchGaleryVideo();
  }
  if (categoryDataIsLoading || !categoryData) {
    return <DefaultSpinner />;
  }

  const handleSubmit = (values) => {
    handleCreateProject(values);
    formik.resetForm();
    formik.values.story = "";
    formik.values.title = "";
    formik.values.description = "";
  }

  return (
    <ScrollView style={{flex: 1, backgroundColor: colors.background}}>
      <ImageView
        images={[
          {
            uri:
              formik.values.picture ||
              "https://i.ibb.co/n0hwfdv/images-upload.png",
          },
        ]}
        imageIndex={0}
        visible={isImageViewVisible}
        onRequestClose={() => setIsImageViewVisible(false)}
      />
      <TouchableOpacity onPress={() => setIsImageViewVisible(true)}>
        <View>
          {cloudStorageIsLoading ? (
            <View style={{marginBottom: 56, marginTop: 30}}>
              <ActivityIndicator animating={true} color={colors.primary} />
            </View>
          ) : (
            <Image
              source={{
                uri:
                  formik.values.picture ||
                  "https://i.ibb.co/n0hwfdv/images-upload.png",
              }}
              style={{width: "100%", height: 200}}
            />
          )}
        </View>
      </TouchableOpacity>
      <TouchableOpacity
        style={{alignItems: "center", padding: 10}}
        onPress={async () =>
          await launchImageLibrary(null, async res => {
            if (!res.didCancel) {
              if (res.assets[0].fileSize > (10 * 1024 * 1024)){
                alert(en.error, "Maximum file size 10 MB");
              } else {
                handleUploadPicture({
                  name: res.assets[0].fileName,
                  uri: res.assets[0].uri,
                });
              }
            }
          })
        }>
        <Button icon="upload" mode="contained" compact color={colors.primary}>
          Cover Image
        </Button>
      </TouchableOpacity>
      <Divider />
      <View style={{flex: 1, padding: 25}}>
        <TextInput
          label="Title"
          mode="outlined"
          style={{height: 50}}
          onChangeText={formik.handleChange("title")}
          onBlur={formik.handleBlur("title")}
          value={formik.values.title}
          error={formik.errors.title && formik.touched.title}
        />
        <HelperText
          type="error"
          visible={Boolean(formik.errors.title && formik.touched.title)}>
          {formik.errors.title}
        </HelperText>
        <PaperDropdown
          listValue={categoryData}
          label="Category"
          setValue={value => formik.setFieldValue("categoryId", value)}
          value={formik.values.categoryId}
          inputProps={{
            error: formik.errors.categoryId && formik.touched.categoryId,
            onBlur: formik.handleBlur("categoryId"),
          }}
        />
        <HelperText
          type="error"
          visible={Boolean(
            formik.errors.categoryId && formik.touched.categoryId,
          )}>
          {formik.errors.categoryId}
        </HelperText>
        <TextInput
          label="Description"
          mode="outlined"
          multiline
          numberOfLines={5}
          style={{height: 100}}
          onChangeText={formik.handleChange("description")}
          onBlur={formik.handleBlur("description")}
          value={formik.values.description}
          error={formik.errors.description && formik.touched.description}
        />
        <HelperText
          type="error"
          visible={Boolean(
            formik.errors.description && formik.touched.description,
          )}>
          {formik.errors.description}
        </HelperText>
        <Modal
          transparent={true}
          animationType={"none"}
          visible={isUploading}
          onRequestClose={() => {}}>
          <View
            style={{
              flex: 1,
              alignItems: "center",
              flexDirection: "column",
              justifyContent: "space-around",
              backgroundColor: "#00000040",
            }}>
            <View
              style={{
                backgroundColor: "#FFFFFF",
                height: 200,
                width: 200,
                borderRadius: 10,
                display: "flex",
                alignItems: "center",
                justifyContent: "space-around",
              }}>
              <Text> Uploading media..... </Text>
              <ActivityIndicator animating={isUploading} size={75} />
            </View>
          </View>
        </Modal>
        <ScrollView>
          <View style={{backgroundColor: "black"}}>
            <RichEditor
              editorStyle={{backgroundColor: "#dce3fa"}}
              initialHeight={300}
              ref={richText}
              initialContentHTML={formik.values.story}
              onChange={descriptionText => {
                handleEditor(descriptionText);
              }}
            />
          </View>
        </ScrollView>
        <RichToolbar
          iconTint={"purple"}
          selectedIconTint={"pink"}
          disabledIconTint={"purple"}
          onPressAddImage={onPressAddImage}
          iconSize={20}
          editor={richText}
          actions={[
            actions.insertImage,
            actions.setBold,
            actions.setItalic,
            actions.insertBulletsList,
            actions.insertOrderedList,
            actions.insertLink,
            actions.keyboard,
            actions.setStrikethrough,
            actions.setUnderline,
            actions.removeFormat,
            actions.checkboxList,
            actions.undo,
            actions.redo,
            actions.insertVideo,
          ]}
          iconMap={{
            [actions.heading1]: ({tintColor}) => (
              <Text style={[{color: tintColor}]}>H1</Text>
            ),
          }}
          insertVideo={insertVideo}
        />
        <Button
          mode="contained"
          onPress={() => handleSubmit(formik.values)}
          loading={createProjectIsLoading}>
          Create Project
        </Button>
      </View>
    </ScrollView>
  );
};

export default withTheme(CreateProjectView);
