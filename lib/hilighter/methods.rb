class Hilighter
    module Methods
        def on_rainbow
            return self if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new
            i = 0

            self.chars.each do |c|
                if (c.match(/\s/))
                    out.push(c)
                else
                    out.push(c.black.send("on_#{clrs[i % clrs.length]}"))
                    i += 1
                end
            end

            return out.join
        end

        def plain
            return self.gsub(/\e\[([0-9;]*m|K)/, "")
        end

        def rainbow_colors
            return [
                "red",
                "green",
                "yellow",
                "blue",
                "magenta",
                "cyan",
                "light_red",
                "light_green",
                "light_yellow",
                "light_blue",
                "light_magenta",
                "light_cyan",
            ]
        end
        private :rainbow_colors

        def rainbow
            return self if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new
            i = 0

            self.chars.each do |c|
                if (c.match(/\s/))
                    out.push(c)
                else
                    out.push(c.send(clrs[i % clrs.length]))
                    i += 1
                end
            end

            return out.join
        end

        def wrap(width = 80)
            lines = Array.new
            line = ""
            self.split(/\s+/).each do |word|
                if ((line.plain.size + word.plain.size) > width)
                    lines.push(line)
                    line = word
                elsif (line.empty?)
                    line = word
                else
                    line += " #{word}"
                end
            end
            lines.push(line) if (!line.empty?)
            return lines.join("\n")
        end
    end
end
