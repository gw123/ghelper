package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSkiplist_Insert(t *testing.T) {
	q := NewSkiplis()

	arr := []int{1, 6, 9, 2, 5, 20, 12, 11, 3, 4, 7, 13, 18, 90, 28, 34, 59, 86, 41, 42, 56}
	for i := -0; i < 100; i++ {
		arr = append(arr, int(rand.Int31n(300)))
	}

	for _, val := range arr {
		newNode := NewNode(val)
		q.Insert(newNode)
	}
	fmt.Printf("display: \n")
	q.Display()

	tests := []map[string]int{
		{
			"input": 1,
			"want":  1,
		},
		{
			"input": 3,
			"want":  3,
		},
		{
			"input": 5,
			"want":  5,
		},
		{
			"input": 8,
			"want":  0,
		},
		{
			"input": 21,
			"want":  0,
		},
	}

	for _, tt := range tests {
		node := q.FindNode(NewNode(tt["input"]))
		if (node != nil && node.data != tt["want"]) || (node == nil && tt["want"] != 0) {
			t.Logf("input :%d ,want: %d , result: %+v", tt["input"], tt["want"], node)
			t.Fatal()
		}
	}

}

func TestLinkedQueue_InsertAfter(t *testing.T) {
	queue := NewLinkedQueue()
	arr := []int{1, 6, 9, 2, 5, 0, 12, -1, 8, 20, 13, 14, 19, 23, 25, 18}
	for _, val := range arr {
		newNode := NewNode(val)
		queue.Insert(newNode)
	}
	queue.Display()

	tests := []map[string]int{
		{
			"input": 1,
			"want":  1,
		},
		{
			"input": 3,
			"want":  2,
		},
		{
			"input": 5,
			"want":  5,
		},
		{
			"input": 8,
			"want":  8,
		},
		{
			"input": 21,
			"want":  20,
		},
	}

	for _, tt := range tests {
		node := queue.FindIndexNode(NewNode(tt["input"]))
		if node.data != tt["want"] {
			t.Logf("input :%d ,want: %d , result: %d", tt["input"], tt["want"], node.data)
			t.Fatal()
		}
	}

	res := queue.FindIndexNode(NewNode(-2))
	if res != nil {
		t.Fail()
	}

	node := queue.FindIndexAfterNode(queue.FindIndexNode(NewNode(1)), NewNode(6))
	if node.data != 6 {
		t.Fail()
	}

	node = queue.FindIndexAfterNode(queue.FindIndexNode(NewNode(3)), NewNode(1))

	if node.data != 2 {
		t.Fail()
	}
}

func TestGetLevel(t *testing.T) {
	type args struct {
		num      uint
		maxLevel uint
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				num:      31,
				maxLevel: 4,
			},
			want: 4,
		},
		{
			name: "t1",
			args: args{
				num:      30,
				maxLevel: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLevel(tt.args.num, tt.args.maxLevel); got != tt.want {
				t.Errorf("GetLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
