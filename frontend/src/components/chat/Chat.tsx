import React, { useState } from "react"
import InputContainer from "./InputContainer"
import Message from "./Message"
import { ChatData } from "./types"
import "./Chat.css"

interface Props {
  chatData: ChatData[]
  onSubmit: (text: string) => void
}

const Chat = (props: Props): JSX.Element => {
  const handleSubmit = (text: string) => {
    props.onSubmit(text)
  }

  return (
    <div className="Chat-space">
      <div className="Message-area">
        {props.chatData.map((data, index) => {
          return (
            <Message
              key={index}
              user={data.user}
              time={data.time}
              table={data.table}
              message={data.message}
            />
          )
        })}
      </div>
      <div className="Message-send">
        <InputContainer onSubmit={handleSubmit} />
      </div>
    </div>
  )
}

export default Chat
