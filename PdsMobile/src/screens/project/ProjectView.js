import {ScrollView, View, Image, TouchableOpacity} from "react-native";
import {
  Title,
  withTheme,
  Divider,
  Card,
  Button,
  Avatar,
  Text,
  Chip,
  Menu,
} from "react-native-paper";
import AwesomeIcon from "react-native-vector-icons/FontAwesome";
import moment from "moment";
import RenderHtml from "react-native-render-html";
import React, {useState} from "react";
import {DefaultSpinner} from "../../components";
import {useWindowDimensions} from "react-native";
import ImageView from "react-native-image-viewing";
import {
  APP_ACCOUNT_DETAILS_PROJECT,
  APP_COMMENT,
  APP_HOME,
  EDIT_PROJECT,
} from "../../config";
import {DEFAULT_PROFILE_PICTURE} from "../../shared";

const ProjectView = ({
  profileDataIsLoading,
  projectDataIsLoading,
  categoryDataIsLoading,
  projectLikeIsLoading,
  projectData,
  profileData,
  totalLikes,
  categoryData,
  theme,
  useNavigation,
  isUserLike,
  likeStart,
  dislikeStart,
  viewIsLoading,
  totalViewData,
  totalComments,
  totalViewIsLoading,
  projectId,
  handleDeleteProject,
  deleteProjectIsLoading,
  accountDetails,
  backStack,
}) => {
  console.log("rendered new");
  const [visible, setVisible] = React.useState(false);
  const openMenu = () => setVisible(true);
  const closeMenu = () => setVisible(false);
  const handleEdit = () => {
    closeMenu();
    navigation.navigate(EDIT_PROJECT, {
      projectPicture: projectData.picture,
      projectTitle: projectData.title,
      projectCategory: projectData.categoryId,
      projectStory: projectData.story,
      projectDescription: projectData.description,
      projectId,
    });
  };
  const {colors, fonts} = theme;
  const navigation = useNavigation();
  const {width} = useWindowDimensions();
  const [isImageViewVisible, setIsImageViewVisible] = useState(false);
  if (
    deleteProjectIsLoading ||
    profileDataIsLoading ||
    projectDataIsLoading ||
    categoryDataIsLoading ||
    projectLikeIsLoading ||
    viewIsLoading ||
    totalViewIsLoading ||
    !projectData ||
    !profileData ||
    !categoryData
  ) {
    return <DefaultSpinner />;
  }
  return (
    <ScrollView>
      <ImageView
        images={[{uri: projectData.picture}]}
        imageIndex={0}
        visible={isImageViewVisible}
        onRequestClose={() => setIsImageViewVisible(false)}
      />
      <TouchableOpacity onPress={() => setIsImageViewVisible(true)}>
        <View>
          <Image
            source={{uri: projectData.picture}}
            style={{width: "100%", height: 200}}
          />
        </View>
      </TouchableOpacity>
      <View style={{backgroundColor: colors.background}}>
        <View
          style={{
            alignItems: "center",
            flexDirection: "row",
            padding: 10,
            paddingLeft: 10,
            paddingRight: 10,
          }}>
          <TouchableOpacity
            onPress={
              backStack === APP_ACCOUNT_DETAILS_PROJECT
                ? () => navigation.replace(APP_ACCOUNT_DETAILS_PROJECT)
                : () => navigation.replace(APP_HOME)
            }
            style={{alignItems: "center", marginRight: 20}}>
            <AwesomeIcon name="arrow-left" size={20} style={{color: "black"}} />
          </TouchableOpacity>
          <View
            style={{
              flexDirection: "row",
              justifyContent: "space-between",
              flex: 1,
              alignItems: "center",
            }}>
            <View style={{flexDirection: "row", alignItems: "center"}}>
              <Avatar.Image
                size={50}
                source={{
                  uri: profileData.picture || DEFAULT_PROFILE_PICTURE,
                }}
              />
              <View style={{marginLeft: 20}}>
                <Text style={{...fonts.medium, fontSize: 17}}>
                  {profileData.fullName}
                </Text>
                <Text>{profileData.roleName}</Text>
                <Text style={{...fonts.light, fontSize: 12, marginTop: 5}}>
                  {moment.utc(projectData.createdAt).fromNow()}
                </Text>
              </View>
            </View>
            <View>
              {projectData.canEdit && projectData.canDelete && (
                <Menu
                  visible={visible}
                  onDismiss={closeMenu}
                  anchor={
                    <Button onPress={openMenu}>
                      <AwesomeIcon
                        name="bars"
                        size={20}
                        style={{color: "black"}}
                      />
                    </Button>
                  }>
                  <Menu.Item
                    title={"Edit Project"}
                    onPress={() => handleEdit()}
                    icon={"pencil"}
                  />
                  <Menu.Item
                    title={"Delete Project"}
                    onPress={() => handleDeleteProject()}
                    icon={"delete"}
                  />
                </Menu>
              )}
            </View>
          </View>
        </View>
        <Divider />
        <Card>
          <Card.Content>
            <View>
              <Title>{projectData.title}</Title>
              <View
                style={{
                  flexDirection: "row",
                  alignItems: "center",
                  justifyContent: "flex-start",
                  marginTop: 10,
                  marginBottom: 10,
                }}>
                <Chip>{categoryData.name}</Chip>
              </View>
            </View>
            <RenderHtml
              source={{html: projectData.story}}
              enableExperimentalMarginCollapsing={true}
              contentWidth={width - 20}
              ignoredDomTags={["iframe", "button"]}
              st
            />
          </Card.Content>
          <Divider />
          <Card.Actions>
            <TouchableOpacity>
              <Button
                icon="comment"
                mode="compact"
                onPress={() =>
                  navigation.navigate(APP_COMMENT, {
                    projectPicture: projectData.picture,
                    userPicture: profileData.picture,
                    userRoleName: profileData.roleName,
                    userFullName: profileData.fullName,
                    projectCreated: projectData.createdAt,
                    projectId,
                  })
                }>
                {totalComments}
              </Button>
            </TouchableOpacity>
            <TouchableOpacity>
              <Button
                onPress={isUserLike ? () => dislikeStart() : () => likeStart()}
                icon={isUserLike ? "thumb-up" : "thumb-up-outline"}
                mode="compact">
                {totalLikes}
              </Button>
            </TouchableOpacity>
            <Button icon="eye" mode="compact">
              {totalViewData.totalViews}
            </Button>
          </Card.Actions>
        </Card>
      </View>
    </ScrollView>
  );
};

export default withTheme(ProjectView);
