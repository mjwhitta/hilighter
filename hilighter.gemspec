Gem::Specification.new do |s|
    s.add_development_dependency("rake", "~> 13.0", ">= 13.0.0")
    s.authors = ["Miles Whittaker"]
    s.date = Time.new.strftime("%Y-%m-%d")
    s.description = [
        "Adds color methods to String class. Also allows for string",
        "wrapping that accounts for color escape codes."
    ].join(" ")
    s.email = "mj@whitta.dev"
    s.executables = Dir.chdir("bin") do
        Dir["*"]
    end
    s.files = Dir["lib/**/*.rb"]
    s.homepage = "https://gitlab.com/mjwhitta/hilighter/tree/ruby"
    s.license = "GPL-3.0"
    s.metadata = {"source_code_uri" => s.homepage}
    s.name = "hilighter"
    s.summary = "Adds color methods to String class"
    s.version = "1.5.2"
end
