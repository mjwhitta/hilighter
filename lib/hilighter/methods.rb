class Hilighter
    module Methods
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
            return self.gsub(/\e\[[0-9;]*m/, "")
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
