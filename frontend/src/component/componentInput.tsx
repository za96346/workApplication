import React,{ Component, ReactNode } from "react";
import { Animated, Easing, Image, ProgressViewIOSBase, TextInput, TouchableHighlight } from "react-native";
import { shadowWrapper } from "../style/styles";
import {  interfaceComponentInputProps, interfaceComponentInputState } from "../type/type";
import { Center, Icon, Text, View, ZStack } from 'native-base';
import { TouchableOpacity } from "react-native-gesture-handler";

export class ComponentInput extends Component <interfaceComponentInputProps, interfaceComponentInputState>{
    placeholder: string;
    change: Function;
    icons!: JSX.Element;
    private mainColor: string;

    constructor(public prop: interfaceComponentInputProps){
        super(prop);
        this.mainColor = 'rgba(255, 255, 255, 1)'
        this.placeholder = prop.placeholder

        this.change=prop.change
        this.state = {
            focusEvent:false,
            focusStyle:{
                backgroundColor:'#rgba(169, 169, 169, 1)',
            },
            blurStyle:{
            }
        }

    }

    onFocusPress() {
        
    }


    render():JSX.Element {
        return(
            <TouchableOpacity
                activeOpacity={0.6}
                onPress={this.onFocusPress}
                style={{
                    display: 'flex',
                    alignItems: 'center',
                    justifyContent: 'flex-start',
                    width: '100%',
                    borderRadius: 10,
                    zIndex: 3,
                    ...shadowWrapper('#000',{width:10,height:10},0.5,10)
            }}>
                <View 
                    style={{
                        display: 'flex',
                        flexDirection:'row',
                        alignItems: 'center',
                        justifyContent: 'space-around',
                        position: 'relative',
                        backgroundColor: this.mainColor,
                        borderRadius: 10,
                        width: '100%'
                    }}>
                    <View style={{width: '15%'}}>
                        {
                            this.props.icons
                        }
                    </View>
                        
                    <View style={{width: '80%', position: 'relative'}}>
                        <Text style={{position: 'absolute', zIndex:3, fontSize: 20, top: 5}}>{this.props.placeholder}</Text>
                        <TextInput
                            onChangeText={(value) => {
                                this.change(value)
                            }}
                            onFocus={() => this.setState({focusEvent: true})}
                            onBlur={() => this.setState({focusEvent: false})}
                            style={{
                                width: '100%',
                                height: 60,
                                borderRadius: 10,
                                fontSize:20,
                                marginTop: 10
                            }}/>
                    </View>
                </View>
            </TouchableOpacity>
        
        )
    }
}