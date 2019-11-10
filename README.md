# Gophercises: Golang exercises to become a gopher

In this repo I am committing my solutions of the awesome exercises from the gophercises website (https://gophercises.com/)

#### Progress

- [x] Quiz Game
- [x] URL Shortener
- [x] Choose Your Own Adventure
- [ ] HTML Link Parser
- [ ] Sitemap Builder
- [ ] Hacker Rank Problem
- [ ] CLI Task Manager
- [ ] Phone Number Normalizer
- [ ] Deck of Cards
- [ ] Blackjack
- [ ] Blackjack AI
- [ ] File Renaming Tool
- [ ] Quiet HN
- [ ] Recover Middleware
- [ ] Recover Middleware w/ Source Code
- [ ] Twitter Retweet Contest CLI
- [ ] Secrets API and CLI
- [ ] Image Transform Service
- [ ] Building Images (png & svg)
- [ ] Building PDFs 


#### Exercises

  - **Quiz Game** - Create a program to run timed quizes via the command line. 
  - **URL Shortener** - Code an http.Handler that forwards paths to other URLs (similar to Bitly).
  - **Choose Your Own Adventure** - Design a choose your own adventure book renderer that shows the story via a web application.
  - **HTML Link Parser** - Build a package to parse links (<a> tags) from an HTML file. 
  - **Sitemap Builder** -  Use the HTML link parser from ex. 4 to build a sitemap of public websites. 
  - **Hacker Rank Problem** - Code the solutions to a few string-related hacker rank problems. 
  - **CLI Task Manager** - Create a command line app to manage a TODO list stored in BoltDB.
  - **Phone Number Normalizer** - Write a program that will normalize an SQL table of phone numbers into a single format. 
  - **Deck of Cards** - Code a package used to build decks of cards with custom options, shuffling, and sorting. 
  - **Blackjack** - Using the deck of cards in ex. 9, create a blackjack game. 
  - **Blackjack AI** - Refactor the blackjack exercise into a package with an exported AI anyone can implement to create a bot that plays blackjack in a simulated game. 
  - **File Renaming Tool** - Build a tool used to rename files with a common pattern. Eg we might want to take many files with names like "Dog (1 of 100).jpg", "Dog (2 of 100).jpg", ... and rename them to "Dog_001.jpg", "Dog_002.jpg", ... 
  - **Quiet HN** - Given an existing web application that displays stories from Hacker News, we will look at ways to add concurrency and caching to the application while looking for race conditions and other potential issues. 
  - **Recover Middleware** -  Build HTTP middleware that will recover from any panics in an application, even if the response write has been partially written to, and then output the stack trace if the application is in development mode. 
  - **Recover Middleware w/ Source Code** - Expand upon the recover middleware and add links to source code along with syntax highlighting of the source code in order to make a useful development tool. 
  - **Twitter Retweet Contest CLI** - Create a CLI to help run a Twitter contest where users retweet a tweet for entry, and after some time you pick one or more of the users who retweeted as the winner. 
  - **Secrets API and CLI** - Create a package that handles storing and retrieving encrypted secrets like API keys. Then use that package to create a CLI that can be used to set and get secrets stored in a file in your home directory. 
  - **Image Transform Service** - Create a web server where a user can upload an image and then go through a guided process of choosing various image transformation options they prefer to get a final version of their image.
  - **Building Images (png & svg)** - Learn to create images in your Go code. First we use the standard library to build a simple PNG bar chart, then we explore how to use an awesome third party library to create a much more complex and compelling chart in SVG format.
  - **Building PDFs** - In this exercise we learn to create almost any PDF in Go. We start off by building an invoice with a dynamic set of line items, and then we move on to creating a course completion certificate for the Gophercises course!
