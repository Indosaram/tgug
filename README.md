# TGUG - Typora-Ghost image Uploader in Go



# Installation

This command will install binary in `$GOBIN` directory. If you haven't set your `$GOBIN`, it will install in `~/go/bin`.

```bash
GOBIN=/usr/local/bin go install
```



# Usage

Type,

```bash
tgug -f=image1.jpg,image2.png,...
```



If you first start the program, it promts you to type information of your Ghost blog.

- Domain
- Username (email)
- Password

This information will be stored in `$GOPATH` as `auth.json`.

# TODOs

- [ ] Documentation
- [ ] Exception control for authentication info with Regex
- [ ] Encryption of `auth.json` file
- [ ] Seamless integration with Typora
  - [ ] Script for automatic integration