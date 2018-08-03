+++
title = "excerpt"
description = "The Excerpt shortcode is used to mark a part of a page's content for re-use."
+++

The Excerpt shortcode is used to mark a part of a page's content for re-use. Defining an excerpt enables other shortcodes, such as the excerpt-include shortcode, to display the marked content elsewhere.

{{%alert warning%}}You can only define one excerpt per page. In other words, you can only add the Excerpt shortcode once to a page.{{%/alert%}}


## Usage

| Parameter | Default | Description |
|:--|:--|:--|
| hidden | "false" | Controls whether the page content contained in the Excerpt shortcode placeholder is displayed on the page.{{%alert warning%}}Note that this option affects only the page that contains the Excerpt shortcode. It does not affect any pages where the content is reused.{{%/alert%}} |

## Demo

	{{%/*excerpt*/%}}
	Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
	tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
	quis nostrud exercitation **ullamco** laboris nisi ut aliquip ex ea commodo
	consequat. Duis aute irure dolor in _reprehenderit in voluptate_
	cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
	proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
	{{%/* /excerpt*/%}}

{{%excerpt%}}
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation **ullamco** laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in _reprehenderit in voluptate_
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{% /excerpt%}}



{{%alert info%}}See re use example with [excerpt-include shortcode]({{%relref "shortcodes/excerpt-include.md"%}}){{%/alert%}}