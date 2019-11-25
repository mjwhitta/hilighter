# Hilighter

## What is this?

This go package provides color methods for strings. It also provides a
method for wrapping strings that accounts for color escape codes.

## How to install

```
$ go get -u gitlab.com/mjwhitta/hilighter/cmd/hl
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

In golang you can do things like the following:

```
package main

import (
    "fmt"

    hl "gitlab.com/mjwhitta/hilighter"
)

func main() {
    // Example 1 (single color)
    var greenStr = hl.Green("1. Hello, %s!\n", "world")
    fmt.Print(greenStr)

    // or

    hl.PrintGreen("1. Hello, %s!\n", "world")

    // or

    hl.PrintlnGreen("1. Hello, world!")

    // Example 2 (multiple colors)
    var multiColored = hl.Hilights(
        []string{"white", "on_green"},
        "2. Hello, %s!\n",
        "world",
    )
    fmt.Print(multiColored)

    // or

    multiColored = hl.Hilights(
        []string{"white", "on_green"},
        "2. Hello, world!",
    )
    fmt.Println(multiColored)

    // or

    hl.PrintHilights(
        []string{"white", "on_green"},
        "2. Hello, %s!\n",
        "world",
    )

    // or

    hl.PrintlnHilights(
        []string{"white", "on_green"},
        "2. Hello, world!",
    )

    // Example 3 (8-bit)
    var eightBit = hl.Color002("3. 8-bit color codes!")
    fmt.Println(eightBit)

    // or

    hl.PrintlnColor002("3. 8-bit color codes!")

    // Example 4 (text wrapping)
    var long_var = "4."
    var word = "hilight"
    for i := 0; i < 32; i++ {
        long_var += " " + word
    }

    var wrapped = hl.Wrap(80, hl.Hilight("green", long_var))
    fmt.Println(wrapped)

    // or

    hl.PrintWrap(80, hl.Green("%s\n", long_var))

    // or

    hl.PrintlnWrap(80, hl.Green(long_var))
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

- [Source](https://gitlab.com/mjwhitta/hilighter)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)

## TODO

- Better README
