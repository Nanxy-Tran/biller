import { useContext } from "react";
import { AppContext } from "../../App";
import { Navigate, useLocation } from "react-router";
import React from "react";

import PropsTypes from "prop-types";

const AuthRequired = ({ children }) => {
  const { auth } = useContext(AppContext);
  const location = useLocation();
  if (!auth.token) {
    return <Navigate to="/login" state={{ from: location }} replace />;
  }
  return children;
};

export default AuthRequired;

AuthRequired.propTypes = {
  children: PropsTypes.element,
};
