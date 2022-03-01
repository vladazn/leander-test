package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	require.True(t, testValidity("23-ab-48-caba-56-haha"))
	require.False(t, testValidity("23--48-caba-56-haha"))
	require.False(t, testValidity("23-34-48-caba-56-haha"))
}

func TestAverage(t *testing.T) {
	require.Equal(t, 42, averageNumber("23-ab-48-caba-56-haha"))
}

func TestStory(t *testing.T) {
	require.Equal(t, "ab caba haha", wholeStory("23-ab-48-caba-56-haha"))
}

func TestStates(t *testing.T) {
	short, long, avg, arr := storyStats("23-ab-48-caba-56-haha")
	require.Equal(t, "ab", short)
	require.Equal(t, "caba", long)
	require.Equal(t, 3, avg)

	validArr := []string{"aba", "cab", "hah"}
	require.Equal(t, len(validArr), len(arr))
	for i, el := range arr {
		require.Equal(t, validArr[i], el)
	}
}

func TestGenerate(t *testing.T) {
	require.True(t, testValidity(generate(true)))
	require.False(t, testValidity(generate(false)))
}
