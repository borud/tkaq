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
