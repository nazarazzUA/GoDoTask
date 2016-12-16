import React from 'react';
import { Link } from 'react-router'

export default class MainMenu extends React.Component {

    constructor(props) {
        super(props);
        this.menuItems = [
            {url:'/tasks', title:'Task'},
            {url:'#', title:'Plans'},
            {url:'#', title:'Statistic'},
            {url:'#', title:'Something else'}
        ];
    }

    render() {

        var listItem = this.menuItems.map(function(item) {
            return (
                <li key={item.title}>
                    <Link to={item.url} >{item.title}</Link>
                </li>
            )
        });

        return (
            <nav className="teal">
                <div className="nav-wrapper">
                    <ul className="left hide-on-med-and-down">
                        {listItem}
                    </ul>
                </div>
            </nav>
        )
    }
}

