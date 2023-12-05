import React, { useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"
import Button from "@mui/material/Button"
import CssBaseline from "@mui/material/CssBaseline"
import TextField from "@mui/material/TextField"
import Grid from "@mui/material/Grid"
import Box from "@mui/material/Box"
import Typography from "@mui/material/Typography"
import Container from "@mui/material/Container"
import { ThemeProvider } from "@mui/material/styles"
import theme from "../../components/common/theme"
import { InputAdornment } from "@mui/material"
import { AccountCircle, Mic, VolumeUp, Videocam } from "@mui/icons-material/"
import MenuItem from "@mui/material/MenuItem"
import { USERNAME_KEY } from "../../utils/const"

const VIDEO_KIND = "videoinput"
const SPEAKER_KIND = "audiooutput"
const MICROPHONE_KIND = "audioinput"

const CALLING_DEVICE_KEY = "CALLING_DEVICE_KEY"

const WelcomePage: React.FC = () => {
  const navigate = useNavigate()

  const buttonHandler: React.ReactEventHandler<Element> = () => {
    localStorage.setItem(
      CALLING_DEVICE_KEY,
      JSON.stringify({ microphone, speaker, video, username })
    )
    navigate("/app/table-view")
  }

  const [username, setUsername] = useState("")

  const [devices, setDevices] = useState([] as MediaDeviceInfo[])
  const [microphone, setMicrophone] = useState("default")
  const [speaker, setSpeaker] = useState("default")
  const [video, setVideo] = useState("default")

  const handleUserName = (event: React.ChangeEvent<HTMLInputElement>) => {
    setUsername(event.target.value)
    localStorage.setItem(USERNAME_KEY, event.target.value)
  }

  useEffect(() => {
    navigator.mediaDevices
      .getUserMedia({ video: true, audio: true })
      .then(function (stream) {
        navigator.mediaDevices.enumerateDevices().then((devices) => {
          setDevices(() => devices)
        })
        stream.getTracks().forEach(function (track) {
          track.stop()
        })
      })
    navigator.mediaDevices.ondevicechange = () => {
      navigator.mediaDevices.enumerateDevices().then((devices) => {
        setDevices(() => devices)
      })
    }
  }, [])

  useEffect(() => {
    const videos = devices.filter((item) => item.kind === VIDEO_KIND)
    if (videos.length < 1) return
    setVideo(() => videos[0].deviceId)

    const speaker = devices.filter((item) => item.kind === SPEAKER_KIND)
    if (speaker.length < 1) return
    setSpeaker(() => speaker[0].deviceId)

    const microphone = devices.filter((item) => item.kind === MICROPHONE_KIND)
    if (microphone.length < 1) return
    setMicrophone(() => microphone[0].deviceId)
  }, [devices])

  const handleChangeMicrophone = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setMicrophone(event.target.value)
    console.log(event.target.value)
  }

  const handleChangeSpeaker = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSpeaker(event.target.value)
    console.log(event.target.value)
  }

  const handleChangeVideo = (event: React.ChangeEvent<HTMLInputElement>) => {
    setVideo(event.target.value)
    console.log(event.target.value)
  }

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Typography variant="h2" color={"#8a0000"} fontWeight={900}>
            FOIP
          </Typography>
          <Box component="form" noValidate sx={{ mt: 3 }}>
            <Grid container spacing={2}>
              <Grid item xs={12}>
                <Typography>Profile</Typography>
                <TextField
                  id="user-name"
                  placeholder="User Name"
                  required
                  fullWidth
                  sx={{ mb: 3 }}
                  onChange={handleUserName}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <AccountCircle />
                      </InputAdornment>
                    ),
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                <Typography>Devices</Typography>
                <TextField
                  required
                  fullWidth
                  id="microphone"
                  placeholder="Device Name"
                  select
                  value={microphone}
                  onChange={handleChangeMicrophone}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <Mic />
                      </InputAdornment>
                    ),
                  }}
                >
                  {devices
                    .filter((item) => item.kind === MICROPHONE_KIND)
                    .map((option, index) => (
                      <MenuItem key={index} value={option.deviceId}>
                        {`${option.label}`}
                      </MenuItem>
                    ))}
                </TextField>
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="speaker"
                  placeholder="Device Name"
                  select
                  value={speaker}
                  onChange={handleChangeSpeaker}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <VolumeUp />
                      </InputAdornment>
                    ),
                  }}
                >
                  {devices
                    .filter((item) => item.kind === SPEAKER_KIND)
                    .map((option, index) => (
                      <MenuItem key={index} value={option.deviceId}>
                        {`${option.label}`}
                      </MenuItem>
                    ))}
                </TextField>
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="video"
                  select
                  value={video}
                  placeholder="Device Name"
                  onChange={handleChangeVideo}
                  InputProps={{
                    startAdornment: (
                      <InputAdornment position="start">
                        <Videocam />
                      </InputAdornment>
                    ),
                  }}
                >
                  {devices
                    .filter((item) => item.kind === VIDEO_KIND)
                    .map((option, index) => (
                      <MenuItem key={index} value={option.deviceId}>
                        {`${option.label}`}
                      </MenuItem>
                    ))}
                </TextField>
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              onClick={buttonHandler}
            >
              JOIN
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  )
}
export default WelcomePage
