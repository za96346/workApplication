import { Box, Center, ScrollView, Text, useColorModeValue, View } from "native-base";
import React from "react";
import { Animated, Dimensions, Pressable, SafeAreaView, StatusBar } from "react-native";
import { SceneMap, TabView } from "react-native-tab-view";
import CO_WorkList from "../../component/CO_WorkList";
import RenderBar from "./RenderBar";
const data = new Array(100).fill({
    imgUrl: '',
    userName: 'jack',
    todayWorkType: '早班',
    todayWorkTime: '9 : 00 ~ 17 : 30'
})

const FirstRoute = () => <CO_WorkList
    styIdx="one"
    data={data}
  />;

const SecondRoute = () => <Center flex={1} my="4">
    This is Tab 2
  </Center>;

const ThirdRoute = () => <Center flex={1} my="4">
    This is Tab 3
  </Center>;

const FourthRoute = () => <Center flex={1} my="4">
    This is Tab 4{' '}
  </Center>;

const initialLayout = {
  width: Dimensions.get('window').width
};
const renderScene = SceneMap({
  first: FirstRoute,
  second: SecondRoute,
  third: ThirdRoute,
  fourth: FourthRoute
});

const Home = (): JSX.Element => {
  const [index, setIndex] = React.useState(0);
  const [routes] = React.useState([{
    key: 'first',
    title: '今日班表'
  }, {
    key: 'second',
    title: 'Tab 2'
  }, {
    key: 'third',
    title: 'Tab 3'
  }, {
    key: 'fourth',
    title: 'Tab 4'
  }]);

  return <TabView navigationState={{
    index,
    routes
  }} renderScene={renderScene} renderTabBar={(p) => RenderBar(p, index, setIndex)} onIndexChange={setIndex} initialLayout={initialLayout} style={{
    marginTop: StatusBar.currentHeight
  }} />;
}
export default Home

