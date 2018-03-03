# *sebatibot*

*sebatibot* is a Telegram bot (written in Go) *slash* brainchild of me and my **#1 favorite person in the world, [@mlsaito](https://github.com/mlsaito).** I won't go saying this is meant to be innovative or that *sebatibot* can do stuff you've never seen before because, honestly, as of writing, *sebatibot* is simply a project Mako and I *(plan to)* bond over during the weekends (aside from *__Monster Hunter: World__* and other stuff).

## Say, what does *sebatibot* do?

Nothing much at the moment :lmao: but here's our *partial* list:
- **remind**s you:

  > remind me to go to the bank and pay my credit card later at 3:30pm

  where *sebatibot* will respond and tell you

  > I will remind you to “go to the bank and pay my credit card later” at 3:30PM today.

- says **hello** to you!

  > hi sebatibot!

  and *sebatibot* will then say

  > hello!

## I want to contribute! (or build my own, similar to *sebatibot*)
### How to install?

Install `glide`, Golang dependency manager

```
### For MacOS
$ brew install glide
```

Clone the project to your Go `src` folder

```
$ cd $GOPATH/src
$ git clone https://github.com/kbleabres/sebatibot.git
$ cd sebatibot
```

Install all dependencies in project vendor folder.

```
$ glide install
```

### How to run?
Build the application to create a binary sebatibot

```
$ go build
```

Run the API server

```
$ GO_ENV_PORT=8000 ./sebatibot
```

OR Run the following docker commands

```
## Build the docker image
$ docker build -t sebatibot .

# Run the API in a container.
$ docker run -p 8000:8000 sebatibot
```

### Check the API output
```
$ curl http://localhost:8000/v1/users/1
```
Expected Output:

```
{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"age":35,"first_name":"bruce","last_name":"wayne"}
```

Special thanks to our reference [boilerplate](https://github.com/sjoshi6/go-rest-api-boilerplate).

### How to contribute?
### How to deploy?

## Dependencies


