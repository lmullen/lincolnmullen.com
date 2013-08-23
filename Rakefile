require "rake/clean"
require "stringex"

desc "New draft post"
task :new_draft do |t|

  branch   = get_stdin("What is the name of the branch? ").to_url
  title    = get_stdin("What is the title of your post? ")
  filename = "source/_posts/#{Time.now.strftime('%Y-%m-%d')}-#{title.to_url}.markdown"

  puts "Checking out a new draft branch: draft/#{branch}"
  `git checkout -b draft/#{branch}`

  puts "Creating new draft: #{filename}" 
  open(filename, "w") do |post|
    post.puts "---"
    post.puts "layout: post"
    post.puts "title: \"#{title.gsub(/&/,'&amp;')}\""
    post.puts "date: #{Time.now.strftime('%Y-%m-%d %H:%M')}"
    post.puts "categories: "
    post.puts "---"
  end

end

desc "Preview the site with Jekyll and Compass"
task :preview do

  puts "Previewing the site locally with Jekyll and Compass."

  system "compass compile"
  jekyllPid  = Process.spawn("jekyll serve --watch")
  compassPid = Process.spawn("compass watch")

  trap("INT") {
    [jekyllPid, compassPid].each { |pid| Process.kill(9, pid) rescue Errno::ESRCH }
    exit 0
  }

  [jekyllPid, compassPid].each { |pid| Process.wait(pid) }

end

desc "Push the site to the development server"
task :push do
  puts "Building the site then pushing it to Amazon S3"
  system "compass compile"
  system "jekyll build --lsi"
  system "s3_website push"
end

def get_stdin(message)
  print message
  STDIN.gets.chomp
end

CLOBBER.include('public/*')
