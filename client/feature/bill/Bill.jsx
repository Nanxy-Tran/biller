import React from "react";
import { formatCurrency, formatDate } from "./utils";
import Proptypes from 'prop-types'

export class Bill extends React.PureComponent {
  constructor(props) {
    super(props);
  }

  render() {
    const { id, name, amount, created_at } = this.props;

    return (
      <div className="row border-bottom border-primary py-4">
        <div className="col align-self-center">{id}</div>
        <div className="col align-self-center">{name}</div>
        <div className="col align-self-center">{formatCurrency(amount)}</div>
        <div className="col">{formatDate(new Date(created_at))}</div>
      </div>
    );
  }
}

Bill.propTypes = {
    id: Proptypes.string.isRequired,
    name: Proptypes.string.isRequired,
    amount: Proptypes.number.isRequired,
    created_at: Proptypes.string.isRequired
}