import React from "react"
import { useNavigate } from "react-router-dom"
import { useRecoilState } from "recoil"
import BaseImage from "../../assets/base.svg"
import HoverImage from "../../assets/hover.svg"
import tableAtom from "../../globalstate/atom/table"
import "./CircleTable.css"

interface Props {
  tableName: string
}

const CircleTable = (props: Props) => {
  const navigate = useNavigate()
  const [table, setTable] = useRecoilState(tableAtom)

  const buttonHandler = () => {
    navigate("/app/calling-view")
    setTable(props.tableName)
  }

  return (
    <div className="circle-table" onClick={buttonHandler}>
      <img className="base" src={BaseImage}></img>
      <div className="hovered">
        <img src={HoverImage}></img>
      </div>
    </div>
  )
}
export default CircleTable
