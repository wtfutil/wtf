+++
title = "alert"
description = "The alert shortcode allows you to highlight information in your page."
+++

The `alert` shortcode allow you to highlight information in your page. They create a colored box surrounding your text, like this:

{{%alert%}}**This is** an alert !{{%/alert%}}
## Usage 

| Parameter | Default | Description |
|:--|:--|:--|
| theme | `info` | `success`, `info`,`warning`,`danger` |

{{%alert info%}}
**Tips :** setting only the theme as argument works too : 
`{{%/*alert warning*/%}}`  instead of `{{%/*alert theme="warning"*/%}}`
{{%/alert%}}

## Basic examples

	{{%/* alert theme="info" */%}}**this** is a text{{%/* /alert */%}}
	{{%/* alert theme="success" */%}}**Yeahhh !** is a text{{%/* /alert */%}}
	{{%/* alert theme="warning" */%}}**Be carefull** is a text{{%/* /alert */%}}
	{{%/* alert theme="danger" */%}}**Beware !** is a text{{%/* /alert */%}}

{{% alert theme="info"%}}**this** is an info{{% /alert %}}
{{% alert theme="success" %}}**Yeahhh !** is an success{{% /alert %}}
{{% alert theme="warning" %}}**Be carefull** is a warning{{% /alert %}}
{{% alert theme="danger" %}}**Beware !** is a danger{{% /alert %}}