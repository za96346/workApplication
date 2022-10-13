import * as React from 'react';

import LinearGradient from "react-native-linear-gradient";
import { SafeAreaView } from "react-native";
import TabNavigator from 'react-native-tab-navigator';
import { language } from "./language";
import Home from "./page/HomePage/Home";
import { createBottomTabNavigator } from '@react-navigation/bottom-tabs';
import { Icon } from 'native-base';
import FontAwesome from 'react-native-vector-icons/FontAwesome5';
import MaterialIcons from 'react-native-vector-icons/MaterialIcons'
import materialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'
import lonicons from 'react-native-vector-icons/Ionicons'
import { TouchableOpacity } from 'react-native-gesture-handler';
import WorkList from './page/WorkListPage/WorkList';
import Setting from './page/SettingPage/Setting';
import CheckIn from './page/CheckInPage/CheckIn';
import StaffChange from './page/StaffChangePage/StaffChange';
import { navigation } from './type/type';

const Tab: any = createBottomTabNavigator();
export default class Main extends React.Component <navigation, any> {
    navigation: any
    constructor(props:any) {
        super(props)
        this.navigation = props.navigation
    }
    render(): JSX.Element {
        return (
            <SafeAreaView style={{flex: 1, justifyContent: 'flex-start', backgroundColor: 'white'}}>
            
            <LinearGradient 
                colors={['red','blue',]}
                start={{ x: 1, y: 0}}
                locations={[0, 0.5]}
                end={{ x: 1, y: 100 }}
                style={{bottom:0,flex:1}}
            >
                    <Tab.Navigator
                        screenOptions={{
                            tabBarStyle: { position: 'absolute', paddingBottom: -30},
                            tabBarActiveTintColor: '#95CACA',
                            tabBarHideOnKeyboard: true,
                            tabBarActiveBackgroundColor: 'rgba(150, 150, 150, 0.1)',
                            tabBarBadgeStyle: {
                                backgroundColor: '#0f0'
                            },
                            tabBarBadge: 3,
                            tabBarLabelPosition: 'below-icon',
                        }}
                        tabBarHideOnKeyboard
                        initialRouteName="Home">
                        <Tab.Screen
                            options={{
                                headerShown: false,
                                tabBarLabel: language.home,
                                tabBarIcon: ( props: any ) => {
                                    return <TouchableOpacity>
                                        <Icon name={props.focused ? 'home-sharp' : 'home-outline'} as={lonicons} color={props.color} size={18}/>
                                    </TouchableOpacity>
                                },
                                
                            }}
                            name="Home"
                            component={Home}/>
                        <Tab.Screen
                            options={{
                                headerShown: false,
                                tabBarLabel: language.workList,
                                tabBarIcon: ( props: any ) => (
                                    <TouchableOpacity>
                                        <Icon name={props.focused ? 'view-list' : 'list'}  as={MaterialIcons} color={props.color} size={18}/>
                                    </TouchableOpacity>
                                ),
                                
                            }}
                            name="WorkList"
                            component={WorkList}/>
                        <Tab.Screen
                            options={{
                                headerShown: false,
                                tabBarLabel: language.checkIn,
                                tabBarIcon: ( props: any ) => (
                                    <TouchableOpacity>
                                        <Icon name={props.focused ? 'clipboard-check-multiple' : 'clipboard-check-multiple-outline'} as={materialCommunityIcons} color={props.color} size={18}/>
                                    </TouchableOpacity>
                                ),
                                
                            }}
                            name="CheckIn"
                            component={CheckIn}/>
                        <Tab.Screen
                            options={{
                                headerShown: false,
                                tabBarLabel: language.staffChange,
                                tabBarIcon: ( props: any ) => (
                                    <TouchableOpacity>
                                        <Icon name={props.focused ? 'people-sharp' : 'people-outline'} as={lonicons} color={props.color} size={18}/>
                                    </TouchableOpacity>
                                ),
                                
                            }}
                            name="StaffChange"
                            component={StaffChange}/>
                        <Tab.Screen
                            options={{
                                headerShown: false,
                                tabBarLabel: language.setting,
                                tabBarIcon: ( props: any ) => (
                                    <TouchableOpacity>
                                        <Icon name={props.focused ? 'settings' : 'settings-outline'} as={lonicons} color={props.color} size={18}/>
                                    </TouchableOpacity>
                                ),
                                
                            }}
                            name="Setting"
                            component={Setting}/>
                    </Tab.Navigator>
            </LinearGradient>
            </SafeAreaView>
        );
    }
}