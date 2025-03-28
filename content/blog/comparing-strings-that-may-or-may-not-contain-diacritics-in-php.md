---
title: Comparing strings that may or may not contain diacritics in PHP
date: 2021-08-06 19:00:00
---
Today I ran into something that really had me scratching my head, I had to compare a string from a form against a string from the database. Clearly that's not where the issue was as it's a pretty simple thing in PHP, what had me scratching my head was that I needed to account for [diacritics](https://en.wikipedia.org/wiki/Diacritic) possibly being in 1 string but not in the other.

I spent quite some time looking online but eventually took to twitter and asked the wondrous PHP community for help
<blockquote class="twitter-tweet"><p lang="en" dir="ltr">Ok, taking a complete blank and need some <a href="https://twitter.com/hashtag/php?src=hash&amp;ref_src=twsrc%5Etfw">#php</a> help. Need to compare 2 strings that may or may not contain diacritics. Example, Seán matches Sean. Don&#39;t know why I can&#39;t figure this one out, anyone any ideas?</p>&mdash; Mark Railton (@markrailton84) <a href="https://twitter.com/markrailton84/status/1423661524883812355?ref_src=twsrc%5Etfw">August 6, 2021</a></blockquote> <script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script> 

Within minutes I had a couple of people offering suggestions and health conversation ensued. I settled on a [solution](https://twitter.com/derickr/status/1423665598832254977) by  [Derick Rethans](https://twitter.com/derickr) That uses the [Collator](http://docs.php.net/manual/en/class.collator.php) class from the `intl` extension. I took the example provided by Derick and tweaked it just a bit to suit how I wanted it, snippet of which is below

```php
$c = new Collator('en');
$c->setStrength(Collator::PRIMARY);

if ($c->getSortKey('Sean') !== $c->getSortKey('Seán') {
    return false;
}
```

To give a bit more context on this, Let's say we have a user called Sean. Sometimes people called Sean will spell it `Sean` but others may spell it `Seán` with the Irish diacritic `Fada`. Both of these people are called Sean and both spellings are seen as correct, however when doing a direct comparison in PHP (or any other language really) you'll end up getting a mismatch if you try using the equals operator. For the task I've been working on, it was important that we allow for the same person possibly having the Fada in their name in the database, but then maybe not entering it another time in a different form.

Thanks to Derick, Ben and the others that posted possible solutions on the twitter thread. It really helped and thankfully I was able to move on with the task.