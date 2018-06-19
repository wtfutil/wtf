Go Trello API
================

[![Trello Logo](https://raw.githubusercontent.com/adlio/trello/master/trello-logo.png)](https://www.trello.com)

[![GoDoc](https://godoc.org/github.com/adlio/trello?status.svg)](http://godoc.org/github.com/adlio/trello)
[![Build Status](https://travis-ci.org/adlio/trello.svg)](https://travis-ci.org/adlio/trello)
[![Coverage Status](https://coveralls.io/repos/github/adlio/trello/badge.svg?branch=master)](https://coveralls.io/github/adlio/trello?branch=master)

A #golang package to access the [Trello API](https://developers.trello.com/v1.0/reference). Nearly 100% of the
read-only surface area of the API is covered, as is creation and modification of Cards.
Low-level infrastructure for features to modify Lists and Boards are easy to add... just not
done yet.

Pull requests are welcome for missing features.

My current development focus is documentation, especially enhancing this README.md with more
example use cases.

## Installation

The Go Trello API has been Tested compatible with Go 1.7 on up. Its only dependency is
the `github.com/pkg/errors` package. It otherwise relies only on the Go standard library.

```
go get github.com/adlio/trello
```

## Basic Usage

All interaction starts with a `trello.Client`. Create one with your appKey and token:

```Go
client := trello.NewClient(appKey, token)
```

All API requests accept a trello.Arguments object. This object is a simple
`map[string]string`, converted to query string arguments in the API call.
Trello has sane defaults on API calls. We have a `trello.Defaults()` utility function
which can be used when you desire the default Trello arguments. Internally,
`trello.Defaults()` is an empty map, which translates to an empty query string.

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}
```

## Client Longevity

When getting Lists from Boards or Cards from Lists, the original `trello.Client` pointer
is carried along in returned child objects. It's important to realize that this enables
you to make additional API calls via functions invoked on the objects you receive:

```Go
client := trello.NewClient(appKey, token)
board, err := client.GetBoard("ID", trello.Defaults())

// GetLists makes an API call to /boards/:id/lists using credentials from `client`
lists, err := board.GetLists(trello.Defaults())

for _, list := range lists {
  // GetCards makes an API call to /lists/:id/cards using credentials from `client`
  cards, err := list.GetCards(trello.Defaults())
}
```

## Get Trello Boards for a User

Boards can be retrieved directly by their ID (see example above), or by asking
for all boards for a member:

```Go
member, err := client.GetMember("usernameOrId", trello.Defaults())
if err != nil {
  // Handle error
}

boards, err := member.GetBoards(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Lists on a Board

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}

lists, err := board.GetLists(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Cards on a Board

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}

cards, err := board.GetCards(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Trello Cards on a List

```Go
list, err := client.GetList("lIsTID", trello.Defaults())
if err != nil {
  // Handle error
}

cards, err := list.GetCards(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Creating a Card

The API provides several mechanisms for creating new cards.

### Creating Cards from Scratch on the Client

This approach requires the most input data on the card:

```Go
card := trello.Card{
  Name: "Card Name",
  Desc: "Card description",
  Pos: 12345.678,
  IDList: "iDOfaLiSt",
  IDLabels: []string{"labelID1", "labelID2"},
}
err := client.CreateCard(card, trello.Defaults())
```


### Creating Cards On a List

```Go
list, err := client.GetList("lIsTID", trello.Defaults())
list.AddCard(trello.Card{ Name: "Card Name", Description: "Card description" }, trello.Defaults())
```

### Creating a Card by Copying Another Card

```Go
err := card.CopyToList("listIdNUmber", trello.Defaults())
```

## Get Actions on a Board

```Go
board, err := client.GetBoard("bOaRdID", trello.Defaults())
if err != nil {
  // Handle error
}

actions, err := board.GetActions(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Get Actions on a Card

```Go
card, err := client.GetCard("cArDID", trello.Defaults())
if err != nil {
  // Handle error
}

actions, err := card.GetActions(trello.Defaults())
if err != nil {
  // Handle error
}
```

## Rearrange Cards Within a List

```Go
err := card.MoveToTopOfList()
err = card.MoveToBottomOfList()
err = card.SetPos(12345.6789)
```


## Moving a Card to Another List

```Go
err := card.MoveToList("listIdNUmber", trello.Defaults())
```


## Card Ancestry

Trello provides ancestry tracking when cards are created as copies of other cards. This package
provides functions for working with this data:

```Go

// ancestors will hold a slice of *trello.Cards, with the first
// being the card's parent, and the last being parent's parent's parent...
ancestors, err := card.GetAncestorCards(trello.Defaults())

// GetOriginatingCard() is an alias for the last element in the slice
// of ancestor cards.
ultimateParentCard, err := card.GetOriginatingCard(trello.Defaults())

```

## Debug Logging

If you'd like to see all API calls logged, you can attach a `.Logger` (implementing `Debugf(string, ...interface{})`)
to your client. The interface for the logger mimics logrus. Example usage:

```Go
logger := logrus.New()
logger.SetLevel(logrus.DebugLevel)
client := trello.NewClient(appKey, token)
client.Logger = logger
```
