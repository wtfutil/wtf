+++
title = "panel"
description = "Allow you to highlight information or put it in a box."
+++



{{% panel theme="success" header="The panel shortcode" %}}Allow you to highlight information or put it in a box. They create a colored box surrounding your text{{% /panel %}}


## Usage 

| Parameter | Default | Description |
|:--|:--|:--|
| header | none | The title of the panel. If specified, this title will be displayed in its own header row. |
| footer | none | the footer of the panel. If specified, this text will be displayed in its own row |
| theme | `primary` | `default`,`primary`,`info`,`success`,`warning`,`danger` |

## Basic example

By default :

	{{%/* panel */%}}this is a panel text{{%/* /panel */%}}

{{%panel%}}this is a panel text{{%/panel%}}

## Panel with heading

Easily add a heading container to your panel with `header` parameter. You may apply any theme.

	{{%/* panel theme="danger" header="panel title" */%}}this is a panel text{{%/* /panel */%}}

{{% panel theme="danger" header="panel title" %}}this is a panel text{{% /panel %}}

	{{%/* panel theme="success" header="panel title" */%}}this is a panel text{{%/* /panel */%}}

{{% panel theme="success" header="panel title" %}}this is a panel text{{% /panel %}}

## Panel with footer
Wrap a secondary text in footer.

	{{%/* panel footer="panel footer" */%}}Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.{{%/* /panel */%}}

{{% panel footer="panel footer" %}}
Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
{{% /panel %}}

## Themes

{{% panel theme="success" header="Success theme" %}}this is a panel text{{% /panel %}}
{{% panel theme="default" header="default theme" %}}this is a panel text{{% /panel %}}
{{% panel theme="primary" header="primary theme" %}}this is a panel text{{% /panel %}}
{{% panel theme="info" header="info theme" %}}this is a panel text{{% /panel %}}
{{% panel theme="warning" header="warning theme" %}}this is a panel text{{% /panel %}}
{{% panel theme="danger" header="danger theme" %}}this is a panel text{{% /panel %}}
