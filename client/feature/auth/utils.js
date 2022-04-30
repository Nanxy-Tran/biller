export const login = (credentials) => {
  return fetch("http://localhost:8080/api/login", {
    method: "POST",
    mode: "cors",
    headers: { Accept: "application/json" },
    referrerPolicy: "no-referrer",
    body: JSON.stringify(credentials),
  }).then((res) => res.json());
};
