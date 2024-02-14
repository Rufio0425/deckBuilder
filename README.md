# DeckBuilder

DeckBuilder is a simple tool to store your Magic The Gathering card.  
The goal of this project is:
- To allow a virtual experience for MTG enthusiasts to store the physical cards they own
- Check a card's legality status (legal, not legal, banned) for certain formats

## Build Program

To build the program, you need to have Go installed. Once installed, you can use the `go build` command to compile the source code into a binary executable.

```bash
go build
```

## Executing Program

After building the program, run the binary executable using the following command:

```bash
./deckBuilder
```

DeckBuilder runs on port 8080 by default, you'll see the following prompt:
```bash
Server started on port 8080
```

Current HTTP requests available
```bash
GET http://localhost:8080/cards/{cardName}
```