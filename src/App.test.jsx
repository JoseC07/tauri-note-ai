import React from 'react';
import { render } from '@testing-library/react'
import { screen } from '@testing-library/dom';
import { describe, expect, it } from 'vitest';
import '@testing-library/jest-dom';
import App from './App.jsx';

describe('App', () => {
  it('renders welcome message', () => {
    render(<App />);
    expect(screen.getByText(/Welcome to Tauri \+ React/i)).toBeInTheDocument();
  });
}); 