# Countdown Server

I have recently started to learn the Go programming language and wanted to make something with it. I decided to make a Countdown word game solver (based on the UK game show). This could also help out in scrabble! The basic ideas is the user provides a combination of letters and the program returns all possible words that could be made with those letters.

This repo contains the server side logic to a react web app which can be found [here](https://wmorris92.github.io/countdown-web/).

## Usage

To start the server use go run command, specifying PORT:
```
PORT=8080 go run main.go
```

By visiting `localhost:8080/countdown/submittedString` where `submittedString` is a random stirng of any letters will return an array of words that can be made with those letters.

## TODOs

* Tests
