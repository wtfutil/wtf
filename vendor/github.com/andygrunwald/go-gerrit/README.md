# go-gerrit

[![GoDoc](https://godoc.org/github.com/andygrunwald/go-gerrit?status.svg)](https://godoc.org/github.com/andygrunwald/go-gerrit)
[![Build Status](https://travis-ci.org/andygrunwald/go-gerrit.svg?branch=master)](https://travis-ci.org/andygrunwald/go-gerrit)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygrunwald/go-gerrit)](https://goreportcard.com/report/github.com/andygrunwald/go-gerrit)
[![codecov](https://codecov.io/gh/andygrunwald/go-gerrit/branch/master/graph/badge.svg)](https://codecov.io/gh/andygrunwald/go-gerrit)

go-gerrit is a [Go(lang)](https://golang.org/) client library for accessing the [Gerrit Code Review](https://www.gerritcodereview.com/) API.

![go-gerrit - Go(lang) client/library for Gerrit Code Review](./img/logo.png "go-gerrit - Go(lang) client/library for Gerrit Code Review")

## Features

* [Authentication](https://godoc.org/github.com/andygrunwald/go-gerrit#AuthenticationService) (HTTP Basic, HTTP Digest, HTTP Cookie)
* Every API Endpoint like Gerrit
	* [/access/](https://godoc.org/github.com/andygrunwald/go-gerrit#AccessService)
	* [/accounts/](https://godoc.org/github.com/andygrunwald/go-gerrit#AccountsService)
	* [/changes/](https://godoc.org/github.com/andygrunwald/go-gerrit#ChangesService)
	* [/config/](https://godoc.org/github.com/andygrunwald/go-gerrit#ConfigService)
	* [/groups/](https://godoc.org/github.com/andygrunwald/go-gerrit#GroupsService)
	* [/plugins/](https://godoc.org/github.com/andygrunwald/go-gerrit#PluginsService)
	* [/projects/](https://godoc.org/github.com/andygrunwald/go-gerrit#ProjectsService)
* Supports optional plugin APIs such as
	* events-log - [About](https://gerrit.googlesource.com/plugins/events-log/+/master/src/main/resources/Documentation/about.md), [REST API](https://gerrit.googlesource.com/plugins/events-log/+/master/src/main/resources/Documentation/rest-api-events.md)


## Installation

go-gerrit requires Go version 1.8 or greater.

It is go gettable ...

```sh
$ go get github.com/andygrunwald/go-gerrit
```

... (optional) to run checks and tests:

**Tests Only**

```sh
$ cd $GOPATH/src/github.com/andygrunwald/go-gerrit
$ go test -v
```

**Checks, Tests, Linters, etc**

```sh
$ cd $GOPATH/src/github.com/andygrunwald/go-gerrit
$ make
```

## API / Usage

Please have a look at the [GoDoc documentation](https://godoc.org/github.com/andygrunwald/go-gerrit) for a detailed API description.

The [Gerrit Code Review - REST API](https://gerrit-review.googlesource.com/Documentation/rest-api.html) was the base document.

### Authentication

Gerrit support multiple ways for [authentication](https://gerrit-review.googlesource.com/Documentation/rest-api.html#authentication).

#### HTTP Basic

Some Gerrit instances (like [TYPO3](https://review.typo3.org/)) has [auth.gitBasicAuth](https://gerrit-review.googlesource.com/Documentation/config-gerrit.html#auth.gitBasicAuth) activated.
With this you can authenticate with HTTP Basic like this:

```go
instance := "https://review.typo3.org/"
client, _ := gerrit.NewClient(instance, nil)
client.Authentication.SetBasicAuth("andy.grunwald", "my secrect password")

self, _, _ := client.Accounts.GetAccount("self")

fmt.Printf("Username: %s", self.Name)

// Username: Andy Grunwald
```

If you get an `401 Unauthorized`, check your Account Settings and have a look at the `HTTP Password` configuration.

#### HTTP Digest

Some Gerrit instances (like [Wikimedia](https://gerrit.wikimedia.org/)) has [Digest access authentication](https://en.wikipedia.org/wiki/Digest_access_authentication) activated.

```go
instance := "https://gerrit.wikimedia.org/r/"
client, _ := gerrit.NewClient(instance, nil)
client.Authentication.SetDigestAuth("andy.grunwald", "my secrect http password")

self, resp, err := client.Accounts.GetAccount("self")

fmt.Printf("Username: %s", self.Name)

// Username: Andy Grunwald
```

If digest auth is not supported by the choosen Gerrit instance, an error like `WWW-Authenticate header type is not Digest` is thrown.

If you get an `401 Unauthorized`, check your Account Settings and have a look at the `HTTP Password` configuration.

#### HTTP Cookie

Some Gerrit instances hosted like the one hosted googlesource.com (e.g. [Go(lang)](https://go-review.googlesource.com/), [Android](https://android-review.googlesource.com/) or [Gerrit](https://gerrit-review.googlesource.com/)) support HTTP Cookie authentication.

You need the cookie name and the cookie value.
You can get them by click on "Settings > HTTP Password > Obtain Password" in your Gerrit instance.

There you can receive your values.
The cookie name will be (mostly) `o` (if hosted on googlesource.com).
Your cookie secret will be something like `git-your@email.com=SomeHash...`.

```go
instance := "https://gerrit-review.googlesource.com/"
client, _ := gerrit.NewClient(instance, nil)
client.Authentication.SetCookieAuth("o", "my-cookie-secret")

self, _, _ := client.Accounts.GetAccount("self")

fmt.Printf("Username: %s", self.Name)

// Username: Andy G.
```

### More more more

In the examples chapter below you will find a few more examples.
If you miss one or got a question how to do something please [open a new issue](https://github.com/andygrunwald/go-gerrit/issues/new) with your question.
We will be happy to answer them.

## Examples

Further a few examples how the API can be used.
A few more examples are available in the [GoDoc examples section](https://godoc.org/github.com/andygrunwald/go-gerrit#pkg-examples).

### Get version of Gerrit instance

Receive the version of the [Gerrit instance used by the Gerrit team](https://gerrit-review.googlesource.com/) for development:

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
)

func main() {
	instance := "https://gerrit-review.googlesource.com/"
	client, err := gerrit.NewClient(instance, nil)
	if err != nil {
		panic(err)
	}

	v, _, err := client.Config.GetVersion()

	fmt.Printf("Version: %s", v)

	// Version: 2.12.2-2512-g0b1bccd
}
```

### Get all public projects

List all projects from [Chromium](https://chromium-review.googlesource.com/):

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
)

func main() {
	instance := "https://chromium-review.googlesource.com/"
	client, err := gerrit.NewClient(instance, nil)
	if err != nil {
		panic(err)
	}

	opt := &gerrit.ProjectOptions{
		Description: true,
	}
	projects, _, err := client.Projects.ListProjects(opt)
	for name, p := range *projects {
		fmt.Printf("%s - State: %s\n", name, p.State)
	}

	// chromiumos/platform/depthcharge - State: ACTIVE
	// external/github.com/maruel/subcommands - State: ACTIVE
	// external/junit - State: ACTIVE
	// ...
}
```

### Query changes

Get some changes of the [kernel/common project](https://android-review.googlesource.com/#/q/project:kernel/common) from the [Android](http://source.android.com/) [Gerrit Review System](https://android-review.googlesource.com/).

```go
package main

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
)

func main() {
	instance := "https://android-review.googlesource.com/"
	client, err := gerrit.NewClient(instance, nil)
	if err != nil {
		panic(err)
	}

	opt := &gerrit.QueryChangeOptions{}
	opt.Query = []string{"project:kernel/common"}
	opt.AdditionalFields = []string{"LABELS"}
	changes, _, err := client.Changes.QueryChanges(opt)

	for _, change := range *changes {
		fmt.Printf("Project: %s -> %s -> %s%d\n", change.Project, change.Subject, instance, change.Number)
	}

	// Project: kernel/common -> android: binder: Fix BR_ERROR usage and change LSM denials to use it. -> https://android-review.googlesource.com/150839
	// Project: kernel/common -> android: binder: fix duplicate error return. -> https://android-review.googlesource.com/155031
	// Project: kernel/common -> dm-verity: Add modes and emit uevent on corrupted blocks -> https://android-review.googlesource.com/169572
	// ...
}
```

## FAQ

### How is the source code organized?

The source code organisation was inspired by [go-github by Google](https://github.com/google/go-github).

Every REST API Endpoint (e.g. [/access/](https://gerrit-review.googlesource.com/Documentation/rest-api-access.html) or [/changes/](https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html)) is coupled in a service (e.g. [AccessService in access.go](./access.go) or [ChangesService in changes.go](./changes.go)).
Every service is part of [gerrit.Client](./gerrit.go) as a member variable.

gerrit.Client can provide basic helper functions to avoid unnecessary code duplications such as building a new request, parse responses and so on.

Based on this structure implementing a new API functionality is straight forwarded. Here is an example of *ChangeService.DeleteTopic* / [DELETE /changes/{change-id}/topic](https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#delete-topic):

```go
func (s *ChangesService) DeleteTopic(changeID string) (*Response, error) {
	u := fmt.Sprintf("changes/%s/topic", changeID)
	return s.client.DeleteRequest(u, nil)
}
```

### What about the version compatibility with Gerrit?

The library was implemented based on the REST API of Gerrit version 2.11.3-1230-gb8336f1 and tested against this version.

This library might be working with older versions as well.
If you notice an incompatibility [open a new issue](https://github.com/andygrunwald/go-gerrit/issues/new) or try to fix it.
We welcome contribution!


### What about adding code to support the REST API of an optional plugin?

It will depend on the plugin, you are welcome to [open a new issue](https://github.com/andygrunwald/go-gerrit/issues/new) first to propose the idea if you wish.
As an example the addition of support for events-log plugin was supported because the plugin itself is fairly
popular and the structures that the REST API uses could also be used by `gerrit stream-events`.


## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
