package org2qmk

import (
	"fmt"
	"strings"

	"github.com/niklasfasching/go-org/org"
)

// Implementation of org.Writer that emits ErgoDone keymap.c
// https://github.com/niklasfasching/go-org/blob/master/org/writer.go
type ErgodoneWriter struct {
	strings.Builder
}

func NewErgodoneWriter() *ErgodoneWriter {
	return &ErgodoneWriter{}
}

func (w *ErgodoneWriter) Before(d *org.Document) {
	w.WriteString("// start of keymap.c\n")
}

func (w *ErgodoneWriter) After(d *org.Document) {
	w.WriteString("// end of keymap.c\n")
}

func (w *ErgodoneWriter) WriterWithExtensions() org.Writer {
	return w
}

func (w *ErgodoneWriter) WriteNodesAsString(nodes ...org.Node) string {
	builder := w.Builder
	w.Builder = strings.Builder{}
	org.WriteNodes(w, nodes...)
	out := w.String()
	w.Builder = builder
	return out
}

func (w *ErgodoneWriter) WriteKeyword(org.Keyword)           {}
func (w *ErgodoneWriter) WriteInclude(org.Include)           {}
func (w *ErgodoneWriter) WriteComment(org.Comment)           {}
func (w *ErgodoneWriter) WriteNodeWithMeta(org.NodeWithMeta) {}
func (w *ErgodoneWriter) WriteNodeWithName(org.NodeWithName) {}

func (w *ErgodoneWriter) WriteHeadline(h org.Headline) {
	// follow the document tree down
	org.WriteNodes(w, h.Children...)
}
func (w *ErgodoneWriter) WriteBlock(b org.Block) {
	// Block should come across verbatim if they're c/c++
	w.WriteString(fmt.Sprintf("WriteBlocks %#v\n", b))
}

func (w *ErgodoneWriter) WriteExample(org.Example)               {}
func (w *ErgodoneWriter) WriteDrawer(org.Drawer)                 {}
func (w *ErgodoneWriter) WritePropertyDrawer(org.PropertyDrawer) {}

func (w *ErgodoneWriter) WriteList(l org.List) {
	org.WriteNodes(w, l.Items...)

}

func (w *ErgodoneWriter) WriteListItem(org.ListItem)                       {}
func (w *ErgodoneWriter) WriteDescriptiveListItem(org.DescriptiveListItem) {}

func (w *ErgodoneWriter) WriteTable(t org.Table) {
	w.WriteString(fmt.Sprintf("WriteTable %#v\n", t))
}

func (w *ErgodoneWriter) WriteHorizontalRule(org.HorizontalRule)         {}
func (w *ErgodoneWriter) WriteParagraph(org.Paragraph)                   {}
func (w *ErgodoneWriter) WriteText(org.Text)                             {}
func (w *ErgodoneWriter) WriteEmphasis(org.Emphasis)                     {}
func (w *ErgodoneWriter) WriteLatexFragment(org.LatexFragment)           {}
func (w *ErgodoneWriter) WriteStatisticToken(org.StatisticToken)         {}
func (w *ErgodoneWriter) WriteExplicitLineBreak(org.ExplicitLineBreak)   {}
func (w *ErgodoneWriter) WriteLineBreak(org.LineBreak)                   {}
func (w *ErgodoneWriter) WriteRegularLink(org.RegularLink)               {}
func (w *ErgodoneWriter) WriteTimestamp(org.Timestamp)                   {}
func (w *ErgodoneWriter) WriteFootnoteLink(org.FootnoteLink)             {}
func (w *ErgodoneWriter) WriteFootnoteDefinition(org.FootnoteDefinition) {}
