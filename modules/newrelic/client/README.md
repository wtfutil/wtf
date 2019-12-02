[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/yfronto/newrelic)
[![Build
status](https://travis-ci.org/yfronto/newrelic.svg)](https://travis-ci.org/yfronto/newrelic)

# New Relic API library for Go

This is a Go library that wraps the [New Relic][1] REST
API. It provides the needed types to interact with the New Relic REST API.

It's still in progress and I haven't finished the entirety of the API, yet. I
plan to finish all GET (read) operations before any POST (create) operations,
and then PUT (update) operations, and, finally, the DELETE operations.

The API documentation can be found from [New Relic][1],
and you'll need an API key (for some operations, an Admin API key is
required).

## USAGE

This library will provide a client object and any operations can be performed
through it. Simply import this library and create a client to get started:

```go
package main

import (
  "github.com/yfronto/newrelic"
)

var api_key = "..." // Required

func main() {
  // Create the client object
  client := newrelic.NewClient(api_key)

  // Get the applciation with ID 12345
  myApp, err := client.GetApplication(12345)
  if err != nil {
    // Handle error
  }

  // Work with the object
  fmt.Println(myApp.Name)

  // Some operations accept options
  opts := &newrelic.AlertEventOptions{
    // Only events with "MyProduct" as the product name
    Filter: newRelic.AlertEventFilter{
      Product: "MyProduct",
    },
  }
  // Get a list of recent events for my product
  events, err := client.GetAlertEvents(opts)
  if err != nil {
    // Handle error
  }
  // Display each event with some information
  for _, e := range events {
    fmt.Printf("%d -- %d (%s): %s\n", e.Timestamp, e.Id, e.Priority, e.Description)
  }
}
```

## Contributing

As I work to populate all actions, bugs are bound to come up. Feel free to
send me a pull request or just file an issue. Staying up to date with an API
is hard work and I'm happy to accept contributors.

**DISCLAIMER:** *I am in no way affiliated with New Relic and this work is
merely a convenience project for myself with no guarantees. It should be
considered "as-is" with no implication of responsibility. See the included
LICENSE for more details.*

[1]: http://www.newrelic.com