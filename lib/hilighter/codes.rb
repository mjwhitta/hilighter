class Hilighter
    module Codes
        def add_methods
            colors.each do |key, val|
                define_method key do
                    return self if (Hilighter.disable?)
                    return "\e[#{val + 30}m#{self}\e[0m"
                end

                define_method "on_#{key}" do
                    return self if (Hilighter.disable?)
                    return "\e[#{val + 40}m#{self}\e[0m"
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
                "black" => 0,
                "red" => 1,
                "green" => 2,
                "yellow" => 3,
                "blue" => 4,
                "magenta" => 5,
                "cyan" => 6,
                "white" => 7,
                "light_black" => 60,
                "light_red" => 61,
                "light_green" => 62,
                "light_yellow" => 63,
                "light_blue" => 64,
                "light_magenta" => 65,
                "light_cyan" => 66,
                "light_white" => 67
            }
        end

        def modes
            return {
                "default" => 0,
                "bold" => 1,
                "underline" => 4,
                "blink" => 5,
                "swap" => 7,
                "hide" => 8
            }
        end
    end
end
