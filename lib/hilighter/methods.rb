class Hilighter
    module Methods
        def on_rainbow
            return self if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new

            self.scan(
                /((\e\[[0-9;]+m)?[^\e](\e\[0m)?)/
            ).each_with_index do |c, i|
                out.push("\e\[#{clrs[i % clrs.length] + 40}m#{c[0]}")
            end
            out.push("\e\[0m")

            return out.join
        end

        def plain
            return self.gsub(/\e\[([0-9;]*m|K)/, "")
        end

        def rainbow_colors
            return [1, 2, 3, 4, 5, 6, 61, 62, 63, 64, 65, 66]
        end
        private :rainbow_colors

        def rainbow
            return self if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new

            self.scan(
                /((\e\[[0-9;]+m)?[^\e](\e\[0m)?)/
            ).each_with_index do |c, i|
                out.push("\e\[#{clrs[i % clrs.length] + 30}m#{c[0]}")
            end
            out.push("\e\[0m")

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
