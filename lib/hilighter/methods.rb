class Hilighter
    module Methods
        def hex_color(hex)
            return self.send("color_#{hex_to_x256(hex)}")
        end

        # Convert hex to xterm-256 8-bit value
        # https://stackoverflow.com/questions/11765623/convert-hex-hex_to_256bitto-closest-x11-color-number
        def hex_to_x256(hex)
            # For simplicity, assume RGB space is perceptually
            # uniform. There are 5 places where one of two outputs
            # needs to be chosen when the input is the exact middle:
            # - The r/g/b channels and the gray value:
            #     - the higher value output is chosen
            # - If the gray and color values have same distance from
            #   the input
            #     - color is chosen

            # Calculate the nearest 0-based color index at 16..231
            r = g = b = 0
            x = "[A-Fa-f0-9]{2}"
            hex.match(/^#?(#{x})(#{x})(#{x})(#{x})?$/) do |m|
                r = m[1].to_i(16)
                g = m[2].to_i(16)
                b = m[3].to_i(16)
            end

            # 0..5 each
            ir = (r < 48) ? 0 : (r < 115) ? 1 : ((r - 35) / 40)
            ig = (g < 48) ? 0 : (g < 115) ? 1 : ((g - 35) / 40)
            ib = (b < 48) ? 0 : (b < 115) ? 1 : ((b - 35) / 40)

            # 0..215 lazy evaluation
            clr_index = (36 * ir) + (6 * ig) + ib + 16

            # Calculate the nearest 0-based gray index at 232..255
            average = (r + g + b) / 3

            # 0..23
            gray_index = (average > 238) ? 23 : (average - 3) / 10

            # Calculate the represented colors back from the index
            i2cv = [0, 0x5f, 0x87, 0xaf, 0xd7, 0xff]

            # r/g/b  0..255 each
            cr = i2cv[ir]
            cg = i2cv[ig]
            cb = i2cv[ib]

            # same value for r/g/b  0..255
            gv = (10 * gray_index) + 8

            # Return the one which is nearer to the original rgb
            # values
            clr_err = ((cr - r) ** 2 + (cg - g) ** 2 + (cb - b) ** 2)
            gray_err = ((gv - r) ** 2 + (gv - g) ** 2 + (gv - b) ** 2)

            if (clr_err <= gray_err)
                return clr_index.to_s.rjust(3, "0")
            else
                return (gray_index + 232).to_s.rjust(3, "0")
            end
        end
        private :hex_to_x256

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
            return self.send("on_color_#{hex_to_x256(hex)}")
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
