import { Icon, Text, View } from "native-base";
import React from "react";
import { FlatList, ScrollView, TouchableOpacity } from "react-native";
import MapView, { Callout, Marker } from "react-native-maps";
import materialCommunityIcons from 'react-native-vector-icons/MaterialCommunityIcons'
import Foundation from 'react-native-vector-icons/Foundation'

const data = new Array(100).fill({
    checkInWork: '9 : 00',
    checkOutWork: Math.floor(Math.random() * 2) === 1 ? '17 : 30' : ''
})
export default class CheckIn extends React.Component <any, any>{
    constructor(props: any) {
        super(props)
        this.state = {
            markers: [
                {
                    latitude : 24.11879264258466,
                    longitude : 120.65799889550874
                },
                {
                    latitude : 24.11879264258466,
                    longitude : 120.65799889559874
                }
            ]
        }
    }
    render(): React.ReactNode {
        return(
            <View style={{
                width: '100%',
                height: '100%',
                flexDirection: 'column',
                alignItems: 'center',
                justifyContent: 'flex-start'}}>
                <MapView          
                    initialRegion={{
                        latitude: 24.11879264258466,
                        longitude: 120.65799889550874,
                        latitudeDelta: 0.0922,
                        longitudeDelta: 0.0421,
                    }}
                    style={{
                        position: 'relative',
                        height: '40%',
                        width: '100%',
                    }}>
                    {this.state.markers.map((item: any, index: any) => (
                        <Marker
                            draggable={false}
                            key={index}
                            coordinate={item}
                            // title={'你的位置'}
                            description={item.description}>
                            
                            <Icon 
                                name='map-marker'
                                as={materialCommunityIcons}
                                color={'#f00'}
                                size={20}
                                />
                            <Callout
                                tooltip={true}
                                >
                                <View 
                                    style={{
                                        height: 20,
                                        width: 40,
                                        backgroundColor: 'white'
                                    }}>
                                    <Text>你的位置</Text>
                                </View>
                            </Callout>
                            
                        </Marker>
                    ))}
                </MapView>
                <View
                    style={{
                        position: 'absolute',
                        top: 200,
                        width: '90%',
                        height: '20%',
                        backgroundColor: 'rgba(255, 255, 255, 0.7)',
                        flexDirection: 'column',
                        alignItems: 'flex-start',
                        justifyContent: 'space-around',
                        paddingHorizontal: 20,
                        borderRadius: 10,
                    }}>
                        <View>
                            <Text style={{fontSize: 20}}>07月16 星期日</Text>
                        </View>
                        <View>
                            <Text>上班卡  09 : 30</Text>
                        </View>
                        <View>
                            <Text>下班卡  - -</Text>
                        </View>
                        <View
                            style={{
                                position: 'absolute',
                                right:20,
                                top: '20%',
                                width: 120,
                                height: 120,
                                borderRadius: 120,
                                borderColor: '#5CADAD',
                                borderWidth: 2,

                            }}>
                            <View
                                style={{
                                    position: 'absolute',
                                    top: 8,
                                    right: 8,
                                    width: 100,
                                    height: 100,
                                    borderRadius: 100,
                                    backgroundColor: '#80FFFF'
                                }}>
                                    <Text style={{position: 'absolute', top: '40%', left: '30%', fontSize: 20,}}>
                                        打卡
                                    </Text>
                                
                            </View>   
                        </View>
                </View>
                <View
                    style={{
                        marginTop: '15%',
                        width: '90%',
                        height: '40%',
                        backgroundColor: 'white',
                        borderRadius: 10,
                        flexDirection: 'column',
                    }}>
                        <FlatList
                            ListHeaderComponent={
                                <View style={{                     
                                        height: 60,
                                        width: '90%',
                                        flexDirection: 'row',
                                        alignItems: 'center',
                                        justifyContent: 'center',

                                    }}>
                                    <Icon name="clipboard-pencil" as={Foundation} color={'#1AFD9C'}/>
                                    <Text style={{ color: '#1AFD9C', fontSize: 20 }}>打卡狀態</Text>
                                </View>
                            }
                            ListFooterComponent={<View style={{ height: 100 }}></View>}
                            data={data}
                            renderItem={({ item, index, separators }) => {
                                return(
                                    <TouchableOpacity
                                        key={index}
                                        style={{
                                            width: '85%',
                                            height: 100,
                                            flexDirection: 'row',
                                            alignItems: 'center',
                                            justifyContent: 'space-around',
                                            marginBottom: 10,
                                            borderColor:
                                                item?.checkOutWork?.length === 0
                                                || item?.checkInWork?.length === 0
                                                    ? 'red'
                                                    : '#66B3FF',
                                            borderWidth: 2,
                                            borderRadius: 10,
                                            alignSelf: 'center'
                                        }}>
                                        <View>
                                            <Text>上班卡  {item.checkInWork}</Text>
                                        </View>
                                        <View>
                                            <Text>下班卡  {item.checkOutWork.length === 0 ? '- -' : item.checkOutWork}</Text>
                                        </View>
                                    </TouchableOpacity>
                                )
                            }}
                        />
                </View>
            </View>

        )
    }
}