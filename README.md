# sloop

![License](https://img.shields.io/github/license/gadd-ferreira/sloop.svg?style=for-the-badge) ![Tag](https://img.shields.io/github/release/gadd-ferreira/sloop.svg?style=for-the-badge) ![Coverage](https://img.shields.io/badge/coverage-42.9%25-green.svg?style=for-the-badge)

Loops the cursor around the screen when it reaches an horizontal edge


## Installation

### Dependencies
Most of these should already be installed

For Ubuntu
```
sudo apt-get install libx11-dev xorg-dev libxtst-dev libpng++-dev   

sudo apt-get install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt-get install libxkbcommon-dev

sudo apt-get install xsel xclip
```

### Direct Download (unix x86-64)
```bash
curl -L https://github.com/gadd-ferreira/sloop/releases/download/v1.2.0/sloop-v1.2.0 -o /tmp/sloop
chmod +x /tmp/sloop
sudo mv /tmp/sloop /usr/local/bin/
```

### Using go
```bash
go get github.com/gadd-ferreira/sloop
go install github.com/gadd-ferreira/sloop
```

## Usage

### Flags
```bash
Usage of sloop:
  -activationMargin int
    	Jump activation margin (default 10)
  -d    Daemonize and return control
  -hz int
    	Pointer position check frequency (default 50)
  -jumpDelta int
    	Jump delta distance from activation margin (default 1)
```

Manually with
```bash
sloop -d&
```

Automatically with X
```bash
echo 'sloop -d &' >> ~/.xinitrc
```
