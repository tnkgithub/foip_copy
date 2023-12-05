import React, { useEffect } from "react"
import "./SelectTablePage.css"
import { useNavigate } from "react-router-dom"
import BaseImage from "../../assets/base.svg"
import HoverImage from "../../assets/hover.svg"
import SideChat from "../../components/chat/SideChat"
import { useRecoilState } from "recoil"
import tableAtom from "../../globalstate/atom/table"
import CircleTable from "../../components/table/CircleTable"

const TableSelect = (): JSX.Element => {
  const navigate = useNavigate()
  const [table, setTable] = useRecoilState(tableAtom)

  return (
    // TODO: display: flexで縦に並べる
    <>
      <SideChat />
      <div className="table-selector">
        <CircleTable tableName="table1" />
        <CircleTable tableName="table2" />
        <CircleTable tableName="table3" />
      </div>
    </>
  )
}
export default TableSelect
