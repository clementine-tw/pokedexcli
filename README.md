# pokedexcli

Play the pokemon game in **command-line**!

This is a project to practice _REPL_ and _HTTP_ communication in Go.

The information of pokemons in game is get from [_pokeapi_](https://pokeapi.co/).

## Requirement

- Go 1.24.5

## Build

```bash
go build
```

## Play

```bash
./pokedexcli
```

### Playing Commands

- map: show next 20 locations
- mapb: show previous 20 locations
- explore <location_name>: show pokemons in the location
- catch <pokemon_name>: catch the pokemon
- inspect <pokemon_name>: show information of the caught pokemon
- pokedex: show caught pokemons
- help: show help information
- exit: quit game
