import {
  Box,
  CardActions,
  TextField,
  IconButton,
} from '@mui/material'
import SendIcon from '@mui/icons-material/Send'
import { useState } from 'react'

interface ChatActionsProps {
  onSend: (message: string) => void
}

export default function ChatActions({ onSend }: ChatActionsProps) {
  const [input, setInput] = useState('')

  const handleSend = () => {
    if (input.trim() !== '') {
      onSend(input.trim())
      setInput('')
    }
  }

  const handleKeyPress = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      handleSend()
    }
  }

  return (
    <CardActions
      sx={{
        px: 2,
        py: 2,
        display: 'flex',
        alignItems: 'center',
        gap: 1,
      }}
    >
      <TextField
        fullWidth
        placeholder="Type your message..."
        variant="outlined"
        value={input}
        onChange={(e) => setInput(e.target.value)}
        onKeyDown={handleKeyPress}
        sx={{ flex: 9.6 }}
      />
      <Box sx={{ flex: 0.4, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
        <IconButton
          color="primary"
          aria-label="send message"
          sx={{ borderRadius: '50%' }}
          onClick={handleSend}
        >
          <SendIcon />
        </IconButton>
      </Box>
    </CardActions>
  )
}
