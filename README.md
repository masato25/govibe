# govibe

## Overview
Automatically detects the `allow_button.png` image on your screen, moves the mouse to the center of the detected image, and clicks it.

## Installation

### 1. Install Homebrew (skip if already installed)
https://brew.sh/

### 2. Install dependencies
```sh
brew install opencv pkg-config
```

### 3. Install Go packages
```sh
go get github.com/go-vgo/robotgo

go get github.com/kbinani/screenshot

go get gocv.io/x/gocv
```
## Intended Usage
- Works with VS Code, Terminal, or any Go-supported development environment
- Code is well-commented and structured for optimal use with GitHub Copilot

### 4. macOS Accessibility Permissions
	- Without this permission, the program cannot perform automatic clicking or mouse movement

## Usage

1. Place `allow_button.png` in the same directory as `main.go`
2. Run:
```sh
go run main.go
```
3. The program will automatically detect the number of screens and continuously scan the primary display (display 0)
4. When the image is detected, the mouse will move to the center of the image and click

## Troubleshooting

### 1. Error: `pkg-config` or `opencv4.pc` not found
```sh
export PKG_CONFIG_PATH="/opt/homebrew/lib/pkgconfig:$PKG_CONFIG_PATH"
```

### 2. No output or panic

### 3. Mouse does not move or click

### 4. Incorrect click position

## Notes
