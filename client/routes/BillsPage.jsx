import React, {lazy, useCallback, useEffect, useMemo} from "react";

import {usePagination} from "../hooks/usePagination";
import Carrier from "../api/apiInstance";

const BillGroup = lazy(() => import("../feature/bill/BillGroup"))
const BillPaginationBar = lazy(() => import("../feature/bill/BillPaginationBar"))
const BillInput = lazy(() => import("../components/BillInput"))
const TotalAmount = lazy(() => import("../feature/bill/TotalAmount"))

const BillsPage = () => {
    const {current_page, total_page, bills, setPayload} = usePagination({
        bills: [],
    });

    const totalAmount = useMemo(() => bills.reduce((acc, cur) => cur.amount + acc, 0), [bills]);

    const fetchBills = useCallback(async (pageIndex = 1) => {
        const response = await Carrier.Get("bills", {
            limit: 10, current_page: pageIndex,
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

    return (<>
            <BillGroup bills={bills}/>
            <BillPaginationBar
                currentPage={current_page}
                totalPage={total_page}
                onSelectPage={fetchArbitraryPage}
            />
            <TotalAmount totalAmount={totalAmount}/>
            <BillInput onCreated={fetchBills}/>
        </>);
};

export default React.memo(BillsPage)