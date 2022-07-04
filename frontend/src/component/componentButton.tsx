import React, { Component } from "react";
import { Button, TouchableOpacity } from "react-native";
import { typeComponentButtonProps } from "../type/type";
export default class ComponentButton extends Component<typeComponentButtonProps, any>{
    btnText: string;
    change: Function
    
    
    constructor(public props:typeComponentButtonProps ) {
        super(props)
        this.btnText= props.btnText
        this.change=props.change

    }
    render(): JSX.Element {
        return(
            <TouchableOpacity style={{backgroundColor: '#00BBFF', width: 100, height: 40, borderRadius: 10}}>
                <Button 
                    title={this.props.btnText}
                    onPress={() => {
                        this.change(false)
                    }}
                />
            </TouchableOpacity>
        
        )
    }
}