import { useEffect, useRef } from 'react'

export function useSocket(onMessage: (msg: string) => void) {
  const socketRef = useRef<WebSocket | null>(null)
  const onMessageRef = useRef(onMessage)

  // Update ref if onMessage callback changes
  useEffect(() => {
    onMessageRef.current = onMessage
  }, [onMessage])

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:9000/ws')
    socketRef.current = socket

    socket.onopen = () => {
      console.log('✅ WebSocket connected')
    }

    socket.onmessage = (event) => {
      onMessageRef.current(event.data)
    }

    socket.onerror = (error) => {
      console.error('❌ WebSocket error:', error)
    }

    socket.onclose = () => {
      console.log('🔌 WebSocket disconnected')
    }

    return () => {
      socket.close()
    }
  }, [])

  const sendMessage = (msg: string) => {
    if (socketRef.current?.readyState === WebSocket.OPEN) {
      socketRef.current.send(msg)
    }
  }

  return { sendMessage }
}
