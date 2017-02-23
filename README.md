Xspec Status
============

- Project:  <http://github.com/c4s4/xspec-status>
- Releases: <http://github.com/c4s4/xspec-status/releases>

When you run [Xspec](https://github.com/expath/xspec), it prints a status on
the console and a report in XML files, but the return value doesn't tell if
tests were on error. Thus it is not possible to stop a build with an error
on test failure.

This tool parses the Xspec test report and returns the number of errors as
status code. Thus the call will be on error and will make your build fail
on Xspec test failure.

Installation
------------

Unzip the archive, copy the binary for your platform (in the *bin* directory)
somewhere in your *PATH* and rename it *xspec-status* (or *xspec-status.exe*
on Windows).

Usage
-----

To check Xspec test suite success, you must pass XML Xspec test result files on
command line. Thus, you might run:

    $ xspec-status xspec/*-result.xml
    Parsing report file 'build/xspec/main-result.xml'
    Parsing report file 'build/xspec/meta-result.xml'
    Errors: 0

This will print parsed files, the number of errors and will exit with the
number of errors as return value.

*Enjoy!*
