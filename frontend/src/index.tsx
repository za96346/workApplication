
import Login from './login';
import Main from './main';

const RootNavigator = ({
    login: {
        screen: Login,
        navigationOptions:{}
    },
    main: {
        screen:Main,
    },
});


export default RootNavigator;