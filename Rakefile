require "rake/clean"
require "stringex"

task :version do
  `ruby --version`
  `which ruby`
end

desc "New draft post"
task :draft do |t|

  title    = get_stdin("What is the title of your post? ")
  filename = "source/_drafts/#{Time.now.strftime('%Y-%m-%d')}-#{title.to_url}.markdown"

  puts "Creating new draft: #{filename}" 
  open(filename, "w") do |post|
    post.puts "---"
    post.puts "layout: post"
    post.puts "title: \"#{title.gsub(/&/,'&amp;')}\""
    post.puts "date: #{Time.now.strftime('%Y-%m-%d %H:%M')}"
    post.puts "categories: "
    post.puts "..."
  end

end

desc "Preview the site with Jekyll"
task :preview do

  puts "Previewing the site locally with Jekyll."

  jekyllPid  = Process.spawn("jekyll build --watch --drafts --limit_posts 5")

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

desc "Copy cartographer vignette"
task :cartographer do
  puts "\nCopying Cartographer vignette..."
  FileUtils.cp("/Users/lmullen/dev/cartographer/inst/doc/cartographer.html",
               "./source/research/cartographer/index.html")
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

