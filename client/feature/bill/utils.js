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
