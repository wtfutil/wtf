# Hugo docDock Theme

This repository contains a theme for [Hugo](https://gohugo.io/), based on 

* [Matcornic Learn theme](https://github.com/matcornic/hugo-theme-learn/).
* [facette.io](https://facette.io/)'s documentation style css (Facette is a great time series data visualization software)

Visit the [theme documentation](http://docdock.netlify.com/) to see what is going on. It is actually built with this theme.

# Main features

- Search
- **Unlimited menu levels**
- RevealJS presentation from markdown (embededed or fullscreen page)
- Attachments files
- List child pages
- Include segment of content from one page in another (Excerpt)
- Automatic next/prev buttons to navigate through menu entries
- Mermaid diagram
- Icons, Buttons, Alerts, Panels, Tip/Note/Info/Warning boxes
- Image resizing, shadow...
- Customizable look and feel


![Overview](https://github.com/vjeantet/hugo-theme-docdock/raw/master/images/tn.png)

## Installation

Check that your Hugo version is minimum `0.30` with `hugo version`. We assume that all changes to Hugo content and customizations are going to be tracked by git (GitHub, Bitbucket etc.). Develop locally, build on remote system.

To start real work:

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

All files from theme will be tracked inside parent repo, to update it, have to override files in theme. Download following zip and extract inside `themes/`.

```
https://github.com/vjeantet/hugo-theme-docdock/archive/master.zip
```
Name of theme in next step will be `hugo-theme-docdock-master`, can rename as you wish.

## Configure

Import sample config from sample site to Hugo root.

```
$ cp themes/docdock/exampleSite/config.toml .
```

Change following `config.toml` line as needed, depending on method above:
```
theme = "<hugo-theme-docdock-dir-name>"
```
Comment out following line, so default `themes/` will be used:

```
# themesdir = "../.."
```

#### (Bonus)
Create empty file `.gitkeep` inside `public/` and add following to `.gitignore`.  This way it will keep repo smaller and won't bring build result files and errors to remote checkout places:
```
/public/*
!/public/.gitkeep
```

### Preview site
```
$ hugo server
```
to browse site on http://localhost:1313

## Usage

- [Visit the documentation](http://docdock.netlify.com/)
- [Hugo docs](https://gohugo.io/getting-started/configuration/)
- [Git submodules](https://git-scm.com/docs/git-submodule)
