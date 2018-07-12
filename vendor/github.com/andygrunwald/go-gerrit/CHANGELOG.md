# Changelog

This is a high level log of changes, bugfixes, enhancements, etc
that have taken place between releases. Later versions are shown
first. For more complete details see
[the releases on GitHub.](https://github.com/andygrunwald/go-gerrit/releases)

## Versions

### Latest

### 0.5.2

* Fix panic in checkAuth() if Gerrit is down #42
* Implement ListVotes(), DeleteVotes() and add missing tests

### 0.5.1

* Added the `AbandonChange`, `RebaseChange`, `RestoreChange` and 
  `RevertChange` functions.

### 0.5.0

**WARNING**: This release includes breaking changes.

* [BREAKING CHANGE] The SetReview function was returning the wrong
  entity type. (#40)

### 0.4.0

**WARNING**: This release includes breaking changes.

* [BREAKING CHANGE] - Added gometalinter to the build and fixed problems 
  discovered by the linters.
    * Comment and error string fixes.
    * Numerous lint and styling fixes.
    * Ensured error values are being properly checked where appropriate.
    * Addition of missing documentation
    * Removed filePath parameter from DeleteChangeEdit which was unused and 
      unnecessary for the request.
    * Fixed CherryPickRevision and IncludeGroups functions which didn't pass
      along the provided input structs into the request.
* Go 1.5 has been removed from testing on Travis. The linters introduced in 
  0.4.0 do not support this version, Go 1.5 is lacking security updates and
  most Linux distros have moved beyond Go 1.5 now.
* Add Go 1.9 to the Travis matrix.
* Fixed an issue where urls containing certain characters in the credentials
  could cause NewClient() to use an invalid url. Something like `/`, which
  Gerrit could use for generated passwords, for example would break url.Parse's
  expectations.

### 0.3.0

**WARNING**: This release includes breaking changes.

* [BREAKING CHANGE] Fix Changes.PublishDraftChange to accept a notify parameter.
* [BREAKING CHANGE] Fix PublishChangeEdit to accept a notify parameter.
* [BREAKING CHANGE] Fix ChangeFileContentInChangeEdit to allow the file content
  to be included in the request.
* Fix the url being used by CreateChange
* Fix type serialization of EventInfo.PatchSet.Number so it's consistent.
* Fix Changes.AddReviewer so it passes along the reviewer to the request.
* Simplify and optimize RemoveMagicPrefixLine

### 0.2.0

**WARNING**: This release includes breaking changes.

* [BREAKING CHANGE] Several bugfixes to GetEvents:
  * Update EventInfo to handle the changeKey field and apply
    the proper type for the Project field
  * Provide a means to ignore marshaling errors
  * Update GetEvents() to return the failed lines and remove
    the pointer to the return value because it's unnecessary.
* [BREAKING CHANGE] In ec28f77 `ChangeInfo.Labels` has been changed to map
  to fix #21.


### 0.1.1

* Minor fix to SubmitChange to use the `http.StatusConflict` constant
  instead of a hard coded value when comparing response codes.
* Updated AccountInfo.AccountID to be omitted of empty (such as when 
  used in ApprovalInfo).
* + and : in url parameters for queries are no longer escaped. This was
  causing `400 Bad Request` to be returned when the + symbol was
  included as part of the query. To match behavior with Gerrit's search
  handling, the : symbol was also excluded.
* Fixed documentation for NewClient and moved fmt.Errorf call from
  inside the function to a `ErrNoInstanceGiven` variable so it's
  easier to compare against.
* Updated internal function digestAuthHeader to return exported errors
  (ErrWWWAuthenticateHeader*) rather than calling fmt.Errorf. This makes
  it easier to test against externally and also fixes a lint issue too.
* Updated NewClient function to handle credentials in the url.
* Added the missing `Submitted` field to `ChangeInfo`.
* Added the missing `URL` field to `ChangeInfo` which is usually included
  as part of an event from the events-log plugin.

### 0.1.0

* The first official release
* Implemented digest auth and several fixes for it.
* Ensured Content-Type is included in all requests
* Fixed several internal bugs as well as a few documentation issues
