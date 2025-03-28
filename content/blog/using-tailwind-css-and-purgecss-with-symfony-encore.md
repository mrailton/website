---
title: Using Tailwind CSS and PurgeCSS with Symfony Encore
date: 2020-08-28 23:45:00
---
Recently I've started playing around with the [Symfony](https://symfony.com) framework as a more lightweight alternative to Laravel (in case you didn't know, Laravel contains a lot of functionality that many apps don't need and has a lot of *magic* in the background).

One thing that I really liked about Laravel was Laravel Mix, and how easy Mix made it to integrate the utility-first CSS framework [Tailwind CSS](https://tailwindcss.com) and Purgecss (used for getting rid of CSS that's not needed). When I first started looking into [Symfony Encore](https://symfony.com/doc/current/frontend.html#webpack-encore) I found a nice guide on the Tailwind docs about getting Tailwind up and running on Encore. This was fantastic, but, there was a massive issue. The default CSS that was being generated was way in excess of 2MB in size, something that's simply far too big. Yes, TailwindCSS has support for PurgeCSS built in, but that means that you need to run your `npm run dev` or similar command each time you make a change that would involve changes to css.

I started to look further and came across a [guide](https://www.phproberto.com/en/41-integrating-purgecss-with-symfony-encore) by Roberto Segura on how to integrate Purgecss with Encore. I'll detail the steps below that will allow you to get TailwindCSS integrated to Symfony Encore using Purgecss to reduce the amount of CSS that's generated. This will assume that you have installed Symfony Encore but nothing else.

&nbsp;

First, let's install the dependencies, TailwindCSS, PostCSS and PurgeCSS:

```shell
npm install -D tailwindcss postcss-loader purgecss-webpack-plugin glob-all path
```

&nbsp;

Create *postcss.config.js* in the project root with the following content:

```js
module.exports = {
    plugins: [
        require('tailwindcss'),
    ],
};
```

&nbsp;

Update the top of your *webpack.config.js* file like so:

```js
const Encore = require('@symfony/webpack-encore');
const PurgeCssPlugin = require('purgecss-webpack-plugin');
const glob = require('glob-all');
const path = require('path');
```

(for some reason Encore defaults to using *var*, so I updated that too)

&nbsp;

Next, update the Encore configuration chain in *webpack.config.js* to include PostCssLoader:

```js
Encore
	...
  .enablePostCssLoader()
;
```



&nbsp;

So that we're not constantly having to run `npm run dev` during development, we will add the PurgeCSS inside an if check where we make sure to only run it on production (`npm run build`)

```js
if (Encore.isProduction()) {
  Encore.addPlugin(new PurgeCssPlugin({
        paths: glob.sync([
            path.join(__dirname, 'templates/**/*.html.twig')
        ]),
        defaultExtractor: (content) => {
            return content.match(/[\w-/:]+(?<!:)/g) || [];
        }
    }));
}
```



Next, you need to update your *assets/css/app.css* to load Tailwind:

```css
@import "tailwindcss/base";

@import "tailwindcss/components";

@import "tailwindcss/utilities";
```



&nbsp;

Now, by running Encore, your CSS will be generated with TailwindCSS and only the required CSS will be created:

```shell
npm run build
```



&nbsp;



You can also use the *watch* command to keep an eye on your CSS and JS files so that you can get fresh copies of your assets built each time you make a change, simply use:

```shell
npm run watch
```