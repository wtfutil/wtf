@echo off
REM - The tests that actually hit the Spotify Web API don't run by default.
REM - Use this script to run them in addition to the standard unit tests.

cmd /C "set FULLTEST=y && go test %*"
