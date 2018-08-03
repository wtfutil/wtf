+++
title = "attachments"
description = "The Attachments shortcode displays a list of files attached to a page."
+++

The Attachments shortcode displays a list of files attached to a page.
Example :
{{%alert success%}}{{%attachments  /%}}{{%/alert%}}

## Usage

The shortcurt lists files found in a **specific folder**.
Currently, it support two implementations for pages

1. If your page is a markdown file, attachements must be place in a **folder** named like your page and ending with **.files**.

    > * content
    >   * _index.md
    >   * page.files
    >      * attachment.pdf
    >   * page.md

2. If your page is a **folder**, attachements must be place in a nested **'files'** folder.

    > * content
    >   * _index.md
    >   * page
    >      * index.md
    >      * files
    >          * attachment.pdf

That's all !

{{%alert info%}}**Tip** : Look at this documentation source code on github{{%/alert%}}

### parameters

| Parameter | Default | Description |
|:--|:--|:--|
| title | "Attachments" | List's title  |
| pattern | ".*" | A regular expressions, used to filter the attachments by file name. <br/><br/>{{%alert warning%}}The **pattern** parameter value must be [regular expressions](https://en.wikipedia.org/wiki/Regular_expression).

For example:

* To match a file suffix of 'jpg', use **.*jpg** (not *.jpg).
* To match file names ending in 'jpg' or 'png', use **.*(jpg|png)**

{{%/alert%}}|


## Demo
### List of attachments ending in pdf or mp4

    {{%/*attachments title="Related files" pattern=".*(pdf|mp4)"/*/%}}

renders as

{{%attachments title="Related files" pattern=".*(pdf|mp4)"/%}}

