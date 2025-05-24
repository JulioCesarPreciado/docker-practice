import MessageBubble from "../MessageBubble";

type Message = {
    text: string;
    from: 'user' | 'system';
};

export default function ChatMessageList({ messages }: { messages: Message[] }) {
    return (
        <>
            {messages.map((msg, idx) => (
                <MessageBubble
                    key={idx}
                    message={msg.text}
                    align={msg.from === 'user' ? 'right' : 'left'}
                />
            ))}
        </>
    );
}