import {
  CardActions,
  TextField,
  Select,
  MenuItem,
  FormHelperText,
  FormControl,
  InputLabel,
} from '@mui/material'

import { useState, useEffect } from 'react'
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
  const [options, setOptions] = useState<{ id: string; name: string }[]>([])

  useEffect(() => {
    if (inputConfig?.reference && inputConfig.type === 'select') {
      fetch(`http://localhost:8080/${inputConfig.reference}`)
        .then((res) => res.json())
        .then((data) => {
          setOptions(data)
        })
        .catch(() => {
          setOptions([])
        })
    }
  }, [inputConfig?.reference, inputConfig?.type])

  const validate = (): boolean => {
    if (!inputConfig) return true
    if (inputConfig.required && (typeof input === 'string' ? input.trim() : input) === '') {
      setError('Este campo es obligatorio.')
      return false
    }
    if (inputConfig.type !== 'select') {
      if (input.length < inputConfig.minLength) {
        setError(`Debe tener al menos ${inputConfig.minLength} caracteres.`)
        return false
      }
      if (input.length > inputConfig.maxLength) {
        setError(`Debe tener menos de ${inputConfig.maxLength + 1} caracteres.`)
        return false
      }
    }
    setError('')
    return true
  }

  const handleSend = () => {
    if (validate()) {
      onSend(typeof input === 'string' ? input.trim() : input)
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
      {inputConfig?.reference && inputConfig.type === 'select' ? (
        <FormControl fullWidth error={!!error} sx={{ flex: 9.6 }}>
          <InputLabel id="select-label">{`Seleccione ${inputConfig.reference}`}</InputLabel>
          <Select
            labelId="select-label"
            value={input}
            label={`Seleccione ${inputConfig.reference}`}
            onChange={(e) => setInput(e.target.value)}
          >
            {options.map((option) => (
              <MenuItem key={option.id} value={option.id}>
                {option.name}
              </MenuItem>
            ))}
          </Select>
          {error && <FormHelperText>{error}</FormHelperText>}
        </FormControl>
      ) : (
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
      )}
      <ChatSendButton handleSend={handleSend} disabled={inputConfig?.required && (typeof input === 'string' ? input.trim() : input) === ''} />
    </CardActions>
  )
}
