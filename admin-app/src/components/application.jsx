import React from 'react';
import ReactDOM from 'react-dom';

import MainMenu from './navigations/main-menu.jsx';
import Clock    from './clock.jsx';

class App extends React.Component {
    render() {
        return (
            <div>
                <MainMenu/>
                <Clock/>
            </div>
        )
    }
}


ReactDOM.render(<App/>, document.getElementById('application'));
