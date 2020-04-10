package org2qmk

import (
	"github.com/niklasfasching/go-org/org"
)

// Implementation of org.Writer that emits ErgoDone keymap.c
// https://github.com/niklasfasching/go-org/blob/master/org/writer.go
type ErgodoneWriter struct {
}

func (w *ErgodoneWriter) Before(d *org.Document) {}

func (w *ErgodoneWriter) After(d *org.Document) {}

func (w *ErgodoneWriter) String() string {
	return "// I am a keymap\n"
}

func (w *ErgodoneWriter) WriterWithExtensions() org.Writer {
	return w
}
func (w *ErgodoneWriter) WriteNodesAsString(...org.Node) string {
	return ""
}
func (w *ErgodoneWriter) WriteKeyword(org.Keyword)                         {}
func (w *ErgodoneWriter) WriteInclude(org.Include)                         {}
func (w *ErgodoneWriter) WriteComment(org.Comment)                         {}
func (w *ErgodoneWriter) WriteNodeWithMeta(org.NodeWithMeta)               {}
func (w *ErgodoneWriter) WriteNodeWithName(org.NodeWithName)               {}
func (w *ErgodoneWriter) WriteHeadline(org.Headline)                       {}
func (w *ErgodoneWriter) WriteBlock(org.Block)                             {}
func (w *ErgodoneWriter) WriteExample(org.Example)                         {}
func (w *ErgodoneWriter) WriteDrawer(org.Drawer)                           {}
func (w *ErgodoneWriter) WritePropertyDrawer(org.PropertyDrawer)           {}
func (w *ErgodoneWriter) WriteList(org.List)                               {}
func (w *ErgodoneWriter) WriteListItem(org.ListItem)                       {}
func (w *ErgodoneWriter) WriteDescriptiveListItem(org.DescriptiveListItem) {}
func (w *ErgodoneWriter) WriteTable(org.Table)                             {}
func (w *ErgodoneWriter) WriteHorizontalRule(org.HorizontalRule)           {}
func (w *ErgodoneWriter) WriteParagraph(org.Paragraph)                     {}
func (w *ErgodoneWriter) WriteText(org.Text)                               {}
func (w *ErgodoneWriter) WriteEmphasis(org.Emphasis)                       {}
func (w *ErgodoneWriter) WriteLatexFragment(org.LatexFragment)             {}
func (w *ErgodoneWriter) WriteStatisticToken(org.StatisticToken)           {}
func (w *ErgodoneWriter) WriteExplicitLineBreak(org.ExplicitLineBreak)     {}
func (w *ErgodoneWriter) WriteLineBreak(org.LineBreak)                     {}
func (w *ErgodoneWriter) WriteRegularLink(org.RegularLink)                 {}
func (w *ErgodoneWriter) WriteTimestamp(org.Timestamp)                     {}
func (w *ErgodoneWriter) WriteFootnoteLink(org.FootnoteLink)               {}
func (w *ErgodoneWriter) WriteFootnoteDefinition(org.FootnoteDefinition)   {}

func NewErgodoneWriter() *ErgodoneWriter {
	return &ErgodoneWriter{}
}
