# gx

⚠️ This is a work in progress and some things might not work as expected. For example extracting whole folders, private repos, etc. ⚠️

![gx](https://github.com/darox/gx/blob/main/assets/gx.gif?raw=true)

It stands for git extraction. The idea behind this CLI is to make it easier to selectively download specific files and folders from a git repository into an existing folder structure.

## Installation

```
git clone https://github.com/darox/gx.git && cd gx && make && mv gx $GOPATH/bin/
```

Check if it worked:
```
gx -h
```

I'm planning to add homebrew support soon and maybe even precompiled binaries.

## Usage

```bash
gx -s <path> -t <path> -u <url> -b <branch>
```

