import { Box, Container } from "@mui/material";

export default function ChatContainer({ children }: { children: React.ReactNode }) {
    return (
        <Container maxWidth="lg">
            <Box
                sx={{
                    height: '100vh',
                    display: 'flex',
                    flexDirection: 'column',
                    p: 2,
                }}
            >
                {children}
            </Box>
        </Container>
    )
}