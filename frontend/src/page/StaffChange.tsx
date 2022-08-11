import { Text, View } from "native-base";
import React from "react";
import CO_FormUI from "../component/CO_FormUI";
import { language } from "../language";
class StaffChange extends React.Component <any, any>{
    constructor(props: any) {
        super(props)
    }

    render(): React.ReactNode {
        return(
                <View 
                    style={{
                        width: '100%',
                        height: '100%',
                        flexDirection: 'column',
                        alignItems: 'center',
                        justifyContent: 'flex-start',
                    }}>
                        <View 
                            style={{
                                flexDirection: 'row',
                                alignItems: 'center',
                                justifyContent: 'space-around',
                                width: '100%',
                                marginTop: 20
                            }}>
                            <CO_FormUI Label={language.applyChangeWork} imgUrl={require('../assert/formChange.png')} />
                            <CO_FormUI Label={language.applyUpWork} imgUrl={require('../assert/workUp.png')} />
                        </View>
                        <View 
                            style={{
                                flexDirection: 'row',
                                alignItems: 'center',
                                justifyContent: 'space-around',
                                width: '100%',
                                marginTop: 20
                            }}>
                            <CO_FormUI Label={language.forgetCheckIn} imgUrl={require('../assert/forget.png')} />
                            <CO_FormUI Label={language.late} imgUrl={require('../assert/late.png')} />
                        </View>
                        <View 
                            style={{
                                flexDirection: 'row',
                                alignItems: 'center',
                                justifyContent: 'space-around',
                                width: '100%',
                                marginTop: 20
                            }}>
                            <CO_FormUI Label={language.applyRequestLeave} imgUrl={require('../assert/sick.png')} />
                            <CO_FormUI Label={language.elseOnWork} imgUrl={require('../assert/else.png')} />
                        </View>
                        
                   
                </View>
        )
    }
}
export default StaffChange;