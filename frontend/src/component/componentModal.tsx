import React from "react";
import { Modal, View, Text, Button, Touchable, TouchableOpacity } from "react-native";
import { interfaceComponentModalProps, typeNavigation } from "../type/type";
import ComponentButton from "./componentButton";

export default class ComponentModal extends React.Component <interfaceComponentModalProps, any>{
    visible:boolean
    titleText:string
    bodyText:string
    btnText:string
    change:Function
    navigation: typeNavigation

    constructor(public props: interfaceComponentModalProps) {
        super(props)
        this.navigation = props.navigation
        this.visible=props.visible
        this.titleText=props.titleText
        this.bodyText=props.bodyText
        this.btnText=props.btnText
        this.change=props.change
    }

    
    render(): JSX.Element {
        return(
            <Modal 
                visible={this.props.visible}
                transparent={true}
                >
                <View style={{width: '100%', height: '100%', alignItems: "center", justifyContent: 'center', flexDirection: 'column'}}>
                    <View 
                        style={{
                            borderRadius: 20,
                            alignItems: 'center',
                            width: '70%', 
                            height: '40%',
                            backgroundColor: 'white'
                        }}>
                        <View style={{width: '100%', height: '20%', justifyContent: 'center', backgroundColor: 'blue', borderTopLeftRadius: 20, borderTopRightRadius :20}}>
                            <Text style={{color: 'white', textAlign: 'center'}}>{this.props.titleText}</Text>
                        </View>
                        <View style={{width: '100%', height: '60%', justifyContent: 'center', backgroundColor: 'white'}}>
                            <Text style={{textAlign: 'center'}}>{this.props.bodyText}</Text>
                        </View>
                        
                        <ComponentButton change={this.props.change} btnText={this.props.btnText}/>

                    </View>
                </View>
            </Modal>
        )
    }
}