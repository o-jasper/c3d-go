[![Stories in Ready](https://badge.waffle.io/project-douglas/c3d-go.png?label=ready&title=Ready)](https://waffle.io/project-douglas/c3d-go)

## Introduction

C3D => Contract Controlled Content Distribution

More coming.

## How to play

To setup, you'll have to do a little mix of `go` and `git`.  Install go, setup your `$GOPATH`, and `go get github.com/ethereum/go-ethereum`. I also recommend adding `$GOPATH` to your `$PATH`, since then you can call up the executables right away from anywhere.

Now, that should put two repos in `$GOPATH/src/github.com/ethereum`.  Go into both, and use git to add the project-douglas versions as remotes:

eg.

```
cd $GOPATH/src/github.com/ethereum/eth-go
git add remote pd git@github.com:project-douglas/eth-go
git checkout -b pd
git pull pd develop
```

That adds the project-douglas repo as a remote, checkouts out a new pd branch, and pulls the pd repo down to that branch.  If you want to switch back to the real ethereum, just do `git checkout develop`

Of course, do the exact same thing for `$GOPATH/src/github.com/ethereum/go-ethereum`.

Basically eth-go is a highly modular library for a full ethereum node, and go-ethereum co-ordinates the library into startup/shutdown routines convenient for the headless and gui clients. We use go-ethereum because it has some nice helper functions for startup/shutdown.

Now, grab c3d-go: `go get github.com/project-douglas/c3d-go`. That will install it.  If you make changes and want to re-install, just hit `go install` in the c3d-go repo. Run it with `c3d-go`, or `$GOPATH/bin/c3d-go` if you must.

## Notes

We're using a custom blockchain with two addresses and lots of funds in each.  The keys are in `keys.txt` and both are loaded.  A simple way to get everyone funds for now is to create a new address, use one of the genesis keys to send yourself some funds, and go from there.  This will all be automated soon

Also, you can kill transmission with `c3d-go --killTransmission`.  I decided that might be convenient. 


## Features

c3d-go doesn't do much yet.  It stores an infohash in a contract, waits for it to be mined, grabs the infohash from the blockchain, and throws it into the torrent client.  You can monitor the torrent client at http://localhost:9091

Much more, coming soon ;)


## Contributing

1. Fork the repository.
2. Create your feature branch (`git checkout -b my-new-feature`).
3. Add Tests (and feel free to help here since I don't (yet) really know how to do that.).
4. Commit your changes (`git commit -am 'Add some feature'`).
5. Push to the branch (`git push origin my-new-feature`).
6. Create new Pull Request.

## License

Modified MIT, see LICENSE.md
