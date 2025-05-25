# Agent Claude

Started from following along with Thorsten Ball: https://ampcode.com/how-to-build-an-agent

# Setup
- You'll need the [Go programming language](https://go.dev/learn/)

Create an .env file to store your api key
```
cp .env.template .env
```

# Usage

Start a chat
```
go run ./src
```

Use Haiku
```
go run ./src haiku
```

Test out its tools
- What version of go are we using?
- Create a javascript file that contains a fizzbizz program that I can run with Nodejs.
- Edit the file so that it only prints to 30.
