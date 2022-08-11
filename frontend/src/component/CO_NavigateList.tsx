import { Icon } from "native-base";
import React from "react";
import { TouchableOpacity, Text, View } from "react-native";
import { interfaceCO_NavigateListProps } from "../type/type";
import Octicons from 'react-native-vector-icons/Octicons'
export default class CO_NavigateList extends React.Component <interfaceCO_NavigateListProps, any>{
    constructor(props: any) {
        super(props)
    }
    render(): JSX.Element {
        return(
            <TouchableOpacity
                onPress={this.props.btnAction}
                style={{
                    width: '100%',
                    height: 60,
                    flexDirection: 'row',
                    alignItems: 'center',
                    justifyContent: 'space-between',
                    backgroundColor: 'white',
                    paddingHorizontal: 20,
                    borderRadius: 30,
                    paddingLeft: 30,
                    ...this.props.style
                }}>
                <View 
                    style={{
                        flexDirection: 'row',
                        alignItems: 'center',
                        width: '40%',
                    }}>
                    {
                        this.props.icons
                    }
                    <Text
                        style={{
                            marginLeft: 10,
                            color: 'rgb(103, 129, 154)'
                        }}>
                        {this.props.name}
                    </Text>
                </View>
                {
                    <Icon name="triangle-right" as={Octicons} color={'rgb(103, 129, 154)'} size={10}/>
                }
            </TouchableOpacity>
        )
    }
}