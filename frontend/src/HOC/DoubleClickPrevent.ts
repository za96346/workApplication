import React, { Component, ReactComponentElement } from 'react';
import { debounce } from "lodash";

const withPreventDoubleClick = (WrappedComponent: any) => {
  
  class PreventDoubleClick extends React.PureComponent {
    onPress: any;
    constructor(onPress: any) {
        super(onPress)
    }

    debouncedOnPress = () => {
      this.props.onPress && this.props.onPress();
    }
    
    onPress = debounce(this.debouncedOnPress, 300, { leading: true, trailing: false });
    
    render() {
      return (<WrappedComponent {...this.props} onPress={this.onPress} />)
    }
  }
  
  PreventDoubleClick.displayName = `withPreventDoubleClick(${WrappedComponent.displayName || WrappedComponent.name})`
  return PreventDoubleClick;
}

export default withPreventDoubleClick;