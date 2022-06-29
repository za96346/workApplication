import React from "react";
import { Modal, View, Text, Button } from "react-native";
import { interfaceComponentModal } from "../type/type";

export default class ComponentModal extends React.Component <interfaceComponentModal,any>{
    visible:boolean
    titleText:string
    bodyText:string
    btnText:string
    change:Function

    constructor(public prop: interfaceComponentModal) {
        super(prop)
        this.visible=prop.visible
        this.titleText=prop.titleText
        this.bodyText=prop.bodyText
        this.btnText=prop.btnText
        this.change=prop.change
        this.state={
            visible:false,
        }
        this.changeVisible=this.changeVisible.bind(this)
    }
    changeVisible(prevProp:boolean) {
        if(prevProp!==this.state.visible){
            console.log('to setState =====> ')
            this.setState({visible:prevProp})
        }
    }
    componentDidUpdate(prevProp:any,prevState:any){
        //prevProp當前傳入的prop(新的)
        //this.prop.visible 原本的prop(舊的)
        if(prevProp.visible!==this.visible){
            console.log('\n')
            console.log('\nprevProp.visible => ',prevProp.visible,'\n','this.state.visible => ',this.prop.visible)
            this.changeVisible(prevProp)
        }
    }
    
    render(): JSX.Element {
        return(
            <Modal 
                visible={this.state.visible}
                transparent={true}
                >
                <View style={{width:'100%',height:'100%',alignItems:'center',justifyContent:'center'}}>
                    <Text>{this.titleText}</Text>
                    <Text>{this.bodyText}</Text>
                    <Button title={this.btnText} onPress={()=>this.change(false)}/>
                </View>
            </Modal>
        )
    }
}