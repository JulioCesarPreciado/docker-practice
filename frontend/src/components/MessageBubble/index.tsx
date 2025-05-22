import { Box, Typography, Paper } from '@mui/material'

interface MessageBubbleProps {
  message: string
  align?: 'left' | 'right'
}

export default function MessageBubble({ message, align = 'left' }: MessageBubbleProps) {
  return (
    <Box
      display="flex"
      justifyContent={align === 'left' ? 'flex-start' : 'flex-end'}
      mb={1}
    >
      <Paper
        elevation={3}
        sx={{
          px: 2,
          py: 1,
          borderRadius: 3,
          maxWidth: '70%',
          boxShadow: 3,
        }}
      >
        <Typography variant="body1">{message}</Typography>
      </Paper>
    </Box>
  )
}
