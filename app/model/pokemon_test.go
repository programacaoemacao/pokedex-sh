package model

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPokemon_GetFormattedID(t *testing.T) {
	t.Run("Test with Pokémon 1", func(t *testing.T) {
		pokemon := Pokemon{ID: 1}
		actual := pokemon.GetFormattedID()

		require.Equal(t, actual, "0001")
	})

	t.Run("Test with Pokémon 23", func(t *testing.T) {
		pokemon := Pokemon{ID: 23}
		actual := pokemon.GetFormattedID()

		require.Equal(t, actual, "0023")
	})

	t.Run("Test with Pokémon 283", func(t *testing.T) {
		pokemon := Pokemon{ID: 283}
		actual := pokemon.GetFormattedID()

		require.Equal(t, actual, "0283")
	})

	t.Run("Test with Pokémon 1010", func(t *testing.T) {
		pokemon := Pokemon{ID: 1010}
		actual := pokemon.GetFormattedID()

		require.Equal(t, actual, "1010")
	})
}
