import React, { lazy, Suspense } from "react";
import { BrowserRouter } from "react-router-dom";
import { Route, Routes } from "react-router";
import AuthRequired from "./feature/auth/AuthRequired";
import Carrier from "./api/apiInstance";

const SignupPage = lazy(() => import("./routes/SignupPage"));
const LoginPage = lazy(() => import("./routes/LoginPage"));
const BillsPage = lazy(() => import("./routes/BillsPage"));
const BaseLayout = lazy(() => import("./components/BaseLayout"));

export const AppContext = React.createContext({ auth: {} });

class App extends React.PureComponent {
  state = {
    auth: {
      username: undefined,
      token: "",
    },
    apiError: "Yo Yo Yo what's up",
  };

  componentDidMount() {
    Carrier.initErrorHandler((err) => this.setRootState({ apiError: err }));
  }

  setRootState = (value, callback) => {
    this.setState(value, callback);
  };

  render() {
    return (
      <AppContext.Provider
        value={{ ...this.state, setRootState: this.setRootState }}
      >
        <BrowserRouter>
          <Suspense fallback={<div>Loading....</div>}>
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
          </Suspense>
        </BrowserRouter>
      </AppContext.Provider>
    );
  }
}

// eslint-disable-next-line no-undef
export default App;
