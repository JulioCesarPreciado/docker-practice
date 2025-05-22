import {
    Box,
    Card,
    Container,
    Typography,
} from '@mui/material'

interface ChatContainerProps {
    title?: string;
    children: React.ReactNode;
}

export default function ChatContainer({ title, children }: ChatContainerProps) {
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
                {title && (
                    <Typography
                        variant="h4"
                        component="h1"
                        textAlign="center"
                        fontWeight="bold"
                        gutterBottom
                    >
                        {title}
                    </Typography>
                )}
                <Card
                    sx={{
                        flexGrow: 1,
                        overflow: 'auto',
                        display: 'flex',
                        flexDirection: 'column',
                        boxShadow: 3,
                        borderRadius: 4,
                    }}
                >
                    {children}
                </Card>
            </Box>
        </Container>
    )
}