# ⚠️ UNDER DEVELOPMENT 

<img src="./assets/autopull.png">

![GitHub contributors](https://img.shields.io/github/contributors/lareii/autopull)
![GitHub Repo stars](https://img.shields.io/github/stars/lareii/autopull?color=yellow)
![GitHub repo size](https://img.shields.io/github/repo-size/lareii/autopull?color=limegreen)

🔄 A real-time Git auto-sync tool.

A simple HTTP server written in Go that monitors pushes to your Git repository in real-time, and pulls them from your Git repository.

## Getting Started
### Prerequisites
- Go >= 1.19
- Git

### Install autopull
You can install it via `go install` command.
```
go install github.com/lareii/autopull
```
### Usage
```
autopull <dir> <repo_url> <secret>
```