require "rake/clean"
require "stringex"
require "yaml"
require "fileutils"

# Configuration

# Set "rake watch" as default task
task :default => :preview

# Load the configuration file
CONFIG = YAML.load_file("_config.yml")

# Get and parse the date
DATE = Time.now.strftime("%Y-%m-%d")

# Directories
POSTS = "source/_posts"
DRAFTS = "source/_drafts"

# Tasks

desc "New draft post"
task :draft do |t|

  title    = get_stdin("What is the title of your post? ")
  filename = "source/_drafts/#{title.to_url}.markdown"

  puts "Creating new draft: #{filename}" 
  open(filename, "w") do |post|
    post.puts "---"
    post.puts "layout: post"
    post.puts "title: \"#{title.gsub(/&/,'&amp;')}\""
    # post.puts "date: #{Time.now.strftime('%Y-%m-%d %H:%M')}"
    post.puts "categories: "
    post.puts "..."
  end

end

desc "Move a post from _drafts to _posts"
task :publish do
  extension = "markdown"
  files = Dir["#{DRAFTS}/*.#{extension}"]
  files.each_with_index do |file, index|
    puts "#{index + 1}: #{file}".sub("#{DRAFTS}/", "")
  end
  print "> "
  number = $stdin.gets
  if number =~ /\D/
    filename = files[number.to_i - 1].sub("#{DRAFTS}/", "")
    FileUtils.mv("#{DRAFTS}/#{filename}", "#{POSTS}/#{DATE}-#{filename}")
    puts "#{filename} was moved to '#{POSTS}'."
  else
    puts "Please choose a draft by the assigned number."
  end
end

desc "Preview the site with Jekyll, show drafts and limit posts"
task :preview do

  puts "Previewing the site locally with Jekyll."

  jekyllPid  = Process.spawn("jekyll build --watch --drafts --limit_posts 10")

  trap("INT") {
    [jekyllPid].each { |pid| Process.kill(9, pid) rescue Errno::ESRCH }
    exit 0
  }

  [jekyllPid].each { |pid| Process.wait(pid) }

end

task :build do
  puts "\nBuilding the production version of the site ..."
  system "jekyll build"
end

desc "Copy vignettes"
task :vignettes do
  puts "\nCopying Cartographer vignette..."
  FileUtils.cp("/Users/lmullen/dev/cartographer/inst/doc/cartographer.html",
               "./source/research/cartographer/index.html")
  puts "\nCopying Internet Archive vignette..."
  FileUtils.cp("/Users/lmullen/dev/internetarchive/inst/doc/internet-archive.html",
               "./source/research/internetarchive/index.html")
  
end

task :rsync do
  puts "\nDeploying the site via rsync..."

  ssh_port       = "22"
  ssh_user       = "lincolnm@lincolnmullen.org"
  rsync_delete   = true
  rsync_options  = "--checksum --stats --info=progress2 -avz -e"
  public_dir     = "public" 
  document_root  = "~/public_html/lincolnmullen.com"

  exclude = ""
  if File.exists?('./rsync-exclude')
    exclude = "--exclude-from '#{File.expand_path('./rsync-exclude')}'"
  end

  system("rsync #{rsync_options} 'ssh -p #{ssh_port}' #{exclude} #{"--delete" unless rsync_delete == false} #{public_dir}/ #{ssh_user}:#{document_root}")
end

desc "Build and deploy the production version of the site"
task :deploy => [:build, :rsync]

def get_stdin(message)
  print message
  STDIN.gets.chomp
end

CLOBBER.include('public/*')

desc "Edit drafts"
task :write do
  sh %[mvim source/_drafts/*]
end
