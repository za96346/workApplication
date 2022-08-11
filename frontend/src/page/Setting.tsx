import { Text, View } from "native-base";
import React from "react";
import CO_NavigateList from "../component/CO_NavigateList";
import { Icon } from "native-base";
import lonicon from 'react-native-vector-icons/Ionicons'
import { language } from "../language";
import FontAwesome from "react-native-vector-icons/FontAwesome";
import { circle } from "../style/styles";
import AntDesign from "react-native-vector-icons/AntDesign"
class Setting extends React.Component <any, any>{
    constructor(props: any) {
        super(props)
        this.btnAction = this.btnAction.bind(this)
    }

    btnAction(name: string):void {
        //pass
    }
    render(): React.ReactNode {
        return(
            <View
                style={{
                    width: '100%',
                    height: '100%',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}>
                <View
                    style={{
                        width: '100%',
                        height: '40%',
                        backgroundColor: 'rgb(103, 129, 154)',
                        borderBottomRightRadius: 55,
                        flexDirection: 'column',
                        alignItems: 'center',
                        justifyContent: 'space-evenly'
                    }}>
                    <View
                        style={{
                            ...circle(120),
                            backgroundColor: '#333'
                        }}>

                    </View>

                    <View
                        style={{
                            flexDirection: 'row',
                            alignItems: 'center',
                        }}>
                        <Text
                            style={{
                                fontSize: 20,
                                color: '#eee',
                                marginRight: 10
                            }}>
                                NAME
                        </Text>
                        <Icon name="form" as={AntDesign} color={'#444'} size={5}/>
                    </View>
                </View>
                <View
                    style={{
                        width: '90%',
                        height: '60%',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}>
                    <CO_NavigateList
                        btnAction={() => this.btnAction(language.personalDataSetting)}
                        style={{ marginTop: 10 }}
                        icons={<Icon name="ios-person-circle-outline" as={lonicon} color={'#444'} size={10}/>}
                        name={language.personalDataSetting}/>
                    <CO_NavigateList
                        btnAction={() => this.btnAction(language.basicSetting)}
                        style={{ marginTop: 10 }}
                        icons={<Icon name="ios-settings-sharp" as={lonicon} color={'#444'} size={10}/>}
                        name={language.basicSetting}/>
                    <CO_NavigateList
                        btnAction={() => this.btnAction(language.personalDataSetting)}
                        style={{ marginTop: 10 }}
                        icons={<Icon name="phone" as={FontAwesome} color={'#444'} size={10}/>}
                        name={language.contactUs}/>
                    <CO_NavigateList
                        btnAction={() => this.btnAction(language.personalDataSetting)}
                        style={{ marginTop: 10 }}
                        icons={<Icon name="commenting" as={FontAwesome} color={'#444'} size={10}/>}
                        name={language.aboutUs}/>
                    <CO_NavigateList
                        btnAction={() => this.btnAction(language.logOut)}
                        style={{ marginTop: 10 }}
                        icons={<Icon name="ios-log-in-outline" as={lonicon} color={'#444'} size={10}/>}
                        name={language.logOut}/>
                </View>
            </View>
        )
    }
}
export default Setting;