# ansiview: view old ANSI art in your terminal

This is a very basic cp437 -> utf8 codec driver through what amounts to UNIX
`cat` otherwise. To use it:

Pass it the ANSI over standard input or pass it a filename.

Check out some art packs at [16colo.rs](https://16colo.rs), unzip, and try this script
or some similar variation on it in your UTF-8-compatible terminal:

```bash
for i in *.ANS; do clear; ansiview $i; echo $i; read; done
```

Then you can walk through the whole directory; hit enter to cycle to the next file.

## Pretty

![](spawn.png)

## Requirements

UTF-8 compatible, minimum 16 color terminal. Your terminal should be at least
`80x25` to see most artwork, however you may benefit from much taller heights
and sometimes width; `130x75` is great.

## Installation

```bash
$ go get github.com/erikh/ansiview
```

## Author

Erik Hollensbe <erik+git@hollensbe.org>, but most of the work was really done
by the person who did the CP437 codec in golang's standard library. Props to
them!

## License

MIT (c) 2020 Erik Hollensbe
