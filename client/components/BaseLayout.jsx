import { Outlet } from "react-router";
import React, { useContext } from "react";
import { AppContext } from "../App";

export const BaseLayout = () => {
  const { username } = useContext(AppContext);
  return (
    <div className={"container-fluid flex px-0"}>
      <nav className="nav-bar shadow bg-body nav-expend-md nav-light bg-light">
        <div className="d-flex h1 text-white bg-info py-3">
          {username || "Biller, be rich already !"}
        </div>
      </nav>
      <div className="mt-5">
        <Outlet />
      </div>
    </div>
  );
};
