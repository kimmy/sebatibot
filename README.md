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

## I want to run this locally! (or build my own, similar to *sebatibot*)
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
Run the application
```
$ go run main.go
```

OR Run the following docker commands

```
## Build the docker image
$ docker build -t sebatibot .

# Run the API in a container.
$ docker run -p 3000:3000 sebatibot
```

### Check the site
```
$ curl http://localhost:3000
```
Expected Output:

```
Welcome!
```

### Setup webhook for local env
Download ngrok [here](https://ngrok.com/download).
Move ngrok to sebatibot folder.
Execute ngrok to generate public URL.
```
$ ./ngrok http 3000
```
Copy URL to main.go on this line:
```
webhookUrl := "https://38e44fc0.ngrok.io/" + os.Getenv("TOKEN")
```
Update your .env with the appropriate TOKEN.
Run main.go and send a message to Sebatibot on Telegram.

Special thanks to our reference [boilerplate](https://github.com/sjoshi6/go-rest-api-boilerplate).

### How to contribute?
Please read our [CONTRIBUTING.md](https://github.com/kbleabres/sebatibot/blob/master/.github/CONTRIBUTING.md) before contributing.

### How to deploy?

## Dependencies

### Notes
When creating new bot, set privacy setting to `false` through `BotFather` so that bot can get data from a channel.  

Send message to `BotFather`:

```
/setprivacy
```

Reference [here](https://stackoverflow.com/questions/38565952/how-to-receive-messages-in-group-chats-using-telegram-bot-api).
