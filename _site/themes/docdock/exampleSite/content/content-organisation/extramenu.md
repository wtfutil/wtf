+++
title = "Extra menu entries"
date = "2017-04-29T18:36:24+02:00"
Weight=2
+++

You can define additional menu entries in the navigation menu without any link to content.

Edit the website configuration `config.toml` and add a `[[menu.shortcuts]]` entry for each link your want to add.


Example from the current website, **note the `pre` param** which allows you to insert HTML code and used here to separate content's menu from this "static" menu 

	[[menu.shortcuts]]
	pre = "<h3>More</h3>"
	name = "<i class='fa fa-github'></i> Github repo"
	identifier = "ds"
	url = "https://github.com/vjeantet/hugo-theme-docdock"
	weight = 1

	[[menu.shortcuts]]
	name = "<i class='fa fa-bookmark'></i> Hugo Documentation"
	identifier = "hugodoc"
	url = "https://gohugo.io/"
	weight = 2


[{{%icon circle-arrow-right%}} Read more about hugo and menu here](https://gohugo.io/extras/menus/)