# GOQU

_WIP, and quite possibly a yak shaving expedition!_

Rewrite of [ASQ](https://github.com/Springest/ASQ) thath should feature
realtime query views that auto-update through DB polling and websockets.

Mostly because I like to learn writing some Go.

## Running locally

I use [Google's Go App Engine](https://cloud.google.com/appengine/docs/go/)
to run this locally. Eventually I want to get rid of that and run it
standalone in a Docker container. For early stage development it's nice
though.

### Install Go App Engine

I use homebrew to install this:

```bash
brew install go-app-engine-64
```

### Clone the repo

```bash
git clone git@github.com:foxycoder/goqu.git
```

### Run the app locally

`cd` into the app directory and run the app:

```bash
cd goqu
goapp serve
```

Open the app in your browser:

```bash
open http://localhost:8080
```

It listens for file changes and rebuilds the app.
