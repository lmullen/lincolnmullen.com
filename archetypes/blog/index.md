---
title: "{{ getenv "BLOG_TITLE" }}"
date: "{{ .Date  }}"
slug: "{{ slicestr .Name 11  }}"
draft: true
# tags:
# - 
# crosspost: 
#   url: ""
#   source: ""
# linkpost: ""
---
