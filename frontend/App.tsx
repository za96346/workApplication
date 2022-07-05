import React,{ Component } from "react";
import { NavigationContainer, ParamListBase, RouteProp } from "@react-navigation/native";
import {createStackNavigator} from '@react-navigation/stack';
import Login from "./src/login";
import Main from "./src/main";
import { Image, Platform, Text, TouchableOpacity } from "react-native";
import { typeNavigation } from "./src/type/type";

const Stack = createStackNavigator();
const headerOptions = ({route,navigation}:{route: RouteProp<ParamListBase, string>,navigation:typeNavigation}) => ({
	title: 'hihi',
	headerTintColor: 'black', // 字體顏色
	headerTitleStyle: {alignSelf: 'center', fontSize: 16}, // header 樣式
	headerStyle: {
	  height: Platform.OS === 'ios' ? 88 : 44,
	}, // 使用裝置來判斷 header 的高度
	headerRight: () => (
		<TouchableOpacity
		onPress={() => {
		  navigation.popToTop();
		}}>
			<Text>home</Text>
	  </TouchableOpacity>
	), // 右邊放入 icon
	headerLeft: () => (
	  <TouchableOpacity
		onPress={() => {
		  navigation.goBack();
		}}>
			<Text>返回</Text>
	  </TouchableOpacity>
	  // 左邊放入icon 並使用 navigation.goBack() 及 backToHome() 回上一頁
	),
  });

export default class App extends Component{
    render(): JSX.Element {
        return(
			<NavigationContainer >
				<Stack.Navigator screenOptions={headerOptions} initialRouteName="Login">
       				<Stack.Screen name="Login" component={Login} />
					<Stack.Screen name="Main" component={Main} />
				</Stack.Navigator>
			</NavigationContainer>
		)
    }
}