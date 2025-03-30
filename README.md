# 🧢 Pokédex REPL

Welcome to the **Pokédex REPL**, a command-line interface for exploring, catching, and managing Pokémon! This tool mimics the classic Pokémon adventure — roam around, discover wild Pokémon, catch them if you can, and build your Pokédex.

---

## 🚀 Getting Started

Once you run the REPL, you'll be able to input commands to interact with the world of Pokémon.

To see the full list of available commands at any time, just type:


---

## 🕹️ Available Commands

| Command    | Description                                                                 |
|------------|-----------------------------------------------------------------------------|
| `exit`     | Exit the Pokédex application.                                               |
| `help`     | Display a list of available commands and their descriptions.                |
| `map`      | Retrieve the next batch of 20 Pokémon locations.                            |
| `mapd`     | Go back and retrieve the previous batch of 20 Pokémon locations.            |
| `explore`  | Explore a specific location to see what Pokémon are currently there.        |
| `catch`    | Attempt to catch a Pokémon. Success depends on its base experience.         |
| `inspect`  | View detailed stats and information of a specific Pokémon you’ve caught.    |
| `pokedex`  | View your personal Pokédex with all the Pokémon you’ve successfully caught. |

---

## 🎮 Example Usage

```shell
> map
# Shows 20 nearby locations with Pokémon

> explore viridian-forest
# Lists Pokémon found in Viridian Forest

> catch pikachu
# Tries to catch Pikachu — success depends on its difficulty

> inspect pikachu
# View Pikachu's stats and details

> pokedex
# List all your caught Pokémon

> exit
# Quits the game