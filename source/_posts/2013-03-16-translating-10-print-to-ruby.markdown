---
layout: post
title: "The Challenge of Translating 10 PRINT to Ruby"
date: 2013-03-17 14:25
comments: true
categories: 
- digital humanities
tags:
- Ruby
- programming
- DHTNG
---

I spent the past two days at "[Digital Humanities: The Next
Generation][]," a conference held at Simmons College. [Patsy Baudoin][]
and [Nick Montfort][] gave a presentation about *[10 PRINT][]*, a book
they co-authored that turns a critical eye to a one-line computer
program written in [BASIC][] for the [Commodore 64][].

  [Digital Humanities: The Next Generation]: http://web.simmons.edu/~fairb/dhsymposium/program.html
  [Patsy Baudoin]: https://twitter.com/pbmit
  [Nick Montfort]: http://nickm.com/
  [10 PRINT]: http://10print.org/
  [BASIC]: http://en.wikipedia.org/wiki/BASIC
  [Commodore 64]: http://en.wikipedia.org/wiki/Commodore_64

```
10 PRINT CHR$(205.5+RND(1)); : GOTO 10
```

The *10 PRINT* authors explain how the code works, but here's a brief
explanation. The `PRINT` command outputs a character to the screen. The
randomization function selects between characters 205 and 206. Then the
`GOTO 10` command returns to the beginning of the line and runs the
command in an endless loop. The output of the program is diagonal maze.

{% img center /postimg/2013-03-16.10PRINT-1.png %}

<!--more-->

As part of the demonstration, Patsy and Nick brought two Commodore 64s
for the audience to try. That was fun: I'd never played with a
Commodore. While we were waiting our turn, Shane Landrum and I decided
to translate the program into another computer language. He chose Perl,
and I chose Ruby. [^1] In about 10 minutes, we tied in reproducing the
output. Here are our solutions:

[^1]: I learned after the fact that the *10 PRINT* book discusses ports
    to other BASIC systems and to Perl and JavaScript. Almost all of
    what I mention below is contained in the authors' dicussion, pp.
    52--58.

<blockquote class="twitter-tweet"><p>Here's the Perl equivalent of that code!my @<a href="https://twitter.com/a">a</a>=qw( / \ ); while (1) {print <a href="https://twitter.com/search/$a">$a</a>[rand @<a href="https://twitter.com/a">a</a>];sleep 0.01;} print "\n";<a href="https://twitter.com/search/%23dhtng">#dhtng</a></p>&mdash; Shane Landrum (@cliotropic) <a href="https://twitter.com/cliotropic/status/312987216248991745">March 16, 2013</a></blockquote>

<blockquote class="twitter-tweet"><p><a href="https://twitter.com/search/%23dhtng">#dhtng</a> Slightly cleaner: in OS X Terminal: perl -e 'my @<a href="https://twitter.com/a">a</a>=qw( / \ ); while (1) {print <a href="https://twitter.com/search/$a">$a</a>[rand @<a href="https://twitter.com/a">a</a>];}';</p>&mdash; Shane Landrum (@cliotropic) <a href="https://twitter.com/cliotropic/status/312987627991224322">March 16, 2013</a></blockquote>

<blockquote class="twitter-tweet"><p>Here's 10PRINT in Ruby as a one-liner. Run it in Terminal. ruby -e 'printf ["/","\\"][Random.new.rand(0..1)] while true' <a href="https://twitter.com/search/%23dhtng">#dhtng</a></p>&mdash; Lincoln Mullen (@lincolnmullen) <a href="https://twitter.com/lincolnmullen/status/312992092882354176">March 16, 2013</a></blockquote>
<script async src="//platform.twitter.com/widgets.js" charset="utf-8"></script>

While these solutions both worked as a port to another language, I
wasn't satisfied with mine as a *translation* into another language. Its
output didn't have the same maze-like appearance as the Commodore 64's,
nor did it have the same aesthetic as the BASIC program. Just as a
translator might attempt different renderings of a poem, I tried several
renderings of the code. I never found a translation that was at once
faithful to everything I liked in the BASIC program and fully idiomatic
[Ruby][]. In the spirit of the *10 PRINT* book and [critical code
studies][], here are some of the translations that I tried along with
commentary.

  [Ruby]: http://www.ruby-lang.org/en/
  [critical code studies]: http://criticalcodestudies.com/wordpress/

### Solution 1

``` ruby
character = ["/","\\"]

while true do
  select = Random.new.rand(0..1)
  print character[select]
end
```

My first inclination was to separate each logical part of the program on
its own line. In this first attempt, I was trying to reproduce only the
output of the Commodore 64 and not the style of the code. This solution
is idiomatic, though not very elegant, [Ruby][].

First, I create an array of characters, the forward slash (`/`) and the
backward slash (`\`). There are a two backslashes as the second element
of the array because Ruby uses that character as an escape sequence.

Second, I use a block of code between the `do` and `end` statements. A
[code block][] is a common element in Ruby programs, so the feel of the
program is more Ruby than BASIC. Because of the `while true` condition,
this block of code will be repeated forever. Notice that the BASIC
program uses the `GOTO 10` command to accomplish the same thing, but
since that is a peculiarity of BASIC, it cannot be literally translated
into Ruby or Perl. Both Shane and I chose to use `while` statements to
accomplish the same effect.

  [code block]: http://www.ruby-doc.org/docs/ProgrammingRuby/html/tut_containers.html

Third, the program assigns a random number between 0 and 1 to a
variable. Because both BASIC and Ruby share the same limitations, the
numbers that are generated are [pseudorandom][], rather than truly
random. Here `Random.new.rand` calls a method on a [class][] `Random` to
make a new object, then calls the function `rand` to generate the random
number. A key difference between Ruby and BASIC is that BASIC is an
imperative programming language, while Ruby is an object-oriented
language. I could have obscured this difference by substituting the
`rand` method for `Random.new.rand`. There is a noticeable etymological
similarity between BASIC's `RND` and Ruby's `rand`, but the apparent
similarity of function is deceptive. In Ruby `rand` is really a method
call on a different [object][], `Kernel`, since in Ruby [everything is
an object][].

  [pseudorandom]: http://en.wikipedia.org/wiki/Pseudorandomness
  [class]: http://ruby-doc.org/core-2.0/Random.html
  [object]: http://ruby-doc.org/core-2.0/Kernel.html#method-i-rand
  [everything is an object]: http://www.ruby-lang.org/en/about/

Fourth, the program uses the variable to select one of the characters
from the array and then prints it. BASIC's `PRINT` and Ruby's `print`
(same in Perl) are the same word and serve the same function, like a
cognate word in human languages. In Ruby, though, the command `print` doesn't
quite seem right. In most cases where I would output text, I would use
the command `puts`, which is more common in Ruby. But since `puts` adds
a new line and since `print` does not, we need the latter option to
produce the right output.

This is the output of solutions 1 through 3:

{% img center /postimg/2013-03-16.10PRINT-2.png %}

### Solution 2

The most obvious defect of the first translation is that the Ruby code
takes five lines to do the same thing that the BASIC code does in one.
Does each Catullan couplet require an entire English sonnet? The next
solution gets the code down to one line.

``` ruby
print ['/','\\'][Random.new.rand(0..1)] while true
```

The code is compressed by eliminating assignments to variables.  We've 
also eliminated the `do` code block. But because we still need the 
`while` logic, we've moved the condition to the end of the line. Though 
we can't use the same mechanism as BASIC to create an endless loop, at 
least by moving the `while` to end we get a structural similarity to 
BASIC.

### Solution 3

``` ruby
print ['/','\\'][rand(0..1)] while true
```

One of the features of the BASIC program is brevity, so `rand` is 
preferable to `Random.new.rand`. We emphasize the cognate, while making 
it more subtle that everything in Ruby is an object.

### Solution 4

The second most obvious defect of the translations so far is that the 
forward slash and backslash characters don't produce the correct output.
Because these glyphs don't occupy the full character width, they do not 
connect to make the maze. As the authors of *10 
PRINT* point out, the Commodore 64 had its own character set 
([PETSCII](http://en.wikipedia.org/wiki/PETSCII)) 
and a different keyboard than is now standard. We'd be better off with 
different characters.

``` ruby
print ["\u2571","\u2572"][rand(0..1)] while true
```

In this solution we use two [Unicode][] characters, [U+2571][] and
[U+2572][], which map more closely to the PETSCII characters. But just
as we have to use method calls to objects in Ruby rather than simple
functions, here we have to use Unicode rather than the limited character
set of the Commodore 64. The output looks much better:

  [Unicode]: http://www.unicode.org/
  [U+2571]: http://www.fileformat.info/info/unicode/char/2571/index.htm
  [U+2572]: http://www.fileformat.info/info/unicode/char/2572/index.htm

{% img center /postimg/2013-03-16.10PRINT-3.png %}

### Solution 5

We could leave our solution well enough alone. It's idiomatic Ruby, and
`print`, `rand`, and the loop logic each gesture to the BASIC program.

But one of the interesting features of the BASIC code is the logic that 
it uses to select the characters. Let's look at this more closely:

```
CHR$(205.5+RND(1))
```

The two PETSCII characters have the index numbers 205 and 206. By
starting with the number 205.5 and adding a random number between 0 and
1, the `CHR` function selects character 205 half the time, and character
206 half the time. This is quite ingenious: more ingenious than our
solution for Ruby needs to be, since Ruby is more powerful at dealing
with strings. Needlessly complicated code is just spaghetti code, but
much of the pleasure of reading code comes from understanding clever
solutions that get around limitations of the language. We can get a lot
closer to the BASIC with this Ruby code:

``` ruby
print (9585+rand(0..1)).chr(Encoding::UTF_8) while true
```

Here we're accessing the same Unicode characters as in solution 4, but
using a decimal rather than a hexadecimal system of enumeration. We want
characters [9585][U+2571] or [9586][U+2572]. We've stopped choosing our
characters from an array, which is a data structure not available in
BASIC. We start with 9585 and add either 0 or 1 to it. Then we call the
`chr` [method][] on that integer. Notice that though we can do without
arrays, once again we can't get away from Ruby's objects, since `chr` is
a method of the class `Integer`. The `chr` method turns the integer into
a character. We've gained another cognate, from `CHR` to `chr`. But
because Ruby, unlike the Commodore 64's BASIC, can handle a number of
encodings, we have to tell the method that we're using UTF-8. As a
result, we've lost some brevity.

  [method]: http://ruby-doc.org/core-2.0/Integer.html#method-i-chr

### Conclusion

None of these solutions reaches my goal of translating the code into
idiomatic Ruby that still preserves the notable features of the BASIC.
I'd be glad to know if you can come up with a better translation. This
exercise isn't just about generating random mazes: it's also a look into
the "elegance" of programming languages. As the *10 PRINT* authors point
out, code is not just the output of a program, but the culture that
surrounds the code as an artifact. Ruby is widely regarded as an elegant
language, which is why it's the first language that I've learned in
depth, with help from other [Rubyist historians][]. I think solution 4,
at least, is elegant Ruby. Solution 5, though, pays better tribute the
ingenuity of the original BASIC.

  [Rubyist historians]: http://hepplerj.github.com/rubyist-historian/

*******************

### Appendix: How to run this code

You can run each of these solutions for yourself in a terminal. The
easiest way to do this is to paste the entire line of code between the
`''` in this command, which interprets the code with Ruby.

``` bash
ruby -e 'CODE GOES HERE'
```

To quit the program, press `CTRL+C`.
