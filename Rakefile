require "fileutils"

# These are no longer symlinks
symlinks = ["bin/hl"]

task :default => :gem

desc "Clean up"
task :clean do
    system("rm -f *.gem #{symlinks.join(" ")}")
    system("chmod -R go-rwx bin lib")
end

desc "Build gem"
task :gem do
    symlinks.each do |symlink|
        FileUtils.cp("bin/hilight", symlink)
    end
    system("chmod -R u=rwX,go=rX bin lib")
    system("gem build *.gemspec")
end

desc "Build and install gem"
task :install => :gem do
    system("gem install *.gem")
end

desc "Push gem to rubygems.org"
task :push => [:clean, :gem] do
    system("gem push *.gem")
end
