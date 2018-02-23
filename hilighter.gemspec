Gem::Specification.new do |s|
    s.name = "hilighter"
    s.version = "1.1.1"
    s.date = Time.new.strftime("%Y-%m-%d")
    s.summary = "Adds color methods to String class"
    s.description =
        "Adds color methods to String class. Also allows for " \
        "string wrapping that accounts for color escape codes."
    s.authors = [ "Miles Whittaker" ]
    s.email = "mjwhitta@gmail.com"
    s.executables = Dir.chdir("bin") do
        Dir["*"]
    end
    s.files = Dir["lib/**/*.rb"]
    s.homepage = "https://mjwhitta.github.io/hilighter"
    s.license = "GPL-3.0"
    s.add_development_dependency("rake", "~> 12.3", ">= 12.3.0")
end
