import React from "react";
import { Text, Touchable } from "react-native"; 
import { TouchableOpacity } from "react-native-gesture-handler";
import {Linking} from 'react-native';

export default class Main extends React.Component{
    navigation: any
    constructor(props:any) {
        super(props)
        this.navigation = props.navigation
    }
    render(): JSX.Element {
        return (
            <>
                <TouchableOpacity onPress={()=>{
                    Linking.openURL(`tel:${+886906930873}`)
                    this.navigation.goBack()}}>
                    <Text >Main page</Text>
                </TouchableOpacity>
            </>
        );
    }
}