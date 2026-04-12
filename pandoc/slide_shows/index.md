---
title: "My Presentation"
author: "John Doe"
date: "June 21, 2017"
theme: "solarized"
transition: "zoom"
parallaxBackgroundImage: ./img/dream.jpeg
title-slide-attributes:
  data-background-image: ./img/dream.jpeg
  data-background-size: contain 
---

# Slide 1
Hello world   


<!--
Source - https://stackoverflow.com/q/75279429
Posted by Sam Hartman
Retrieved 2026-02-16, License - CC BY-SA 4.0
-->

# Slide Title

- This appears initially.

- This appears later


<!--
Source - https://stackoverflow.com/q/75279429
Posted by Sam Hartman
Retrieved 2026-02-16, License - CC BY-SA 4.0
-->

::: {.fragment}

- This appears later

:::



<!--
Source - https://stackoverflow.com/a/75279448
Posted by Sam Hartman
Retrieved 2026-02-16, License - CC BY-SA 4.0
-->

# Slide Title

- This appears initially.

::: incremental

- And this appears later.
- This appears even later???

:::



# Slide 1

- Bullet point 1
- Bullet point 2

## Slide 2

```python
print("Hello, world!")
```

### Generate Slides
Use this command to convert Markdown to HTML:
```bash
pandoc -t revealjs -s -o slides.html slides.md \
  -V revealjs-url=https://unpkg.com/reveal.js@4.3.1/ \
  -V theme=sky \
  -V transition=cube \
  --css=custom.css   
```

::: notes
This is my note.
- It can contain Markdown
- like this list
:::


## Slide One

Slide 1 has background_image.png as its background.

## {background-image="./img/dream.jpeg"}

Slide 2 has a special image for its background, even though the heading has no content.