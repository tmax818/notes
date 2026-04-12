---
title: Pandoc's Markdown
subtitle: This is my sweet subtitle
author: Tyler Maxwell
fontsize: 18pt
---

# test

    # heading

```go
qsort [] = []
```

:smile: 

This text is _emphasized with underscores_

H~2~O

```{=ms}
.MYMACRO
```

An inline ![image](foo.jpg){#id .class width=30 height=20px}
and a reference ![image][ref] with attributes.
[ref]: foo.jpg "optional title" {#id .class key=val key2="val 2"}

Here is a footnote reference,[^1] and another.[^longnote]
[^1]: Here is the footnote.
[^longnote]: Here's one with multiple blocks.
Subsequent paragraphs are indented to show that they
belong to the previous footnote.
{ some.code }
The whole paragraph can be indented, or just the first
line.
In this way, multi-paragraph footnotes work like
multi-paragraph list items.
This paragraph won

Here is an inline note.^[Inline notes are easier to write, since
you don't have to pick an identifier and move down to type the
note.]


Blah blah [@doe99; @smith2000; @smith2004].

[[URL|title]]