import { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import { Box, Button, Textarea, Spinner, Heading } from "@chakra-ui/react";
import { BASE_URL } from "./App";



const Summarizer = () => {
  const [inputText, setInputText] = useState("");
  const [outputText, setOutputText] = useState("");
  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ["summary"],
    queryFn: async () => {
      try {
        const response = await fetch(BASE_URL + "/summarize", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ text: inputText }), 
        });

        if (!response.ok) throw new Error("Failed to fetch summary");
        return response.json();
      } catch (error) {

      }
    },
    enabled: false, 
  });

  // Update textarea when data is available
  useEffect(() => {
    if (data) {
      setOutputText(data.summary);
    }
  }, [data]);

  return (
    <Box maxW="600px" mx="auto" mt={8}>
      <Heading as='h3' size='lg' mb={4}>
        AI Powered Text Summanizer
      </Heading>
      <Textarea
        value={inputText}
        placeholder="Input text here..."
        onChange={(e) => setInputText(e.target.value)}
        size="lg"
        rows={6}
      />
      <Button ml={"auto"} mt={4} onClick={() => refetch()} colorScheme="blue" mb={4}>
        Generate Summary
      </Button>

      {isLoading && <Spinner size="lg" color="blue.500" />}

      {isError && <Box color="red.500">Error fetching summary</Box>}

      <Textarea
        value={outputText}
        onChange={(e) => setOutputText(e.target.value)}
        placeholder="Your summarized text will appear here..."
        size="lg"
        rows={6}
      />
    </Box>
  );
};

export default Summarizer;