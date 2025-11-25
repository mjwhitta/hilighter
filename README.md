# Hilighter

[![Yum](https://img.shields.io/badge/-Buy%20me%20a%20cookie-blue?labelColor=grey&logo=cookiecutter&style=for-the-badge)](https://www.buymeacoffee.com/mjwhitta)

[![Go Report Card](https://goreportcard.com/badge/github.com/mjwhitta/hilighter?style=for-the-badge)](https://goreportcard.com/report/github.com/mjwhitta/hilighter)
[![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/mjwhitta/hilighter/ci.yaml?style=for-the-badge)](https://github.com/mjwhitta/hilighter/actions)
![License](https://img.shields.io/github/license/mjwhitta/hilighter?style=for-the-badge)

## What is this?

This go package provides color methods for strings. It also provides a
method for line-wrapping strings that accounts for color escape codes.

## How to install

Open a terminal and run the following:

```
$ go get -u github.com/mjwhitta/hilighter
$ go install github.com/mjwhitta/hilighter/cmd/hl@latest
```

## Usage

In a terminal you can do things like the following:

```
$ cat some_file | hl green on_blue
$ cat some_file | hl rainbow on_white dim
$ cat some_file | hl rainbow on_rainbow
$ hl rainbow on_rainbow <some_file
$ echo "Hex color codes!" | hl ffffff on_ff0000
$ cat some_file | hl wrap
$ cat some_file | hl wrap_64
```

In Go you can do things like the following:

```
package main

import (
    "fmt"
    "slices"
    "strings"

    hl "github.com/mjwhitta/hilighter"
)

func main() {
    // Example 1 (single color)
    greenStr := hl.Green("Hello, world!")
    fmt.Println("1. " + greenStr)

    // Example 2 (multiple colors)
    multiColored := hl.Hilights(
        []string{"white", "on_green"},
        "Hello, world!",
    )
    fmt.Println("2. " + multiColored)

    // Example 3 (8-bit)
    eightBit := hl.Hilight("002", "8-bit color codes!")
    fmt.Println("3. " + eightBit)

    // Example 4 (hex)
    hexColorStr := hl.Hex("5f8700", "Hex color codes!")
    fmt.Println("4. " + hexColorStr)

    // Example 5 (text wrapping)
    longVar := strings.Join(slices.Repeat([]string{"word"}, 32), " ")
    fmt.Println("5.\n" + hl.Wrap(70, hl.Green(longVar)))

    // Example 6 (rainbow)
    fmt.Println("6.\n" + hl.Wrap(61, hl.Rainbow(longVar)))

    // Example 7 (double rainbow)
    fmt.Println(
        "7.\n" + hl.Wrap(61, hl.Rainbow(hl.OnRainbow(longVar))),
    )
}
```

The following colors are supported:

Foreground           | Background
----------           | ----------
black                | on_black
red                  | on_red
green                | on_green
yellow               | on_yellow
blue                 | on_blue
magenta              | on_magenta
cyan                 | on_cyan
white                | on_white
light_black          | on_light_black
light_red            | on_light_red
light_green          | on_light_green
light_yellow         | on_light_yellow
light_blue           | on_light_blue
light_magenta        | on_light_magenta
light_cyan           | on_light_cyan
light_white          | on_light_white
default              | on_default
color000 to color255 | on_color000 to on_color255
000000 to ffffff     | on_000000 to on_ffffff

The following modes are supported:

On            | Off              | Description
---           | ---              | -----------
normal        |                  | Same as default
reset         |                  | Same as default
bold          | no_bold          | Turn on/off bold
dim           | no_dim           | Turn on/off dim. Not widely supported
faint         | no_faint         | Same as dim
italic        | no_italic        | Turn on/off italics. Sometimes treated as inverse. Not widely supported.
underline     | no_underline     | Turn on/off underline
blink         | no_blink         | Turn on/off blink. Less than 150/min.
blink_slow    | no_blink_slow    | Same as blink
blink_rapid   | no_blink_rapid   | Same as blink. 150+/min. Not widely supported.
inverse       | no_inverse       | Reverse foreground/background
negative      | no_negative      | Same as inverse
swap          | no_swap          | Same as inverse
conceal       | no_conceal       | Turn on/off conceal. Useful for passwords. Not widely supported.
hide          | no_hide          | Same as conceal
crossed_out   | no_crossed_out   | Characters legible, but marked for deletion. Not widely supported.
strikethrough | no_strikethrough | Same as CrossedOut
[fraktur]     | no_fraktur       | Hardly ever supported

[fraktur]: https://en.wikipedia.org/wiki/Fraktur

## Links

- [Source](https://github.com/mjwhitta/hilighter)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)

## TODO

- Better README
