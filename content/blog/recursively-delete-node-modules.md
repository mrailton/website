---
title: Recursively delete node_modules
date: 2019-09-12 23:00:00
---
Every now and then I like to take a backup of my main code directory. When doing this I generally will delete all the `node_modules` and `vendor` directories to vastly reduce the amount of space taken up. The only issue with this approach is that I can never remember the full command to use, so, this post is nothing more than a brain dump for me to be able to find the command quickly.

From within my primary code directory, I simply run the snippet below to remove the node modules directories.

```shell
find . -name "node_modules" -exec rm -rf '{}' +
find . -name "vendor" -exec rm -rf '{}' +
```