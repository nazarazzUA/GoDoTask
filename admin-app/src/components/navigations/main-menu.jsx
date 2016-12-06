import React from 'react';

export default class MainMenu extends React.Component {

    constructor(props) {
        super(props);
        this.menuItems = [
            {url:'#', title:'Task'},
            {url:'#', title:'Plans'},
            {url:'#', title:'Statistic'},
            {url:'#', title:'Something else'}
        ];
    }

    render() {

        var listItem = this.menuItems.map(function(item) {
            return (
                <li key={item.title}>
                    <a href={item.url} >{item.title}</a>
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

