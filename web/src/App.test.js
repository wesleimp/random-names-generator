import React from 'react'
import { render } from '@testing-library/react'

import App from './App'

test('renders the application', () => {
  const text = 'Hello App'
  const { queryByText } = render(<App />)
  const element = queryByText(text)
  expect(element.textContent).toBe(text)
})