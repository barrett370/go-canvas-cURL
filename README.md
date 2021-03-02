# [Canvas cURL](https://github.com/Chasbob/Canvas-cURL) rewritten in golang

![Build Status](https://github.com/barrett370/go-canvas-cURL/workflows/Build/badge.svg)

## Description 

This program, once compiled (see [usage](# Usage)), simply extracts all files from your canvas courses. You can restrict the courses and file extensions using different flags and files also described in  [usage](# Usage))

This is a rewriting of [Canvas cURL](https://github.com/Chasbob/Canvas-cURL) by Charlie de Freitas using golang as I felt this language was a better fit (more performant & easier to read)

## Usage

to use build using:

```bash 
just 
```

with your Canvas API token defined as AuthToken in a yaml file (default is config.yaml in same dir as exe, can point to alternative locations with `--config` flag)

To scrape a specific list of modules simply follow `download` with the names of the modules

```bash
./scrape download mod1 mod2 ...
```
Alternatively, `./scrape download all` will download all found modules

You can also restrict the file extensions which you are interested in by typing the extensions you wish to ignore in a `.scrapeignore` file in the toplevel directory

