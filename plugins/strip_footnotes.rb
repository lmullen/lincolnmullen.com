# Borrowed completely from
# http://penguindreams.org/blog/removing-footnotes-from-excerpts-in-jekyll/
# and modified for Pandoc footnotes
require 'nokogiri'

module Jekyll
  module StripFootnotesFilter

    def strip_footnotes(raw)
      doc = Nokogiri::HTML.fragment(raw.encode('UTF-8', :invalid => :replace, :undef => :replace, :replace => ''))

      for block in ['div', 'sup', 'a', 'section'] do
        doc.css(block).each do |ele|
          ele.remove if (ele['class'] == 'footnoteRef' or ele['class'] == 'footnotes')
        end
      end

      doc.inner_html

    end
  end
end

Liquid::Template.register_filter(Jekyll::StripFootnotesFilter)
