import { Box, IconButton } from "@mui/material";

import SendIcon from '@mui/icons-material/Send';

interface ChatSendButtonProps {
    handleSend: () => void;
    disabled?: boolean;
}

export default function ChatSendButton({ handleSend, disabled }: ChatSendButtonProps) {
    return (
        <Box sx={{ flex: 0.4, display: 'flex', justifyContent: 'center', alignItems: 'center' }}>
            <IconButton
                color="primary"
                aria-label="send message"
                sx={{ borderRadius: '50%' }}
                onClick={handleSend}
                disabled={disabled}
            >
                <SendIcon />
            </IconButton>
        </Box>
    )
}