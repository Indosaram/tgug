# TGUG - Typora-Ghost image Uploader in Go

> This app is being actively developed. Currently tgug only supports Mac OS and Linux, but will support Windows.



In any chance if you want to know why I developed this, check [this](https://devbull.xyz/tgug/) out. (It's written in Korean.)



## Installation

`tgug` will store authentication file `auth.json` in your home directory, i.e. `~/.tgug` for mac and linux. If you start the program for the first time, it promts you to type information of your Ghost blog.

- Domain
- Username (email)
- Password

Clone the repo and go to the directory.

```bash
git clone https://github.com/Indosaram/tgug
cd tgug
```

### MacOS and Linux

This command will install binary in `$GOBIN` directory. If you haven't set your `$GOBIN`, it will install in `~/go/bin`.  

```bash
# make a directory reserved for tgug and install in GOBIN path
sudo mkdir ~/.tgug
GOBIN=/usr/local/bin go install
```

### Windows

In Windows, you need to `.tgug` folder in your home directory. (i.e. `C:\Users\USERNAME\.tgug`) Next, type following command in shell.
```cmd
go install
```


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

