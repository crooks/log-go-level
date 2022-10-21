# log-go-level
Convert loglevel strings to associated intergers.

## Overview
This package does nothing more than converting common loglevel strings into integers.  This enables loglevels to be specified in config files as strings such as "debug", "info", "warn" and converted in the background to the integers expected by Go's standard log package and the [Masterminds] (https://github.com/Masterminds/log-go) package.