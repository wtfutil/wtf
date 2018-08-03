+++
title = "Installation"
description = ""
weight = 1
+++

{{% alert theme="warning" %}}HUGO **v0.32** minimum required to use this theme{{%/alert%}}

The following steps are here to help you initialize your new website. If you don’t know Hugo at all, we strongly suggest you to train by following this [great documentation for beginners](https://gohugo.io/overview/quickstart/).
<!--more-->

## Installation

We assume that all changes to Hugo content and customizations are going to be tracked by git (GitHub, Bitbucket etc.). Develop locally, build on remote system.

Before start real work:

1. Initialize Hugo
2. Install DocDock theme
3. Configure DocDock and Hugo

### Prepare empty Hugo site

Create empty directory, which will be root of your Hugo project. Navigate there and let Hugo to create minimal required directory structure:
```
$ hugo new site .
```
AFTER that, initialize this as git directory where to track further changes
```
$ git init
```

Next, there are at least three ways to install DocDock (first recommended):

1. **As git submodule**
2. As git clone
3. As direct copy (from ZIP)

Navigate to your themes folder in your Hugo site and use perform one of following scenarios.

### 1. Install DocDock as git submodule

DocDock will be added like a dependency repo to original project. When using CI tools like Netlify, Jenkins etc., submodule method is required, or you will get `theme not found` issues. Same applies when building site on remote server trough SSH.

If submodule is no-go, use 3rd option.

On your root of Hugo execute:

```
$ git submodule add https://github.com/vjeantet/hugo-theme-docdock.git themes/docdock
```
Next initialize submodule for parent git repo:

```
$ git submodule init
$ git submodule update
```

Now you are ready to add content and customize looks. Do not change any file inside theme directory.

If you want to freeze changes to DocDock theme itself and use still submodules, fork private copy of DocDock and use that as submodule. When you are ready to update theme, just pull changes from origin to your private fork.

### 2. Install DocDock simply as git clone

This method results that files are checked out locally, but won't be visible from parent git repo. Probably you will build site locally with `hugo` command and use result from `public/` on your own.

```
$ git clone https://github.com/vjeantet/hugo-theme-docdock.git themes/docdock
```


### 3. Install DocDock from ZIP

All files from theme will be tracked inside parent repo, to update it, have to override files in theme. [{{%icon download%}} download following zip](https://github.com/vjeantet/hugo-theme-docdock/archive/master.zip) and extract inside `themes/`.

```
https://github.com/vjeantet/hugo-theme-docdock/archive/master.zip
```
Name of theme in next step will be `hugo-theme-docdock-master`, can rename as you wish.

## Configuration

[Follow instructions here]({{%relref "configuration.md"%}})
