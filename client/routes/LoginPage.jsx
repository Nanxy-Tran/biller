import React, {useCallback, useContext, useState} from "react";
import {useLocation, useNavigate} from "react-router";
import {AppContext} from "../App";
import Carrier, {Tokenizer} from "../api/apiInstance";

const LoginPage = () => {
    const navigate = useNavigate();
    const location = useLocation();
    const appContext = useContext(AppContext);

    const [form, setForm] = useState({
        email: "",
        password: "",
    });

    const onInput = (field) => (e) => {
        e.preventDefault();
        setForm((prev) => ({...prev, [field]: e.target.value}));
    };

    const handleLogin = useCallback(async () => {
        const response = await Carrier.Post("login", form);
        if (response.error) {
            appContext.setRootState({apiError: response.error});
        }

        if (response) {
            Tokenizer.setToken(response.token);
            appContext.setRootState({auth: response}, () =>
                navigate(location?.state?.pathname || "/", {replace: true})
            );
        }
    }, [form, navigate, appContext, location]);

    return (
        <div className="container col-6 border rounded-3 d-flex flex-column align-items-center py-4 mx-auto mt-5">
            <div className="container">
                <h3 className="text-center">Don't forget yourself !</h3>
            </div>

            <div className="container col-10 align-items-center">
                <div className="input-group flex-nowrap">
                    <span className="input-group-text">Email</span>
                    <input
                        type="email"
                        className="form-control"
                        placeholder="Email.."
                        onChange={onInput("email")}
                    />
                </div>

                <div className="input-group flex-nowrap mt-4">
                    <span className="input-group-text">Password</span>
                    <input
                        type="password"
                        className="form-control"
                        placeholder="***"
                        onChange={onInput("password")}
                    />
                </div>
            </div>

            <div className="container row mt-4">
                <button
                    className="btn btn-success rounded-pill col-6 mx-auto"
                    type="submit"
                    onClick={handleLogin}
                >
                    Login
                </button>
            </div>

            <div className="container row mt-4">
                <button
                    className="btn btn-success rounded-pill col-6 mx-auto"
                    type="submit"
                    onClick={() => navigate("/signup")}
                >
                    Signup
                </button>
            </div>
        </div>
    );
};

export default React.memo(LoginPage)