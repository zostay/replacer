package replacer_test

import (
	"log"
	"strings"
	"testing"

	"github.com/zostay/replacer"
)

func TestMain(m *testing.M) {
	log.Println("Testing github.com/zostay/replacer Version", replacer.Version)
}

func TestReplacer_Replace_Grow(t *testing.T) {
	t.Parallel()

	r := replacer.New(
		"modbyid", "mod_by_id",
		"modby", "mod_by",
		"modbydate", "mod_by_date",
		"", "ignored",
		"mod", "m_o_d",
		"mega", strings.Repeat("mega", 100))

	if r == nil {
		t.Errorf("replacer.New should return an object")
	}

	testCases := []struct {
		in  string
		exp string
	}{
		{in: "mod", exp: "m_o_d"},
		{in: "moddy", exp: "m_o_ddy"},
		{in: "meh", exp: "meh"},
		{in: "modbx", exp: "m_o_dbx"},
		{in: "modby", exp: "mod_by"},
		{in: "mmmmmod", exp: "mmmmm_o_d"},
		{in: "modbyid", exp: "mod_by_id"},
		{in: "modbydate", exp: "mod_by_date"},
		{in: "megamod", exp: strings.Repeat("mega", 100) + "m_o_d"},
	}

	for _, tc := range testCases {
		s := r.Replace(tc.in)
		if s != tc.exp {
			t.Errorf("r.Replace(%q): expected %q but got %q", tc.in, tc.exp, s)
		}
	}
}

func TestReplacer_Replace_Shrink(t *testing.T) {
	t.Parallel()

	r := replacer.New(
		"modbyid", "mbi",
		"modby", "mb",
		"modbydate", "mbd",
		"mod", "")

	if r == nil {
		t.Errorf("replacer.New should return an object")
	}

	testCases := []struct {
		in  string
		exp string
	}{
		{in: "mod", exp: ""},
		{in: "moddy", exp: "dy"},
		{in: "meh", exp: "meh"},
		{in: "modbx", exp: "bx"},
		{in: "modby", exp: "mb"},
		{in: "mmmmmod", exp: "mmmm"},
		{in: "modbyid", exp: "mbi"},
		{in: "modbydate", exp: "mbd"},
	}

	for _, tc := range testCases {
		s := r.Replace(tc.in)
		if s != tc.exp {
			t.Errorf("r.Replace(%q): expected %q but got %q", tc.in, tc.exp, s)
		}
	}
}
