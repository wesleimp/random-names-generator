import React, { useState } from "react";
import { Container, Button, Name } from "./styles"
import api from "../../services/api"

function Main() {
  const [name, setName] = useState("")
  const opts = [
    {value: "animals", description: "Animals" },
    {value: "fruits", description: "Fruits" },
    {value: "heroes", description: "Heroes" },
    {value: "vegetables", description: "Vegetables" },
  ]

  function handleButtonClick(value) {
    api.get(`/${value}`)
      .then(res => setName(res.data))
  }

  return (
    <Container>
      <Name>{name}</Name>

      <div>
        {opts.map((o, i) => 
          <Button key={i} onClick={() => handleButtonClick(o.value)}>
            {o.description}
          </Button>
        )}
      </div>
    </Container>
  )
}

export default Main