class Hilighter
    module Methods
        def on_rainbow
            return self if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new

            self.scan(
                /((\e\[[0-9;]+m)?[^\e](\e\[0m)?)/
            ).each_with_index do |c, i|
                out.push("\e\[#{clrs[i % clrs.length] + 10}m#{c[0]}")
            end
            out.push("\e\[0m")

            return out.join
        end

        def plain
            return self.gsub(/\e\[([0-9;]*m|K)/, "")
        end

        def rainbow_colors
            return [31, 32, 33, 34, 35, 36, 91, 92, 93, 94, 95, 96]
        end
        private :rainbow_colors

        def rainbow
            return self if (Hilighter.disable?)

            clrs = rainbow_colors
            out = Array.new

            self.scan(
                /((\e\[[0-9;]+m)?[^\e](\e\[0m)?)/
            ).each_with_index do |c, i|
                out.push("\e\[#{clrs[i % clrs.length]}m#{c[0]}")
            end
            out.push("\e\[0m")

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
