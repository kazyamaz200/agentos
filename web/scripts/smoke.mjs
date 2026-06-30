import { mkdir } from 'node:fs/promises'
import { chromium } from 'playwright'
import { preview } from 'vite'

const port = 4173

async function checkViewport(browser, name, width, height) {
  const page = await browser.newPage({ viewport: { width, height } })
  await page.goto(`http://127.0.0.1:${port}/`, { waitUntil: 'networkidle' })
  const overflow = await page.evaluate(() => document.documentElement.scrollWidth - window.innerWidth)
  if (overflow > 1) {
    throw new Error(`${name} has horizontal overflow of ${overflow}px`)
  }
  await page.screenshot({ path: `../tmp/webui-${name}.png`, fullPage: true })
  await page.close()
}

let server
let browser
try {
  await mkdir('../tmp', { recursive: true })
  server = await preview({ preview: { host: '127.0.0.1', port, strictPort: true } })
  browser = await chromium.launch()
  await checkViewport(browser, 'mobile', 390, 844)
  await checkViewport(browser, 'desktop', 1440, 1000)
} finally {
  if (browser) await browser.close()
  await new Promise((resolve, reject) => {
    const httpServer = server?.httpServer
    if (!httpServer) {
      resolve()
      return
    }
    httpServer.close((err) => (err ? reject(err) : resolve()))
  })
}
