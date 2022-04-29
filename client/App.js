import React from "react";
import {hot} from "react-hot-loader";
import {BillsPage} from "./routes/BillsPage";
import {BaseLayout} from "./components/BaseLayout";
import {BrowserRouter} from "react-router-dom";
import {Route, Routes} from "react-router";

export const AppContext = React.createContext({});

class App extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {
            userName: "Nanxy-Tran",
            bills : [],
            totalAmount: 0
        }
    }
    componentDidMount() {
        console.log("I'm here dude !")
    }

    render() {
        return (
            <AppContext.Provider value={{session: "fdjashfkdhak"}}>
                <h1>F*ck the world</h1>
                <BrowserRouter>
                    <Routes>
                        <Route path={"/"} component={BillsPage} />
                    </Routes>
                </BrowserRouter>
            </AppContext.Provider>
        )
    }
}

// eslint-disable-next-line no-undef
export default hot(module)(App)