import React, { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import { LiveKitRoom, DisplayContext, DisplayOptions } from "livekit-react"
import "livekit-react/dist/index.css"
import "react-aspect-ratio/aspect-ratio.css"
import { USERNAME_KEY } from "../../utils/const"
import SideChat from "../../components/chat/SideChat"
import { useRecoilState } from "recoil"
import tableAtom from "../../globalstate/atom/table"

const LiveKitServerURL = `ws://${window.location.hostname}:7880/`

const CallingView = (): JSX.Element => {
  const navigate = useNavigate()
  const [table, setTable] = useRecoilState(tableAtom)

  const [displayOptions, setDisplayOptions] = useState<DisplayOptions>({
    stageLayout: "grid",
    showStats: true,
  })

  const [token, setToken] = useState("")

  const onLeave = () => {
    navigate("/app/table-view")
    setTable("table0")
  }

  useEffect(() => {
    const randomInfo =
      localStorage.getItem(USERNAME_KEY) || `aassdd${Math.random() * 100}`
    fetch(`/api/v1/token?room=${table || "stark-tower"}&user=${randomInfo}`)
      .then((res) => res.json())
      .then((r) => {
        setToken(r.token)
        console.log(r.token)
      })
      .catch((err) => console.log(`error info${err}`))
  }, [table])

  return (
    <>
      <SideChat />
      <DisplayContext.Provider value={displayOptions}>
        {token && (
          <LiveKitRoom url={LiveKitServerURL} token={token} onLeave={onLeave} />
        )}
      </DisplayContext.Provider>
    </>
  )
}
export default CallingView
