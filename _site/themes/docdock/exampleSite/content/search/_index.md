+++
title = "About the Search Engine"
description = ""
weight = 40
+++

{{%excerpt-include filename="getting-start/configuration.md" /%}}


Docdock theme uses the last improvement available in hugo version 20+ to generate a json index file ready to be consumed by lunr.js javascript search engine.


{{%note%}}hugo generate lunrjs index.json at the root of `public` folder if the site only has one language or within each language subfolder. <br/>When you build the site with `hugo server`, hugo generates it internally and of course it don't show up in the filesystem{{%/note%}}