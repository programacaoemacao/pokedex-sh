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

[![](https://mermaid.ink/img/pako:eNp9UstOwzAQ_JWVT0EiByjPHJBoK6FKICrK0ZdtsmlN_ZK94SHEB_Ed_BhOUxcOiJwymRl7djbvonYNiUqUZSktK9ZUwZSeSTtvyDK02r1Iu2Wl7UG9xsDwOJYW0rM4Kh4pMvCaYJIYA-NuudTEhHCrlgHDG7QYqxbLVmPcHOxsxz-2mcEVwfViMpvBDVkKyC78ax4Vd7ghQJi7zddnQ68wv5_Ai-I1RGdo-Gycjdmt0bPzZT9qPuOkGHdKN-mQWAf0FLI2etVQyLLTYrGlt0kbZPxbdlZMnH2mVEy-G6IPiikCu91sO6NHZblchi6us_u8mHcMigG1TvoVpcv2efrQSY623usvigdqse5r-iXK7GUxdXW3XV4f2gf3RDXvi1SasjyvEMryKq0kr2aAo1z2AE9ybwM8zf0M8Cz3MMDzPNgAL3LuAV5KKw6FoWBQNenXe-9pKVJYQ1JU6bXBsJFC2o-kw47d4s3WouLQ0aHofFoDTRWuAhqRRtKRPr4BVZnhHw?type=png)](https://mermaid.live/edit#pako:eNp9UstOwzAQ_JWVT0EiByjPHJBoK6FKICrK0ZdtsmlN_ZK94SHEB_Ed_BhOUxcOiJwymRl7djbvonYNiUqUZSktK9ZUwZSeSTtvyDK02r1Iu2Wl7UG9xsDwOJYW0rM4Kh4pMvCaYJIYA-NuudTEhHCrlgHDG7QYqxbLVmPcHOxsxz-2mcEVwfViMpvBDVkKyC78ax4Vd7ghQJi7zddnQ68wv5_Ai-I1RGdo-Gycjdmt0bPzZT9qPuOkGHdKN-mQWAf0FLI2etVQyLLTYrGlt0kbZPxbdlZMnH2mVEy-G6IPiikCu91sO6NHZblchi6us_u8mHcMigG1TvoVpcv2efrQSY623usvigdqse5r-iXK7GUxdXW3XV4f2gf3RDXvi1SasjyvEMryKq0kr2aAo1z2AE9ybwM8zf0M8Cz3MMDzPNgAL3LuAV5KKw6FoWBQNenXe-9pKVJYQ1JU6bXBsJFC2o-kw47d4s3WouLQ0aHofFoDTRWuAhqRRtKRPr4BVZnhHw)

## Attention points

**You don't need to run the scraper and collect any data. The repository already contains the necessary data to run the Pokédex!**

If you would like to run the scraper, check this [page](https://github.com/programacaoemacao/pokedex-sh/tree/main/cmd/data_scraper).

If you would like to run the ASCII images generator, check this [page](https://github.com/programacaoemacao/pokedex-sh/tree/main/cmd/ascii_images_generator).