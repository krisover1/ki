import { test, expect, Page } from '@playwright/test';

const uniqueEmail = () => `test-${Date.now()}-${Math.random().toString(36).slice(2, 8)}@example.com`;
const TEST_PASSWORD = 'testpass123';

async function loginAsNewUser(page: Page): Promise<void> {
  const email = uniqueEmail();
  await page.goto('/register');
  await page.fill('#email', email);
  await page.fill('#password', TEST_PASSWORD);
  await page.fill('#passwordConfirm', TEST_PASSWORD);
  await page.click('button[type="submit"]');
  await page.waitForURL('/timer');
}

test('timervisning teller opp uten sideoppdatering', async ({ page }) => {
  await loginAsNewUser(page);

  await page.click('button:has-text("Start")');
  await expect(page.locator('#active-timer')).toBeVisible();

  const before = await page.locator('#active-timer').textContent();

  // Vent 3 sekunder og sjekk at visningen er endret
  await page.waitForTimeout(3000);

  const after = await page.locator('#active-timer').textContent();
  expect(after).not.toBe(before);
});
