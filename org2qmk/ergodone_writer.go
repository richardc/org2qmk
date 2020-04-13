package org2qmk

import (
	"fmt"
	"strings"

	"github.com/johncgriffin/yogofn"
	"github.com/niklasfasching/go-org/org"
)

// Implementation of org.Writer that emits ErgoDone keymap.c
// https://github.com/niklasfasching/go-org/blob/master/org/writer.go
type ErgodoneWriter struct {
	inKeymap bool

	strings.Builder
}

func NewErgodoneWriter() *ErgodoneWriter {
	return &ErgodoneWriter{}
}

// Basic keymappings, built up by =init()= also
// TODO(richardc): this would probably be better off in an org-table that's included?
var mappings = map[string]string{
	// https://docs.qmk.fm/#/keycodes?id=basic-keycodes
	"noop":        "KC_NO",
	"\"\"":        "KC_TRNS",
	"enter":       "KC_ENT",
	"escape":      "KC_ESC",
	"esc":         "KC_ESC",
	"backspace":   "KC_BSPC",
	"tab":         "KC_TAB",
	"space":       "KC_SPACE",
	"-":           "KC_MINS",
	"=":           "KC_EQL",
	"[":           "KC_LBRC",
	"]":           "KC_RBRC",
	"\\":          "KC_BSLS",
	"#":           "KC_NONUS_HASH",
	";":           "KC_SCOLON",
	"'":           "KC_QUOT",
	"`":           "KC_GRAVE",
	",":           "KC_COMM",
	".":           "KC_DOT",
	"/":           "KC_SLSH",
	"capslock":    "KC_CAPSLOCK",
	"printscreen": "KC_PSCR",
	"scrolllock":  "KC_SCROLLOCK",
	"pause":       "KC_PAUSE",
	"insert":      "KC_INS",
	"home":        "KC_HOME",
	"pgup":        "KC_PGUP",
	"delete":      "KC_DEL",
	"end":         "KC_END",
	"pgdn":        "KC_PGDOWN",
	"pgdown":      "KC_PGDOWN",
	"left":        "KC_LEFT",
	"right":       "KC_RGHT",
	"down":        "KC_DOWN",
	"up":          "KC_UP",
	"numlock":     "KC_NUMLOCK",

	// skipped a bunch because lazy

	"ctrl_l":  "KC_LCTRL",
	"shift_l": "KC_LSFT",
	"alt_l":   "KC_LALT",
	"gui_l":   "KC_LGUI",
	"cmd_l":   "KC_LCMD",
	"win_l":   "KC_LWIN",
	"ctrl_r":  "KC_RCTRL",
	"shift_r": "KC_RSFT",
	"alt_r":   "KC_RALT",
	"gui_r":   "KC_RGUI",
	"cmd_r":   "KC_RCMD",
	"win_r":   "KC_RWIN",

	"mute":           "KC_MUTE",
	"volume_up":      "KC_VOLU",
	"volume_down":    "KC_VOLD",
	"next_track":     "KC_MNXT",
	"prev_track":     "KC_MPRV",
	"previous_track": "KC_MPRV",
	"stop_track":     "KC_MSTP",
	"play":           "KC_MPLY",

	"back": "KC_WWW_BACK",

	// https://docs.qmk.fm/#/keycodes?id=us-ansi-shifted-symbols
	"~":         "KC_TILDE",
	"!":         "KC_EXCLAIM",
	"@":         "KC_AT",
	"shifted_#": "KC_HASH",
	"$":         "KC_DOLLAR",
	"%":         "KC_PERCENT",
	"^":         "KC_CIRCUMFLEX",
	"&":         "KC_AMPERSAND",
	"*":         "KC_ASTR",
	"(":         "KC_LEFT_PAREN",
	")":         "KC_RIGHT_PAREN",
	"_":         "KC_UNDS",
	"+":         "KC_PLUS",
	"{":         "KC_LCBR",
	"}":         "KC_RCBR",
	"\\vert":    "KC_PIPE",

	// https://docs.qmk.fm/#/keycodes?id=mouse-keys
	"mouse_up":      "KC_MS_UP",
	"mouse_down":    "KC_MS_DOWN",
	"mouse_left":    "KC_MS_LEFT",
	"mouse_right":   "KC_MS_RIGHT",
	"mouse_button1": "KC_MS_BTN1",
	"mouse_button2": "KC_MS_BTN2",
	"mouse_button3": "KC_MS_BTN3",
	"mouse_button4": "KC_MS_BTN4",
	"mouse_button5": "KC_MS_BTN5",

	// synthetic
	"alt_shift_l": "LALT(KC_LSFT)",
	"meh":         "MEH_T(KC_NO)",
	"hyper":       "ALL_T(KC_NO)",
}

func init() {
	for c := 'a'; c <= 'z'; c++ {
		mappings[string(c)] = strings.ToUpper("KC_" + string(c))
	}
	for c := '0'; c <= '9'; c++ {
		mappings[string(c)] = strings.ToUpper("KC_" + string(c))
	}
	for f := 1; f <= 24; f++ {
		mappings[fmt.Sprintf("f%d", f)] = fmt.Sprintf("KC_F%d", f)
	}
}

func (w *ErgodoneWriter) TranslateKeycode(descriptive string) string {
	if keycode, ok := mappings[strings.ToLower(descriptive)]; ok {
		return keycode
	}

	// Must've been something raw
	return descriptive
}

func (w *ErgodoneWriter) Before(d *org.Document) {}
func (w *ErgodoneWriter) After(d *org.Document)  {}

func (w *ErgodoneWriter) WriterWithExtensions() org.Writer { return w }

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

func (w *ErgodoneWriter) startKeymap(s string) {
	if w.inKeymap {
		w.WriteString("),\n")
	} else {
		w.WriteString("const uint16_t PROGMEM keymaps[][MATRIX_ROWS][MATRIX_COLS] = {\n")
	}

	w.WriteString("[")
	w.WriteString(s)
	w.WriteString("] = LAYOUT_ergodox(\n")
	w.inKeymap = false
}

func (w *ErgodoneWriter) exitKeymap() {
	if !w.inKeymap {
		return
	}
	w.WriteString(")\n};\n")
	w.inKeymap = false
}

func (w *ErgodoneWriter) WriteHeadline(h org.Headline) {
	if h.Properties != nil {
		if layer, ok := h.Properties.Get("LAYER"); ok {
			w.startKeymap(layer)
		}
	}

	// follow the document tree down
	org.WriteNodes(w, h.Children...)
}

func (w *ErgodoneWriter) WriteBlock(b org.Block) {
	// Block should come across verbatim if they're c/c++
	if b.Name == "SRC" && b.Parameters[0] == "c" {
		w.exitKeymap()
		org.WriteNodes(w, b.Children...)
	}
}

func (w *ErgodoneWriter) WriteExample(org.Example)               {}
func (w *ErgodoneWriter) WriteDrawer(d org.Drawer)               {}
func (w *ErgodoneWriter) WritePropertyDrawer(org.PropertyDrawer) {}

func (w *ErgodoneWriter) WriteList(l org.List) {
	org.WriteNodes(w, l.Items...)
}

func (w *ErgodoneWriter) WriteListItem(org.ListItem)                       {}
func (w *ErgodoneWriter) WriteDescriptiveListItem(org.DescriptiveListItem) {}

func nonEmptyString(s string) bool { return s != "" }

func (w *ErgodoneWriter) WriteTable(t org.Table) {
	// Translate tables into keycodes
	if w.inKeymap {
		w.WriteString(",\n")
	}

	w.WriteString(strings.Join(yogofn.Map(func(row org.Row) string {
		return strings.Join(yogofn.Filter(nonEmptyString, yogofn.Map(func(column org.Column) string {
			content := w.WriteNodesAsString(column.Children...)
			keycode := w.TranslateKeycode(content)
			return keycode
		}, row.Columns).([]string)).([]string), ", ")
	}, t.Rows).([]string), ",\n"))
	w.WriteString("\n")

	w.inKeymap = true
}

func (w *ErgodoneWriter) WriteHorizontalRule(org.HorizontalRule) {}
func (w *ErgodoneWriter) WriteParagraph(org.Paragraph)           {}

func (w *ErgodoneWriter) WriteText(t org.Text) { w.WriteString(t.Content) }

func (w *ErgodoneWriter) WriteEmphasis(e org.Emphasis) {
	org.WriteNodes(w, e.Content...)
}

func (w *ErgodoneWriter) WriteLatexFragment(org.LatexFragment)   {}
func (w *ErgodoneWriter) WriteStatisticToken(org.StatisticToken) {}

func (w *ErgodoneWriter) WriteExplicitLineBreak(org.ExplicitLineBreak) {}

func (w *ErgodoneWriter) WriteLineBreak(org.LineBreak) { w.WriteString("\n") }

func (w *ErgodoneWriter) WriteRegularLink(org.RegularLink)               {}
func (w *ErgodoneWriter) WriteTimestamp(org.Timestamp)                   {}
func (w *ErgodoneWriter) WriteFootnoteLink(org.FootnoteLink)             {}
func (w *ErgodoneWriter) WriteFootnoteDefinition(org.FootnoteDefinition) {}
