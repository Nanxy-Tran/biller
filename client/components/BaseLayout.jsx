import {Outlet} from "react-router";
import React, {useContext} from "react";
import {AppContext} from "../App";

export const BaseLayout = () => {
    const {username} = useContext(AppContext)
    return (
        <div className={"container"}>
            <nav className="nav-bar nav-expend-md nav-light bg-light">
                <div className="d-flex h1 text-white bg-info">
                    {username || "Biller, be rich already !"}
                </div>
            </nav>
            <Outlet />
        </div>

    )
}
