import './App.css'
import ChatActions from './components/ChatActions'
import ChatContainer from './components/ChatContainer'

import { useState } from 'react'
import { useSocket } from './hooks/useSocket'
import MessageBubble from './components/MessageBubble'
import { CardContent } from '@mui/material'

type Message = {
  text: string
  from: 'user' | 'system'
}

function App() {
  const [messages, setMessages] = useState<Message[]>([])

  const { sendMessage } = useSocket((newMsg: string) => {
    const parsed = JSON.parse(newMsg)
    setMessages((prev) => [...prev, { text: parsed.message, from: 'system' }])
  })

  const handleSend = (msg: string) => {
    sendMessage(msg)
    setMessages((prev) => [...prev, { text: msg, from: 'user' }])
  }

  return (
    <ChatContainer title="Realtime Chat App 📱">
      <CardContent
        sx={{
          flex: 1,
          backgroundColor: 'background.paper',
        }}
      >
         {messages.map((msg, idx) => (
        <MessageBubble
          key={idx}
          message={msg.text}
          align={msg.from === 'user' ? 'right' : 'left'}
        />
      ))}
      </CardContent>
      <ChatActions onSend={handleSend} />
    </ChatContainer>
  )
}

export default App
