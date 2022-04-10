import React from "react";
import moment from "moment";
import {
  FlatList,
  RefreshControl,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import {
  ActivityIndicator,
  Card,
  List,
  Text,
  withTheme,
} from "react-native-paper";
import {SafeAreaView} from "react-native-safe-area-context";
import {DefaultSpinner} from "../../../components";
import AwesomeIcon from "react-native-vector-icons/FontAwesome";
import {goToAccountDetails} from "../../../navigator";

const AccountDetailsActivityView = ({
  activities,
  profileActivityIsLoading,
  setIsEndReachRefresh,
  isEndReachRefresh,
  isPullRefresh,
  setIsPullRefresh,
}) => {
  const renderFooter = () => {
    if (isEndReachRefresh && profileActivityIsLoading) {
      return <ActivityIndicator size="large" />;
    } else {
      return null;
    }
  };
  if (!isPullRefresh && !isEndReachRefresh && profileActivityIsLoading) {
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
        <Text style={styles.subTitle}>Your Activity</Text>
      </View>
      <View style={{flexDirection: "row", alignItems: "center"}}>
        <View style={{flex: 1, height: 1, backgroundColor: "silver"}} />
      </View>

      <FlatList
        data={activities}
        renderItem={({item}) => {
          return (
            <Card
              style={{
                borderRadius: 20,
                shadowRadius: 10,
                marginTop: 10,
                marginBottom: 10,
                shadowOpacity: 10,
                borderColor: "black",
              }}>
              {item.Type === "LIKE" ? (
                <List.Icon color="silver" icon="thumb-up" />
              ) : item.Type === "DISLIKE" ? (
                <List.Icon color="silver" icon="thumb-down" />
              ) : (
                <List.Icon color="silver" icon="comment" />
              )}
              <List.Item title={item.Header} description={item.Body} />
              <Text style={{marginLeft: 20}}>
                {moment.utc(item.CreatedAt).fromNow()}
              </Text>
            </Card>
          );
        }}
        keyExtractor={item => item.ID}
        onEndReached={() => setIsEndReachRefresh(true)}
        ListFooterComponent={renderFooter}
        refreshControl={
          <RefreshControl
            refreshing={isPullRefresh && profileActivityIsLoading}
            onRefresh={() => setIsPullRefresh(true)}
          />
        }
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

export default withTheme(AccountDetailsActivityView);
