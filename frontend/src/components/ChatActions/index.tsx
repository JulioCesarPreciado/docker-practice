import {
  CardActions,
  TextField,
} from '@mui/material'

import { useState } from 'react'
import ChatSendButton from '../ChatSendButton'

interface InputConfig {
  maxLength: number
  minLength: number
  reference?: string
  required: boolean
  type: string
}
interface ChatActionsProps {
  onSend: (message: string) => void
  inputConfig?: InputConfig
}

export default function ChatActions({ onSend, inputConfig }: ChatActionsProps) {
  const [input, setInput] = useState('')
  const [error, setError] = useState('')

  const validate = (): boolean => {
    if (!inputConfig) return true
    if (inputConfig.required && input.trim() === '') {
      setError('Este campo es obligatorio.')
      return false
    }
    if (input.length < inputConfig.minLength) {
      setError(`Debe tener al menos ${inputConfig.minLength} caracteres.`)
      return false
    }
    if (input.length > inputConfig.maxLength) {
      setError(`Debe tener menos de ${inputConfig.maxLength + 1} caracteres.`)
      return false
    }
    setError('')
    return true
  }

  const handleSend = () => {
    if (validate()) {
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
        error={!!error}
        helperText={error}
        sx={{ flex: 9.6 }}
        slotProps={{
          input: {
            inputProps: {
              minLength: inputConfig?.minLength,
              maxLength: inputConfig?.maxLength,
            }
          }
        }}
      />
      <ChatSendButton handleSend={handleSend} />
    </CardActions>
  )
}
