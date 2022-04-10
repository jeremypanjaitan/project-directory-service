import {useNavigation} from "@react-navigation/core";
import React from "react";
import {
  FlatList,
  RefreshControl,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import {
  ActivityIndicator,
  Button,
  Card,
  Paragraph,
  Text,
  withTheme,
} from "react-native-paper";
import {SafeAreaView} from "react-native-safe-area-context";
import {DefaultSpinner} from "../../../components";
import {APP_ACCOUNT_DETAILS_PROJECT, APP_PROJECT} from "../../../config";
import {goToAccountDetails} from "../../../navigator";
import AwesomeIcon from "react-native-vector-icons/FontAwesome";

const AccountDetailsProjectView = ({
  projects,
  profileProjectIsLoading,
  setIsEndReachRefresh,
  isEndReachRefresh,
  isPullRefresh,
  setIsPullRefresh,
}) => {
  const navigation = useNavigation();
  console.log("account detail project")

  const renderItemProjects = ({item}) => {
    return (
      <Card style={{margin: 10}}>
        <TouchableOpacity
          onPress={() =>
            navigation.navigate(APP_PROJECT, {
              projectId: item.ID,
              backStack: APP_ACCOUNT_DETAILS_PROJECT,
            })
          }>
          <Card.Cover source={{uri: item.picture}} />
          <Card.Title title={item.title} titleNumberOfLines={2} />
          <Card.Content>
            <Paragraph>{item.description}</Paragraph>
          </Card.Content>
        </TouchableOpacity>
        <Card.Actions>
          <Button icon="comment">{item.totalComments}</Button>
          <Button icon="thumb-up">{item.totalLikes}</Button>
          <Button icon="eye" mode="compact">
            {item.totalViews}
          </Button>
        </Card.Actions>
      </Card>
    );
  };
  const renderFooter = () => {
    if (isEndReachRefresh && profileProjectIsLoading) {
      return <ActivityIndicator size="large" />;
    } else {
      return null;
    }
  };
  if (!isPullRefresh && !isEndReachRefresh && profileProjectIsLoading) {
    return <DefaultSpinner />;
  }
  return (
    <SafeAreaView style={{flex: 1}}>
      <View
        style={{
          alignItems: "center",
          flexDirection: "row",
          padding: 10,
          paddingLeft: 10,
          paddingRight: 10,
        }}>
        <TouchableOpacity
          onPress={() => goToAccountDetails()}
          style={{alignItems: "center", marginRight: 20, marginTop: 20}}>
          <AwesomeIcon name="arrow-left" size={20} style={{color: "black"}} />
        </TouchableOpacity>
        <Text style={styles.subTitle}>Your Project</Text>
      </View>
      <View style={{flexDirection: "row", alignItems: "center"}}>
        <View style={{flex: 1, height: 1, backgroundColor: "silver"}} />
      </View>
      <FlatList
        data={projects}
        renderItem={renderItemProjects}
        keyExtractor={item => item.ID}
        refreshControl={
          <RefreshControl
            refreshing={isPullRefresh && profileProjectIsLoading}
            onRefresh={() => setIsPullRefresh(true)}
          />
        }
        onEndReached={() => setIsEndReachRefresh(true)}
        ListFooterComponent={renderFooter}
      />
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "space-evenly",
    alignItems: "center",
  },
  title: {
    fontSize: 30,
    color: "black",
    fontWeight: "bold",
    marginTop: 5,
    marginLeft: 15,
  },
  subTitle: {
    marginTop: 20,
    fontSize: 22,
    color: "black",
    fontWeight: "700",
  },
});

export default withTheme(AccountDetailsProjectView);
