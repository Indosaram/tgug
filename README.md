# TGUG - Typora-Ghost image Uploader in Go

> This app is being actively developed. Currently tgug only supports Mac OS and Linux, but will support Windows.



In any chance if you want to know why I developed this, check [this](https://devbull.xyz/tgug/) out. (It's written in Korean.)



## Installation

This command will install binary in `$GOBIN` directory. If you haven't set your `$GOBIN`, it will install in `~/go/bin`.  

```bash
git clone https://github.com/Indosaram/tgug
cd tgug

# make a directory reserved for tgug and install in GOBIN path
sudo mkdir ~/.tgug
GOBIN=/usr/local/bin go install
```

`tgug` will store authentication file `auth.json` in `~/.tgug`.  If you first start the program, it promts you to type information of your Ghost blog.

- Domain
- Username (email)
- Password



## Usage

Check out [this page](https://devbull.xyz/how-to-setup-tgug-with-typora/) for setup tgug for your Typora.

For command line usage, you can use this by typing,

```bash
tgug -f=image1.jpg,image2.png,...
```



## TODOs

- [ ] Documentation
- [ ] Exception control
  - [ ] authentication info with Regex
  - [ ] Invalid file path
- [ ] Encryption of `auth.json` file
- [ ] Seamless integration with Typora
  - [ ] Script for automatic integration

