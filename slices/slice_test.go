package slices

import (
	"collections/maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	expected := []int{1, 2, 3}
	actual := YieldAll[[]int](Filter([]int{1, 2, 3, 4, 5, 6}, func(x int) bool { return x <= 3 }))
	assert.ElementsMatch(t, expected, actual)
	expected1 := []string{"test"}
	actual1 := YieldAll[[]string](Filter([]string{"asdasd", "123123", "test", "test123123"}, func(x string) bool { return x == "test" }))
	assert.ElementsMatch(t, expected1, actual1)
}

func TestFirst(t *testing.T) {
	expected := 1
	actual := FirstMatch([]int{-3, -2, -1, 0, 1, 2, 3}, func(x int) bool { return x > 0 })
	assert.Equal(t, expected, actual)
	expected = 0
	actual = FirstMatch([]int{1, 2, 3}, func(x int) bool { return x < 0 })
	assert.Equal(t, expected, actual)
}

func TestLast(t *testing.T) {
	expected := 3
	actual := LastMatch([]int{-3, -2, -1, 0, 1, 2, 3}, func(x int) bool { return x > 0 })
	assert.Equal(t, expected, actual)
	expected = 0
	actual = LastMatch([]int{1, 2, 3}, func(x int) bool { return x < 0 })
	assert.Equal(t, expected, actual)
}

func TestAny(t *testing.T) {
	assert.True(t, AnyMatch([]int{1, 1, 2, 3, 5, 8, 13}, func(x int) bool { return x > 10 }))
	assert.False(t, AnyMatch([]int{1, 1, 2, 3, 5, 8, 13}, func(x int) bool { return x < 0 }))
}

func TestAll(t *testing.T) {
	assert.True(t, AllMatch([]int{1, 2, 3}, func(x int) bool { return x > 0 }))
	assert.False(t, AllMatch([]int{-1, 0, 1, 2, 3}, func(x int) bool { return x > 0 }))
}

func TestMap(t *testing.T) {
	assert.ElementsMatch(t,
		[]int{2, 4, 6, 8},
		Map[[]int]([]int{1, 2, 3, 4}, func(x int) int { return 2 * x }),
	)
	assert.ElementsMatch(t,
		[]int{4, 1, 6},
		Map[[]string]([]string{"test", "a", "golang"}, func(x string) int { return len(x) }),
	)
}

func TestIndexOf(t *testing.T) {
	assert.Equal(t, 2, IndexOf([]int{1, 2, 3, 4, 5}, 3))
	assert.Equal(t, -1, IndexOf([]int{1, 2, 3, 4, 5}, 6))
}

func TestContains(t *testing.T) {
	assert.True(t, Contains([]int{1, 2, 3, 4}, 3))
	assert.False(t, Contains([]string{"a", "b", "test"}, "wednesday"))
}

type person struct {
	ID        int
	Name      string
	ManagerID int
}

func TestGroupBy(t *testing.T) {
	alice := person{1, "Alice", 0}
	bob := person{2, "Bob", 1}
	carol := person{9, "Carol", 1}
	david := person{11, "David", 9}
	eugene := person{13, "Eugene", 2}
	francene := person{12, "Francene", 9}
	employees := []person{alice, bob, carol, david, eugene, francene}

	assert.Equal(t,
		map[int][]person{
			0: {alice},
			1: {bob, carol},
			2: {eugene},
			9: {david, francene},
		},
		GroupBy(employees, func(p person) int { return p.ManagerID }, func(p person) person { return p }),
	)
}

func TestToSet(t *testing.T) {
	people := []person{
		{
			ID:   1,
			Name: "Alice",
		},
		{
			ID:   2,
			Name: "Bob",
		},
		{
			ID:   9,
			Name: "Carol",
		},
	}

	assert.Equal(t,
		map[int]struct{}{1: {}, 2: {}, 9: {}},
		maps.YieldAll[map[int]struct{}](ToSet(people, func(p person) int { return p.ID })),
	)

	assert.Equal(t,
		map[string]struct{}{
			"Alice": {},
			"Bob":   {},
			"Carol": {},
		},
		maps.YieldAll[map[string]struct{}](ToSet(people, func(p person) string { return p.Name })),
	)
}

func TestToMap(t *testing.T) {
	alice := person{1, "Alice", 0}
	bob := person{2, "Bob", 1}
	carol := person{9, "Carol", 1}
	people := []person{alice, bob, carol}

	assert.Equal(t,
		map[int]person{
			1: alice,
			2: bob,
			9: carol,
		},
		maps.YieldAll[map[int]person](ToMap(people, func(p person) int { return p.ID }, func(p person) person { return p })),
	)

	assert.Equal(t,
		map[string]person{
			"Alice": alice,
			"Bob":   bob,
			"Carol": carol,
		},
		maps.YieldAll[map[string]person](ToMap(people, func(p person) string { return p.Name }, func(p person) person { return p })),
	)
}
