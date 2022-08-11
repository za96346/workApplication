import { View, Image, Text } from "native-base";
import React from "react";
import { Icon } from "native-base";
import { shadowWrapper } from "../style/styles";
import { TouchableOpacity } from "react-native";
import { interfaceCO_FormUIProps } from "../type/type";

export default class CO_FormUI extends React.Component <interfaceCO_FormUIProps, any>{
    constructor(props: interfaceCO_FormUIProps) {
        super(props)
    }
    render(): React.ReactNode {
        return (
            <TouchableOpacity
                style={{
                    width: 150,
                    height: 150,
                    borderRadius: 30,
                    backgroundColor: 'rgba(250, 250, 250, 0.8)',
                    flexDirection: 'column',
                    alignItems: 'center',
                    justifyContent: 'space-around',
                    ...shadowWrapper()
                }}>
                <Image
                style={{
                    width: '60%',
                    height: '60%'
                }}
                source={this.props.imgUrl}
                />
                <Text 
                    style={{
                        fontSize: 20
                    }}>{this.props.Label}</Text>
                
            </TouchableOpacity>
        )
    }
}