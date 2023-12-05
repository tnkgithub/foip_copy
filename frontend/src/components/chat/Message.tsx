import React from "react"
import "./Message.css"
import Box from "@mui/material/Box"

interface PropType {
  user: string
  time: string
  table: string
  message: string
}

const lbToBr = (txt: string): (string | JSX.Element)[] => {
  return txt.split(/(\n)/g).map((t) => (t === "\n" ? <br /> : t))
}

const Message = (props: PropType) => {
  return (
    <Box sx={{ border: 1, p: "8px", my: "8px" }}>
      <Box sx={{ display: "flex", justifyContent: "space-between" }}>
        <Box sx={{ textAlign: "left" }}>{props.user}</Box>
        <Box sx={{ textAlign: "right" }}>{props.time}</Box>
      </Box>
      <Box>{lbToBr(props.message)}</Box>
    </Box>
  )
}

export default Message
