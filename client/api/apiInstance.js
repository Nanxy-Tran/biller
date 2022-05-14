export const Tokenizer = (() => {
    let token = "";
    return {
        setToken: (val) => (token = val),
        get() {
            return `Bearer ${token}`;
        },
    };
})();


const BaseAPI = "http://localhost:8080/api/";

export const ApiConfig = () => ({
    mode: "cors",
    referrerPolicy: "origin",
    headers: {
        Accept: "application/json",
        Authorization: Tokenizer.get(),
    }
})

const parseQuery = (param) =>
    Object.keys(param).reduce(
        (acc, current) =>
            `${acc}${acc === "?" ? "" : "&"}${current}=${param[current]}`,
        "?"
    );

const apiGet = (endpoint, param) => {
    const url = BaseAPI + endpoint + parseQuery(param);
    return fetch(url, {
        method: "GET",
        ...ApiConfig(),
    }).then((res) => res.json());
};

const apiPost =  (endpoint, body) => {
    const url = BaseAPI + endpoint;
    return fetch(url, {
        ...ApiConfig(),
        method: "POST",
        body: JSON.stringify(body),
    }).then((res) => res.json())
}

const apiPut =  (endpoint, body) => {
    const url = BaseAPI + endpoint;
    return fetch(url, {
        ...ApiConfig(),
        method: "PUT",
        body: JSON.stringify(body),
    }).then((res) => res.json())
}

const apiDelete =  (endpoint, params) => {
    const url = BaseAPI + endpoint + parseQuery(params);
    return fetch(url, {
        ...ApiConfig(),
        method: "DELETE",
    }).then((res) => res.json())
}

export {apiGet, apiPost, apiPut, apiDelete }
