# hyde-hyde

__`hyde-hyde`__ is a [Hugo](https://gohugo.io)'s theme derived from @spf13's [Hyde](https://github.com/spf13/hyde.git) which is in turn ported from @mdo Jekyll's [Hyde](https://github.com/poole/hyde). 

## Notable Changes
* Restructuring/Modularising the layouts (`baseof.html`, `single.html`, `list.html` and the partials)
* Using [highlight.js](https://highlightjs.org) for highlighting code
* Using [Font-Awesome 5](https://fontawesome.com) for icons
* Using color tones inspired by [Nate Finch's blog](https://npf.io)
* Using main font [Fira-Sans](https://fonts.google.com/specimen/Fira+Sans) + fixed width font [Roboto Mono](https://fonts.google.com/specimen/Roboto+Mono)
* Adding [GraphComment](https://graphcomment.com) for replacing the built-in [Disqus](https://disqus.com)

A real site in action [here](https://htr3n.github.io).

![hyde-hyde main screen](https://github.com/htr3n/hyde-hyde/blob/master/images/main.png)

![hyde-hyde main screen](https://github.com/htr3n/hyde-hyde/blob/master/images/posts.png)

## Installation

`hyde-hyde` can be easily installed as many other Hugo's themes:

```sh
$ cd HUGO_SITE
# then clone hyde-hyde
$ git clone https://github.com/htr3n/hyde-hyde.git themes/hyde-hyde
# or add hyde-hyde as a submodule
$ git submodule add https://github.com/htr3n/hyde-hyde.git themes/hyde-hyde
```

Then indicate `hyde-hyde` as the main theme

* `config.toml` 

```tomp
theme = "hyde-hyde"
```

* `config.yaml`

```yaml
theme : "hyde-hyde"
```

## Options

* `hyde-hyde` essentially inherits [all options](https://github.com/spf13/hyde#options) from Hyde.

## Customisations

* Most of the newly added customisations are in the file `hyde-hyde/static/css/custom.css`
* The layouts for a single post or a list/table of content in `hyde-hyde/layouts` are modularised and easily to changed

## Author(s)
### Original Developed by Mark Otto

- <https://github.com/mdo>
- <https://twitter.com/mdo>

### Hugo's Hyde Ported by Steve Francia
- <https://github.com/spf13>
- <https://twitter.com/spf13>

### Color Theme Inspired By

* [Nate Finch's blog](https://npf.io)

## License

Open sourced under the [MIT license](LICENSE.md).
