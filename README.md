# dyndns 
[![Go Report Card](https://goreportcard.com/badge/github.com/henryleduc/dyndns)](https://goreportcard.com/report/github.com/henryleduc/dyndns)
[![Documentation](https://godoc.org/github.com/henryleduc/dyndns?status.svg)](http://godoc.org/github.com/henryleduc/dyndns)
[![Coverage Status](https://coveralls.io/repos/github/henryleduc/dyndns/badge.svg?branch=master)](https://coveralls.io/github/henryleduc/dyndns?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/henryleduc/dyndns.svg)](https://github.com/henryleduc/dyndns/issues)
[![license](https://img.shields.io/github/license/henryleduc/dyndns.svg?maxAge=2592000)](https://github.com/henryleduc/dyndns/LICENSE)

Cloudflare DynamicDNS updater for A records.

<img alt="Cloudflare Logo" src="https://www.cloudflare.com/img/logo-cloudflare-dark.svg" width="250">

If you are hosting a Home Server and have a Dnymaic IP this is an essential library to use.

You will need a free account with Cloudflare.
Then you can set this this either in a cron job, or have it running in a docker container.

DynDNS will then update your specified A record(s) to the given public IP address that is given by your ISP.

This is all done through the Cloudflare APIs so no need to worry about the stability of 3rd Part Services.

This is a similar self hosted alternative service to something like DuckDNS.org
