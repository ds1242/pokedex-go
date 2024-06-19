# Pokedex CLI
This is a CLI program that a uses the [Pokemon API](https://pokeapi.co/). 
A user can use several commands to explore the various regions from the game.  Users can then explore to see the list of Pokemon that are in any of the regions.  Using the catch command throws a Pokeball to see if the user can catch a Pokemon.  If successful they are added the user's Pokedex which can be displayed.  All of the information is stored in a local cache created when starting up the program.

## Available commands:
- map
- mapb
- explore
- inspect
- pokedex
- help
- exit


## Future Enhancement Ideas
- Update the CLI to support the "up" arrow to cycle through previous commands
- Simulate battles between pokemon
- Add more unit tests
- Refactor your code to organize it better and make it more testable
- Keep pokemon in a "party" and allow them to level up
- Allow for pokemon that are caught to evolve after a set amount of time 
- Persist a user's Pokedex to disk so they can save progress between sessions
- Use the PokeAPI to make exploration more interesting. For example, rather than typing the names of areas, maybe you are given choices of areas and just type "left" or "right"
- Random encounters with wild pokemon
- Adding support for different types of balls (Pokeballs, Great Balls, Ultra Balls, etc), which have different chances of catching pokemon