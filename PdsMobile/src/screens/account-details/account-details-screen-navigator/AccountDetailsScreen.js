import React from "react";
import {
  Text,
  SafeAreaView,
  StyleSheet,
  TouchableOpacity,
  View,
} from "react-native";
import {
  goToAccountDetailsActivity,
  goToAccountDetailsProject,
} from "../../../navigator";
import AwesomeIcon from "react-native-vector-icons/FontAwesome";
import {List} from "react-native-paper";
import {useAuth} from "../../../services";

const AccountDetailsScreen = () => {
  const auth = useAuth();
  return (
    <SafeAreaView>
      <Text style={styles.title}>Hello {auth?.userData?.data?.fullName}</Text>
      <View>
        <TouchableOpacity onPress={() => goToAccountDetailsActivity()}>
          <List.Item
            title={<Text style={styles.subTitle}>Your Activity</Text>}
            right={props => (
              <AwesomeIcon
                {...props}
                name="angle-right"
                size={40}
                style={{
                  color: "#679DDE",
                  marginLeft: 10,
                  marginTop: 20,
                }}
              />
            )}
          />
        </TouchableOpacity>

        <View
          style={{
            flexDirection: "row",
            alignItems: "center",
            marginBottom: 10,
          }}>
          <View style={{flex: 1, height: 1, backgroundColor: "silver"}} />
        </View>
        <TouchableOpacity onPress={() => goToAccountDetailsProject()}>
          <List.Item
            title={<Text style={styles.subTitle}>Your Project</Text>}
            right={props => (
              <AwesomeIcon
                {...props}
                name="angle-right"
                size={40}
                style={{
                  color: "#679DDE",
                  marginLeft: 10,
                  marginTop: 20,
                }}
              />
            )}
          />
        </TouchableOpacity>
        <View
          style={{
            flexDirection: "row",
            alignItems: "center",
            marginBottom: 10,
          }}>
          <View style={{flex: 1, height: 1, backgroundColor: "silver"}} />
        </View>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  title: {
    fontSize: 30,
    color: "black",
    fontWeight: "bold",
    marginTop: 10,
    marginLeft: 15,
    marginBottom: 30,
  },
  subTitle: {
    fontSize: 22,
    color: "black",
    fontWeight: "700",
    marginLeft: 15,
  },
});

export default AccountDetailsScreen;
