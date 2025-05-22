package tree_sitter_vb_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_vb "github.com/gabriel-gubert/tree-sitter-vb/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_vb.Language())
	if language == nil {
		t.Errorf("Error loading no grammar")
	}
}
