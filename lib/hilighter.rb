class Hilighter
    @@disable = false

    def self.disable(val = true)
        @@disable = val
    end

    def self.disable?
        return @@disable
    end

    def self.sample
        String.colors.keys.each do |fg|
            String.colors.keys.each do |bg|
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
