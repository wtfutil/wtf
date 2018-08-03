+++
title = "Create Page"
description = ""
date = "2017-04-24T18:36:24+02:00"
creatordisplayname = "Valere JEANTET"
creatoremail = "valere.jeantet@gmail.com"
lastmodifierdisplayname = "Valere JEANTET"
lastmodifieremail = "valere.jeantet@gmail.com"
tags = ["tag1","tag2"]
weight = 10
pre ="<i class='fa fa-edit' ></i> "

+++


Hugo-theme-docdock defines two types of pages. _Default_ and _Slide_.

* **Default** is the common page like the current one you are reading.
* **Slide** is a page that use the full screen to display its markdown content as a [reveals.js presentation](http://lab.hakim.se/reveal-js/).
* **HomePage** is a special content that will be displayed as home page content.

To tell Hugo-theme-docdock to consider a page as a slide, just add a `type="slide"`in then frontmatter of your file. [{{%icon circle-arrow-right%}}read more on page as slide]({{%relref "page-slide.md"%}})


Hugo-theme-docdock provides archetypes to help you create this kind of pages.


## Front Matter
Each Hugo page has to define a Front Matter in yaml, toml or json.

Hugo-theme-docdock uses the following parameters on top of the existing ones :

	+++
	# Type of content, set "slide" to display it fullscreen with reveal.js
	type="page"

	# Creator's Display name
	creatordisplayname = "Valere JEANTET"
	# Creator's Email
	creatoremail = "valere.jeantet@gmail.com"
	# LastModifier's Display name
	lastmodifierdisplayname = "Valere JEANTET"
	# LastModifier's Email
	lastmodifieremail = "valere.jeantet@gmail.com"
	+++


## Ordering

Hugo provides a flexible way to handle order for your pages.

The simplest way is to use `weight` parameter in the front matter of your page. 

[{{%icon circle-arrow-right%}}Read more on content organization]({{%relref "content-organisation/_index.md"%}})
