require "rake/clean"
require "stringex"

desc "New draft post in its own branch"
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

desc "Create new post on master branch and open Vim"
task :new_post do |t|

  title    = get_stdin("What is the title of your post? ")
  filename = "source/_posts/#{Time.now.strftime('%Y-%m-%d')}-#{title.to_url}.markdown"

  puts "Creating new draft: #{filename}" 
  open(filename, "w") do |post|
    post.puts "---"
    post.puts "layout: post"
    post.puts "title: \"#{title.gsub(/&/,'&amp;')}\""
    post.puts "date: #{Time.now.strftime('%Y-%m-%d %H:%M')}"
    post.puts "categories: "
    post.puts "---"
  end

  editor = fork do
    exec "mvim #{filename}"
  end

  Process.detach(editor)

end

desc "Preview the site with Jekyll"
task :preview do

  puts "Previewing the site locally with Jekyll."

  jekyllPid  = Process.spawn("jekyll build --watch")

  trap("INT") {
    [jekyllPid].each { |pid| Process.kill(9, pid) rescue Errno::ESRCH }
    exit 0
  }

  [jekyllPid].each { |pid| Process.wait(pid) }

end

desc "Build the production version of the site"
task :build do
  puts "\nBuilding the production version of the site ..."
  ok_failed system "jekyll build"
end

desc "Copy CV"
task :cv do
  puts "\nCopying the CV..."
  FileUtils.cp("/home/lmullen/acad/cv/Mullen-cv.pdf", 
               "./source/downloads/docs/Mullen-cv.pdf")
end

desc "Deploy the site via rsync"
task :rsync do
  puts "\nDeploying the site via rsync..."

  ssh_port       = "22"
  ssh_user       = "lincolnm@lincolnmullen.org"
  rsync_delete   = true
  rsync_options  = "--checksum --stats -avze"
  public_dir     = "public" 
  document_root  = "~/public_html/lincolnmullen.com"

  exclude = ""
  if File.exists?('./rsync-exclude')
    exclude = "--exclude-from '#{File.expand_path('./rsync-exclude')}'"
  end

  ok_failed system("rsync #{rsync_options} 'ssh -p #{ssh_port}' #{exclude} #{"--delete" unless rsync_delete == false} #{public_dir}/ #{ssh_user}:#{document_root}")
end

desc "Build and deploy the production version of the site"
task :deploy => [:build, :rsync]

def get_stdin(message)
  print message
  STDIN.gets.chomp
end

def ok_failed(condition)
  if (condition)
    puts "OK"
  else
    puts "FAILED"
  end
end

CLOBBER.include('public/*')

