# Hilighter

## What is this?

This ruby gem adds color methods to the String class. It also provides
a method for wrapping strings that accounts for color escape codes.

## How to install

```bash
$ gem install hilighter
```

## Usage

```ruby
require "hilighter"

puts "Hello, world!".white
puts Array.new(100, "hilight").join(" ").wrap
puts Array.new(100, "hilight").join(" ").green.wrap
```

## Links

- [Homepage](https://mjwhitta.github.io/hilighter)
- [Source](https://gitlab.com/mjwhitta/hilighter)
- [Mirror](https://github.com/mjwhitta/hilighter)
- [RubyGems](https://rubygems.org/gems/hilighter)

## TODO

- Better README
- RDoc
