import {
  Box,
  CardActions,
  TextField,
  IconButton,
} from '@mui/material'

import SendIcon from '@mui/icons-material/Send';

export default function ChatActions() {
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
              sx={{ flex: 9.6 }}
            />
            <Box sx={{ flex: 0.4, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
              <IconButton
                color="primary"
                aria-label="send message"
                sx={{ borderRadius: '50%' }}
              >
                <SendIcon />
              </IconButton>
            </Box>
          </CardActions>
    )
}
