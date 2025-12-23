# 🕒 aclock
A colorful analog clock for the terminal, built with Golang.

## 💾 Download 
Pre-built binaries are available for Windows, macOS, and Linux.

👉 Get the latest release here:  
https://github.com/y-hatano-github/aclock/releases/latest


## 🚀 Quick start
```bash
wget https://github.com/y-hatano-github/aclock/releases/latest/download/aclock_linux_amd64.tar.gz
tar -xzvf aclock_linux_amd64.tar.gz
mv aclock /usr/local/bin/
aclock
```
## ⚙️ Features

- ⏰ Colorful analog clock rendered directly in your terminal
- 🎨 Customizable color scheme for all clock elements
- 🧩 No dependencies — just a single binary
- 🖥️ Cross-platform binaries (Windows, macOS, Linux)

## 📘 Usage
```text
Usage: aclock [flags]

You can customize the clock's appearance by specifying colors for:

    background
    face
    frame
    hour/minute/second hands
    pivot point
    tick marks

Colors available:

    black, red, green, yellow, blue, magenta, cyan, white
    gray, purple, brown, pink, orange
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