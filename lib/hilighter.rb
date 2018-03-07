class Hilighter
    def self.disable(val = true)
        @@disable = val
    end

    def self.disable?
        @@disable ||= false
        return @@disable
    end

    # Convert hex to xterm-256 8-bit value
    # https://stackoverflow.com/questions/11765623/convert-hex-hex_to_256bitto-closest-x11-color-number
    def self.hex_to_x256(hex)
        @@cached_codes ||= Hash.new
        cache = @@cached_codes[hex]
        return "color_#{cache}" if (!cache.nil?)

        # For simplicity, assume RGB space is perceptually uniform.
        # There are 5 places where one of two outputs needs to be
        # chosen when the input is the exact middle:
        # - The r/g/b channels and the gray value:
        #     - the higher value output is chosen
        # - If the gray and color values have same distance from the
        #   input
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
        cidx = (36 * ir) + (6 * ig) + ib + 16

        # Calculate the nearest 0-based gray index at 232..255
        average = (r + g + b) / 3

        # 0..23
        gidx = (average > 238) ? 23 : (average - 3) / 10

        # Calculate the represented colors back from the index
        i2cv = [0, 0x5f, 0x87, 0xaf, 0xd7, 0xff]

        # r/g/b  0..255 each
        cr = i2cv[ir]
        cg = i2cv[ig]
        cb = i2cv[ib]

        # same value for r/g/b  0..255
        gv = (10 * gidx) + 8

        # Return the one which is nearer to the original rgb
        # values
        clr_err = ((cr - r) ** 2 + (cg - g) ** 2 + (cb - b) ** 2)
        gray_err = ((gv - r) ** 2 + (gv - g) ** 2 + (gv - b) ** 2)

        if (clr_err <= gray_err)
            @@cached_codes[hex] = cidx.to_s.rjust(3, "0")
        else
            @@cached_codes[hex] = (gidx + 232).to_s.rjust(3, "0")
        end

        return "color_#{@@cached_codes[hex]}"
    end

    def self.sample
        String.colors.keys[0, 16].each do |fg|
            String.colors.keys[0, 16].each do |bg|
                print " test ".send(fg).send("on_#{bg}")
            end
            puts
        end
    end

    def self.table
        (0..15).each do |i|
            bg = i.to_s.rjust(3, "0")
            print " ".send("on_color_#{bg}")
            print bg.black.send("on_color_#{bg}")
            print " ".send("on_color_#{bg}")
            print bg.white.send("on_color_#{bg}")
            print " ".send("on_color_#{bg}")
            puts if (((i + 1) % 8) == 0)
        end

        (16..255).each do |i|
            bg = i.to_s.rjust(3, "0")
            print " ".send("on_color_#{bg}")
            print bg.black.send("on_color_#{bg}")
            print " ".send("on_color_#{bg}")
            print bg.white.send("on_color_#{bg}")
            print " ".send("on_color_#{bg}")
            puts if (((i - 15) % 6) == 0)
        end
    end
end

require "hilighter/codes"
require "hilighter/methods"
require "string"
