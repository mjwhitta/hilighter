class Hilighter
    module Codes
        def add_methods
            colors.each do |key, val|
                define_method key do
                    return self if (Hilighter.disable?)
                    return "\e[#{val}m#{self}\e[0m"
                end

                define_method "on_#{key}" do
                    return self if (Hilighter.disable?)
                    return "\e[#{val + 10}m#{self}\e[0m"
                end
            end

            256.times.each do |i|
                clr = i.to_s.rjust(3, "0")

                define_method "color_#{clr}" do
                    return self if (Hilighter.disable?)
                    return "\e[38;5;#{clr}m#{self}\e[0m"
                end

                define_method "on_color_#{clr}" do
                    return self if (Hilighter.disable?)
                    return "\e[48;5;#{clr}m#{self}\e[0m"
                end
            end

            modes.each do |key, val|
                define_method key do
                    return self if (Hilighter.disable?)
                    return "\e[#{val}m#{self}\e[0m"
                end
            end
        end

        def colors
            return {
                "black" => 30,
                "red" => 31,
                "green" => 32,
                "yellow" => 33,
                "blue" => 34,
                "magenta" => 35,
                "cyan" => 36,
                "white" => 37,
                "light_black" => 90,
                "light_red" => 91,
                "light_green" => 92,
                "light_yellow" => 93,
                "light_blue" => 94,
                "light_magenta" => 95,
                "light_cyan" => 96,
                "light_white" => 97
            }
        end

        def modes
            return {
                "default" => 0,
                "reset" => 0,
                "normal" => 0,
                "bold" => 1,
                "dim" => 2,
                "faint" => 2,
                "italic" => 3,
                "underline" => 4,
                "blink" => 5,
                "blink_slow" => 5,
                "blink_rapid" => 6,
                "swap" => 7,
                "negative" => 7,
                "hide" => 8,
                "conceal" => 8,
                "crossed_out" => 9,
                "strikethrough" => 9,
                "fraktur" => 20,
                "no_bold" => 21,
                "no_dim" => 22,
                "no_faint" => 22,
                "no_italic" => 23,
                "no_underline" => 24,
                "no_blink" => 25,
                "no_blink_slow" => 25,
                "no_blink_rapid" => 26,
                "no_swap" => 27,
                "no_negative" => 27,
                "no_hide" => 28,
                "no_conceal" => 28,
                "no_crossed_out" => 29,
                "no_strikethrough" => 29
            }
        end
    end
end
