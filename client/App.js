import React from "react";
import {hot} from "react-hot-loader";
import BillInput from "./components/BillInput";

class App extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {
            userName: "Nanxy-Tran",
            bills : [],
            totalAmount: 0
        }
    }

    getBills = async () => {
        const bills = await getBills()
        if (bills) {
            const totalAmount = bills.data.reduce((acc, cur) => cur.amount + acc, 0)
            this.setState({bills: bills.data, totalAmount})
        }
    }

    componentDidMount() {
        this.getBills()
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
                <BillInput onCreated={this.getBills} />
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
                    {formatDate(new Date(createdAt))}
                </div>
            </div>
        )
    }
}


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

const formatCurrency = (amount) => {
    return new Intl.NumberFormat("vi-VN", {style: "currency", currency: "VND"}).format(amount)
}

const formatDate = (date) => {
   return new Intl.DateTimeFormat('vi-VN', {timeZone: "UTC",dateStyle: "full", timeStyle: "short"}).format(date)
}

export default hot(module)(App)