package graph

import (
	"testing"

	"github.com/erikh/turtle/assert"
)

var sanitizesTestCases = map[string]struct {
	str                string
	expected           string
	predicate_expected string
	typ                string
}{
	"empty_string": {
		str:      "",
		expected: "",
		typ:      "iri",
	},
	"iri": {
		str:      "http://www.w3.org/1999/02/22-rdf-syntax-ns#type",
		expected: "<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>",
		typ:      "iri",
	},
	"blank_node": {
		str:      "_:b23",
		expected: "_:b23",
		typ:      "blank",
	},
	"literal": {
		str:      "this is a literal",
		expected: `"this is a literal"`,
		typ:      "literal",
	},
	"multiline literal": {
		str: `this is a
literal`,
		expected: `'''this is a
literal'''`,
		typ: "literal",
	},
	"multiline_literal_apostrophe": {
		str: `this is 'a
literal`,
		expected: `"""this is 'a
literal"""`,
		typ: "literal",
	},
	"multiline_literal_quotation": {
		str: `this is "a
literal`,
		expected: `'''this is "a
literal'''`,
		typ: "literal",
	},
}

func TestSanitize(t *testing.T) {
	g := New()
	for name, tc := range sanitizesTestCases {
		t.Run(name, func(t *testing.T) {
			actual := g.sanitize(tc.str, tc.typ)
			assert.Equal(t, tc.expected, actual, "function should have returned correctly sanitized string")
		})
	}
}
