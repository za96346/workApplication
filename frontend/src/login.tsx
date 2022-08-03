import React,{Component} from "react";
import { ActivityIndicator, Animated, Button, Modal, ProgressViewIOSBase, Text, TextInput, Touchable, TouchableHighlight, View } from "react-native";
import { ComponentInput } from "./component/componentInput";
import { styleLogin,shadowWrapper } from "./style/styles";
import { StyleSheet } from "react-native";
import { interfaceLoginState, typeNavigation } from "./type/type";
import ComponentModal from "./component/componentModal";
import { login } from "./config/api";
import { store } from "../App";
import { getLoginData } from "./action/action";
import { language } from "./language";
import Ionicons from 'react-native-vector-icons/Ionicons';
import FontAwesome from 'react-native-vector-icons/FontAwesome5';
import { Icon } from "native-base";


export default class Login extends Component< typeNavigation, interfaceLoginState >{
    navigation: typeNavigation;
    constructor(props:any){
        super(props);
        this.navigation = props.navigation
        this.changeAccount = this.changeAccount.bind(this)
        this.changePassword = this.changePassword.bind(this)
        this.changeModalVisible = this.changeModalVisible.bind(this)
        this.state={
            account: '',
            password: '',
            modalVisible: false,
        }
    }

    public changeAccount( value:string ) {
        this.setState( {account:value} )
    }
    public changePassword( value:string ) {
        this.setState( { password:value} )
    }
    public changeModalVisible() {
        this.setState( {modalVisible: this.state.modalVisible?false:true} )
    }
    public componentDidUpdate() {
        console.log('Login state => ',this.state)
        console.log('store loginData => ', store.getState().loginData)
    }
    

    render():JSX.Element{
        
        return(
            <>
            
                <ComponentModal 
                    navigation={this.navigation}
                    visible={this.state.modalVisible}
                    titleText={language.warning}
                    bodyText={language.loginOrNot}
                    btnText={language.confirm}
                    change={this.changeModalVisible}
                    />
                <View style={styles.wrapper}>
                    <ActivityIndicator size="small" color="#0000ff"/>
                    <View> 
                        <Text>here is the login logo</Text>
                    </View>
                    <View style={styles.allInputWrapper}>
                        
                        <View style={styles.inputWrapper}>
                            <ComponentInput 
                                icons={
                                    <Icon
                                        style={{color: '#C7C7E2'}}
                                        size={10}
                                        name="user-alt"
                                        as={FontAwesome}
                                    />
                                }
                                change={this.changeAccount}  
                                placeholder={language.pleaseInputAccount}/>
                        </View>
                        <View style={styles.inputWrapper}>
                            <ComponentInput
                                icons={
                                    <Icon
                                        style={{color: '#C7C7E2'}}
                                        size={10}
                                        name="lock"
                                        as={FontAwesome}
                                    />
                                }
                                change={this.changePassword}  
                                placeholder={language.pleaseInputPassword}/>
                        </View>
                    </View>
                    <TouchableHighlight
                        activeOpacity={0.6}
                        underlayColor="#DDDDDD"
                        
                        style={{
                            ...styles.buttonWrapper,
                            ...shadowWrapper('#000', { width: 10, height: 10 }, 0.5, 10)
                        }}>
                        <Button
                            title={language.login} 
                            onPress={() => {
                                //api to fetching data
                                login(this.state)
                                store.dispatch(getLoginData())
                                this.navigation.navigate('Main', {params:'',navigation:this.navigation})
                                this.changeModalVisible()
                            }}/>
                    </TouchableHighlight>
                    
                </View>
            
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
    buttonWrapper:{
        width:'90%',
        backgroundColor:'#ccc',
        borderRadius:10,

    }
    
})
