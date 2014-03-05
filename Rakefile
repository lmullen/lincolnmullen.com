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

desc "Preview the site with Jekyll"
task :preview do

  puts "Previewing the site locally with Jekyll."

  jekyllPid  = Process.spawn("jekyll serve --watch")

  trap("INT") {
    [jekyllPid].each { |pid| Process.kill(9, pid) rescue Errno::ESRCH }
    exit 0
  }

  [jekyllPid].each { |pid| Process.wait(pid) }

end

desc "Build the production version of the site"
task :build do
  puts "Building the production version of the site ..."
  system "jekyll build"
end

desc "Copy CV"
task :cv do
  puts "Copying the CV..."
  FileUtils.cp("/home/lmullen/acad/cv/Mullen-cv.pdf", 
               "./source/downloads/docs/Mullen-cv.pdf")
end

desc "Deploy the site to Amazon S3"
task :amazon_s3 do
  puts "Deploying the site to Amazon S3..."
  system "s3_website push --site public"
end

desc "Deploy the site via rsync"
task :rsync do
  puts "Deploying the site via rsync..."

  ssh_port       = "22"
  ssh_user       = "lincolnm@lincolnmullen.org"
  rsync_delete   = true
  public_dir      = "public" 
  document_root  = "~/public_html/lincolnmullen.com"

  exclude = ""
  if File.exists?('./rsync-exclude')
    exclude = "--exclude-from '#{File.expand_path('./rsync-exclude')}'"
  end

  system("rsync -avze 'ssh -p #{ssh_port}' #{exclude} #{"--delete" unless rsync_delete == false} #{public_dir}/ #{ssh_user}:#{document_root}")
end

desc "Build and deploy the production version of the site"
task :production => [:cv, :build, :rsync]

def get_stdin(message)
  print message
  STDIN.gets.chomp
end

CLOBBER.include('public/*')

