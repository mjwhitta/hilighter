class Hilighter
    module Codes
        def add_methods
            colors.each do |key, val|
                if (key.start_with?("on_"))
                    define_method key do
                        return self if (Hilighter.disable?)
                        return [
                            "\e[#{val}m",
                            self.plain_bg.gsub(
                                /\n/,
                                "\e[49m\n\e[#{val}m"
                            ),
                            "\e[49m"
                        ].join
                    end
                else
                    define_method key do
                        return self if (Hilighter.disable?)
                        return "\e[#{val}m#{self.plain_fg}\e[39m"
                    end
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
            @valid_colors ||= {
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
                "light_white" => 97,

                "on_black" => 40,
                "on_red" => 41,
                "on_green" => 42,
                "on_yellow" => 43,
                "on_blue" => 44,
                "on_magenta" => 45,
                "on_cyan" => 46,
                "on_white" => 47,
                "on_light_black" => 100,
                "on_light_red" => 101,
                "on_light_green" => 102,
                "on_light_yellow" => 103,
                "on_light_blue" => 104,
                "on_light_magenta" => 105,
                "on_light_cyan" => 106,
                "on_light_white" => 107,

                "default" => 39,
                "on_default" => 49
            }
            if (@valid_colors.length < 256)
                256.times.each do |i|
                    clr = i.to_s.rjust(3, "0")
                    @valid_colors["color_#{clr}"] = "38;5;#{clr}"
                    @valid_colors["on_color_#{clr}"] = "48;5;#{clr}"
                end
            end
            return @valid_colors
        end

        def modes
            return {
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
                "no_fraktur" => 23,
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
