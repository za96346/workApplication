import debounce from 'lodash.debounce'; // 4.0.8
import React from 'react';

const withPreventDoubleClick = (WrappedComponent: React.Component) => {

  class PreventDoubleClick extends React.PureComponent<any, any> {

    debouncedOnPress = () => {
      this.props.onPress && this.props.onPress();
    }

    onPress = debounce(this.debouncedOnPress, 300, { leading: true, trailing: false });

    render() {
      return <WrappedComponent {...this.props} onPress={this.onPress} />;
    }
  }

  PreventDoubleClick.displayName = `withPreventDoubleClick(${WrappedComponent.displayName || WrappedComponent.name})`
  return PreventDoubleClick;
}