# powersurge

A simple tool that monitors a web host and power cycles an energenie power strip if the check fails.

## Usage

```
Usage of ./powersurge:
  -host="http://www.google.com": Address of site to use to test connectivity
  -interval="5m": Interval in which to run the test. Example: 10s, 5m, 1h
  -socket=1: Number of the socket the router is connected to
  -strip="10.1.1.213": IP address of the power strip
```

## Requirements

A power strip with LAN connection made by [EnerGenie](http://www.energenie.com/products.aspx?sg=210).

## How to build

1. Set a proper `GOPATH`
2. Run `go build`
3. There is no step 3.
