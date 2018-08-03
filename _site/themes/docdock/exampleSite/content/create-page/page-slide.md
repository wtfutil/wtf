+++
title = "Present a Slide"
description = ""
date = "2017-04-24T18:36:24+02:00"
+++

A basic md content page can be rendered as a reveal.js presentation full screen.

{{%alert info%}}You can, also, **embed presentation in a page** as a small box, using the [revealjs]({{% relref "shortcodes/revealjs.md"%}}) shortcode in your md file.{{%/alert%}}


## Formating
Use your common Markdown syntax you use in Hugo, don't forget, you can put html tags too.

{{%notice info %}} Special syntax (in html comment) is available for adding attributes to Markdown elements. This is useful for fragments, amongst other things.
{{%/notice%}}

Please read the [{{%icon book%}} doc from hakimel](https://github.com/hakimel/reveal.js/#instructions)


## Options
In the frontmatter of your page file, set **type** and **revealOptions** params

Your content will be served as a fullscreen revealjs presentation and revealOptions will be used to ajust its behaviour.

	+++
	title = "Test slide"
	type="slide"

	theme = "league"
	[revealOptions]
	transition= 'concave'
	controls= true
	progress= true
	history= true
	center= true
	+++

[read more about reveal options here](https://github.com/hakimel/reveal.js/#configuration)


## Slide Delimiters
When creating the content for your slideshow presentation within content markdown file you need to be able to distinguish between one slide and the next. This is achieved very simply using a  convention within Markdown that indicates the start of each new slide.

As both horizontal and vertical slides are supported by reveal.js each has it's own unique delimiter.

To denote the start of a horizontal slide simply add the following delimiter in your Markdown:

	---


To denote the start of a vertical slide simply add the following delimiter in your Markdown:
	
	___

By using a combination of horizontal and vertical slides you can customize the navigation within your slideshow presentation. Typically vertical slides are used to present information below a top-level horizontal slide.



For example, a very simple slideshow presentation can be created as follows

```
+++

title = "test"
date = "2017-04-24T18:36:24+02:00"
type="slide"

theme = "league"
[revealOptions]
transition= 'concave'
controls= true
progress= true
history= true
center= true
+++

# In the morning

___

## Getting up

- Turn off alarm
- Get out of bed

___

## Breakfast

- Eat eggs
- Drink coffee

---

# In the evening

___

## Dinner

- Eat spaghetti
- Drink wine

___

## Going to sleep

- Get in bed
- Count sheep

```

[{{%icon expand%}}click here to view this page rendered]({{%relref "myslide.md"%}})