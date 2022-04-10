import {
  Image,
  TouchableOpacity,
  View,
  TextInput,
  FlatList,
  RefreshControl,
} from "react-native";
import {
  withTheme,
  Avatar,
  Text,
  List,
  Divider,
  ActivityIndicator,
} from "react-native-paper";
import ImageView from "react-native-image-viewing";
import React, {useState, useRef} from "react";
import moment from "moment";
import {useNavigation} from "@react-navigation/native";
import AwesomeIcon from "react-native-vector-icons/FontAwesome";
import {useFormik} from "formik";
import * as Yup from "yup";
import {DEFAULT_PROFILE_PICTURE, en} from "../../shared";

const CommentView = ({
  projectPicture,
  userPicture,
  userRoleName,
  userFullName,
  projectCreated,
  theme,
  comments,
  handlePostComment,
  alert,
  commentDataIsLoading,
  setIsEndReachRefresh,
  isEndReachRefresh,
  isPullRefresh,
  setIsPullRefresh,
  userProfilePicture,
}) => {
  const navigation = useNavigation();
  const {colors, fonts} = theme;
  const [isImageViewVisible, setIsImageViewVisible] = useState(false);
  const textInputRef = useRef();
  const formik = useFormik({
    initialValues: {
      body: "",
    },
    validationSchema: Yup.object({
      body: Yup.string()
        .required("please input your comment")
        .max(300, "300 max character length"),
    }),
    onSubmit: (values, {resetForm}) => {
      handlePostComment(values);
      resetForm();
      textInputRef.current.blur();
    },
  });
  return (
    <View style={{flex: 1}}>
      <ImageView
        images={[{uri: projectPicture}]}
        imageIndex={0}
        visible={isImageViewVisible}
        onRequestClose={() => setIsImageViewVisible(false)}
      />
      <TouchableOpacity onPress={() => setIsImageViewVisible(true)}>
        <View>
          <Image
            source={{
              uri: projectPicture || DEFAULT_PROFILE_PICTURE,
            }}
            style={{width: "100%", height: 200}}
          />
        </View>
      </TouchableOpacity>
      <View style={{backgroundColor: colors.background, flex: 1}}>
        <View
          style={{
            alignItems: "center",
            flexDirection: "row",
            padding: 10,
            paddingLeft: 10,
            paddingRight: 10,
          }}>
          <TouchableOpacity
            onPress={() => navigation.goBack()}
            style={{alignItems: "center", marginRight: 20}}>
            <AwesomeIcon name="arrow-left" size={20} style={{color: "black"}} />
          </TouchableOpacity>
          <View style={{flexDirection: "row", alignItems: "center"}}>
            <Avatar.Image
              size={50}
              source={{uri: userPicture || DEFAULT_PROFILE_PICTURE}}
            />
            <View style={{marginLeft: 20}}>
              <Text style={{...fonts.medium, fontSize: 17}}>
                {userFullName}
              </Text>
              <Text>{userRoleName}</Text>
              <Text style={{...fonts.light, fontSize: 12, marginTop: 5}}>
                {moment.utc(projectCreated).fromNow()}
              </Text>
            </View>
          </View>
        </View>
        <Divider />
        <View style={{flex: 1}}>
          {!isPullRefresh && !isEndReachRefresh && commentDataIsLoading ? (
            <View
              style={{
                flex: 1,
                alignItems: "center",
                justifyContent: "center",
              }}>
              <ActivityIndicator animating />
            </View>
          ) : (
            <FlatList
              data={comments}
              renderItem={({item}) => {
                return (
                  <List.Item
                    title={item.from}
                    description={item.body}
                    right={() => (
                      <View
                        style={{
                          alignItems: "center",
                          justifyContent: "center",
                        }}>
                        <Text style={{...fonts.light, fontSize: 12}}>
                          {moment.utc(item.createdAt).fromNow()}
                        </Text>
                      </View>
                    )}
                    left={() => (
                      <Avatar.Image
                        size={50}
                        source={{uri: item.picture || DEFAULT_PROFILE_PICTURE}}
                      />
                    )}
                  />
                );
              }}
              onEndReached={() => setIsEndReachRefresh(true)}
              ListFooterComponent={() => {
                if (isEndReachRefresh && commentDataIsLoading) {
                  return <ActivityIndicator size="large" />;
                } else {
                  return null;
                }
              }}
              refreshControl={
                <RefreshControl
                  refreshing={isPullRefresh && commentDataIsLoading}
                  onRefresh={() => setIsPullRefresh(true)}
                />
              }
            />
          )}
        </View>
        <Divider />
        <View
          style={{
            flexDirection: "row",
            alignItems: "center",
            padding: 5,
            paddingLeft: 10,
            paddingRight: 10,
          }}>
          <Avatar.Image
            size={40}
            source={{uri: userProfilePicture || DEFAULT_PROFILE_PICTURE}}
          />
          <TextInput
            placeholder="add comment..."
            style={{flex: 1, marginLeft: 10}}
            onBlur={formik.handleBlur("body")}
            onChangeText={formik.handleChange("body")}
            value={formik.values.body}
            ref={textInputRef}
          />
          <TouchableOpacity
            onPress={() => {
              if (formik.values.body === "") {
                alert(en.warning, "please input your comment");
              }
              formik.handleSubmit();
            }}>
            <Text style={{color: colors.primary}}>post</Text>
          </TouchableOpacity>
        </View>
      </View>
    </View>
  );
};

export default withTheme(CommentView);
