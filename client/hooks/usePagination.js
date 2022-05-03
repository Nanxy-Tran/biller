import { useState } from "react";

export const usePagination = (initialData) => {
  const [payload, setPayload] = useState({
    current_page: 0,
    total_page: 0,
  });
  return { ...initialData, ...payload, setPayload };
};
