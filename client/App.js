import React from "react";
import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router";
import { BaseLayout } from "./components/BaseLayout";
import { LoginPage } from "./routes/LoginPage";
import AuthRequired from "./feature/auth/AuthRequired";
import { BillsPage } from "./routes/BillsPage";
import {SignupPage} from "./routes/SignupPage";

export const AppContext = React.createContext({ auth: {} });

class App extends React.PureComponent {
  state = {
    auth: {
      username: undefined,
      token: "",
    },
    apiError: "Yo Yo Yo what's up"
  };

  setRootState = (value, callback) => {
    this.setState(value, callback);
  };

  render() {
    return (
      <AppContext.Provider
        value={{ ...this.state, setRootState: this.setRootState }}
      >
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<BaseLayout />}>
              <Route
                path="/"
                element={
                  <AuthRequired>
                    <BillsPage />
                  </AuthRequired>
                }
              />
              <Route path="/login" element={<LoginPage />} />
              <Route path="/signup" element={<SignupPage />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </AppContext.Provider>
    );
  }
}

// eslint-disable-next-line no-undef
export default App;
