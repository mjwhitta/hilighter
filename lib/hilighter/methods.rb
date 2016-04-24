class Hilighter
    module Methods
        def plain
            return self.gsub(/\e\[([0-9;]*m|K)/, "")
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
