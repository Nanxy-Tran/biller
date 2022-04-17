import React from "react";

class BillInput extends React.PureComponent {
  state = {
    bill: {
      name: "",
      amount: 0,
      description: "",
    },
    errors: [],
    loading: false,
  };

  updateBillValue = (field) => (e) => {
    this.setState({ bill: { [field]: e.target.value } });
  };

  isValidBill = () => {
    const errors = [];
    Object.keys(this.state.bill).forEach((field) => {
      if (field === "amount") {
        if (this.state.amount < 0) {
          errors.push("amount.invalid");
        }
      }

      if (!this.state[field]) {
        errors.push(`${field}.invalid`);
      }
    });

    this.setState({ errors });
    return errors.length < 1;
  };

  submitBill = () => {
    const isValidBill = this.isValidBill();
    if (isValidBill) {
      this.setState({});
    }
  };

  render() {
    const { errors } = this.state;
    console.log(errors.some((error) => error?.includes("name")));
    return (
      <div className="container">
        <button
          className="btn btn-primary"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#billInputForm"
          aria-expanded
          aria-controls="billInputForm"
        >
          <h5>Create new bill ✔</h5>
        </button>

        <div className="collapse" id="billInputForm">
          <div className="container p-3">
            <label className="row">Name:</label>
            <input
              className="form-control row border-gray border"
              type="text"
              id="validation01"
              placeholder="Cat food or chicken dinner...."
              onChange={this.updateBillValue("name")}
            />
            <ErrorIndicator errorKey="name" errors={errors} />

            <label className="row mt-2">Amount:</label>
            <input
              className="form-control row border-gray border"
              type="number"
              onChange={this.updateBillValue("amount")}
            />
            <ErrorIndicator errorKey="amount" errors={errors} />

            <label className="row mt-2 ">Description:</label>
            <input
              className="form-control row border-gray border"
              type="text"
              onChange={this.updateBillValue("description")}
            />
          </div>

          <div className="d-flex flex-column mx-5">
            <button
              className="btn btn-success"
              type="submit"
              onClick={this.submitBill}
            >
              Create
            </button>
          </div>
        </div>
      </div>
    );
  }
}

const ErrorIndicator = ({errorKey, errors}) => {
   const visible = errors.some(error => error?.includes(errorKey))
    if(!visible) {
        return null
    }
    return (
            <span className="text-danger">
                {`⛔ Invalid ${errorKey}`}
            </span>
    )
}

const createBill = async (bill) => {
    return await fetch("http://localhost:8080/api/bill", {
        method: 'POST',
        mode: 'cors',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(bill),
        referrerPolicy: 'no-referrer'
    }).then(res => res.json())
}

export default BillInput