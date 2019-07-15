# TransmissionRPC

[![GoDoc](https://godoc.org/github.com/hekmon/transmissionrpc?status.svg)](https://godoc.org/github.com/hekmon/transmissionrpc) [![Go Report Card](https://goreportcard.com/badge/github.com/hekmon/transmissionrpc)](https://goreportcard.com/report/github.com/hekmon/transmissionrpc)

Golang bindings to Transmission (bittorent) RPC interface (Work in Progress).

Even if there is some high level wrappers/helpers, the goal of this lib is to stay close to the original API in terms of methods and payloads while enhancing certain types to be more "golangish": timestamps are converted from/to time.Time, numeric durations in time.Duration, booleans in numeric form are converted to real bool, etc...

Also payload generation aims to be precise: when several values can be added to a payload, only instanciated values will be forwarded (and kept !) to the final payload. This means that the default JSON marshalling (with omitempty) can't always be used and therefor a manual, reflect based, approach is used to build the final payload and accurately send what the user have instanciated, even if a value is at its default type value.

This lib follow the [transmission v15 RPC specification](https://github.com/transmission/transmission/blob/2.9x/extras/rpc-spec.txt#L639).

## Getting started

First the main client object must be instantiated with [New()](https://godoc.org/github.com/hekmon/transmissionrpc#New). In its basic form only host/ip, username and password must be provided. Default will apply for port (`9091`) rpc URI (`/transmission/rpc`) and others values.

```golang
transmissionbt := transmissionrpc.New("127.0.0.1", "rpcuser", "rpcpass", nil)
```

But advanced values can also be configured to your liking using [AdvancedConfig](https://godoc.org/github.com/hekmon/transmissionrpc#AdvancedConfig).
Each value of `AdvancedConfig` with a type default value will be replaced by the lib default value, so you can set only the ones you want:

```golang
transmissionbt := transmissionrpc.New("bt.mydomain.net", "rpcuser", "rpcpass",
    &transmissionrpc.AdvancedConfig{
        HTTPS: true,
        Port:  443,
    })
```

The remote RPC version can be checked against this library before starting to operate:

```golang
ok, serverVersion, serverMinimumVersion, err := transmission.RPCVersion()
if err != nil {
    panic(err)
}
if !ok {
    panic(fmt.Sprintf("Remote transmission RPC version (v%d) is incompatible with the transmission library (v%d): remote needs at least v%d",
        serverVersion, transmissionrpc.RPCVersion, serverMinimumVersion))
}
fmt.Printf("Remote transmission RPC version (v%d) is compatible with our transmissionrpc library (v%d)\n",
    serverVersion, transmissionrpc.RPCVersion)
```

## Features

* [Torrent Requests](#torrent-requests)
  * [Torrent Action Requests](#torrent-action-requests)
    * [x] torrent-start
    * [x] torrent-start-now
    * [x] torrent-stop
    * [x] torrent-verify
    * [x] torrent-reannounce
  * [Torrent Mutators](#torrent-mutators)
    * [x] torrent-set
  * [Torrent Accessors](#torrent-accessors)
    * [x] torrent-get
  * [Adding a Torrent](#adding-a-torrent)
    * [x] torrent-add
  * [Removing a Torrent](#removing-a-torrent)
    * [x] torrent-remove
  * [Moving a Torrent](#moving-a-torrent)
    * [x] torrent-set-location
  * [Renaming a Torrent path](#renaming-a-torrent-path)
    * [x] torrent-rename-path
* [Session Requests](#session-requests)
  * [Session Arguments](#session-arguments)
    * [x] session-set
    * [x] session-get
  * [Session Statistics](#session-statistics)
    * [x] session-stats
  * [Blocklist](#blocklist)
    * [ ] blocklist-update
  * [Port Checking](#port-checking)
    * [x] port-test
  * [Session Shutdown](#session-shutdown)
    * [ ] session-close
  * [Queue Movement Requests](#queue-movement-requests)
    * [ ] queue-move-top
    * [ ] queue-move-up
    * [ ] queue-move-down
    * [ ] queue-move-bottom
  * [Free Space](#free-space)
    * [x] free-space

### Torrent Requests

#### Torrent Action Requests

Each rpc methods here can work with ID list, hash list or `recently-active` magic word. Therefor, there is 3 golang method variants for each of them.

```golang
transmissionbt.TorrentXXXXIDs(...)
transmissionbt.TorrentXXXXHashes(...)
transmissionbt.TorrentXXXXRecentlyActive()
```

* torrent-start

Check [TorrentStartIDs()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStartIDs), [TorrentStartHashes()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStartHashes) and [TorrentStartRecentlyActive()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStartRecentlyActive).

Ex:

```golang
err := transmissionbt.TorrentStartIDs([]int64{55})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println("yay")
}
```

* torrent-start-now

Check [TorrentStartNowIDs()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStartNowIDs), [TorrentStartNowHashes()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStartNowHashes) and [TorrentStartNowRecentlyActive()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStartNowRecentlyActive).

Ex:

```golang
err := transmissionbt.TorrentStartNowHashes([]string{"f07e0b0584745b7bcb35e98097488d34e68623d0"})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println("yay")
}
```

* torrent-stop

Check [TorrentStopIDs()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStopIDs), [TorrentStopHashes()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStopHashes) and [TorrentStopRecentlyActive()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentStopRecentlyActive).

Ex:

```golang
err := transmissionbt.TorrentStopIDs([]int64{55})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println("yay")
}
```

* torrent-verify

Check [TorrentVerifyIDs()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentVerifyIDs), [TorrentVerifyHashes()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentVerifyHashes) and [TorrentVerifyRecentlyActive()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentVerifyRecentlyActive).

Ex:

```golang
err := transmissionbt.TorrentVerifyHashes([]string{"f07e0b0584745b7bcb35e98097488d34e68623d0"})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println("yay")
}
```

* torrent-reannounce

Check [TorrentReannounceIDs()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentReannounceIDs), [TorrentReannounceHashes()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentReannounceHashes) and [TorrentReannounceRecentlyActive()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentReannounceRecentlyActive).

Ex:

```golang
err := transmissionbt.TorrentReannounceRecentlyActive()
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println("yay")
}
```

#### Torrent Mutators

* torrent-set

Mapped as [TorrentSet()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentSet).

Ex: apply a 1 MB/s limit to a torrent.

```golang
uploadLimited := true
uploadLimitKBps := int64(1000)
err := transmissionbt.TorrentSet(&transmissionrpc.TorrentSetPayload{
    IDs:           []int64{55},
    UploadLimited: &uploadLimited,
    UploadLimit:   &uploadLimitKBps,
})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println("yay")
}
```

There is a lot more [mutators](https://godoc.org/github.com/hekmon/transmissionrpc#TorrentSetPayload) available.

#### Torrent Accessors

* torrent-get

All fields for all torrents with [TorrentGetAll()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentGetAll):

```golang
torrents, err := transmissionbt.TorrentGetAll()
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println(torrents) // meh it's full of pointers
}
```

All fields for a restricted list of ids with [TorrentGetAllFor()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentGetAll):

```golang
torrents, err := transmissionbt.TorrentGetAllFor([]int64{31})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    fmt.Println(torrents) // meh it's still full of pointers
}
```

Some fields for some torrents with the low level accessor [TorrentGet()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentGet):

```golang
torrents, err := transmissionbt.TorrentGet([]string{"status"}, []int64{54, 55})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    for _, torrent := range torrents {
        fmt.Println(torrent.Status) // the only instanciated field, as requested
    }
}
```

Some fields for all torrents, still with the low level accessor [TorrentGet()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentGet):

```golang
torrents, err := transmissionbt.TorrentGet([]string{"id", "name", "hashString"}, nil)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    for _, torrent := range torrents {
        fmt.Println(torrent.ID)
        fmt.Println(torrent.Name)
        fmt.Println(torrent.HashString)
    }
}
```

Valid fields name can be found as JSON tag on the [Torrent](https://godoc.org/github.com/hekmon/transmissionrpc#Torrent) struct.

#### Adding a Torrent

* torrent-add

Adding a torrent from a file (using [TorrentAddFile](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentAddFile) wrapper):

```golang
filepath := "/home/hekmon/Downloads/ubuntu-17.10.1-desktop-amd64.iso.torrent"
torrent, err := transmissionbt.TorrentAddFile(filepath)
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    // Only 3 fields will be returned/set in the Torrent struct
    fmt.Println(*torrent.ID)
    fmt.Println(*torrent.Name)
    fmt.Println(*torrent.HashString)
}
```

Adding a torrent from an URL (ex: a magnet) with the real [TorrentAdd](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentAdd) method:

```golang
magnet := "magnet:?xt=urn:btih:f07e0b0584745b7bcb35e98097488d34e68623d0&dn=ubuntu-17.10.1-desktop-amd64.iso&tr=http%3A%2F%2Ftorrent.ubuntu.com%3A6969%2Fannounce&tr=http%3A%2F%2Fipv6.torrent.ubuntu.com%3A6969%2Fannounce"
torrent, err := btserv.TorrentAdd(&transmissionrpc.TorrentAddPayload{
    Filename: &magnet,
})
if err != nil {
    fmt.Fprintln(os.Stderr, err)
} else {
    // Only 3 fields will be returned/set in the Torrent struct
    fmt.Println(*torrent.ID)
    fmt.Println(*torrent.Name)
    fmt.Println(*torrent.HashString)
}
```

Which would output:

```raw
55
ubuntu-17.10.1-desktop-amd64.iso
f07e0b0584745b7bcb35e98097488d34e68623d0
```

#### Removing a Torrent

* torrent-remove

Mapped as [TorrentRemove()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentRemove).

#### Moving a Torrent

* torrent-set-location

Mapped as [TorrentSetLocation()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentSetLocation).

#### Renaming a Torrent path

* torrent-rename-path

Mapped as [TorrentRenamePath()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.TorrentRenamePath).

### Session Requests

#### Session Arguments

* session-set

Mapped as [SessionArgumentsSet()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.SessionArgumentsSet).

* session-get

Mapped as [SessionArgumentsGet()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.SessionArgumentsGet).

#### Session Statistics

* session-stats

Mapped as [SessionStats()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.SessionStats).

#### Blocklist

* blocklist-update _(to do)_

#### Port Checking

* port-test

Mapped as [PortTest()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.PortTest).

Ex:

```golang
    st, err := transmissionbt.PortTest()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
    if st {
        fmt.Println("Open!")
    }
```

#### Session Shutdown

* session-close _(to do)_

#### Queue Movement Requests

* queue-move-top _(to do)_
* queue-move-up _(to do)_
* queue-move-down _(to do)_
* queue-move-bottom _(to do)_

#### Free Space

* free-space

Mappped as [FreeSpace()](https://godoc.org/github.com/hekmon/transmissionrpc#Client.FreeSpace).

Ex: Get the space available for /data.

```golang
    freeSpace, err := transmissionbt.FreeSpace("/data")
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    } else  {
        fmt.Printd("%s | %d | %v", freeSpace, freeSpace, freeSpace)
    }
}
```

For more information about the freeSpace type, check the [ComputerUnits](https://github.com/hekmon/cunits) library.
