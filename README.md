# Canvas cURL rewritten in golang

![Build Status](https://github.com/barrett370/go-canvas-cURL/workflows/Build/badge.svg)

to use build using:

```bash
make build
```

with your Canvas API token exported on your path as CANVAS_TOKEN

To scrape a specific list of modules simply enter the course names into a file separated by newlines. Then to run:

```bash
./Scrape -requirementsFile=<path-to-file>
```

To scrape a single specific module run:

```bash 
./Scrape -module=<module-name>
```