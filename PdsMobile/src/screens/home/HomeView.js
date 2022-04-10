import React, {useState} from "react";
import {
  Card,
  Button,
  Paragraph,
  ActivityIndicator,
  Modal,
  Searchbar,
  Text,
  List,
} from "react-native-paper";
import {
  FlatList,
  SafeAreaView,
  TouchableOpacity,
  RefreshControl,
  StyleSheet,
  View,
} from "react-native";
import {DefaultSpinner, PaperDropdown} from "../../components";
import {useNavigation} from "@react-navigation/native";
import {APP_HOME, APP_HOME_SCREEN, APP_PROJECT} from "../../config";
import AwesomeIcon from "react-native-vector-icons/FontAwesome";
import {navigationRef} from "../../navigator";

const HomeView = ({
  allProjectData,
  allProjectDataIsLoading,
  isPullRefresh,
  setCurrentPage,
  setIsPullRefresh,
  setIsEndReachRefresh,
  isEndReachRefresh,
  handleSearch,
  handleCategoryId,
  handleLike,
  allCategoriesData,
  handleStartSearch,
}) => {
  console.log("rendered home")
  const [searchQuery, setSearchQuery] = useState("");
  const [likeQuery, setLikeQuery] = useState("");
  const [categoryIdQuery, setCategoryIdQuery] = useState(null);
  const onChangeSearch = query => setSearchQuery(query);
  const [visibleSearch, setVisibleSearch] = useState(false);
  const [visibleFilterAndSort, setVisibleFilterAndSort] = useState(false);
  const showModalSearch = () => setVisibleSearch(true);
  const hideModalSearch = () => setVisibleSearch(false);
  const showModalFilterAndSort = () => setVisibleFilterAndSort(true);
  const hideModalFilterAndSort = () => setVisibleFilterAndSort(false);
  const navigation = useNavigation();
  const [header, setHeader] = useState("");
  const [subHeader, setSubHeader] = useState("");
  const sorting = [
    {
      value: "DESC",
      label: "Most likes",
    },
    {
      value: "ASC",
      label: "Least likes",
    },
  ];
  const setSubHeaderLike = likeQuery => {
    if (likeQuery === "DESC") {
      setSubHeader("Most Likes");
    } else if (likeQuery === "ASC") {
      setSubHeader("Least Likes");
    } else {
      setSubHeader("");
    }
  };
  const setHeaderCategory = categoryIdQuery => {
    if (categoryIdQuery === null) {
      setHeader("");
    } else {
      setHeader(allCategoriesData[categoryIdQuery - 1].label);
    }
  };

  const renderItem = ({item}) => {
    return (
      <Card style={{margin: 10}}>
        <TouchableOpacity
          onPress={() =>
            navigationRef.navigate(APP_HOME_SCREEN, {
              screen: APP_PROJECT,
              params: {projectId: item.ID, backStack: APP_HOME},
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
    if (isEndReachRefresh && allProjectDataIsLoading) {
      return <ActivityIndicator size="large" />;
    } else {
      return null;
    }
  };
  if (!isPullRefresh && !isEndReachRefresh && allProjectDataIsLoading) {
    return <DefaultSpinner />;
  }

  return (
    <SafeAreaView style={{flex: 1}}>
      <FlatList
        data={allProjectData}
        renderItem={renderItem}
        keyExtractor={item => item.ID}
        refreshControl={
          <RefreshControl
            refreshing={isPullRefresh && allProjectDataIsLoading}
            onRefresh={() => setIsPullRefresh(true)}
          />
        }
        onEndReached={() => setIsEndReachRefresh(true)}
        ListFooterComponent={renderFooter}
        ListHeaderComponent={
          allProjectData.length === 0 && !allProjectDataIsLoading ? (
            <>
              {header !== "" || subHeader !== "" ? (
                <>
                  <List.Item
                    title={
                      <Text style={styles.header}>Category: {header}</Text>
                    }
                    description={
                      <Text style={styles.subHeader}>{subHeader}</Text>
                    }
                    right={props => (
                      <>
                        <AwesomeIcon
                          {...props}
                          name="search"
                          size={20}
                          style={{
                            color: "silver",
                            marginRight: 20,
                            marginTop: 10,
                          }}
                          onPress={showModalSearch}
                        />
                        <AwesomeIcon
                          {...props}
                          name="bars"
                          size={20}
                          style={{
                            color: "silver",
                            marginRight: 10,
                            marginTop: 10,
                          }}
                          onPress={showModalFilterAndSort}
                        />
                      </>
                    )}
                  />
                  <View
                    style={{
                      alignItems: "center",
                      flexDirection: "row",
                      padding: 5,
                      paddingLeft: 10,
                      paddingRight: 10,
                      marginTop: 5,
                    }}>
                    <TouchableOpacity
                      onPress={() => navigation.replace(APP_HOME)}
                      style={{
                        alignItems: "center",
                        marginRight: 20,
                      }}>
                      <AwesomeIcon
                        name="arrow-left"
                        size={20}
                        style={{color: "black"}}
                      />
                    </TouchableOpacity>
                  </View>
                  <Text style={styles.noData}>No Data...</Text>
                </>
              ) : (
                <>
                  <List.Item
                    title={<Text style={styles.header}>All Projects</Text>}
                    description={
                      <Text style={styles.subHeader}>{subHeader}</Text>
                    }
                    right={props => (
                      <>
                        <AwesomeIcon
                          {...props}
                          name="search"
                          size={20}
                          style={{
                            color: "silver",
                            marginRight: 20,
                            marginTop: 10,
                          }}
                          onPress={showModalSearch}
                        />
                        <AwesomeIcon
                          {...props}
                          name="bars"
                          size={20}
                          style={{
                            color: "silver",
                            marginRight: 10,
                            marginTop: 10,
                          }}
                          onPress={showModalFilterAndSort}
                        />
                      </>
                    )}
                  />
                  <View
                    style={{
                      alignItems: "center",
                      flexDirection: "row",
                      padding: 5,
                      paddingLeft: 10,
                      paddingRight: 10,
                      marginTop: 5,
                    }}>
                    <TouchableOpacity
                      onPress={() => navigation.replace(APP_HOME)}
                      style={{
                        alignItems: "center",
                        marginRight: 20,
                      }}>
                      <AwesomeIcon
                        name="arrow-left"
                        size={20}
                        style={{color: "black"}}
                      />
                    </TouchableOpacity>
                  </View>
                  <Text style={styles.noData}>No Data...</Text>
                </>
              )}
            </>
          ) : (
            <View>
              {header === "" ? (
                <View>
                  {searchQuery !== "" ? (
                    <>
                      <>
                        <List.Item
                          title={
                            <Text style={styles.header}>All Projects</Text>
                          }
                          description={
                            <Text style={styles.subHeader}>{subHeader}</Text>
                          }
                          right={props => (
                            <>
                              <AwesomeIcon
                                {...props}
                                name="search"
                                size={20}
                                style={{
                                  color: "silver",
                                  marginRight: 20,
                                  marginTop: 20,
                                }}
                                onPress={showModalSearch}
                              />
                              <AwesomeIcon
                                {...props}
                                name="bars"
                                size={20}
                                style={{
                                  color: "silver",
                                  marginRight: 10,
                                  marginTop: 20,
                                }}
                                onPress={showModalFilterAndSort}
                              />
                            </>
                          )}
                        />
                        <View
                          style={{
                            alignItems: "center",
                            flexDirection: "row",
                            padding: 5,
                            paddingLeft: 10,
                            paddingRight: 10,
                          }}>
                          <TouchableOpacity
                            onPress={() => navigation.replace(APP_HOME)}
                            style={{
                              alignItems: "center",
                              marginRight: 20,
                            }}>
                            <AwesomeIcon
                              name="arrow-left"
                              size={20}
                              style={{color: "black"}}
                            />
                          </TouchableOpacity>
                        </View>
                      </>
                    </>
                  ) : (
                    <>
                      {subHeader !== "" ? (
                        <>
                          <List.Item
                            title={
                              <Text style={styles.header}>All Projects</Text>
                            }
                            description={
                              <Text style={styles.subHeader}>{subHeader}</Text>
                            }
                            right={props => (
                              <>
                                <AwesomeIcon
                                  {...props}
                                  name="search"
                                  size={20}
                                  style={{
                                    color: "silver",
                                    marginRight: 20,
                                    marginTop: 10,
                                  }}
                                  onPress={showModalSearch}
                                />
                                <AwesomeIcon
                                  {...props}
                                  name="bars"
                                  size={20}
                                  style={{
                                    color: "silver",
                                    marginRight: 10,
                                    marginTop: 10,
                                  }}
                                  onPress={showModalFilterAndSort}
                                />
                              </>
                            )}
                          />
                          <View
                            style={{
                              alignItems: "center",
                              flexDirection: "row",
                              padding: 5,
                              paddingLeft: 10,
                              paddingRight: 10,
                            }}>
                            <TouchableOpacity
                              onPress={() => navigation.replace(APP_HOME)}
                              style={{
                                alignItems: "center",
                                marginRight: 20,
                              }}>
                              <AwesomeIcon
                                name="arrow-left"
                                size={20}
                                style={{color: "black"}}
                              />
                            </TouchableOpacity>
                          </View>
                        </>
                      ) : (
                        <List.Item
                          title={
                            <Text style={styles.header}>All Projects</Text>
                          }
                          description={
                            <Text style={styles.subHeader}>{subHeader}</Text>
                          }
                          right={props => (
                            <>
                              <AwesomeIcon
                                {...props}
                                name="search"
                                size={20}
                                style={{
                                  color: "silver",
                                  marginRight: 20,
                                  marginTop: 10,
                                }}
                                onPress={showModalSearch}
                              />
                              <AwesomeIcon
                                {...props}
                                name="bars"
                                size={20}
                                style={{
                                  color: "silver",
                                  marginRight: 10,
                                  marginTop: 10,
                                }}
                                onPress={showModalFilterAndSort}
                              />
                            </>
                          )}
                        />
                      )}
                    </>
                  )}
                </View>
              ) : (
                <>
                  <List.Item
                    title={
                      <Text style={styles.header}>Category: {header}</Text>
                    }
                    description={
                      <Text style={styles.subHeader}>{subHeader}</Text>
                    }
                    right={props => (
                      <>
                        <AwesomeIcon
                          {...props}
                          name="search"
                          size={20}
                          style={{
                            color: "silver",
                            marginRight: 20,
                            marginTop: 20,
                          }}
                          onPress={showModalSearch}
                        />
                        <AwesomeIcon
                          {...props}
                          name="bars"
                          size={20}
                          style={{
                            color: "silver",
                            marginRight: 10,
                            marginTop: 20,
                          }}
                          onPress={showModalFilterAndSort}
                        />
                      </>
                    )}
                  />
                  <View
                    style={{
                      alignItems: "center",
                      flexDirection: "row",
                      padding: 5,
                      paddingLeft: 10,
                      paddingRight: 10,
                    }}>
                    <TouchableOpacity
                      onPress={() => navigation.replace(APP_HOME)}
                      style={{
                        alignItems: "center",
                        marginRight: 20,
                      }}>
                      <AwesomeIcon
                        name="arrow-left"
                        size={20}
                        style={{color: "black"}}
                      />
                    </TouchableOpacity>
                  </View>
                </>
              )}
            </View>
          )
        }
      />

      <Modal
        visible={visibleSearch}
        onDismiss={hideModalSearch}
        contentContainerStyle={styles.modal}>
        <View style={{padding: 10, marginTop: 20}}>
          <Searchbar
            placeholder="Search by title..."
            onChangeText={onChangeSearch}
            value={searchQuery}
          />
          <Button
            style={{
              alignItems: "center",
              marginTop: 30,
              flexDirection: "row",
              justifyContent: "center",
            }}
            mode="contained"
            onPress={() => {
              handleSearch(searchQuery);
              handleStartSearch({
                searchByTitle: searchQuery,
                sortByLikes: likeQuery,
                categoryId: categoryIdQuery,
              });

              setCurrentPage(1);
              hideModalSearch();
            }}>
            Search
          </Button>
        </View>
      </Modal>

      <Modal
        visible={visibleFilterAndSort}
        onDismiss={hideModalFilterAndSort}
        contentContainerStyle={styles.modal}>
        <View style={{padding: 10, marginTop: 20}}>
          <PaperDropdown
            listValue={allCategoriesData}
            label="Search by category"
            setValue={value => setCategoryIdQuery(value)}
            value={categoryIdQuery}
          />
          <PaperDropdown
            listValue={sorting}
            label="Sorting by likes"
            setValue={value => setLikeQuery(value)}
            value={likeQuery}
          />
          <Button
            style={{
              alignItems: "center",
              marginTop: 30,
              flexDirection: "row",
              justifyContent: "center",
            }}
            mode="contained"
            onPress={() => {
              handleCategoryId(categoryIdQuery);
              handleLike(likeQuery);

              setHeaderCategory(categoryIdQuery);
              setSubHeaderLike(likeQuery);
              handleStartSearch({
                searchByTitle: searchQuery,
                sortByLikes: likeQuery,
                categoryId: categoryIdQuery,
              });

              setCurrentPage(1);
              hideModalFilterAndSort();
            }}>
            Submit
          </Button>
        </View>
      </Modal>
    </SafeAreaView>
  );
};
const styles = StyleSheet.create({
  fab: {
    position: "absolute",
    margin: 20,
    right: 10,
    bottom: 0,
    backgroundColor: "#66A8E6",
  },
  modal: {
    backgroundColor: "white",
    padding: 20,
    width: 300,
    alignSelf: "center",
  },
  header: {
    fontSize: 18,
    color: "black",
    fontWeight: "bold",
    marginTop: 5,
    marginLeft: 10,
  },
  subHeader: {
    fontSize: 15,
    color: "black",
    fontWeight: "bold",
    marginTop: 5,
    marginLeft: 10,
  },
  noData: {
    fontSize: 30,
    color: "black",
    fontWeight: "bold",
    textAlign: "center",
    marginTop: 220,
  },
});
export default HomeView;
