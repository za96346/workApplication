import React,{ Component, ReactNode } from "react";
import { Animated, Easing, Image, ProgressViewIOSBase, TextInput } from "react-native";
import { shadowWrapper } from "../style/styles";
import {  interfaceComponentInput } from "../type/type";
export class ComponentInput extends Component <interfaceComponentInput,any>{
    placeholder: string;
    styles: any;
    change:Function;
    img:NodeRequire


    constructor(public prop:interfaceComponentInput){
        super(prop);
        this.placeholder = prop.placeholder
        this.styles = { ...prop.styles,position:'relative' }
        this.img = prop.require
        this.change=prop.change
        this.state = {
            focusEvent:false,
            focusStyle:{
                backgroundColor:'#f8f',
                ...shadowWrapper('#fff',{width:5,height:5},0.5,5)
            },
            blurStyle:{
                ...shadowWrapper('#000',{width:5,height:5},0.5,5)
            }
        }

    }


    render():JSX.Element {
        return(
            <>
                <TextInput
                    onChangeText={(value)=>{
                        this.change(value)  
                    }}
                    onFocus={() => this.setState({focusEvent:true})}
                    onBlur={() => this.setState({focusEvent:false})}
                    style={{
                    ...this.styles,
                    ...this.state.focusEvent
                        ?this.state.focusStyle
                        :this.state.blurStyle
                    }}
                    placeholder={this.placeholder}/>
                <Image
                    style={{width:40,
                        height: '90%',
                        position: 'absolute',
                        top: (this.styles.height-(this.styles.height*0.9))/2,
                        left: 10
                    }} 
                    source={this.img}></Image>
            </>
        
        )
    }
}