# Matrix-Appservice-Go
This is a Matrix Application Service framework written in Golang.
The Framework is in Alpha nothing does yet work!

## Support Room
Join the discussion at [#matrix-appservice-go:matrix.ffslfl.net](https://matrix.to/#/#matrix-appservice-go:matrix.ffslfl.net)

This can be used to quickly setup performant application services for almost
anything you can think of in a web framework agnostic way.

To create an app service registration file:

// TODO Add Example

You only need to generate a registration once, provided the registration info does not
change. Once you have generated a registration, you can run the app service like so:

// TODO Add Example

TLS Connections
===============
If `MATRIX_AS_TLS_KEY` and `MATRIX_AS_TLS_CERT` environment variables are
defined and point to valid tls key and cert files, the AS will listen using
an HTTPS listener.