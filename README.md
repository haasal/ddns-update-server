# ddns-update-server

This repo is a web server for updating a dyndns service as described in [this](https://dev.to/haasal/set-up-your-own-ddns-server-with-bind9-and-go-193o) blogpost.

## Security

If you spot any vulnerabilities please start an issue.

## How to run

Follow the tutorial mentioned above before doing the follwoing steps.

The service doesn't have the feature to set a password yet.
So create the file `/secrets/ddns-web-passwd` on your system and paste in a **hashed** and secure PSK.
The PSK authorizes a client to make changes to the DNS table.

Copy an SSL certificate to the location specified at the end of `main.go`. Make sure your user has sufficient **read** permissions for these files.
