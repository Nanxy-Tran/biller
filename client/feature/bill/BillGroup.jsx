import React from "react";
import PropTypes from "prop-types"
import {Bill} from "./Bill"

export class BillGroup extends React.PureComponent {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="container rounded border border-success">
                {this.props.bills.map(bill => <Bill key={bill.ID}  {...bill} />)}
            </div>
        )
    }
}
BillGroup.propTypes = {
    bills: PropTypes.array.isRequired
}