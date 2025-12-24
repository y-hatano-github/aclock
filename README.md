# 🕒 aclock
A colorful analog clock for the terminal, built with Golang.  

![aclock](aclock.gif)

## ⚙️ Features

- ⏰ Colorful analog clock rendered directly in your terminal
- 🎨 Customizable color scheme for all clock elements
- 🧩 No dependencies — just a single binary
- 🖥️ Cross-platform binaries (Windows, macOS, Linux)

## 💾 Download 
Pre-built binaries are available for Windows, macOS, and Linux.

👉 Get the latest release here:  
https://github.com/y-hatano-github/aclock/releases/latest

## 🚀 Quick start
### 🐧 Linux
```bash
wget https://github.com/y-hatano-github/aclock/releases/latest/download/aclock_linux_amd64.tar.gz
tar -xzvf aclock_linux_amd64.tar.gz
mv aclock /usr/local/bin/
aclock
```
### 🍎 macOS
```bash
curl -LO https://github.com/y-hatano-github/aclock/releases/latest/download/aclock_darwin_amd64.tar.gz
tar -xzvf aclock_darwin_amd64.tar.gz
sudo mv aclock /usr/local/bin/
aclock
```
### 🪟 Windows
```powershell
Invoke-WebRequest -OutFile aclock_windows_amd64.zip https://github.com/y-hatano-github/aclock/releases/latest/download/aclock_windows_amd64.zip
Expand-Archive aclock_windows_amd64.zip
.\aclock.exe
```

## 📘 Usage
```text
Usage: aclock [flags]

A colorful analog clock rendered in your terminal.

You can customize the clock's appearance by specifying colors for:

    background
    face
    frame
    hour/minute/second hands
    pivot point
    tick marks

Colors available:

    red, orange, yellow, lime, green, cyan, sky, blue, indigo,
    purple, magenta, pink, scarlet, brown, gray, black, white,
    system (uses terminal's background color)

Example:

    aclock --face blue --frame white --hour yellow --min green --sec red

Controls:

    ESC, Ctrl+C    Exit the application

Flags:
  -h, --help                   Show context-sensitive help.
      --background="system"    Background color of the terminal area surrounding the clock.
      --face="gray"            Color of the clock face. This is the filled area inside the frame.
      --frame="white"          Color of the outer frame of the clock.
      --hour="blue"            Color of the hour hand.
      --min="green"            Color of the minute hand.
      --sec="cyan"             Color of the second hand.
      --piv="white"            Color of the pivot point.
      --tick="red"             Color of the tick marks.
```