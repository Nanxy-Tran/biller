import React from "react";
import PropTypes from 'prop-types';

export class TotalAmount extends React.PureComponent {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="d-flex flex-row-reverse text-primary container" >
                {`Tổng chi tiêu: ${formatCurrency(this.props.totalAmount)}`}
            </div>
        )
    }
}

TotalAmount.propTypes = {
    totalAmount: PropTypes.number
}