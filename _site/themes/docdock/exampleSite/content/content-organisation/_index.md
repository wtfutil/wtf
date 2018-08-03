+++
title = "Content Organisation"
description = ""
weight = 20
+++

With **Hugo**, pages are the core of your site. Organize your site like any other Hugo project. **Magic occurs with the nested sections implemention done in v0.22 of hugo (congrats @bep)**.

With docdock, **Each content page composes the menu**, they shape the structure of your website.

To link pages to each other, place them in a folders hierarchy

```
content
├── level-one
│   ├── level-two
│   │   ├── level-three
│   │   │   ├── level-four
│   │   │   │   ├── _index.md
│   │   │   │   ├── page-4-a.md
│   │   │   │   ├── page-4-b.md
│   │   │   │   └── page-4-c.md
│   │   │   ├── _index.md
│   │   │   ├── page-3-a.md
│   │   │   ├── page-3-b.md
│   │   │   └── page-3-c.md
│   │   ├── _index.md
│   │   ├── page-2-a.md
│   │   ├── page-2-b.md
│   │   └── page-2-c.md
│   ├── _index.md
│   ├── page-1-a.md
│   ├── page-1-b.md
│   └── page-1-c.md
├── _index.md
└── page-top.md
```


{{%alert info %}} **_index.md** is required in each folder, it's your "folder home page"{{%/alert%}}

### Add header to a menu entry

in the page frontmatter, add a `head` param to insert any HTML code before the menu entry:

example to display a "Hello"

	+++
	title = "Github repo"
	head ="<label>Hello</label> "
	+++



### Add icon to a menu entry

in the page frontmatter, add a `pre` param to insert any HTML code before the menu label:

example to display a github icon 

	+++
	title = "Github repo"
	pre ="<i class='fa fa-github'></i> "
	+++

![dsf](/menu-entry-icon.png?height=40px&classes=shadow)

<!-- ### Customize menu entry label

Add a `name` param next to `[menu.main]`

	+++
	[menu.main]
	parent = ""
	identifier = "repo"
	pre ="<i class='fa fa-github'></i> "
	name = "Github repo"
	+++ -->

<!-- ### Create a page redirector
Add a `url` param next to `[menu.main]`

	+++
	[menu.main]
	parent = "page"
	identifier = "page-images"
	weight = 23
	url = "/shortcode/image/"
	+++

{{%alert info%}}Look at the menu "Create Page/About images" which redirects to "Shortcodes/image{{%/alert%}}
 -->
### Order sibling menu/page entries

in your frontmatter add `weight` param with a number to order.

	+++
	title="My page"
	weight = 4
	+++

{{%info%}}add `ordersectionsby = "title"` in your config.toml to order menu entries by title{{%/info%}}


### Hide a menu entry

in your frontmatter add `hidden=true` param.

	+++
	title="My page"
	hidden = true
	+++


### Unfolded menu entry by default

One or more menuentries can be displayed unfolded by default. (like the "Getting start" menu entry  in this website)

in your frontmatter add `alwaysopen=true` param.
example :

```
title = "Getting start"
description = ""
weight = 1
alwaysopen = true
```

### Folder structure and file name

Content organization **is** your `content` folder structure.

### Homepage

Find out how to [customize homepage]({{%relref "homepage.md"%}}) 



