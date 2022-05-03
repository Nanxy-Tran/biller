import React, { useCallback, useMemo } from "react";

export const BillPaginationBar = ({
  currentPage = 1,
  totalPage = 1,
  onSelectPage,
}) => {
  const items = useMemo(
    () => Array.from({ length: totalPage }, (_, i) => i + 1),
    [totalPage, currentPage]
  );
  const uiItems = ["previous", ...items, "next"];

  const changePage = useCallback((pageIndex) => {
    if (["previous", "next"].includes(pageIndex)) return;
    onSelectPage(pageIndex);
  }, []);

  return (
    <div className="container d-flex justify-content-center mt-3">
      <ul className="pagination">
        {uiItems.map((item) => (
          <PaginationItem
            key={item}
            isActive={currentPage === item}
            onSelect={changePage}
            pageIndex={item}
          />
        ))}
      </ul>
    </div>
  );
};

const PaginationItem = ({ pageIndex, isActive, onSelect }) => {
  const className = "page-link" + (isActive ? " active" : "");
  return (
    <li className="page-item">
      <a className={className} onClick={() => onSelect(pageIndex)} href="#">
        {pageIndex}
      </a>
    </li>
  );
};
