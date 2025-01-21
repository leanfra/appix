package biz_test

import (
	"appix/internal/biz"
	"reflect"
	"testing"
)

// TestDiffSliceUint32 测试 DiffSliceUint32 函数的各种情况
func TestDiffSliceUint32(t *testing.T) {
	tests := []struct {
		name     string
		s1       []uint32
		s2       []uint32
		expected []uint32
	}{
		{"EmptySlices", []uint32{}, []uint32{}, []uint32{}},
		{"EmptyS1", []uint32{}, []uint32{1, 2, 3}, []uint32{}},
		{"EmptyS2", []uint32{1, 2, 3}, []uint32{}, []uint32{1, 2, 3}},
		{"NoCommonElements", []uint32{1, 2, 3}, []uint32{4, 5, 6}, []uint32{1, 2, 3}},
		{"AllCommonElements", []uint32{1, 2, 3}, []uint32{1, 2, 3}, []uint32{}},
		{"PartialOverlap", []uint32{1, 2, 3, 4}, []uint32{3, 4, 5, 6}, []uint32{1, 2}},
		{"DuplicatesInS1", []uint32{1, 2, 2, 3}, []uint32{2, 3}, []uint32{1}},
		{"DuplicatesInS2", []uint32{1, 2, 3}, []uint32{2, 2, 3}, []uint32{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := biz.DiffSliceUint32(tt.s1, tt.s2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DiffSliceUint32(%v, %v) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestIntersectSliceUint32(t *testing.T) {
	tests := []struct {
		name     string
		s1       []uint32
		s2       []uint32
		expected []uint32
	}{
		{"EmptySlices", []uint32{}, []uint32{}, []uint32{}},
		{"EmptyS1", []uint32{}, []uint32{1, 2, 3}, []uint32{}},
		{"EmptyS2", []uint32{1, 2, 3}, []uint32{}, []uint32{}},
		{"NoCommonElements", []uint32{1, 2, 3}, []uint32{4, 5, 6}, []uint32{}},
		{"AllCommonElements", []uint32{1, 2, 3}, []uint32{1, 2, 3}, []uint32{1, 2, 3}},
		{"PartialOverlap", []uint32{1, 2, 3, 4}, []uint32{3, 4, 5, 6}, []uint32{3, 4}},
		{"DuplicatesInS1", []uint32{1, 2, 2, 3}, []uint32{2, 3}, []uint32{2, 3}},
		{"DuplicatesInS2", []uint32{1, 2, 3}, []uint32{2, 2, 3}, []uint32{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := biz.IntersectSliceUint32(tt.s1, tt.s2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DiffSliceUint32(%v, %v) = %v, want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
