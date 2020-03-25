# Go DynDNS
[![Build](https://api.travis-ci.com/henryleduc/dyndns.svg?branch=master)](https://travis-ci.com/github/henryleduc/dyndns)
[![Go Report Card](https://goreportcard.com/badge/github.com/henryleduc/dyndns)](https://goreportcard.com/report/github.com/henryleduc/dyndns)
[![Documentation](https://godoc.org/github.com/henryleduc/dyndns?status.svg)](http://godoc.org/github.com/henryleduc/dyndns)
[![Coverage Status](https://coveralls.io/repos/github/henryleduc/dyndns/badge.svg?branch=master)](https://coveralls.io/github/henryleduc/dyndns?branch=master)
[![license](https://img.shields.io/github/license/henryleduc/dyndns.svg?maxAge=2592000)](https://github.com/henryleduc/dyndns/LICENSE)

Cloudflare DynamicDNS Tool for using Cloudflare DynamicDNS as a DDNS provider.

<img alt="Cloudflare Logo" src="https://www.cloudflare.com/img/logo-cloudflare-dark.svg" width="250">

If you are hosting a Home Server and have a Dynamic IP this is an essential library to use.

You will need a free account with Cloudflare.
Then you can set this either in a cron job, or have it running in a Docker container.

DynDNS will then update your specified A record(s) to the given public IP address that is given by your ISP.

This is all done through the Cloudflare APIs so no need to worry about the stability of 3rd Part Services.

# Getting Started
To install this project use:
`go get github.com/henryleduc/dyndns/cmd/ddns-updater`

You can also use this project as a DDNS client inside your project or application.

Just simply import:
`import "github.com/henryleduc/dyndns/cloudflare"`

# Contributing
Feel free to contribute to this project, I plan on adding other services in the future such as DuckDNS and Google Domains.
I'm also open to functionality requests.
