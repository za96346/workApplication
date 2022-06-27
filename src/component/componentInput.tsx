import React,{ Component, ReactNode } from "react";
import { TextInput } from "react-native";
import {  interfaceComponentInput } from "../type/type";
export class ComponentInput extends Component <interfaceComponentInput,any>{
    placeholder: string;
    styles: any;


    constructor( prop:interfaceComponentInput){
        super(prop);
        this.placeholder=prop.placeholder
        this.styles=prop.styles
        this.state={
            focusEvent:false
        }
    }
    render():JSX.Element {
        console.log(this.placeholder,this.styles)
        return(
            <>
                <TextInput
                    onFocus={()=>this.setState({focusEvent:true})}
                    onBlur={()=>this.setState({focusEvent:false})}
                    style={{...this.styles,backgroundColor:this.state.focusEvent?'#fff':this.styles.backgroundColor}}
                    placeholder={this.placeholder}/>
                
            </>
        
        )
    }
}