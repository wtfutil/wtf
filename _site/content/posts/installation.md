---
title: "Installation"
date: 2018-05-18T09:59:40-07:00
draft: false
---

There are two ways to install WTF:

## From Source

Get this repo and install the dependencies:

```bash
go get -u github.com/senorprogrammer/wtf
cd $GOPATH/src/github.com/senorprogrammer/wtf
make install
make run
```
and that should probably do it.

## As a Binary

Grab the latest version from here:

```bash
https://github.com/senorprogrammer/wtf/releases
```

expand it, and `cd` into the resulting directory. Then run:

```bash
./wtf
```

and that should also do it.
