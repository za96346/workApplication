import React,{ Component } from "react";
import { NavigationContainer, ParamListBase, RouteProp } from "@react-navigation/native";
import {createStackNavigator} from '@react-navigation/stack';
import Login from "./src/login";
import Main from "./src/main";
import { Image, Platform, Text, TouchableOpacity } from "react-native";
import { typeNavigation } from "./src/type/type";
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
import storage from 'redux-persist/lib/storage';

//開發工具
import { composeWithDevTools } from 'redux-devtools-extension';
import { PersistGate } from "redux-persist/integration/react";
import thunk from "redux-thunk";


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

const reducer = combineReducers({
    loginData
})

const persistConfig = {
	key: 'root',
	storage: storage,
	whitelist: ['loginData'], // only member will be persisted
};


// 持久化根reducers
const persistedReducer = persistReducer(persistConfig, reducer);

const middleware = [thunk];
  
export const store = createStore(
  persistedReducer,
  composeWithDevTools(applyMiddleware(...middleware)),
);
export const persisStore = persistStore(store);

// if you are not using the middleWare , you can choosing option below.
//export const store = createStore(reducer,window.__REDUX_DEVTOOLS_EXTENSION__ && window.__REDUX_DEVTOOLS_EXTENSION__())


export class App extends Component{
    render(): JSX.Element {
        return(
			<Provider store={store}>
				<PersistGate loading={null} persistor={persisStore}>
					<NavigationContainer >
						<Stack.Navigator screenOptions={headerOptions} initialRouteName="Login">
							<Stack.Screen name="Login" component={Login} />
							<Stack.Screen name="Main" component={Main} />
						</Stack.Navigator>
					</NavigationContainer>
				</PersistGate>
			</Provider>
		)
    }
}


// 將store中的items值傳綁到props上
const mapStateToProps = (store: any):object => (
	{ items: store.items }
  )
  
  // 連接Redux store
  // 並把store.items綁到props.items，
  // actionCreators裡面的方法也綁到props上
export default connect(mapStateToProps, actionCreators)(App)