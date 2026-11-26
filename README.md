<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:0f172a,100:0f766e&height=170&section=header&text=go%20http%20mux%20lab&fontSize=36&fontColor=ffffff&animation=twinkling" alt="Animated wave header"/></p>
<p align="center"><img src="https://readme-typing-svg.demolab.com?font=Fira+Code&weight=600&size=20&duration=2800&pause=900&color=14B8A6&center=true&vCenter=true&width=760&lines=go%20http%20mux%20lab%20%7C%20inspect%2C%20understand%2C%20improve" alt="Animated title"/></p>
<p align="center"><img src="https://img.shields.io/github/license/CyberLoop-dev/go-http-mux-lab?style=for-the-badge" alt="License"/> <img src="https://img.shields.io/github/languages/top/CyberLoop-dev/go-http-mux-lab?style=for-the-badge" alt="Language"/> <img src="https://img.shields.io/github/stars/CyberLoop-dev/go-http-mux-lab?style=for-the-badge" alt="Stars"/> <img src="https://img.shields.io/github/last-commit/CyberLoop-dev/go-http-mux-lab?style=for-the-badge" alt="Last commit"/></p>
<p align="center"><img src="assets/mark.svg" alt="go-http-mux-lab project mark" width="520"/></p>

## 🧭 Overview

`go-http-mux-lab` is a tiny localhost HTTP server built around Go 1.22's method-aware `http.ServeMux` patterns. Explore exact paths, `{id}` wildcards, trailing `{path...}` wildcards, and automatic method matching.

## ✨ Features

- 🛣️ Demonstrates method-aware route patterns such as `GET /users/{id}`
- 🔎 Reads wildcard values with `Request.PathValue`
- 🎯 Includes an exact root route using `GET /{$}`
- 🚦 Shows method-not-allowed behavior for mismatched methods
- 🧪 Includes `httptest` coverage for the route handlers

## 🧪 Demo

Start the server and request `/health`, `/users/42`, or `/files/docs/guide.txt`. The handlers return small JSON responses. No hosted demo is published.

## 🚀 Quick Start

```bash
go run .
# In another terminal:
curl http://127.0.0.1:8080/users/42
curl http://127.0.0.1:8080/files/docs/guide.txt
go test ./...
```
Requires Go 1.22 or newer. The server binds to localhost by default.

## 🧰 Tech Stack

- **Stack:** Go 1.22+, net/http, encoding/json, and net/http/httptest
- **Approach:** small, local-first tooling with clear output

## 🗂️ Folder Structure

```text
go-http-mux-lab/
├── assets/mark.svg
├── main.go
├── main_test.go
├── go.mod
├── .gitignore
├── LICENSE
└── README.md
```

## 🛣️ Roadmap

- Add a route precedence visualization mode
- Add a small custom 404 handler and request logging example

## 🤝 Contributing

Issues and focused pull requests are welcome. Explain the use case, keep changes small, and add tests for behavior changes. Run the project using the Quick Start instructions before submitting.

## 📄 License

Released under the [MIT License](LICENSE).

<p align="center"><img src="https://capsule-render.vercel.app/api?type=waving&color=0:0f766e,100:0f172a&height=120&section=footer&animation=twinkling" alt="Animated wave footer"/></p>
