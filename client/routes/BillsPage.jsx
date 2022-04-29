import BillInput from "../components/BillInput";
import React, {useCallback, useEffect, useMemo, useState} from "react";
import {BillGroup} from "../feature/bill/BillGroup";
import {TotalAmount} from "../feature/bill/TotalAmount";

export const BillsPage = () => {
    const [bills, setBills] = useState([])
    const totalAmount = useMemo(() => bills.data.reduce((acc, cur) => cur.amount + acc, 0) ,[bills])

    const getBills = useCallback(async () => {
        const response = await getBills()
        if (bills) {
            setBills(response)
        }
    }, [])

    useEffect(() => {
        getBills()
    }, [])

    return (
        <>
            <BillGroup bills={this.state.bills} />
            <TotalAmount totalAmount={totalAmount} />
            <BillInput onCreated={getBills} />
        </>
    )
}
