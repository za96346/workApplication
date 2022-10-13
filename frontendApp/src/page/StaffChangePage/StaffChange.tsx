import { Text, View } from "native-base";
import React from "react";
import CO_FormUI from "../../component/CO_FormUI";
import { language } from "../../language";
import { interfaceCO_FormUIProps } from "../../type/type";

const form: Array < Array < interfaceCO_FormUIProps> > = [
    [
        {
            styIdex: 'one',
            Label: language.applyChangeWork,
            imgUrl: require('../../assert/formChange.png'),
            btnAction: () => {},
        },
        {
            styIdex: 'one',
            Label: language.applyUpWork,
            imgUrl: require('../../assert/workUp.png'),
            btnAction: () => {},
        },
    ],
    [
        {
            styIdex: 'one',
            Label: language.forgetCheckIn,
            imgUrl: require('../../assert/forget.png'),
            btnAction: () => {},
        },
        {
            styIdex: 'one',
            Label: language.late,
            imgUrl: require('../../assert/late.png'),
            btnAction: () => {},
        },
    ],
    [
        {
            styIdex: 'one',
            Label: language.applyRequestLeave,
            imgUrl: require('../../assert/sick.png'),
            btnAction: () => {},
        },
        {
            styIdex: 'one',
            Label: language.elseOnWork,
            imgUrl: require('../../assert/else.png'),
            btnAction: () => {},
        },
    ]
]

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
                        {
                            form.map((item, index) => {
                                return (
                                    <View 
                                        style={{
                                            flexDirection: 'row',
                                            alignItems: 'center',
                                            justifyContent: 'space-around',
                                            width: '100%',
                                            marginTop: 20
                                        }}>
                                        <CO_FormUI
                                            styIdex={item[0].styIdex}
                                            btnAction={item[0].btnAction}
                                            Label={item[0].Label}
                                            imgUrl={item[0].imgUrl} />
                                        <CO_FormUI
                                            styIdex={item[1].styIdex}
                                            btnAction={item[1].btnAction}
                                            Label={item[1].Label}
                                            imgUrl={item[1].imgUrl} />
                                    </View>
                                )
                            })
                        }
                   
                </View>
        )
    }
}
export default StaffChange;