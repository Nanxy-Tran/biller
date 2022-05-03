export const Tokenizer = (() => {
  let token =
    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXNzd29yZCI6ImNvbmNhYmlldGJheSIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImV4cCI6MTY1MTU4MjQ3NH0.Ik806riH7GGOnK_SWlYF1-QLYOzRAzflHqd1JalYreI";
  return {
    setToken: (val) => (token = val),
    get() {
      return `Bearer ${token}`;
    },
  };
})();

const BaseAPI = "http://localhost:8080/api/";

const parseQuery = (param) =>
  Object.keys(param).reduce(
    (acc, current) =>
      `${acc}${acc === "?" ? "" : "&"}${current}=${param[current]}`,
    "?"
  );

export const apiGet = (endpoint, param) => {
  const url = BaseAPI + endpoint + parseQuery(param);
  console.log(parseQuery(param));
  return fetch(url, {
    method: "GET",
    mode: "cors",
    referrerPolicy: "origin",
    headers: {
      Accept: "application/json",
      Authorization: Tokenizer.get(),
    },
  }).then((res) => res.json());
};

export const getBills = () => {
  return fetch("http://localhost:8080/api/bills", {
    method: "GET",
    mode: "cors",
    referrerPolicy: "origin",
    headers: {
      Accept: "application/json",
      Authorization: Tokenizer.get(),
    },
  }).then((res) => res.json());
};

export const formatCurrency = (amount) => {
  return new Intl.NumberFormat("vi-VN", {
    style: "currency",
    currency: "VND",
  }).format(amount);
};

export const formatDate = (date) => {
  return new Intl.DateTimeFormat("vi-VN", {
    timeZone: "UTC",
    dateStyle: "full",
    timeStyle: "short",
  }).format(date);
};
