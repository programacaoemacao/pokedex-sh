# Pokédex SH

[![language](https://img.shields.io/badge/Language-Go-blue)](https://golang.org/)
[![license](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/programacaoemacao/pokedex-sh/blob/main/LICENSE)


A Pokédex made in Golang using the [Charm BubbleTea](https://github.com/charmbracelet/bubbletea) to be used in the terminal, displaying information about Pokémon and their respective ASCII arts.

This repository has some bugs in the Pokémon visualization due to unknown causes.

![Pokédex](./pokedex.gif "Pokédex app showcase")

## Running the Pokédex

You need to have [Golang >= 1.21](https://go.dev/doc/install) installed on your machine.

To run the Pokédex in your terminal, run this command:

> **Important**: It's highly recommended to run this repository using the VSCode intergrated terminal. It's going to show more accurate ASCII arts.


1. Install the dependencies
   ```bash
   make download-depedencies
   ```
1. Run the app
   ```bash
   make run-pokedex
   ```

If you don't have [Make](https://www.gnu.org/software/make/) installed, go to `Makefile` and copy the the commands corresponding to the desired command and execute them in the project root.

## Development flow

```mermaid
---
title: Development flow
---

flowchart TB
   S1(Test the Charm Bubbletea Library fas:fa-flask)
   S2(Test the Image ASCII Generator Library fas:fa-flask)
   S3(Make a Pokédex POC with some Pokémons fas:fa-laptop-code)
   S4(Build a scraper fas:fa-spider)
   S5(Scrape the data fas:fa-spider)
   S6(Convert Pokémon sprites to ASCII fas:fa-paint-brush)
   S7(Put it all together fas:fa-code-branch)
   S8(Refactor fas:fa-code)
   S9(Document the project fas:fa-file-code)

   S1 --> S2
   S2 --> S3
   S3 --> S4
   S4 --> S5
   S5 --> S6
   S6 --> S7
   S7 --> S8
   S8 --> S9
```

## Attention points

**You don't need to run the scraper and collect any data. The repository already contains the necessary data to run the Pokédex!**

If you would like to run the scraper, check this [page](https://github.com/programacaoemacao/pokedex-sh/tree/main/cmd/data_scraper).

If you would like to run the ASCII images generator, check this [page](https://github.com/programacaoemacao/pokedex-sh/tree/main/cmd/ascii_images_generator).