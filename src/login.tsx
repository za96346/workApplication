import React,{Component} from "react";
import { TextInput, View } from "react-native";
import { ComponentInput } from "./component/componentInput";
import { styleLogin } from "./style/styles";
import { StyleSheet } from "react-native";
export default class Login extends Component{
    constructor(props:any){
        super(props);
    }
    render():JSX.Element{
        return(
            <View style={styles.wrapper}>
                <View style={styles.inputWrapper}>
                    <ComponentInput styles={styleLogin.Input} placeholder={'請輸入帳號'}/>
                    <ComponentInput styles={styleLogin.Input} placeholder={'請輸入密碼'}/>
                </View>
            </View>
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
        justifyContent:'flex-start'

    },
    inputWrapper:{
        marginTop:130,
        height:200,
        width:'90%',
        flexDirection:'column',
        alignItems:'center',
        justifyContent:'space-evenly',
    }
})
