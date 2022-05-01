import BillInput from "../components/BillInput";
import React, { useCallback, useEffect, useMemo, useState } from "react";
import { BillGroup } from "../feature/bill/BillGroup";
import { TotalAmount } from "../feature/bill/TotalAmount";
import { getBills } from "../feature/bill/utils";

export const BillsPage = () => {
  const [bills, setBills] = useState([]);

  const totalAmount = useMemo(
    () => bills.reduce((acc, cur) => cur.amount + acc, 0),
    [bills]
  );

  const fetchBills = useCallback(async () => {
    const response = await getBills();
    if (response) {
      setBills(response);
    }
  }, []);

  useEffect(() => {
    fetchBills();
  }, []);

  return (
    <>
      <BillGroup bills={bills} />
      <TotalAmount totalAmount={totalAmount} />
      <BillInput onCreated={fetchBills} />
    </>
  );
};
