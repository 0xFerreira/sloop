# sloop

![License](https://img.shields.io/github/license/gadd-ferreira/sloop.svg?style=for-the-badge) ![Tag](https://img.shields.io/github/release/gadd-ferreira/sloop.svg?style=for-the-badge) ![Coverage](https://img.shields.io/badge/coverage-35.7%25-yellow.svg?style=for-the-badge)

Loops the cursor around the screen when it reaches an horizontal edge


## Installation

### Dependencies
```
sudo apt-get install libx11-dev xorg-dev libxtst-dev libpng++-dev   

sudo apt-get install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt-get install libxkbcommon-dev

sudo apt-get install xsel xclip
```

### Download
```bash
curl -L https://github.com/gadd-ferreira/sloop/releases/download/v1.1.0/sloop-v1.1.0 -o /tmp/sloop
chmod +x /tmp/sloop
sudo mv /tmp/sloop /usr/local/bin/
```


## Usage



Manually with



```bash
sloop&
```



Automatically with .xinitrc
```bash
echo 'sloop&' >> ~/.xinitrc
```
