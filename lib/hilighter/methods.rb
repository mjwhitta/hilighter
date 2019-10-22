class Hilighter
    module Methods
        def color_regex
            return /\e\[([0-9;]*m|K)/
        end
        private :color_regex

        def hex_color(hex)
            return self.send(Hilighter.hex_to_x256(hex))
        end

        def method_missing(method_name, *args)
            method_name.to_s.match(/^([A-Fa-f0-9]{6})$/) do |m|
                return self.hex_color(m[1])
            end
            method_name.to_s.match(/^on_([A-Fa-f0-9]{6})$/) do |m|
                return self.on_hex_color(m[1])
            end
            super
        end

        def on_hex_color(hex)
            return self.send("on_#{Hilighter.hex_to_x256(hex)}")
        end

        def on_rainbow
            return self.plain if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new

            self.plain_bg.each_line do |line|
                line.chomp!
                colorized = line.scan(
                    /((#{color_regex})*[^\e](#{color_regex})*)/
                ).map.with_index do |c, i|
                    "\e[#{clrs[i % clrs.length] + 10}m#{c[0]}"
                end.join
                out.push(colorized)
                out.push("\e[49m") if (!line.empty?)
                out.push("\n")
            end

            return out.join
        end

        def plain
            return self.gsub(color_regex, "")
        end

        def plain_bg
            return self.gsub(/\e\[(4|10)[0-9;]+m/, "")
        end

        def plain_fg
            return self.gsub(/\e\[[39][0-9;]+m/, "")
        end

        def rainbow_colors
            return [31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96]
        end
        private :rainbow_colors

        def rainbow
            return self.plain if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new

            self.plain_fg.each_line do |line|
                line.chomp!
                colorized = line.scan(
                    /((#{color_regex})*[^\e](#{color_regex})*)/
                ).map.with_index do |c, i|
                    "\e[#{clrs[i % clrs.length]}m#{c[0]}"
                end.join
                out.push("#{colorized}\n")
            end

            return out.join
        end

        def wrap(width = 80)
            lines = Array.new
            line = ""
            self.split(/\s+/).each do |word|
                if ((line.plain.size + word.plain.size) > width)
                    lines.push("#{line}\n")
                    line = word
                elsif (line.empty?)
                    line = word
                else
                    line += " #{word}"
                end
            end
            lines.push("#{line}\n") if (!line.empty?)
            return lines.join
        end
    end
end
