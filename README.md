# grep-using-go
A simplified version of `grep` implemented in Go, broken out into development stages

## Objective & Approach
The end goal is to build command line application that supports a subset of functionality found in the standard unix utility `grep`.

The project starts with a base implementation and gradually adds features through several stages.  In this way a standard software development methodology is followed: 

  1. Start with a simple implementation that provides base functionality
  2. Iterate, gradually improving:
     - Functionality
     - Design & organization
     - Test cases and coverage
  3. Repeat step _#2_ until project is complete


## Specification

The resulting binary executable is to be named `gogrep`.

`gogrep` should mirror the behavior of regular `grep` to the extent which features are copied. For example, `gogrep` should support matching content from files specified as command line arguments or from STDIN if no file arguments are supplied.

    gogrep PATTERN [[file1] [file2] [file3] ...]

`gogrep` should exit with `0` status if at least one line match is made, otherwise exit with `1`.



**The following flags are to be supported:**

    -h    Show help/usage

    -H    Prints the file name followed by a ':' and then the matching content.  This is default
          behavior is matches are performed against more than one file.

    -n    Prints the line number of the match followed by a ':' and then the matching content
          If the file name is also printed, the line number information is printed after
          the formatted file name.

    -c    Prints the count of matching lines, but not the matching content lines.

    -l    Prints only the file name if the file contains a match. Does not print matching lines.

    -v    Inverts the match selection.  Content lines that do not match the PATTERN are selected.

    -q    Quiet mode.  Does not print any output.  Only the exit code is set.


**Optional flags**

    -e    Interpret the PATTERN as a regular expression.

    -r    Recursively search subdirectories listed
   