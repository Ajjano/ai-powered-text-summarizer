import { useState } from 'react'
import { Container, Stack } from '@chakra-ui/react'
import Form from './Form'

export const BASE_URL = "http://localhost:8080"
function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <Stack h="100vh">
       <Container>
          <Form/>
       </Container>
    </Stack>
    </>
  )
}

export default App
