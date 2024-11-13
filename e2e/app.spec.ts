import { test, expect } from '@playwright/test';

test('microphone permission can be requested', async ({ page }) => {
  await page.goto('http://localhost:1420');
  
  // Wait for the permission request to complete
  await page.waitForResponse('http://localhost:8080/api/audio/permission');
  
  // Check if permission was granted
  await expect(page.getByRole('button', { name: /start recording/i })).toBeEnabled();
});

test('recording can be started and stopped', async ({ page }) => {
  await page.goto('http://localhost:1420');
  
  // Start recording
  const startButton = page.getByRole('button', { name: /start recording/i });
  await startButton.click();
  
  // Verify recording state
  await expect(page.getByText(/recording in progress/i)).toBeVisible();
  
  // Stop recording
  const stopButton = page.getByRole('button', { name: /stop recording/i });
  await stopButton.click();
  
  // Verify recording stopped
  await expect(page.getByText(/recording in progress/i)).not.toBeVisible();
});

test('backend health check', async ({ request }) => {
  const response = await request.get('http://localhost:8080/health');
  expect(response.ok()).toBeTruthy();
  const body = await response.json();
  expect(body.status).toBe('ok');
}); 