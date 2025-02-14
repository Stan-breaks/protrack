# 🚀 NimbleStack

_A Modern Go + Templ + Tailwind CSS Starter Template with HTMX & Alpine.js_

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)  
![Templ](https://img.shields.io/badge/Templ-0.2+-blue)  
![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3.3+-06B6D4?logo=tailwind-css)  
![HTMX](https://img.shields.io/badge/HTMX-1.9+-red)  
![Alpine.js](https://img.shields.io/badge/Alpine.js-3.13+-8BC0D0)

**NimbleStack** is a lightning-fast, full-stack starter template designed for developers who want to build modern web apps with minimal boilerplate. It combines Go's performance, Templ's server-side rendering, Tailwind CSS's styling magic, and the dynamic interactivity of HTMX + Alpine.js.

---

## 🌟 Features

- **Go Backend**: Blazing-fast API and server logic with Go.
- **Templ Templates**: Clean, type-safe HTML templating for server-rendered UIs.
- **Tailwind CSS**: Utility-first styling with JIT compilation for rapid design.
- **HTMX**: Dynamic content updates without writing JavaScript.
- **Alpine.js**: Lightweight reactivity for client-side interactions.
- **Zero JS Framework Overhead**: Keep bundles tiny and performance high.

---

## 🛠️ Why NimbleStack?

- **Rapid Development**: Start coding features instead of configuring tools.
- **Minimal Boilerplate**: Focus on your app, not setup.
- **Full-Stack Ready**: Backend + frontend in one cohesive stack.
- **Modern UI/UX**: Tailwind + Alpine.js for polished, responsive interfaces.

---

## 🚀 Getting Started

### Prerequisites

- Go 1.21+
- Node.js 18+ & pnpm
- Git

### Installation

1. Clone the repo:

   ```bash
   git clone https://github.com/stan-breaks/nimblestack.git
   cd nimblestack
   ```

2. Install dependencies:

   ```bash
   pnpm install
   pnpm run build:css
   ```

3. Generate Templ views:

   ```bash
   temp generate ./views/
   ```

4. Start the dev server:

   ```bash
   go run .
   ```

   Open http://localhost:8080 in your browser.

🛠️ Development Workflow

- Run the server: go run main.go (auto-reload with air if configured)

- Build Tailwind: pnpm build:css (or pnpm watch:css for live updates)

- Templ Templates: Write .templ files in views/ and regenerate with templ generate ./views/

- HTMX + Alpine.js: Add interactivity directly in HTML with hx-\* attributes and x-data

📂 Project Structure

```
nimblestack/
├── public/ # Static assets (Tailwind CSS, Alpine.js, HTMX)
├── views/ # Templ components and pages
├── handlers/ # Go HTTP handlers
├── go.mod # Go dependencies
├── package.json # Frontend tooling
├── main.go # Server entrypoint
└── tailwind.config.js # tailwind css config

```

📈 Roadmap

- Add authentication example (Go + HTMX)

- Dockerize for easy deployment

- Prebuilt UI components (modals, forms, etc.)

- Deployment guides for Fly.io/Vercel

🤝 Contributing

Found a bug or have an idea? Open an issue or submit a PR!

📜 License

MIT License - see LICENSE.

Built with ❤️ by Stan Breaks
