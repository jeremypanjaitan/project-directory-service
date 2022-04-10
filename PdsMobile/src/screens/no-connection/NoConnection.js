import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import React from "react";

const NoConnectionScreen = ({onRetry, isRetrying}) => (
  <View style={styles.container}>
    <Image style={styles.image} source={require("../../assets/NOT_FOUND.png")} />
    <Text style={styles.modalTitle}>Connection Error</Text>
    <Text style={styles.modalText}>
      Oops! Looks like your device is not connected to the Internet, please turn on your wifi/phone data.
    </Text>
  </View>
);

const styles = StyleSheet.create({
  image: {
    width: "100%",
    height: "60%",
    marginTop: 30,
  },
  container: {
    backgroundColor: "#fff",
    paddingHorizontal: 16,
    paddingTop: 20,
    paddingBottom: 40,
    alignItems: "center",
    justifyContent: "center",
  },
  modalTitle: {
    fontSize: 22,
    fontWeight: "600",
  },
  modalText: {
    fontSize: 18,
    color: "#555",
    marginTop: 24,
    textAlign: "center",
    marginBottom: 10,
  },
});


export default NoConnectionScreen;
