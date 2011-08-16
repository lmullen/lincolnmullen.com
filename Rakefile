# Modified from Jason Heppler:
# https://github.com/hepplerj/jekyll-blog/blob/master/Rakefile
# and from Jeff McFadden:
# http://jeffmcfadden.com/blog/2011/04/13/rsync-your-jekyll/

desc 'copy current CV, build site, deploy to lincolnmullen.com'
task :default => [:copycv, :build, :deploy] do
  puts 'COPYING CV, BUILDING SITE, DEPLOYING SITE'
  # runs copy cv and deploy tasks
  puts 'Successfully updated website!'
end

desc 'build site using Jekyll'
task :build do
  # builds the site using Jekyll
  # Jekyll will get is configuration from _config.yaml
  puts 'BUILDING SITE USING JEKYLL'
  sh "time jekyll"
  puts 'Successfully built site!'
end

desc 'deploy to lincolnmullen.com via rsync'
task :deploy do
  # uploads changes via SSH and rsync
  # does NOT delete other files on the server
  puts 'BUILDING AND DEPLOYING TO LINCOLNMULLEN.COM'
  sh "time rsync --progress -artze ssh _site/* lam:/home/lincolnm/public_html/"
  puts 'Successfully deployed site!'
end

desc 'copy current CV to development directory'
task :copycv do
  puts 'COPYING CURRENT CV'
  sh "time cp ~/Documents/academic/CV/Mullen-cv.pdf ~/dev/lincolnmullen.com/docs/Mullen-cv.pdf"
  puts 'Successfully copied the CV.'
end
