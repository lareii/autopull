<img width=200 src="./assets/autopull.png">

<img alt="GitHub contributors" src="https://img.shields.io/github/contributors/lareii/autopull">
<img alt="GitHub Repo stars" src="https://img.shields.io/github/stars/lareii/autopull?color=yellow">
<img alt="GitHub repo size" src="https://img.shields.io/github/repo-size/lareii/autopull?color=limegreen">

🔄 A real-time Git auto-sync tool.

<hr>

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