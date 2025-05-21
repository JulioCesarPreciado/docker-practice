import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import ThemeWrapper from './theme/ThemeWrapper'
import './main.css'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeWrapper />
  </StrictMode>
)
