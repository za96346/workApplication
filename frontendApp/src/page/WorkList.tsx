import { Text, View } from "native-base";
import React from "react";
import CO_ChangeWorkTime from "../component/CO_ChangeWorkTIme";
class WorkList extends React.Component <any, any>{
    constructor(props: any) {
        super(props)
    }
    render(): React.ReactNode {
        return(
            <View
                style={{
                    width: '100%',
                    height: '100%',
                }}>
                <CO_ChangeWorkTime />
            </View>
        )
    }
}
export default WorkList;