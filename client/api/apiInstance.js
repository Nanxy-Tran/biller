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

//TODO: try to write in async / await with a curry onError dispatcher ?
export const Carrier = (() => {
    const handler = {
        onSuccess: undefined,
        onError: () => null,
    }

    const errorGuard = async (response) => {
        if (!response.ok) {
            response = await response.json()
            let errorMessage = response?.error || "Unknown error"
            return handler.onError(errorMessage)
        }
        return await response.json()
    }

    const Get = async (endpoint, param) => {
        const url = BaseAPI + endpoint + parseQuery(param);
        const response = await fetch(url, {
            method: "GET",
            ...ApiConfig(),
        })
        return errorGuard(response)

    };

    const Post = async (endpoint, body) => {
        const url = BaseAPI + endpoint;
        const response = await fetch(url, {
            ...ApiConfig(),
            method: "POST",
            body: JSON.stringify(body),
        })
        return errorGuard(response)
    }

    const Put = (endpoint, body) => {
        const url = BaseAPI + endpoint;
        return fetch(url, {
            ...ApiConfig(),
            method: "PUT",
            body: JSON.stringify(body),
        }).then((res) => res.json())
    }

    const Delete = (endpoint, params) => {
        const url = BaseAPI + endpoint + parseQuery(params);
        return fetch(url, {
            ...ApiConfig(),
            method: "DELETE",
        }).then((res) => res.json())
    }

    return {
        Get, Post, Put, Delete, initErrorHandler: (fnc) => (handler.onError = fnc)
    }
})()
