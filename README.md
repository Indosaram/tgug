# TGUG - Typora-Ghost image Uploader in Go



## Installation

This command will install binary in `$GOBIN` directory. If you haven't set your `$GOBIN`, it will install in `~/go/bin`.  

```bash
sudo mkdir ~/.tgug
GOBIN=/usr/local/bin go install
```

`tgug` will store authentication file `auth.json` in `~/.tgug`.  If you first start the program, it promts you to type information of your Ghost blog.

- Domain
- Username (email)
- Password



## Usage

Refer https://support.typora.io/Upload-Image/#use-current-filename--filepath-in-custom-commands for setting custom cli.

For command line usage, you can use this by typing,

```bash
tgug -f=image1.jpg,image2.png,...
```



## TODOs

- [ ] Documentation
- [ ] Exception control for authentication info with Regex
- [ ] Encryption of `auth.json` file
- [ ] Seamless integration with Typora
  - [ ] Script for automatic integration