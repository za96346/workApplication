import { ScrollView, Text, View } from "native-base";
import React from "react";
import { SafeAreaView } from "react-native";
import CO_WorkList from "../component/CO_WorkList";
const data = new Array(100).fill({
        imgUrl: '',
        userName: 'jack',
        todayWorkType: '早班',
        todayWorkTime: '9 : 00 ~ 17 : 30'
    })

export default class Home extends React.Component <any, any>{
    constructor(props: any) {
        super(props)
        console.log('home props => ', props)
    }
    render(): React.ReactNode {
        return(
                <CO_WorkList
                    styIdx="one"
                    data={data}
                />

        )
    }
}