# Pokedex CLI

A small Go CLI that talks to the PokeAPI and lets you browse locations, explore areas, catch Pokémon, and view your collection. Caching is built in to keep repeated requests fast.

## Commands

- `help` — show available commands.
- `exit` — quit the CLI.
- `mapf` — show the next page of 20 location areas.
- `mapb` — show the previous page of 20 location areas.
- `explore <location-area>` — list Pokémon encounters for the given area.
- `catch <pokemon>` — attempt to catch a Pokémon by name; uses base experience to determine difficulty.
- `inspect <pokemon>` — show details for a caught Pokémon (name, height, weight, stats, types).
- `pokedex` — list all caught Pokémon.

## How It Works

- **PokeAPI client** (`internal/pokeapi`): wraps HTTP calls for location lists, location areas, and Pokémon details.
- **Caching** (`internal/pokecache`): in-memory cache with timed reaping; reused across API calls to speed navigation and re-queries.
- **State**: the REPL keeps pagination URLs and a `pokedex` map of caught Pokémon for offline inspect.

## Development

- Format: `gofmt -w .`
- Tests: `go test ./...`
- Run: `go run .`
