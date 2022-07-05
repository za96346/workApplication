import React from "react";
import { Text, Touchable } from "react-native"; 
import { TouchableOpacity } from "react-native-gesture-handler";

export default class Main extends React.Component{
    navigation: any
    constructor(props:any) {
        super(props)
        this.navigation = props.navigation
    }
    render(): JSX.Element {
        return (
            <>
                <TouchableOpacity onPress={()=>this.navigation.goBack()}>
                    <Text >Main page</Text>
                </TouchableOpacity>
            </>
        );
    }
}