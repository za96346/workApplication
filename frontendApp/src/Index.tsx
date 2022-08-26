
import Login from './Login';
import Main from './Main';

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