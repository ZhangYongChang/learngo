package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcbb", 3},
		{"一二三三二一", 3},
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}

	for _, tt := range tests {
		actual := lengthOfNonRepeatingSubStr2(tt.s)
		if actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubStr(b *testing.B) {
	s := "为是年解放萨的房间内克的思考国家考虑过开放嗯规划法国当局的黑色风格世纪的法国三大和"
	ans := 20
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}

func BenchmarkSubStr2(b *testing.B) {
	s := "为是年解放萨的房间内克的思考国家考虑过开放嗯规划法国当局的黑色风格世纪的法国三大和"
	ans := 20
	for i := 0; i < b.N; i++ {
		actual := lengthOfNonRepeatingSubStr2(s)
		if actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
