import React,{ Component, Key, LegacyRef, ReactNode } from "react";
import { Animated, Easing, Image, ProgressViewIOSBase, TextInput, TouchableHighlight } from "react-native";
import { shadowWrapper } from "../style/styles";
import {  interfaceCO_InputProps, interfaceCO_InputState } from "../type/type";
import { Center, Icon, Text, View, ZStack } from 'native-base';
import { TouchableOpacity } from "react-native-gesture-handler";

type typeMainStyle = {
    one: {
        viewBgColor: string
        textColor: string
    },
    two: {
        viewBgColor: string,
        textColor: string
    },
}

const mainStyles: typeMainStyle = {
    one: {
        viewBgColor: 'rgba(255, 255, 255, 1)',
        textColor: 'rgba(255, 155, 255)'
    },
    two: {
        viewBgColor: 'rgba(255, 255, 255, 1)',
        textColor: 'rgba(100, 100, 100)'
    },
}

export class CO_Input extends Component <interfaceCO_InputProps, interfaceCO_InputState>{
    placeholder: string;
    change: Function;
    icons!: JSX.Element;
    private mainStyle: any;
    textInputRef: any;

    constructor(public props: interfaceCO_InputProps){
        super(props);
        //console.log(this.props.styIdx)
        this.mainStyle = mainStyles[this.props.styIdx as keyof typeof mainStyles]
        this.placeholder = props.placeholder
        this.textInputRef = React.createRef();
        this.change = props.change
        this.state = {
            focusEvent:false,
            focusStyle: {
                borderColor: '#C4E1FF',
                borderWidth: 1,
                borderStyle: 'solid',
                borderRadius: 10,
            },
            blurStyle: {
            },
            isPress: false,
        }
        this.onFocusPress = this.onFocusPress.bind(this)
    }

    componentDidUpdate(prevProps: interfaceCO_InputProps, prevState: interfaceCO_InputState) {
        if (prevState.isPress !== this.state.isPress && this.state.isPress) {this.textInputRef.current.focus()}
        console.log(prevProps, prevState)
    }

    onFocusPress() {
        //console.log(this.textInputRef)
        this.setState({isPress: true})
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
                    ...(this.state.focusEvent ? this.state.focusStyle : this.state.blurStyle),
                    ...shadowWrapper()
            }}>
                <View 
                    style={{
                        display: 'flex',
                        flexDirection:'row',
                        alignItems: 'center',
                        justifyContent: 'space-around',
                        position: 'relative',
                        backgroundColor: this.mainStyle!.viewBgColor,
                        borderRadius: 10,
                        height: 80,
                        width: '100%'
                    }}>
                    <View style={{width: '15%'}}>
                        {
                            this.props.icons
                        }
                    </View>
                        
                    <View style={{width: '80%', position: 'relative'}}>
                        <Text 
                            style={{
                                position: 'absolute', 
                                zIndex: 3, 
                                fontSize: 20,
                                top: this.state.isPress ? 5 : 0,
                                color: this.mainStyle!.textColor,
                                opacity: this.state.isPress ? 1 : 0.5
                            }}>
                                {this.props.placeholder}
                        </Text>
                        {this.state.isPress && (<TextInput
                            ref={this.textInputRef}
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
                                marginTop: 10,
                            }}/>)}
                    </View>
                </View>
            </TouchableOpacity>
        
        )
    }
}