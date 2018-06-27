/*
Package gerrit provides a client for using the Gerrit API.

Construct a new Gerrit client, then use the various services on the client to
access different parts of the Gerrit API. For example:

	instance := "https://go-review.googlesource.com/"
	client, _ := gerrit.NewClient(instance, nil)

	// Get all public projects
	projects, _, err := client.Projects.ListProjects(nil)

Set optional parameters for an API method by passing an Options object.

	// Get all projects with descriptions
	opt := &gerrit.ProjectOptions{
		Description: true,
	}
	projects, _, err := client.Projects.ListProjects(opt)

The services of a client divide the API into logical chunks and correspond to
the structure of the Gerrit API documentation at
https://gerrit-review.googlesource.com/Documentation/rest-api.html#_endpoints.

Authentication

The go-gerrit library supports various methods to support the authentication.
This methods are combined in the AuthenticationService that is available at client.Authentication.

One way is an authentication via HTTP cookie.
Some Gerrit instances hosted like the one hosted googlesource.com (e.g. https://go-review.googlesource.com/,
https://android-review.googlesource.com/ or https://gerrit-review.googlesource.com/) support HTTP Cookie authentication.

You need the cookie name and the cookie value.
You can get them by click on "Settings > HTTP Password > Obtain Password" in your Gerrit instance.
There you can receive your values.
The cookie name will be (mostly) "o" (if hosted on googlesource.com).
Your cookie secret will be something like "git-your@email.com=SomeHash...".

	instance := "https://gerrit-review.googlesource.com/"
	client, _ := gerrit.NewClient(instance, nil)
	client.Authentication.SetCookieAuth("o", "my-cookie-secret")

	self, _, _ := client.Accounts.GetAccount("self")

	fmt.Printf("Username: %s", self.Name)

	// Username: Andy G.

Some other Gerrit instances (like https://review.typo3.org/) has auth.gitBasicAuth activated.
With this you can authenticate with HTTP Basic like this:

	instance := "https://review.typo3.org/"
	client, _ := gerrit.NewClient(instance, nil)
	client.Authentication.SetBasicAuth("andy.grunwald", "my secrect password")

	self, _, _ := client.Accounts.GetAccount("self")

	fmt.Printf("Username: %s", self.Name)

	// Username: Andy Grunwald

Additionally when creating a new client, pass an http.Client that supports further actions for you.
For more information regarding authentication have a look at the Gerrit documentation:
https://gerrit-review.googlesource.com/Documentation/rest-api.html#authentication

*/
package gerrit
