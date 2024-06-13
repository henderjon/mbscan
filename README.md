```
NAME
  mbscan - compare byte and rune counts from stdin

SYNOPSIS
  $ mbscan
  $ mbscan [-h|help]

DESCRIPTION
  mbscan scans stdin and counts the number of bytes and runes to detect
  multi-byte characters. Be default, it will output the difference of
  bytes vs runes.

EXAMPLES
  $ mbscan -v < file.txt
  $ cat file.txt | mbscan -v

OPTIONS
  -path string
    	an opaque string that is used to identify the source of the input stream
  -s	silent; no output only exit codes
  -v	verbose; shows multi-byte characters found in the input stream and the byte and rune counts

```
