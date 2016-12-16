import React from 'react';
import MainMenu from './navigations/main-menu.jsx';

export default class App extends React.Component {
    render() {
        return (
            <div>
                <MainMenu/>

                <div className="container">
                    {this.props.children}
                </div>
            </div>
        )
    }
}



