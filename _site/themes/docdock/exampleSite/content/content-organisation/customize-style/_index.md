+++
title = "Customize website look and feel"
Weight=3
+++

You can change the style and behavior of the theme without touching it.

* inject your own html,  css or js into the page
* overide existing css or js with your own files

{{%notice note %}}
No needs to copy the entire theme to customize some parts 
Bellow are solutions to avoid copying the entire theme into your own codebase.
{{%/notice%}}

## Add custom CSS and JS or HTML into the \<head\> part of each page :

Create a custom header partial `layouts/partials/custom-head.html` 

> * content/
> * layouts/
>   * partials/
>      * custom-head.html

write your own content like (an example from @nzbart):
```html
<link rel="stylesheet" href="/css/custom.css">
<script src="/js/custom.js"></script>
```

Then overrode the style your want to change in `static/css/custom.css` (in this case, to avoid altering the casing of titles):
```css
h2 {
    text-transform: none;
}
```

And executed some additional JavaScript from `static/js/custom.js` (note that jQuery is already loaded by the theme):
```javascript
function tweakPage() {
    // make some changes here
}

$(tweakPage)
```


now feel free to add the JS, CSS, HTML code you want :)

## Add custom HTML at the end of the body part of each page :

Create a `custom-footer.html` into a `layouts/partials` folder next to the content folder

> * content/
> * layouts/
>   * partials/
>      * custom-footer.html

now feel free to add the JS, CSS, HTML code you want :)

## Overide existing CSS or JS

Create the matching file in your static folder, hugo will use yours instead of the theme's one.
Example : 

create a theme.css and place it into `static/css/` to fully overide docdock's theme.css
