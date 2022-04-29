import React from "react";

export class Bill extends React.PureComponent {
    constructor(props) {
        super(props);
    }
    render() {
        const {id, name, amount, createdAt, description} = this.props;
        const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };

        return (
            <div className="row border-bottom border-primary">
                <div className="col align-self-center">
                    {id}
                </div>
                <div className="col align-self-center">
                    {name}
                </div>
                <div className="col align-self-center">
                    {description}
                </div>
                <div className="col align-self-center">
                    {formatCurrency(amount)}
                </div>
                <div className="col">
                    {formatDate(new Date(createdAt))}
                </div>
            </div>
        )
    }
}