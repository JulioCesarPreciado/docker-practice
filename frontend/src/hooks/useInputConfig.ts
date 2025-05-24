// hooks/useInputConfig.ts
import { useState } from 'react'

export type InputConfig = {
  maxLength: number
  minLength: number
  reference?: string
  required: boolean
  type: string
}

export function useInputConfig() {
  const [inputConfig, setInputConfig] = useState<InputConfig | undefined>()

  const updateConfig = (config: InputConfig) => {
    const { minLength, maxLength, reference, required, type } = config
    setInputConfig({ minLength, maxLength, reference, required, type })
  }

  return { inputConfig, updateConfig }
}