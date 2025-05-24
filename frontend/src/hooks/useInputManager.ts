import { useState } from "react";

type Message = {
  text: string
  from: 'user' | 'system'
};

type InputConfig = {
  maxLength: number
  minLength: number
  reference?: string
  required: boolean
  type: string
};

// hooks/useInputManager.ts
export function useInputManager() {
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputConfig, setInputConfig] = useState<InputConfig>();

  const addMessage = (text: string, from: 'user' | 'system') => {
    setMessages((prev) => [...prev, { text, from }]);
  };

  return { messages, addMessage, inputConfig, setInputConfig };
}