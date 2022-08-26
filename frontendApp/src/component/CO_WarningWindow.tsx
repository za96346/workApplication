import React from "react";
import { Modal, View, Text, Touchable, TouchableOpacity } from "react-native";
import { interfaceCO_WarningWindowProps } from "../type/type";
import CO_Button from "./CO_Button";

export default class CO_WarningWindow extends React.Component <interfaceCO_WarningWindowProps, any>{

    constructor(public props: interfaceCO_WarningWindowProps) {
        super(props)
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
                        
                        <CO_Button btnText={this.props.bodyText} btnAction={this.props.navigation} styIdx="one"/>

                    </View>
                </View>
            </Modal>
        )
    }
}