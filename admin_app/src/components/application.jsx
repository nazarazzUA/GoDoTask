import React from 'react';
import ReactDOM from 'react-dom';

import MainMenu from './navigations/main-menu.jsx';

class App extends React.Component {
    render() {
        return <div><MainMenu/></div>
    }
}


ReactDOM.render(<App/>, document.getElementById('application'));
