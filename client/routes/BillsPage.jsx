import BillInput from "../components/BillInput";
import React, { useCallback, useEffect, useMemo } from "react";
import { BillGroup } from "../feature/bill/BillGroup";
import { TotalAmount } from "../feature/bill/TotalAmount";
import { apiGet } from "../feature/bill/utils";
import { usePagination } from "../hooks/usePagination";
import { BillPaginationBar } from "../feature/bill/BillPaginationBar";

export const BillsPage = () => {
  const { current_page, total_page, bills, setPayload } = usePagination({
    bills: [],
  });

  const totalAmount = useMemo(
    () => bills.reduce((acc, cur) => cur.amount + acc, 0),
    [bills]
  );

  const fetchBills = useCallback(async (pageIndex = 1) => {
    const response = await apiGet("bills", {
      limit: 10,
      current_page: pageIndex,
    });
    if (response) setPayload(response);
  }, []);

  const fetchArbitraryPage = (pageIndex) => {
    const availablePage = 0 < pageIndex <= total_page;
    if (availablePage) fetchBills(pageIndex);
  };

  useEffect(() => {
    fetchBills();
  }, []);

  return (
    <>
      <BillGroup bills={bills} />
      <BillPaginationBar
        currentPage={current_page}
        totalPage={total_page}
        onSelectPage={fetchArbitraryPage}
      />
      <TotalAmount totalAmount={totalAmount} />
      <BillInput onCreated={fetchBills} />
    </>
  );
};
