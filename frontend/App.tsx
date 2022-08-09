import React,{ Component } from "react";
import { NavigationContainer, ParamListBase, RouteProp } from "@react-navigation/native";
import {createStackNavigator, Header} from '@react-navigation/stack';
import Login from "./src/Login";
import Main from "./src/Main";
import { Image, Platform, Text, TouchableOpacity } from "react-native";
import { connect, Provider } from "react-redux";
import {
    legacy_createStore as createStore,
    combineReducers,
	applyMiddleware
} from 'redux';
import { loginData } from './src/reducer/Reducer';
import * as actionCreators from './src/action/action'

// 持久化存储 state
import { persistStore, persistReducer } from 'redux-persist';
//import storage from 'redux-persist/lib/storage';
import AsyncStorage from "@react-native-community/async-storage";

//開發工具
import { composeWithDevTools } from 'redux-devtools-extension';
import { PersistGate } from "redux-persist/integration/react";
import thunk from "redux-thunk";
import { extendTheme, Icon, NativeBaseProvider } from "native-base";
import FontAwesome from 'react-native-vector-icons/FontAwesome5';
import AntDesign from 'react-native-vector-icons/AntDesign';


const Stack = createStackNavigator();
const headerOptions = ({ route, navigation } : { route: any , navigation: any }) => ({
	title: route?.params?.title || '',
	headerTintColor: route?.params?.headerTintColor || 'black', // 字體顏色
	headerTitleStyle: route?.params?.headerTitleStyle || {alignSelf: 'center', fontSize: 16}, // header 樣式
	headerStyle: {
	  height: Platform.OS === 'ios' ? 88 : 44,
	}, // 使用裝置來判斷 header 的高度
	headerRight: () => (
		<TouchableOpacity
			onPress={() => {
			navigation.popToTop();
			}}>
			<Icon
				style={{marginRight: 10}}
				color={'#C7C7E2'}
				size={10}
				name="home"
				as={FontAwesome}
			/>
	  </TouchableOpacity>
	), // 右邊放入 icon
	headerLeft: () => (
	  <TouchableOpacity
		onPress={() => {
			try {
				navigation.goBack();	
			} catch {
				//pass
			}
		}}>
			<Icon
				style={{marginLeft: 10}}
				color={'#C7C7E2'}
				size={10}
				name="back"
				as={AntDesign}
			/>
	  </TouchableOpacity>
	  // 左邊放入icon 並使用 navigation.goBack() 及 backToHome() 回上一頁
	),
});

const reducer = combineReducers({
    loginData: loginData
})

const persistConfig = {
	key: 'root',
	storage: AsyncStorage,
	whitelist: ['loginData'], // only member will be persisted
};


// 持久化根reducers
const persistedReducer = persistReducer(persistConfig, reducer);

const middleware = [thunk];
  
export const store = createStore(
  persistedReducer,
  composeWithDevTools(applyMiddleware(...middleware)),
);
const persisStore = persistStore(store);

// if you are not using the middleWare , you can choosing option below.
//export const store = createStore(reducer,window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__())

const newColorTheme = {
	brand: {
	  900: '#8287af',
	  800: '#7c83db',
	  700: '#b3bef6',
	},
  };

export default class App extends Component{
    render(): JSX.Element {
        return(
			<Provider store={store}>
				<PersistGate loading={null} persistor={persisStore}>
					<NativeBaseProvider theme={extendTheme({ colors: newColorTheme })}>
						<NavigationContainer >
							<Stack.Navigator
								screenOptions={headerOptions}
								initialRouteName="Login">
								<Stack.Screen 
									name="Login"
									component={Login}
									options={{headerShown: false}}//隱藏頭
								/>
								<Stack.Screen
									options={{headerShown: false}}
									name="Main" 
									component={Main} 
								/>
							</Stack.Navigator>
						</NavigationContainer>
					</NativeBaseProvider>
				</PersistGate>
			</Provider>
		)
    }
}


// 將store中的items值傳綁到props上
//const mapStateToProps = (store: any):object => (
	//{ items: store.items }
 // )
  
  // 連接Redux store
  // 並把store.items綁到props.items，
  // actionCreators裡面的方法也綁到props上
//export default connect(mapStateToProps, actionCreators)(App)