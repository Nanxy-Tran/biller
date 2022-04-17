import React from "react";
import {hot} from "react-hot-loader";

class App extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {
            userName: "Nanxy-Tran",
            bills : [],
            totalAmount: 0
        }
    }

    async componentDidMount() {
        const bills = await getBills()
        if (bills) {
            const totalAmount = bills.data.reduce((acc, cur) => cur.amount + acc, 0)
            this.setState({bills: bills.data, totalAmount})
        }
    }

    render() {
        return (
            <div>
                <nav className="nav-bar nav-expend-md nav-light bg-light">
                    <div className="d-flex h1 text-white bg-info">
                        {this.state.userName}
                    </div>
                </nav>
                <BillGroup bills={this.state.bills} />
                <TotalAmount totalAmount={this.state.totalAmount} />
            </div>
        )
    }

}

class BillGroup extends React.PureComponent {
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

class Bill extends React.PureComponent {
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
                    {new Date(createdAt).toLocaleDateString(undefined, options)}
                </div>
            </div>
        )
    }
}

// class BillInput extends React.Component {
//     constructor(props) {
//         super(props);
//         this.state = {
//             name: '',
//             amount: 0,
//             description: ''
//         }
//     }
//     render() {
//         return e('input', {className: ''})
//     }
//
// }


class TotalAmount extends React.PureComponent {
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

const getBills = async () => {
    return await fetch("http://localhost:8080/api/bills", {
        method: 'GET',
        mode: 'cors',
        headers: { 'Accept': 'application/json'},
        referrerPolicy: 'no-referrer'
    }).then(res => res.json())
}

const createBill = async (bill) => {
    return await fetch("http://localhost:8080/api/bills", {
        method: 'POST',
        mode: 'cors',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(bill),
        referrerPolicy: 'no-referrer'
    }).then(res => res.json())
}

const formatCurrency = (amount) => {
    return new Intl.NumberFormat("vi-VN", {style: "currency", currency: "VND"}).format(amount)
}

export default hot(module)(App)