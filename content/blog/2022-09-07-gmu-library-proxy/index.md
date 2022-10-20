---
title: 'Bookmarklet for the GMU library proxy'
date: 2022-09-07
draft: false
tags:
- library
---

If you are a GMU student, faculty, or staff member, you can access library resources from off campus using a library proxy. Suppose you want to access this article:

Rosenzweig, Roy. “Scarcity or Abundance? Preserving the Past in a Digital Era.” _American Historical Review_ 108, no. 3 (2003): 735–62. <https://doi.org/10.1086/ahr/108.3.735>.

Following that DOI will take you to [this page](https://academic.oup.com/ahr/article/108/3/735/22504) at the Oxford University Press website. This version is behind a paywall, however: we can't see the full article.

{{< figure src="01-blocked.png" caption="A paywalled version of the article." >}}

<!--more-->

We can send that URL through the GMU proxy and get to a version where we have access by adding the following prefix:

```
http://mutex.gmu.edu/login?url=https://academic.oup.com/ahr/article/108/3/735/22504
```

But who wants to remember to do that? Instead we can create a bookmarklet that will send our URL through the proxy automatically. In your browser, you can create a new bookmark. You can name the bookmark whatever you want: I call mine `+GMU Proxy`. Then in the address field add the following JavaScript snippet.

```js {linenos=false}
javascript:(function()%20%7Bwindow.location=%22http://mutex.gmu.edu/login?url=%22%20+%20window.location.toString();%7D)()
```

Here is what adding the bookmark looks like in FireFox.

{{< figure src="02-bookmarklet.png" caption="Adding the bookmarklet in FireFox." >}}

If you are on the URL that is paywalled, you can then click the bookmarklet, which will take you to the proxied URL. (Note that unlike in the video below, you will likely have to sign in with your GMU password.)

{{< video src="03-proxy.mp4" caption="Clicking on the bookmarklet will take you to the proxied version." autoplay=true >}}
