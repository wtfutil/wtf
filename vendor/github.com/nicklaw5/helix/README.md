# helix

A Twitch Helix API client written in Go (Golang).

[![Build Status](https://travis-ci.org/nicklaw5/helix.svg?branch=master)](https://travis-ci.org/nicklaw5/helix)
[![Coverage Status](https://coveralls.io/repos/github/nicklaw5/helix/badge.svg)](https://coveralls.io/github/nicklaw5/helix)

## Package Status

This project is a work in progress. Twitch has not finished all available endpoints/features for the Helix API, but as these get released they are likely to be implemented in this package.

## Documentation & Examples

All documentation and usage examples for this package can be found in the [docs directory](docs). If you are looking for the Twitch API docs, see the [Twitch Developer website](https://dev.twitch.tv/docs/api).

## Supported Endpoints & Features

**Authentication:**

- [x] Generate Authorization URL
- [x] Get App Access Tokens (OAuth Client Credentials Flow)
- [x] Get User Access Tokens (OAuth Authorization Code Flow)
- [x] Refresh User Access Tokens
- [x] Revoke User Access Tokens

**API Endpoint:**

- [x] Get Bits Leaderboard
- [x] Get Clip
- [x] Create Clip
- [x] Create Entitlement Grants Upload URL
- [x] Get Games
- [x] Get Top Games
- [x] Get Extension Analytics
- [x] Get Game Analytics
- [x] Get Streams
- [x] Get Streams Metadata
- [x] Create Stream Marker
- [x] Get Stream Markers
- [x] Get Users
- [x] Get Users Follows
- [x] Update User
- [x] Get Videos
- [x] Get Webhook Subscriptions

## Quick Usage Example

This is a quick example of how to get users. Note that you don't need to provide both a list of ids and logins, one or the other will suffice.

```go
client, err := helix.NewClient(&helix.Options{
    ClientID: "your-client-id",
})
if err != nil {
    // handle error
}

resp, err := client.GetUsers(&helix.UsersParams{
    IDs:    []string{"26301881", "18074328"},
    Logins: []string{"summit1g", "lirik"},
})
if err != nil {
    // handle error
}

fmt.Printf("Status code: %d\n", resp.StatusCode)
fmt.Printf("Rate limit: %d\n", resp.GetRateLimit())
fmt.Printf("Rate limit remaining: %d\n", resp.GetRateLimitRemaining())
fmt.Printf("Rate limit reset: %d\n\n", resp.GetRateLimitReset())

for _, user := range resp.Data.Users {
    fmt.Printf("ID: %s Name: %s\n", user.ID, user.DisplayName)
}
```

Output:

```txt
Status code: 200
Rate limit: 30
Rate limit remaining: 29
Rate limit reset: 1517695315

ID: 26301881 Name: sodapoppin
ID: 18074328 Name: destiny
ID: 26490481 Name: summit1g
ID: 23161357 Name: lirik
```

## Contributions

PRs are very much welcome. Where possible, please write tests for any code that is introduced by your PRs.

## Contributors

Thanks to all of the following people (ordered alphabetically) for contributing their time for improving this library:

- Y.Horie ([@u5surf](https://github.com/u5surf))
- Philipp Führer ([@flipkick](https://github.com/flipkick))

## License

This package is distributed under the terms of the [MIT](License) License.
