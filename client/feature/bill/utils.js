export const getBills = () => {
  return fetch("http://localhost:8080/api/bills", {
    method: "GET",
    mode: "cors",
    headers: { Accept: "application/json" },
    referrerPolicy: "no-referrer",
  })
    .then((res) => res.json())
    .then((res) => res.data);
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
