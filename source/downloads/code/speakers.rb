#!/usr/bin/env ruby
# encoding: utf-8

# Name::          speakers.rb
# Author::        Lincoln Mullen (mailto:lincoln@lincolnmullen.com)
# Copyright::     Copyright (c) 2013 Lincoln Mullen 
# License::       MIT License | http://lmullen.mit-license.org/

# This program finds all of the speakers in a TEI file and lists them 
# by the number of times that they speak.
# Usage: ./speakers.rb my-tei-file.xml

require 'nokogiri'          # for xml parsing
require 'pp'                # for a nicer output

# Get the file name to open
filename = ARGV[0]

# Open a hash to store our data
speakers = Hash.new

begin 
  # Open the file and parse it with Nokogiri
  doc = Nokogiri::XML(open(filename))
  # Find each instance of a <speaker> tag
  doc.search('speaker').each do |speaker|
    # Clean up the line breaks in the speaker's name
    name = speaker.content.gsub(/\n/,"")
    if speakers.has_key?(name)
      # If the speaker is already in our hash then add 1 to the count 
      # of utterances
      speakers[name] += 1
    else
      # If the speaker is not already in the hash then add the speaker 
      # with a count of 1
      speakers[name] = 1
    end
  end
rescue Errno::ENOENT
  # If the file we've been passed doesn't exist, catch the error
  puts "That file does not exist."
end

# Sort the hash of speakers by the number of times they speak, in 
# descending order, then print the output
pp speakers.sort_by { |name, lines| lines }.reverse
