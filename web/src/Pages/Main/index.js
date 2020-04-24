import React from "react";
import { Container, Button, Name } from "./styles"

const Main = () => {
  const opts = [
    {value: "animals", description: "Animals" },
    {value: "fruits", description: "Fruits" },
    {value: "heroes", description: "Heroes" },
    {value: "vegetables", description: "Vegetables" },
  ]

  function handleButtonClick(endpoint) {
    alert(endpoint)
  }

  return (
    <Container>
      <Name>lkdfjsdjfsdklfjsldkfj</Name>

      <div>
        {opts.map(o => 
          <Button onClick={() => handleButtonClick(o.value)}>
            {o.description}
          </Button>
        )}
      </div>
    </Container>
  )
}

export default Main