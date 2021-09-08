# About

This is a simple utility to run a tftp server in a chrooted directory. This is useful for interacting with things like managed network switches where you don't want to set up a tftp server like tftpd-hpa.

# Installing

    $ cd <build directory>
    $ go mod init build
    $ go get -d -v github.com/korylprince/tftphere@<version>
    $ go build -o tftphere github.com/korylprince/tftphere
    $ ./tftphere -h
    $ cp tftphere <$PATH>

# Usage

    $ tftphere -h
    Usage of tftphere:
      -force
            allow files to be overwritten
      -root string
            root file directory for tftp server (default ".")
