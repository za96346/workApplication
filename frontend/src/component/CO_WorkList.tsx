import { Hidden, Icon } from "native-base";
import React from "react";
import { TouchableOpacity, View, Text, ScrollView, FlatList, Image } from "react-native";
import FontAwesome from 'react-native-vector-icons/FontAwesome5';
import { shadowWrapper } from "../style/styles";
import { interfaceCO_WorkListProps } from "../type/type";

type typeMainStyle = {
    one: {
        nameStyle: object,
        todayStyle: object,
        todayLabelStyle: object,
    },
}

const mainStyles: typeMainStyle = {
    one: {
        nameStyle: {
            fontSize: 25,
            fontWight: 1000,
            TextAlign: 'center',
            color: '#000',
        },
        todayStyle: {
            fontSize: 15,
            TextAlign: 'center',
            color: '#00E3E3'
        },
        todayLabelStyle: {
            fontSize: 17,
            TextAlign: 'center',
            color: '#999'
        }
    },
}

export default class CO_WorkList extends React.Component<interfaceCO_WorkListProps, any> {
    private mainStyle: any;
    constructor(props: any) {
        super(props)
        this.mainStyle = mainStyles[this.props.styIdx as keyof typeof mainStyles] || 'one'
    }
    render(): JSX.Element {
        return(
                <View
                    style={{
                        width: '100%'
                    }}
                    >
                    <FlatList
                        ListHeaderComponent={
                            
                            <View 
                                style={{
                                    height: 50,
                                    backgroundColor: '#95CACA',
                                    flexDirection: 'row',
                                    alignItems: 'center',
                                    justifyContent: 'center',
                                    opacity: 0.5
                                }}>
                                <Text style={{fontSize: 20, color: '#333'}}>今日班表</Text>
                            </View>
                        }
                        ListFooterComponent={<View style={{ height: 100 }}></View>}
                        ListEmptyComponent={<Text style={{textAlign: 'center'}}>none solution</Text>}
                        data={this.props.data}
                        renderItem={({ item, index, separators }) => (
                            <TouchableOpacity
                                style={{
                                    width: '100%',
                                    height: 100,
                                    marginVertical: 10,
                                    display: 'flex',
                                    flexDirection:'row',
                                    alignItems: 'center',
                                    justifyContent: 'flex-start',
                                    backgroundColor: 'white',
                                    ...shadowWrapper(),
                                    borderRadius: 10,
                                    paddingHorizontal: '4%'
                                }}
                                activeOpacity={0.2}
                                key={index}
                            //   onPress={() => this.onPress(item)}
                                >
                                    <View style={{
                                        alignItems: 'center',
                                        justifyContent: 'center',
                                        height: 80,
                                        width: 80 }}>
                                        {/* <Icon
                                            color={'#C7C7E2'}
                                            size={10}
                                            name="lock"
                                            as={FontAwesome}
                                        /> */}
                                        <Image
                                            style={{
                                                width: 80,
                                                height: 80,
                                                borderRadius: 80,
                                               
                                            }}
                                            fadeDuration={400}
                                            source={require('../assert/dog.jpg')}
                                            resizeMode='contain'
                                            />
                                    </View>
                                    <View
                                        style={{
                                            width: '78%',
                                            height: '100%',
                                            paddingLeft: 30,
                                            flexDirection: 'column',
                                            justifyContent: 'space-evenly',
                                            alignItems: 'flex-start',
                                        }}>
                                        <Text style={this.mainStyle.nameStyle}>{item.userName}</Text>
                                        <View style={{flexDirection: 'row', alignContent: 'center', justifyContent: 'center'}}>
                                            <Text style={this.mainStyle.todayLabelStyle}>今日班次:    </Text>
                                            <Text style={[this.mainStyle.todayStyle, {marginRight: 10}]}>{item.todayWorkType}</Text>
                                            <Text style={this.mainStyle.todayStyle}>{item.todayWorkTime}</Text>
                                        </View>
                                    </View>
                            </TouchableOpacity>
                        )}
                    />
                </View>
        )
    }
}