require "stringex"

# Concept borrowed from Octopress:
# https://github.com/imathis/octopress/blob/master/Rakefile
desc "New draft post"
task :new_draft, :title do |t, args|
  if args.title
    title = args.title
  else
    title = get_stdin("What is the title of your post? ")
  end

  filename = "source/_drafts/#{Time.now.strftime('%Y-%m-%d')}-#{title.to_url}.markdown"
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

def get_stdin(message)
  print message
  STDIN.gets.chomp
end
