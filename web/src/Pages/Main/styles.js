import styled from "styled-components";

export const Container = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  flex-direction: column;

  div {
    width: 800px;
  }
`
export const Name = styled.span`
  font-size: 42px;
  margin: 55px 0;
`
export const Button = styled.button`
  width: 200px;
  height: 45px;
  margin: 15px 15px;
  border: 1px solid #ffffffe6;
  border-radius: 2px;
  font-size: 16px;
  font-weight: bold;
  color: #1f2833;
  font-family: Roboto, sans-serif;
  cursor: pointer;
`