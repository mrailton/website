---
title: Set default Node.js version in nvm
date: 2019-10-22 22:00:00
---
Today seen Node.js 12.13.0 get flagged as the new LTS release. As someone who uses nvm to manage having different versions of node installed this can potentially pose a bit of an issue, you see, I tend to build my projects against the LTS release, but when a new LTS comes out, nvm will switch to it. In order to keep the default version of Node.js as 10.16.3 (current release on the 10.x branch) as not 12.13.0 you need to alias the default install to 10.16.3.

This is actually a very simple step.

First, confirm the current version of Node in use is 12.13.0

```shell
╭─mark@ZeroCool ~
╰─$ node -v
v12.13.0
```

Now, set the default alias to 10.16.3

```shell
╭─mark@ZeroCool ~
╰─$ nvm alias default 10.16.3
default -> 10.16.3 (-> v10.16.3)
```

Then close the terminal window and re-open to confirm that Node.js is now using 10.16.3 as default

```shell
╭─mark@ZeroCool ~
╰─$ node -v
v10.16.3
```

Nice and simple, now when I open a new terminal I will always be using Node.js 10.16.3