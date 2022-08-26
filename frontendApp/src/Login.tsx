import React,{Component} from "react";
import { ActivityIndicator, Animated, Button, ProgressViewIOSBase, Text, TextInput, Touchable, TouchableHighlight, View } from "react-native";
import { styleLogin,shadowWrapper } from "./style/styles";
import { StyleSheet } from "react-native";
import { interfaceLoginState, navigation } from "./type/type";
import { login } from "./config/api";
import { store } from "../App";
import { getLoginData } from "./action/action";
import { language } from "./language";
import Ionicons from 'react-native-vector-icons/Ionicons';
import FontAwesome from 'react-native-vector-icons/FontAwesome5';
import { Icon } from "native-base";
import CO_WarningWindow from "./component/CO_WarningWindow";
import { CO_Input } from "./component/CO_Input";
import { KeyboardAwareScrollView } from "react-native-keyboard-aware-scroll-view";
import CO_Button from "./component/CO_Button";
import Modal from 'react-native-modal';
import CO_Loading from "./component/CO_Loading";

export default class Login extends Component<navigation, interfaceLoginState >{
    constructor(props: any){
        console.log('login props => ', props)
        super(props);
        this.changeAccount = this.changeAccount.bind(this)
        this.changePassword = this.changePassword.bind(this)
        this.changeModalVisible = this.changeModalVisible.bind(this)
        this.loginbtnAction = this.loginbtnAction.bind(this)
        this.state={
            account: '',
            password: '',
            modalVisible: false,
            isLoading: false,
        }
    }

    private changeAccount( value:string ) {
        this.setState( {account:value} )
    }
    private changePassword( value:string ) {
        this.setState( { password:value} )
    }
    private changeModalVisible() {
        this.setState( {modalVisible: this.state.modalVisible ? false : true} )
    }
    public loginbtnAction() {
        //api to fetching data
        login(this.state)
        store.dispatch(getLoginData())
        this.props.navigation.navigate('Main', {title: 'Main'})
        this.changeModalVisible()
        console.log('=>>>>>>>>>>>>>>>>>', this.props.navigation.canGoBack(false))
    }
    componentDidUpdate() {
        console.log('Login state => ',this.state)
        console.log('store loginData => ', store.getState().loginData)
    }
    

    render():JSX.Element{
        
        return(
            <>
                {/* <CO_WarningWindow
                    styIdx="one"
                    navigation={this.props.navigation}
                    visible={this.state.modalVisible}
                    titleText={language.warning}
                    bodyText={language.loginOrNot}
                    btnText={language.confirm}
                    btnAction={this.changeModalVisible}
                /> */}
                <CO_Loading isVisible={this.state.isLoading} styIdx='one'/>
                <KeyboardAwareScrollView
                        resetScrollToCoords={{ x: 0, y: 0 }}
                        contentContainerStyle={styles.wrapper}
                        scrollEnabled={false}
                    >
                    <View> 
                        <Text>here is the login logo</Text>
                    </View>
                    <View style={styles.allInputWrapper}>
                        
                        <View style={styles.inputWrapper}>
                            <CO_Input
                                icons={
                                    <Icon
                                        color={'#C7C7E2'}
                                        size={10}
                                        name="user-alt"
                                        as={FontAwesome}
                                    />
                                }
                                styIdx='two'
                                change={this.changeAccount}  
                                placeholder={language.pleaseInputAccount}/>
                        </View>
                        <View style={styles.inputWrapper}>
                            <CO_Input
                                icons={
                                    <Icon
                                        color={'#C7C7E2'}
                                        size={10}
                                        name="lock"
                                        as={FontAwesome}
                                    />
                                }
                                styIdx='two'
                                change={this.changePassword}  
                                placeholder={language.pleaseInputPassword}/>
                        </View>
                    </View>

                    <CO_Button styIdx="one" btnText={language.login} btnAction={this.loginbtnAction}/>
                    
                </KeyboardAwareScrollView>
            </>
        )
    }
}

const styles=StyleSheet.create({
    wrapper:{
        width:'100%',
        backgroundColor:'rgba(0,100,100,0.3)',
        height:'100%',
        flexDirection:'column',
        alignItems:'center',
        justifyContent:'center'

    },
    allInputWrapper:{
        marginTop:130,
        height:200,
        width:'90%',
        flexDirection:'column',
        alignItems:'center',
        justifyContent:'space-evenly',
    },
    inputWrapper:{
        width:'100%',
    },
    
})
