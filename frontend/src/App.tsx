import React from "react"
import "./App.css"
import { BrowserRouter as Router, Routes, Route } from "react-router-dom"
import WelcomePage from "./pages/WelcomePage/WelcomePage"
import SelectTablePage from "./pages/SelectTablePage/SelectTablePage"
import CallingView from "./pages/CallingPage/CallingPage"
import { RecoilRoot } from "recoil"
import { display, flexbox } from "@mui/system"

function App() {
  return (
    <div className="wrapper">
      <RecoilRoot>
        <Router>
          <Routes>
            <Route path="/" element={<WelcomePage />} />
            <Route path="/app/table-view" element={<SelectTablePage />} />
            <Route path="/app/calling-view" element={<CallingView />} />
          </Routes>
        </Router>
      </RecoilRoot>
    </div>
  )
}

export default App
