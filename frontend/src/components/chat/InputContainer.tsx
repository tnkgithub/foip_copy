import React, { useState } from "react"
import "./InputContainer.css"
import { LengthMax } from "./types"
import TextField from "@mui/material/TextField"

interface Props {
  onSubmit: (text: string) => void
}

const InputContainer = (props: Props): JSX.Element => {
  const [textValue, setTextValue] = useState("")

  const handleKeyDown = (ev: React.KeyboardEvent<HTMLDivElement>) => {
    if (ev.code === "Enter") {
      if (ev.shiftKey) return
      if (ev.nativeEvent.isComposing) return
      ev.preventDefault()

      if (textValue !== "") {
        props.onSubmit(textValue)
        setTextValue("")
      }
    }
  }

  const handleTextChange = (ev: React.ChangeEvent<HTMLTextAreaElement>) => {
    setTextValue(ev.target.value)
  }

  return (
    <>
      <TextField
        className="InputContainer-text"
        onKeyDown={handleKeyDown}
        inputProps={{ maxLength: LengthMax.MESSAGE }}
        onChange={handleTextChange}
        value={textValue}
        placeholder="Message..."
      ></TextField>
    </>
  )
}

export default InputContainer
