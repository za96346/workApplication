import React from "react"
import Modal from 'react-native-modal';
import { ActivityIndicator, Text } from "react-native";
import { interfaceCO_LoadingProps, interfaceCO_LoadingState } from "../type/type";
import { language } from "../language";

type typeMainStyles = {
    one: {
        loadingColor: string,
        textStyle: Object,
        text: string
    }
}

const mainStyles = {
    one: {
        loadingColor: '#fff',
        textStyle: {
            color: '#fff', 
            fontSize: 20, 
            marginTop: 10
        },
        text: language.loading,
    }
}


export default class CO_Loading extends React.Component <interfaceCO_LoadingProps, interfaceCO_LoadingState> {
    private mainStyle: any;
    private interval: any;
    constructor(props: any) {
        super(props)
        this.mainStyle = mainStyles[this.props.styIdx as keyof typeof mainStyles] || 'one'
        this.state = {
            dot: ''
        }
    }
    
    componentDidMount() {
        this.interval = setInterval(() => {
            this.setState({
                dot: this.state.dot === '...' ? '' : this.state.dot + '.'
            })
        }, 500)
    }

    componentWillUnmount() {
        if (this.interval) clearInterval(this.interval);
    }

    render(): JSX.Element {
        return(
            <Modal
                animationIn={'slideInUp'}
                animationOut={'slideOutDown'}
                isVisible={this.props.isVisible}
                // onBackButtonPress={() => {}}
                style={{justifyContent: 'center', alignItems: 'center'}}
                >
                <ActivityIndicator size="large" color={this.mainStyle.loadingColor}/>
                <Text style={this.mainStyle.textStyle}>{this.mainStyle.text + this.state.dot}</Text>
            </Modal>
        )
    }
}