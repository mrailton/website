---
title: Getting Started With PHP Development, A Guide For Beginners.
date: 2014-07-13 23:00:00
---

##### **TL:DR, scroll to last paragraph**

When starting a new PHP project there is a world of decisions that you need to make including if you are going to use a framework, what version of PHP you are going to target and who is going to use and also maintain this app. These decisions get easier overtime, but as a new developer you may feel completely overwhelmed, especially with points number one and two, and this is what I am going to focus on now. At the end of 2013 I was asked to create a web application to record details of visitors to a church. I was provided with the scope of the project which was to basically digitise their welcome card and provide a way to automatically subscribe visitors to the email list if they wished. In this article I am going to explain my decision making process on how I came to write the application the way I did and hope that it helps other new developers.

### CMS vs no CMS

As this was my first development gig my initial thought was to build something out using Joomla as the church’s website was based on Joomla and this would have made for an easy integration, but that just did not sit right with me. After much debating I decided that I was not going to use Joomla (or any other CMS for that matter), instead I was going to use this as a tool to learn PHP development.

### To framework or not

Once I had decided on using PHP without a CMS I then started looking into different frameworks. All frameworks promise more or less the same, to make building out a web application a more painless process. This excited me as I wanted to build this out real quick and not get into too much of the nitty gritty. It quickly became evident that different frameworks had their own supporters and asking for recommendations was like asking for recommendations on the best pizza topping, everyone has their favourites and often are not willing to try anything else. I also realised that if I started with a PHP framework it was going to be much harder to break out of using the framework if I did not know much of the underlying PHP code.

### Framework vs Framework

During my decision making I tried comparing various frameworks against each other and realised that many of them approach similar tasks in completely different ways, some tried to stay close to pure PHP whilst making it easier (hint, easier is rarely better) whilst others went the complete opposite direction and require the developer has at least a decent knowledge of how to use the linux command line. Eventually, after about a week or so of research I decided that the best way forward was to use PHP in its purest form, without any frameworks at all. Whilst this meant that everything would have to be written from scratch, I found some fantastic tutorials online that helped me understand the build process and one that provided the backbone of the whole project.

### PHP version selection

Now that I knew that a framework was not the way to go, I was left with a decision as to what version of PHP I was going to target. For such a small application this is generally an easy choice to make, you go with the currently available version of PHP that is supported on the customers web server. In this case however it was not as simple, the church that I was working with is hosted with a company that offers 3 versions of PHP, 5.3 5.4 and 5.5. In the end I cheated a little, I just went with 5.3 as this is the default version on their existing server, however having now completed the application I would have used PHP 5.5 (explanation to come).

### Summary

So, the decisions had been made, I was going to build out a web application written purely in PHP targeting version 5.3 and had found several tutorials that helped along with the invaluable assistance of the PHP community at large. Now all that was left with was to put in some hard graft and a lot of head scratching. I honestly believe that by starting learning PHP and NOT a framework that even as a relative beginner I have a much better grasp of the PHP development cycle than if I had gone with a framework (or CMS for that matter) right from the start and I am eternally thankful for the people that had a role in steering me away from using such a framework.

### Points to ponder

This leads me to the main point of this post. The PHP development community is very large and everyone has their own takes on how you should start out (including me). There are an alarmingly high number of people that have the opinion that learning PHP without starting on a framework is a waste of time and that you can be up and running quicker if you just use a framework. Whilst there may be some truth to that (I admit you may get your application up quicker when using a framework and third party libraries) I strongly believe it is a very bad idea to jump straight to using a framework. I realize that this reads as an attack on frameworks but note, it is not. I believe frameworks like Codeigniter, Laravel and Zend (to name but a few of the great many) have a great role to play in the web development community, however these should not be used until you know what you are doing. If you do not know how to create a simple login system, then using something as complex as laravel can cause you a world of pain. To finish, I want to leave you with a comment that I have had spoken over me many, many times in my (so far) 30 short years on this earth, WALK before you start to run. Learn the fundamentals of PHP, that way once you start using a framework you are going to know what is going on under the hood and also appreciate the ease of development a lot more.
