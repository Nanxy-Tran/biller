import App from './App';
import React from "react";
import {createRoot} from "react-dom/client";

const domContainer = document.querySelector('#root');
const root = createRoot(domContainer);
root.render(<App />);