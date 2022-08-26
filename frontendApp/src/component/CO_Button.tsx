import { View } from "native-base";
import React, { Component } from "react";
import { Button, TouchableOpacity } from "react-native";
import { typeCO_ButtonProps } from "../type/type";

type typeMainStyles = {
    one: {
        touchStyle: any,
        btnColor: string
    }
}

const mainStyles: typeMainStyles = {
    one: {
        touchStyle: {
            backgroundColor: '#00BBFF',
            width: 100,
            height: 40,
            borderRadius: 10
        },
        btnColor: '#c9f4e3',

    }
}

export default class CO_Button extends Component<typeCO_ButtonProps, any>{
    private mainStyle: any
    
    
    constructor(public props:typeCO_ButtonProps ) {
        super(props)
        this.mainStyle = mainStyles[this.props.styIdx as keyof typeof mainStyles] || 'one'
        this.state = {

        }
    }

    render(): JSX.Element {
        return(
            <View>
                <TouchableOpacity
                    activeOpacity={0.1}
                    style={{...this.mainStyle.touchStyle}}>
                    <Button 
                        onPress={() => this.props.btnAction()}
                        color={this.mainStyle.btnColor}
                        title={this.props.btnText}/>
                </TouchableOpacity>
            </View>
        )
    }
}