# addcol-go
Command line tool to add columns to an existing CSV file.

usage: addcol-go.exe <column> <value> <infile> [<outfile>]


Flags:
  --[no-]help  Show context-sensitive help (also try --help-long and
               --help-man).

Args:
  <column>     Column number where to insert Value (9999 to add column at the
               end)
  <value>      Value to insert
  <infile>     Input CSV file name
  [<outfile>]  Output CSV file name (defaults to stdout)