package maps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeysValues(t *testing.T) {
	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	assert.ElementsMatch(t, []int{1, 2, 3}, Keys(m))
	assert.ElementsMatch(t, []string{"one", "two", "three"}, Values(m))
}

func TestUnion(t *testing.T) {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	b := map[int]string{
		3: "threeb",
		4: "four",
		5: "five",
	}

	expected := map[int]string{
		1: "one",
		2: "two",
		3: "threeb",
		4: "four",
		5: "five",
	}

	assert.Equal(t, expected, Union(a, b))
	assert.NotEqual(t, expected, a)
	inplace := UnionInPlace(a, b)
	assert.Equal(t, expected, inplace)
	assert.Equal(t, inplace, a)
}

func TestIntersect(t *testing.T) {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	b := map[int]string{
		3: "threeb",
		4: "four",
		5: "five",
		1: "one",
	}

	expected := map[int]string{
		1: "one",
		3: "threeb",
	}

	assert.Equal(t, expected, Intersect(a, b))
}

func TestDifference(t *testing.T) {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	b := map[int]string{
		3: "threeb",
		4: "four",
		5: "five",
		1: "one",
	}

	expected := map[int]string{
		2: "two",
	}

	assert.Equal(t, expected, Difference(a, b))

}
