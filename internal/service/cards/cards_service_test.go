package cards

import (
	"testing"
)

func TestService_GetCardArchetypeByToken(t *testing.T) {
	tests := []struct {
		name        string
		token       string
		withEdition bool
		res         string
		want        string
	}{
		{name: "1003006001002028000000003 -> 006-001-002-028", token: "1003006001002028000000003", want: "006-001-002-028"},
		{name: "1003006001001004000000014 -> 006-001-001-004", token: "1003006001001004000000014", want: "006-001-001-004"},
		{name: "1003006001005019000000001 -> 006-001-005-019", token: "1003006001005019000000001", want: "006-001-005-019"},
		{name: "1003006001001003000000007 -> 006-001-001-003", token: "1003006001001003000000007", want: "006-001-001-003"},
		{name: "1003006001004011000000001 -> 006-001-004-011", token: "1003006001004011000000001", want: "006-001-004-011"},
		{name: "1003006001002028000000003 -> 006-001-002-028", token: "1003006001002028000000003", want: "003-006-001-002-028", withEdition: true},
		{name: "1003006001001004000000014 -> 006-001-001-004", token: "1003006001001004000000014", want: "003-006-001-001-004", withEdition: true},
		{name: "1003006001005019000000001 -> 006-001-005-019", token: "1003006001005019000000001", want: "003-006-001-005-019", withEdition: true},
		{name: "1003006001001003000000007 -> 006-001-001-003", token: "1003006001001003000000007", want: "003-006-001-001-003", withEdition: true},
		{name: "1003006001004011000000001 -> 006-001-004-011", token: "1003006001004011000000001", want: "003-006-001-004-011", withEdition: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// s := &Service{}
			if got := GetCardArchetypeByToken(tt.token, tt.withEdition); got != tt.want {
				t.Errorf("GetCardArchetypeByToken(%v) = %v, want %v", tt.token, got, tt.want)
			}
		})
	}
}

func TestService_GetUniqueIdFromArchetypeId(t *testing.T) {
	tests := []struct {
		name        string
		archetypeID string
		res         string
		want        string
	}{
		{name: "006-001-002-028 -> 001-028", archetypeID: "006-001-002-028", want: "001-028"},
		{name: "006-001-001-004 -> 001-004", archetypeID: "006-001-001-004", want: "001-004"},
		{name: "006-001-005-019 -> 001-019", archetypeID: "006-001-005-019", want: "001-019"},
		{name: "006-001-001-003 -> 001-003", archetypeID: "006-001-001-003", want: "001-003"},
		{name: "006-001-004-011 -> 001-011", archetypeID: "006-001-004-011", want: "001-011"},
		{name: "003-006-001-002-028 -> 001-028", archetypeID: "003-006-001-002-028", want: "001-028"},
		{name: "003-006-001-001-004 -> 001-004", archetypeID: "003-006-001-001-004", want: "001-004"},
		{name: "003-006-001-005-019 -> 001-019", archetypeID: "003-006-001-005-019", want: "001-019"},
		{name: "003-006-001-001-003 -> 001-003", archetypeID: "003-006-001-001-003", want: "001-003"},
		{name: "003-006-001-004-011 -> 001-011", archetypeID: "003-006-001-004-011", want: "001-011"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{}
			if got := s.GetUniqueIdFromArchetypeId(tt.archetypeID); got != tt.want {
				t.Errorf("GetUniqueIdFromArchetypeId(%v) = %v, want %v", tt.archetypeID, got, tt.want)
			}
		})
	}
}

func BenchmarkService_GetCardArchetypeByToken(b *testing.B) {
	// s := &Service{}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetCardArchetypeByToken("1003006001002028000000003", true)
	}
}
