# Modified from Jason Heppler:
# https://github.com/hepplerj/jekyll-blog/blob/master/Rakefile
# and from Jeff McFadden:
# http://jeffmcfadden.com/blog/2011/04/13/rsync-your-jekyll/

desc 'build and deploy to lincolnmullen.com via rsync'
task :deploy do
  # uploads changes via SSH and rsync
  # does NOT delete other files on the server
  puts 'BUILDING AND DEPLOYING TO LINCOLNMULLEN.COM'
  sh "time jekyll && rsync --progress -artze ssh _site/* lam:/home/lincolnm/public_html/"
  puts 'Successfully built and deployed!'
end