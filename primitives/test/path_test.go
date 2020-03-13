package test

import (
	"fmt"
	//"fmt"
	"testing"
	//"github.com/stretchr/testify/require"
	. "github.com/l1va/goosli/primitives"
	"github.com/stretchr/testify/require"
)

func TestPath_Join(t *testing.T) { //TODO: add more and improve
	var cases = []struct {
		in  []Path
		out []Path
	}{
		{
			in: []Path{
				{Points: []Point{{1, 1, 1}, {1, 1, 2}}},
				{Points: []Point{{1, 1, 2}, {1, 1, 3}}},
				{Points: []Point{{1, 1, 3}, {1, 1, 4}}},
				{Points: []Point{{1, 1, 4}, {1, 1, 1}}},
			},
			out: []Path{{Points: []Point{{1, 1, 1}, {1, 1, 2}, {1, 1, 3},
				{1, 1, 4}, {1, 1, 1}}}},
		},
		{
			in: []Path{
				{Points: []Point{{1, 1, 1}, {1, 2, 1}}},
				{Points: []Point{{1, 4, 1}, {1, 1, 1}}},
				{Points: []Point{{1, 3, 1}, {1, 4, 1}}},
				{Points: []Point{{1, 2, 1}, {1, 3, 1}}},
			},
			out: []Path{{Points: []Point{{1, 3, 1}, {1, 4, 1}, {1, 1, 1},
				{1, 2, 1}, {1, 3, 1}}}},
		},
		{
			in: []Path{
				{Points: []Point{{1, 1, 1}, {1, 2, 1}, {1, 3, 1}}},
				{Points: []Point{{1, 1, 1}, {1, 2, 1}}},
				{Points: []Point{{1, 1, 1}, {1, 2, 1}, {1, 3, 1}}},
				{Points: []Point{{1, 2, 1}, {1, 3, 1}}},
			},
			out: []Path{{Points: []Point{{1, 1, 1}, {1, 2, 1}, {1, 3, 1},
				{1, 2, 1}, {1, 1, 1}, {1, 2, 1}, {1, 3, 1}}}},
		},
		{
			in: []Path{
				{Points: []Point{{1, 1, 1}, {2, 1, 1}, {3, 1, 1}}},
				{Points: []Point{{4, 1, 1}, {5, 1, 1}}},
			},
			out: []Path{
				{Points: []Point{{1, 1, 1}, {2, 1, 1}, {3, 1, 1}}},
				{Points: []Point{{4, 1, 1}, {5, 1, 1}}},},
		},
	}
	for i, row := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := JoinPaths2(row.in)
			require.Equal(t, len(row.out), len(res))
			for _, p := range res {
				found := false
				for _, p2 := range row.out {
					if p.Equal(p2) {
						found = true
					}
				}
				require.True(t, found, "not found: ", p)
			}
		})
	}
}

func TestPath_FindCentroid(t *testing.T) {
	cases := []struct {
		in  Path
		out Point
	}{
		{
			in:  Path{Points: []Point{{1, 1, 1}, {1, 1, 2}, {1, 2, 2}, {X: 1, Y: 2, Z: 1}, {1, 1, 1}}},
			out: Point{X: 1, Y: 1.5, Z: 1.5},
		},
		{
			in:  Path{Points: []Point{{2, 1, 1}, {2, 1, 2}, {1, 2, 2}, {1, 2, 1}, {2, 1, 1}}},
			out: Point{1.5, 1.5, 1.5},
		},
		{
			in:  Path{Points: []Point{{2, 1, 1}, {2, 2, 2}, {1, 3, 1}, {1, 2, 3}, {2, 1, 1}},},
			out: Point{1.5, 2, 2},
		},
		{
			in:  Path{Points: []Point{{15, 3, 2}, {3, 2, 12}, {1, 13, 10}, {12, 11, 3}, Point{15, 3, 2}},},
			out: Point{X: 7.472222222222222, Y: 7.138888888888889, Z: 7.444444444444445},
		},
	}
	for i, row := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			require.Equal(t, row.out, FindCentroid(row.in))
		})
	}
}
