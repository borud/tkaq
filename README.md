# tkaq

This is a sample program showing how you can extract data from the
Horde server for NB-IoT and LTE-M based devices.  Since Horde doesn't
care much about the payloads it processes the data from the device are
in an application specific format, so for this tool we included some
code for decoding the data from the sensor so you can easily print it
out in any format you like.  Here we just produce a simple CSV file.

Feel free to clone the program and adapt it for your purposes.  If you
feel like doing something exciting with this program (like add output
formats, feel free to send me a pull request).

# Paging

Note that the data is sorted in reverse chronological order from the
server, meaning that the newest dates are first.  This means when
paging through the dataset N elements at a time, we have to go
backwards in time.

# Building

## Prerequisites

You need to have Go version 1.12.6 or newer since this project uses
modules.

## Building for different operating systems

Building `tkaq` is quite straight forward.  Below are instructions for
building it for various platforms.  The binary will turn up in the
`bin` directory of the project.

### For OSX

    make
	
### For Linux

    GOARCH=amd64 GOOS=linux make	

### For Windows

	GOARCH=amd64 GOOS=windows make
	
# Running

In order to run this program you will need an API token from
`https://nbiot.engineering/api-tokens-overview` in order to be allowed to access the air quality
collection.  Once you have created this token you can either put this
into a configuration file named `.telenor-nbiot` which looks like
this:

    address=https://api.nbiot.telenor.io
    token=<your API token>

or you can make the API address and token available via environment
variables:

    export TELENOR_NBIOT_ADDRESS="https://api.nbiot.telenor.io/"
	export TELENOR_NBIOT_TOKEN="<your API token>"

The build process produces one binary `bin/tkaq`.  If you invoke it
with the `-h` flag it will display the command line options.

    $ bin/tkaq -h
    Usage:
      tkaq [OPTIONS]

    Application Options:
      -c, --collection-id= Collection ID (default: 17dh0cf43jfi2f)
      -p, --pagesize=      Number of datapoints to return per page (default: 500)
      -s, --start-time=    Start date and time in RFC3339 format

    Help Options:
      -h, --help           Show this help message
	  
So in order to list every datapoint since midnight 2019-10-01 UTC you
issue the command:

    bin/tkaq -s 2019-10-01T00:00:00Z
	  




