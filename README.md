# sloop

![CocoaPods](https://img.shields.io/github/license/gadd-ferreira/sloop.svg?style=for-the-badge) ![GitHub tag](https://img.shields.io/github/tag/gadd-ferreira/sloop.svg?style=for-the-badge)

Loops the cursor around the screen when it reaches an horizontal edge


## Installation
```bash
curl -L https://github.com/gadd-ferreira/sloop/releases/download/v1.0.0/sloop-v1.0.0 -o /tmp/sloop
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
