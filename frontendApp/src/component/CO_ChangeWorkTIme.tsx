import { View, Text, TouchableOpacity } from "react-native";
import React from "react";

export default class CO_ChangeWorkTime extends React.Component <any, any> {
    constructor(props: any) {
        super(props)
    }
    render(): JSX.Element {
        return(
            <View
                style={{
                    width: '100%',
                    height: 200,
                    backgroundColor: 'white',
                    flexDirection: 'column',
                }}>
                <View
                    style={{
                        width: '100%',
                        height: '15%',
                        backgroundColor: 'rgba(160, 160, 160, 0.4)',
                    }}>
                </View>
                <View
                    style={{
                        width: '100%',
                        height: '85%',
                        flexDirection: 'row',
                        alignItems: 'center',
                    }}>
                        <View
                            style={{
                                width: '20%',
                                height: '100%',
                                alignItems: 'center',
                                justifyContent: 'center'
                            }}>
                            <Text>25 {'\n'} 週二</Text>

                        </View>
                        <View
                            style={{
                                width: '60%',
                                height: '100%',
                                alignItems: 'center',
                                justifyContent: 'center'
                            }}>
                                <Text>9 : 00 - 18 : 00</Text>

                        </View>
                        <View
                            style={{
                                width: '20%',
                                height: '100%',
                                alignItems: 'center',
                                justifyContent: 'center'
                            }}>
                            <TouchableOpacity>
                                <Text>同意</Text>
                            </TouchableOpacity>
                        </View>

                </View>
            </View>
        )
    }
}
