package goosli

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/require"
	"math"
)

func TestVector_Cross(t *testing.T) {
	cases := []struct {
		in1, in2, out Vector
	}{
		{
			in1: Vector{1, 1, 1},
			in2: Vector{1, 1, 1},
			out: Vector{0, 0, 0},
		},
		{
			in1: Vector{2, 1, 0},
			in2: Vector{1, 1, 1},
			out: Vector{1, -2, 1},
		},
		{
			in1: Vector{1, 1, 1},
			in2: Vector{2, 1, 0},
			out: Vector{-1, 2, -1},
		},
		{
			in1: Vector{0, 0, 0},
			in2: Vector{1, -1, 2},
			out: Vector{0, 0, 0},
		},
	}

	for i, row := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := row.in1.Cross(row.in2)
			require.Equal(t, row.out, res)
		})
	}
}

func TestVector_CodirectedWith(t *testing.T) {
	cases := []struct {
		in1, in2 Vector
		out      bool
	}{
		{
			in1: Vector{1, 1, 1},
			in2: Vector{1, 1, 1},
			out: true,
		},
		{
			in1: Vector{2, 1, 0},
			in2: Vector{1, 1, 1},
			out: true,
		},
		{
			in1: Vector{1, 1, 1},
			in2: Vector{2, 1, 0},
			out: true,
		},
		{
			in1: Vector{0, 0, 0},
			in2: Vector{1, -1, 2},
			out: true,
		},
		{
			in1: Vector{-1, -1, -1},
			in2: Vector{1, 1, 1},
			out: false,
		},
		{
			in1: Vector{1, 0, 0},
			in2: Vector{0, 0, 1},
			out: true,
		},
	}

	for i, row := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := row.in1.CodirectedWith(row.in2)
			require.Equal(t, row.out, res)
		})
	}
}

func TestVector_Normalize(t *testing.T) {
	cases := []struct {
		in, out Vector
	}{
		{
			in: Vector{1, 1, 1},
			out: Vector{1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3)},
		},
		{
			in: Vector{-1, 1, 1},
			out: Vector{-1/math.Sqrt(3), 1/math.Sqrt(3), 1/math.Sqrt(3)},
		},
		{
			in: Vector{0, 1, 0},
			out: Vector{0, 1, 0},
		},
		{
			in: Vector{0, 0, 0},
			out: Vector{0, 0, 0},
		},

	}
	for i, row := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := row.in.Normalize()
			require.Equal(t, row.out, res)
		})
	}
}


func TestVector_ProjectOn(t *testing.T) {
	cases := []struct {
		in1, in2, out Vector
	}{
		{
			in1: Vector{1, 1, 1},
			in2: Vector{1, 1, 1},
			out: Vector{1, 1, 1},
		},
		{
			in1: Vector{2, 1, 0},
			in2: Vector{1, 1, 1},
			out: Vector{1, 1, 1},
		},
		{
			in1: Vector{1, 1, 1},
			in2: Vector{2, 1, 0},
			out: Vector{1.2, 0.6, 0},
		},
		{
			in1: Vector{0, 0, 0},
			in2: Vector{1, -1, 2},
			out: Vector{0, 0, 0},
		},
		{
			in1: Vector{1, -1, 2},
			in2: Vector{0, 0, 0},
			out: Vector{0, 0, 0},
		},
		{
			in1: Vector{1, 1, 1},
			in2: Vector{0, 0, 1},
			out: Vector{0, 0, 1},
		},
		{
			in1: Vector{1, 1, 1},
			in2: Vector{0, 0, 2},
			out: Vector{0, 0, 1},
		},
	}
	for i, row := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			res := row.in1.ProjectOn(row.in2)
			require.Equal(t, row.out, res)
		})
	}
}