/**
 * @format
 */

import {AppRegistry} from 'react-native';
import Login from './src/login';
import {name as appName} from './app.json';
import { StackNavigator } from 'react-navigation';
const navigation = StackNavigator({
    login: {screen: Login}
})
AppRegistry.registerComponent(appName, () => navigation);
