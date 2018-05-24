---
title: "Configuration: iTerm2"
date: 2018-05-24T09:57:40-07:00
draft: false
---

Many terminal apps don't properly display multibyte emoji characters
properly. This **may** fix the issue for you in iTerm2, it also may not.

By default iTerm2 uses a unicode rendering format
that is not comletely compatible with some emoji characters. Instead what you'll
see is the emoji over-lapping normal text characters, or drawing outside
the bounds of where they should be.

In iTerm2 open:

```bash
Preferences -> Profiles -> Text
```
and check **on** the "Use Unicode Version 9 Widths" checkbox. Then
restart WTF.

<img src="/imgs/iterm2prefs.png" width="800" height="437" alt="iTerm2
Prefs" />

(*Note:* This issue is not unique to iTerm2. As of this writing it also
affects <a href="https://en.wikipedia.org/wiki/Terminal_(macOS)">Terminal</a>, and <a href="https://hyper.is">Hyper</a>.)
