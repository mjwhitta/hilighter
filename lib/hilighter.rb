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
end

require "hilighter/codes"
require "hilighter/methods"
require "string"
