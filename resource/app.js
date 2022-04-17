const e = React.createElement;

const getBills = async () => {
   return await fetch("http://localhost:8080/api/bills", {
       method: 'GET',
       mode: 'cors',
       headers: { 'Content-Type': 'application/json'}
   }).then(res =>  res.json())
}

const createBill = async (bill) => {
    return await fetch("http://localhost:8080/api/bills", {
        method: 'POST',
        mode: 'cors',
        headers: { 'Content-Type': 'application/json'},
        body: JSON.stringify(bill)
    }).then(res => res.json())
}

const formatCurrency = (amount) => {
    return new Intl.NumberFormat("vi-VN", {style: "currency", currency: "VND"}).format(amount)
}

class App extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            userName: "Nanxy Tran",
            bills: [],
            total: 0
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
        const {bills, totalAmount} = this.state;
        return e(
            'div', null,
             [e('div', {className: "d-flex h1 text-white bg-info"}, this.state.userName),
                 e(BillGroup, {key: "bills", bills}),
                 e(TotalAmount, {totalAmount})]
        );
    }
}

class BillGroup extends React.PureComponent {
    constructor(props) {
        super(props);
    }
    render() {
        return e(
            'div',
            {className: "container rounded border border-success"},
            [this.props.bills.map(bill => e(Bill, {key: bill.ID, id: bill.ID, billName: bill.name, billAmount: bill.amount, billCreatedTime: bill.createdAt, billDescription: bill.description}))]
        )
    }
}

class Bill extends React.PureComponent {
    constructor(props) {
        super(props);
    }
    render() {
        const {id, billName, billAmount, billCreatedTime, billDescription} = this.props;
        const options = { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' };
        return e(
            'div',
            {className: "row border-bottom border-primary", key: id},
            [e('div', {className: "col align-self-center"}, id),
                e('div', {className: "col align-self-center"}, billName),
                e('div', {className: "col align-self-center"}, billDescription),
                e('div', {className: "col align-self-center"}, formatCurrency(billAmount)),
                e('div', {className: "col"}, new Date(billCreatedTime).toLocaleDateString('vi-VN', options)),
            ]
        )
    }
}

class BillInput extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            name: '',
            amount: 0,
            description: ''
        }
    }
    render() {
        return e('input', {className: ''})
    }

}


class TotalAmount extends React.PureComponent {
    constructor(props) {
        super(props);
    }
    render() {
        return e('div', {className: 'd-flex flex-row-reverse text-primary container'}, `Tổng chi tiêu: ${formatCurrency(this.props.totalAmount)}`)
    }
}

