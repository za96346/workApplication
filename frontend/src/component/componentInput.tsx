import React,{ Component, ReactNode } from "react";
import { Image, ProgressViewIOSBase, TextInput } from "react-native";
import {  interfaceComponentInput } from "../type/type";
export class ComponentInput extends Component <interfaceComponentInput,any>{
    placeholder: string;
    styles: any;
    propSetState:React.Dispatch<React.SetStateAction<object>>


    constructor(private prop:interfaceComponentInput){
        super(prop);
        this.placeholder=prop.placeholder
        this.styles=prop.styles
        this.propSetState=this.setState
        this.state={
            focusEvent:false
        }
    }
    render():JSX.Element {
        console.log()
        return(
            <>
                <TextInput
                    onChange={(e)=>{
                        this.propSetState((prev)=>({...prev,account:e.currentTarget.valueOf}))
                    }}
                    
                    onFocus={()=>this.setState({focusEvent:true})}
                    onBlur={()=>this.setState({focusEvent:false})}
                    style={{...this.styles,backgroundColor:this.state.focusEvent?'#fff':this.styles.backgroundColor}}
                    placeholder={this.placeholder}/>
                <Image source={require('')}></Image>
            </>
        
        )
    }
}