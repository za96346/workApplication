import React,{Component} from "react";
import { ActivityIndicator, Animated, Button, Text, TextInput, Touchable, TouchableHighlight, View } from "react-native";
import { ComponentInput } from "./component/componentInput";
import { styleLogin,shadowWrapper } from "./style/styles";
import { StyleSheet } from "react-native";
export default class Login extends Component{
    constructor(props:any){
        super(props);
        this.state={
            account: '',
            password: ''
        }
        this.changeAccount=this.changeAccount.bind(this)
        this.changePassword=this.changePassword.bind(this)
    }
    changeAccount( value:string ) {
        this.setState( (prev) => ({...prev, account:value}) )
    }
    changePassword( value:string ) {
        this.setState( (prev) => ({...prev, password:value}) )
    }
    componentDidUpdate() {
        console.log('sideEffect',this.state)
    }

    render():JSX.Element{
        console.log(this.state)
        return(
            <>  
            <View style={styles.wrapper}>
                <ActivityIndicator size="small" color="#0000ff"/>
                <View>
                    <Text>here is the login logo</Text>
                </View>
                <View style={styles.allInputWrapper}>
                    
                    <View style={styles.inputWrapper}>
                        <ComponentInput require={require('./static/account.png')} change={this.changeAccount} styles={styleLogin.Input} placeholder={'請輸入帳號'}/>
                    </View>
                    <View style={styles.inputWrapper}>
                        <ComponentInput require={require('./static/password.png')} change={this.changePassword} styles={styleLogin.Input} placeholder={'請輸入密碼'}/>
                    </View>
                </View>
                <TouchableHighlight
                    activeOpacity={0.6}
                    underlayColor="#DDDDDD"
                    style={{
                    ...styles.buttonWrapper,
                    ...shadowWrapper('#000',{width:10,height:10},0.5,10)
                    }}>
                    <Button title='登入' onPress={()=>console.log('on press')}/>
                </TouchableHighlight>
                
            </View>
            </>
        )
    }
}

const styles=StyleSheet.create({
    wrapper:{
        marginTop:40,
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
