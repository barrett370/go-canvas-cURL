# [Canvas cURL](https://github.com/Chasbob/Canvas-cURL) rewritten in golang

![Build Status](https://github.com/barrett370/go-canvas-cURL/workflows/Build/badge.svg)

## Description 

This program, once compiled (see [usage](# Usage)), simply extracts all files from your canvas courses. You can restrict the courses and file extensions using different flags and files also described in  [usage](# Usage))

This is a rewriting of [Canvas cURL](https://github.com/Chasbob/Canvas-cURL) by Charlie de Freitas using golang as I felt this language was a better fit (more performant & easier to read)

## Usage

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

You can also restrict the file extensions that you are interested in by typing the extensions you wish to ignore in a `.scrapeignore` file in the `./main` directory
