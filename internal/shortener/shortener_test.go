package shortener

import (
	"testing"
	"github.com/mradigen/short/internal/storage"
)

func TestUrlShortening(t* testing.T) {

	tests := []struct {
		name	string
		url		string
		wantErr	bool
	} {
		{"Valid URL",	"https://aadivishnu.com", false},
		{"Valid URL",	"http://tevatel.com", false},
		{"Valid URL", "example.com", true},
		//{"Invalid URL", "htp://example.com", true}
	}

	m := storage.NewMemory()
	s := New(m)

	for _, test := range tests {
		t.Run(test.url, func(t* testing.T) {
			got, err := s.Shorten(test.url)
			if (err != nil) != test.wantErr {
				t.Errorf("Error occured: %v", err)
			} else if !test.wantErr && len(got) != 4 {
				t.Errorf("Unexpected length")
			}
		})
	}

}
