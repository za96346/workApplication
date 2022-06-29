import React from "react";
import { Text } from "react-native"; 

export default class Main extends React.Component{
    constructor(props:any) {
        super(props)
    }
    render(): JSX.Element {
        return (
            <>
                <Text>Main page</Text>
            </>
        );
    }
}