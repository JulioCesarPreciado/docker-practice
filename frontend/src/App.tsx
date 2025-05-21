import {
  Card,
  CardContent,
  Typography,
} from '@mui/material'

import './App.css'
import ChatActions from './components/ChatActions'
import ChatContainer from './components/ChatContainer'

import { useState } from 'react'
import { useSocket } from './hooks/useSocket'

function App() {
  const [messages, setMessages] = useState<string[]>([])

  const { sendMessage } = useSocket((newMsg: string) => {
    setMessages((prev) => [...prev, newMsg])
  })

  return (
    <ChatContainer>
      <Typography
        variant="h4"
        component="h1"
        textAlign="center"
        fontWeight="bold"
        gutterBottom
      >
        Realtime Chat App 📱
      </Typography>

      <Card
        sx={{
          flexGrow: 1,
          overflow: 'auto',
          display: 'flex',
          flexDirection: 'column',
          boxShadow: 3,
          borderRadius: 4,
        }}
      >
        <CardContent
          sx={{
            flex: 1,
            backgroundColor: 'background.paper',
          }}
        >
          {/* Mensajes u otros contenidos */}
        </CardContent>
        <ChatActions />
      </Card>
    </ChatContainer>
  )
}

export default App
