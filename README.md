# Hilighter

## What is this?

This ruby gem adds color methods to the String class. It also provides
a method for wrapping strings that accounts for color escape codes.

## How to install

```
$ gem install hilighter
```

## Usage

In a terminal you can do things like the following:

```
$ cat some_file | hilight green on_blue
$ cat some_file | hilight rainbow on_white dim
$ cat some_file | hl rainbow on_rainbow
$ hl rainbow on_rainbow <some_file
$ echo "Hex color codes!" | hl ffffff on_ff0000
```

Technically this is just calling methods from the String class on each
line of input, so you can also do:

```
$ cat some_bin | hilight dump
$ cat some_file | hl downcase
$ cat some_file | hl reverse
```

In a script you can do things like the following:

```ruby
require "hilighter"

puts("Hello, world!".white)
puts(Array.new(100, "hilight").join(" ").wrap)
puts(Array.new(100, "hilight").join(" ").green.wrap)
puts("Hex color codes!".hex_color("#ffffff").on_hex_color("#ff0000"))
puts("Hex color codes!".ffffff.on_ff0000)
```

The following colors are supported:

Foreground             | Background
----------             | ----------
black                  | on_black
red                    | on_red
green                  | on_green
yellow                 | on_yellow
blue                   | on_blue
magenta                | on_magenta
cyan                   | on_cyan
white                  | on_white
light_black            | on_light_black
light_red              | on_light_red
light_green            | on_light_green
light_yellow           | on_light_yellow
light_blue             | on_light_blue
light_magenta          | on_light_magenta
light_cyan             | on_light_cyan
light_white            | on_light_white
default                | on_default
color_000 to color_255 | on_color_000 to on_color_255
hex_color(hex)         | on_hex_color(hex)
000000 to ffffff       | on_000000 to on_ffffff

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
strikethrough | no_strikethrough | Same as crossed_out
[fraktur]     | no_fraktur       | Hardly ever supported. Use no_italic to turn off.

[fraktur]: https://en.wikipedia.org/wiki/Fraktur

## Links

- [Source](https://gitlab.com/mjwhitta/hilighter)
- [RubyGems](https://rubygems.org/gems/hilighter)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)

## TODO

- Better README
- RDoc
