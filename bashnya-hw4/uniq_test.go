package main

import (
	"reflect"
	"testing"
)

type InputTestStrings struct {
	test1to4, test5, test6, test7to8and12, test9to10, test11 []string
}
type OutputTestStrings struct {
	one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve []string
}

func TestProcessing(t *testing.T) {
	inputStrings := InputTestStrings{
		test1to4:      []string{"I love music.", "I love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
		test5:         []string{"I LOVE MUSIC.", "I love music.", "I LoVe MuSiC.", "", "I love MuSIC of Kartik.", "I love music of kartik.", "Thanks."},
		test6:         []string{"We love music.", "I love music.", "They love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		test7to8and12: []string{"I love music.", "A love music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		test9to10:     []string{"I lovE music.", "A lovE music.", "C love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		test11:        []string{"I Love music.", "I Love music.", "I love music.", "", "I love music of Kartik.", "I love music of Kartik.", "Thanks."},
	}
	outputStrings := OutputTestStrings{
		one:    []string{"I love music.", "", "I love music of Kartik.", "Thanks."},
		two:    []string{"3 I love music.", "1 ", "2 I love music of Kartik.", "1 Thanks."},
		three:  []string{"I love music.", "I love music of Kartik."},
		four:   []string{"", "Thanks."},
		five:   []string{"I LOVE MUSIC.", "", "I love MuSIC of Kartik.", "Thanks."},
		six:    []string{"We love music.", "", "I love music of Kartik.", "Thanks."},
		seven:  []string{"I love music.", "", "I love music of Kartik.", "We love music of Kartik.", "Thanks."},
		eight:  []string{"I love music.", "", "I love music of Kartik.", "Thanks."},
		nine:   []string{"3 I lovE music.", "1 ", "1 I love music of Kartik.", "1 We love music of Kartik.", "1 Thanks."},
		ten:    []string{"3 I lovE music.", "1 ", "2 I love music of Kartik.", "1 Thanks."},
		eleven: []string{"", "Thanks."},
		twelve: []string{"I love music.", "", "I love music of Kartik.", "Thanks."},
	}
	tests := []struct {
		name  string
		in    []string
		out   []string
		flags Flags
	}{
		{name: "1: No parameters",
			in:    inputStrings.test1to4,
			out:   outputStrings.one,
			flags: Flags{},
		}, {
			"2: -c flag (count)",
			inputStrings.test1to4,
			outputStrings.two,
			Flags{count: true},
		}, {
			"3: -d flag (repeated)",
			inputStrings.test1to4,
			outputStrings.three,
			Flags{repeated: true},
		}, {
			"4: -u flag (uniq)",
			inputStrings.test1to4,
			outputStrings.four,
			Flags{uniq: true},
		}, {
			"5: -i flag (register)",
			inputStrings.test5,
			outputStrings.five,
			Flags{register: true},
		}, {
			"6: -f flag (numFields: 1)",
			inputStrings.test6,
			outputStrings.six,
			Flags{numFields: 1},
		}, {
			"7: -s flag (numChars: 1)",
			inputStrings.test7to8and12,
			outputStrings.seven,
			Flags{numChars: 1},
		}, {
			"8: -f flag (numFields: 2)",
			inputStrings.test7to8and12,
			outputStrings.eight,
			Flags{numFields: 2},
		}, {
			"9: -ics flags (register + count + numChars: 2)",
			inputStrings.test9to10,
			outputStrings.nine,
			Flags{count: true, register: true, numChars: 2},
		}, {
			"10: -cif flags (count + register + numFields: 1)",
			inputStrings.test9to10,
			outputStrings.ten,
			Flags{count: true, register: true, numFields: 1},
		}, {
			"11: -ui flags (uniq + register)",
			inputStrings.test11,
			outputStrings.eleven,
			Flags{uniq: true, register: true},
		}, {
			"12: -fs flags (numFields: 1 + numChars: 1)",
			inputStrings.test7to8and12,
			outputStrings.twelve,
			Flags{numFields: 1, numChars: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.flags.Processing(tt.in)
			if !reflect.DeepEqual(actual, tt.out) {
				t.Errorf("\n%v: got %v, want %v", tt.name, actual, tt.out)
			}
		})
	}
}
