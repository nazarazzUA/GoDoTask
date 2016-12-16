import React from 'react';
import ReactDOM from 'react-dom';
import Applicatiion from './components/application.jsx';
import { Router, Route, Link, browserHistory } from 'react-router';
import TaskListComponent from './components/tasts/list.jsx';
import NoMatch from './components/errors/noMatch.jsx';

ReactDOM.render((
    <Router history={browserHistory}>
        <Route path="/" component={Applicatiion}>
            <Route path="tasks" component={TaskListComponent}/>

            <Route path="*" component={NoMatch}/>
        </Route>
    </Router>
),document.getElementById('application'));