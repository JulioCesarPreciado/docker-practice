import './App.css';
import ChatActions from './components/ChatActions';
import ChatContainer from './components/ChatContainer';

import { useSocket } from './hooks/useSocket';
import { CardContent } from '@mui/material';
import ChatMessageList from './components/ChatMessageList';
import { parseSocketMessage } from './utils/parseSocketMessage';
import { useInputManager } from './hooks/useInputManager';

function App() {
  const { messages, addMessage, inputConfig, setInputConfig } = useInputManager();

  const { sendMessage } = useSocket((newMsg: string) => {
    const { message, inputConfig } = parseSocketMessage(newMsg);
    addMessage(message, 'system');
    setInputConfig(inputConfig);
  });

  const handleSend = (msg: string) => {
    sendMessage(msg);
    addMessage(msg, 'user');
  };

  return (
    <ChatContainer title="Realtime Chat App 📱">
      <CardContent
        sx={{
          flex: 1,
          backgroundColor: 'background.paper',
        }}
      >
        <ChatMessageList messages={messages} />
      </CardContent>
      <ChatActions onSend={handleSend} inputConfig={inputConfig} />
    </ChatContainer>
  )
}

export default App
