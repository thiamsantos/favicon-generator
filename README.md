# favicon-generator

Command-line to generate favicons from a image.

## Installation

Install [libvips](https://github.com/libvips/libvips):

```shell
$ apt install libvips-dev
```

You need `go` installed and `GOBIN` in your `PATH`. Once that is done, run the
command:

```shell
$ go get -u github.com/thiamsantos/favicon-generator
```

## Usage

```
Usage
  $ favicon-generator [options]

Options
  -i string
        Input image (default "logo.png")
  -o string
        Folder to output the favicons (default "favicons")

Examples
  $ favicon-generator
  $ favicon-generator -i path/to/logo.jpg
  $ favicon-generator -o /tmp/favicons
  $ favicon-generator -i path/to/logo.jpg -o /tmp/favicons
```
