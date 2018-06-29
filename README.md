# MlDict
Program for search in dictionary database.

## Owerview

### Base execution

```
Search program for dictionary DB.

Usage:
  dict [command]

Available Commands:
  help        Help about any command
  list        List of dictionaries
  search      Search dictionaries for a word

Flags:
      --config string   config file (default is $HOME/.mldict.yaml)
      --dbfile string   database file (default "/home/martin/.dictionary.db")
  -h, --help            help for dict

Use "dict [command] --help" for more information about a command.
```

### Command list

```
List of dictionaries

Usage:
  dict list [flags]

Aliases:
  list, dicts

Flags:
  -h, --help   help for list

Global Flags:
      --config string   config file (default is $HOME/.mldict.yaml)
      --dbfile string   database file (default "/home/martin/.dictionary.db")
```

### Command search

```
Search dictionary database for a word. For searching use fulltext index in sqlite3 (FT4). 
Implicit is search for "term*", it is for key starting with term. You can owerride this by '-e' option and search for exact term.

Usage:
  dict search [flags]

Aliases:
  search, find, s

Flags:
  -d, --dict string   Select one dictionary
  -e, --exact         Search for exact term
  -h, --help          help for search

Global Flags:
      --config string   config file (default is $HOME/.mldict.yaml)
      --dbfile string   database file (default "/home/martin/.dictionary.db")
```

### Example usage

```
$ dist search -e hello
 * en-cz * 
hello -  ahoj; haló!; nazdar!
hello, everybody. -  dobrý den vespolek.

 * jargon * 
hello sailor! interj. -  Occasional West Coast equivalent of hello world; seems to have originated at SAIL, later associated with the game Zork (which also included "hello, aviator" and "hello, implementor"). Originally from the traditional hooker's greeting to a swabbie fresh off the boat, of course. The standard response is "Nothing happens here."; of all the Zork/Dungeon games, only in Infocom's Zork 3 is "Hello, Sailor" actually useful (excluding the unique situation where _knowing_ this fact is important in Dungeon...). 
hello world interj. -  1. The canonical minimal test message in the C/Unix universe. 2. Any of the minimal programs that emit this message. Traditionally, the first program a C coder is supposed to write in a new environment is one that just prints "hello, world" to standard output (and indeed it is the first example program in K&R). Environments that generate an unreasonably large executable for this trivial test or which require a hairy compiler-linker invocation to generate it are considered to lose (see X). 3. Greeting uttered by a hacker making an entrance or requesting information from anyone present. "Hello, world! Is the LAN back up yet?" 
hello, wall! excl. -  See wall. 

 * webster * 
Hello -  See Halloo.	interj. & n.

```

## BUGS

This is early development version and may be with many bugs.

## TODO

- GUI (via web on localhost)
- documentation for creating DB
- initialize DB from program
- export/import 

## Install

Compile from sources with [go](https://golang.org/) and [just](https://github.com/casey/just)
or copy executable binary (static linked) from [releases](https://github.com/martinlebeda/mldict/releases) where you want and run.