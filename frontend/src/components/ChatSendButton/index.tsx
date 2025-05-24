import { Box, IconButton } from "@mui/material";

import SendIcon from '@mui/icons-material/Send';

interface ChatSendButtonProps {
    handleSend: () => void;
}

export default function ChatSendButton({ handleSend }: ChatSendButtonProps) {
    return (
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
    )
}