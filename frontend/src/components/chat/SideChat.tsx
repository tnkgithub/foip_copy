import React, { useEffect, useState } from "react"
import Chat from "./Chat"
import "./SideChat.css"
import { ChatData } from "./types"
import { LengthMax } from "./types"
import tableAtom from "../../globalstate/atom/table"
import { useRecoilValue } from "recoil"

const SideChat: React.FC = () => {
  const table = useRecoilValue(tableAtom)
  const url = `ws://${window.location.hostname}/api/v1/chat`
  const [socket, setSocket] = useState(new WebSocket(url))
  const [chatData, setChatData] = useState([] as ChatData[])

  useEffect(() => {
    const listener_soket_message = function (event: any) {
      const data: ChatData = JSON.parse(event.data)
      if (table === data.table) {
        setChatData((chatData) => {
          const tmp = [...chatData, data]
          localStorage.setItem("chatData", JSON.stringify(tmp))
          return tmp
        })
        console.log("socket receive")
      }
    }
    const listener_soket_connect = function () {
      console.log("connect")
    }

    socket.addEventListener("open", listener_soket_connect)
    socket.addEventListener("message", listener_soket_message)
    console.log(table)

    return function cleanup() {
      console.log("clean up function")
      socket.removeEventListener("message", listener_soket_message)
      socket.removeEventListener("open", listener_soket_connect)
    }
  }, [socket, table])

  useEffect(() => {
    setChatData(JSON.parse(localStorage.getItem("chatData") ?? "[]"))
  }, [])
  
  const username: string = localStorage.getItem('USERNAME_KEY') || "";
  
  const handleChatSubmit = (message: string) => {
    const data: ChatData = {
      user: username,
      time: "",
      table,
      message,
    }

    if (data.user.length > LengthMax.USER) {
      console.log("User name size exceeds fixed maximum user name size")
      return
    }
    if (data.table.length > LengthMax.TABLE) {
      console.log("Table name size exceeds fixed maximum table name size")
      return
    }
    if (data.message.length > LengthMax.MESSAGE) {
      console.log("Message size exceeds fixed maximum message size")
      return
    }

    const json = JSON.stringify(data)

    if (json.length > LengthMax.TOTAL) {
      console.log("JSON size exceeds fixed maximum JSON size")
    } else {
      socket.send(json)
    }
  }

  return (
    <div className="sidechat-wrapper">
      <div className="sidechat">
        <Chat chatData={chatData} onSubmit={handleChatSubmit} />
      </div>
    </div>
  )
}
export default SideChat
