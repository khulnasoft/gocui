package gocui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextArea(t *testing.T) {
	tests := []struct {
		actions         func(*TextArea)
		expectedContent string
		expectedCursor  int
	}{
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('c')
			},
			expectedContent: "abc",
			expectedCursor:  3,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('\n')
				textarea.TypeRune('c')
			},
			expectedContent: "a\nc",
			expectedCursor:  3,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("abcd")
			},
			expectedContent: "abcd",
			expectedCursor:  4,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("a字cd")
			},
			expectedContent: "a字cd",
			expectedCursor:  4,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.BackSpaceChar()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.BackSpaceChar()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.BackSpaceChar()
			},
			expectedContent: "a",
			expectedCursor:  1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.DeleteChar()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.DeleteChar()
			},
			expectedContent: "a",
			expectedCursor:  1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.MoveCursorLeft()
				textarea.DeleteChar()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('c')
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.DeleteChar()
			},
			expectedContent: "ac",
			expectedCursor:  1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.MoveCursorLeft()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.MoveCursorLeft()
			},
			expectedContent: "a",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.MoveCursorLeft()
			},
			expectedContent: "ab",
			expectedCursor:  1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.MoveCursorRight()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.MoveCursorRight()
			},
			expectedContent: "a",
			expectedCursor:  1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.MoveCursorLeft()
				textarea.MoveCursorRight()
			},
			expectedContent: "ab",
			expectedCursor:  2,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('漢')
				textarea.TypeRune('字')
				textarea.MoveCursorLeft()
			},
			expectedContent: "漢字",
			expectedCursor:  1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.ToggleOverwrite()
				textarea.TypeRune('a')
				textarea.TypeRune('b')
			},
			expectedContent: "ab",
			expectedCursor:  2,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('c')
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.ToggleOverwrite()
				textarea.TypeRune('d')
			},
			expectedContent: "adc",
			expectedCursor:  2,
		},
		{
			actions: func(textarea *TextArea) {
				// overwrite mode acts same as normal mode when cursor is at the end
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('c')
				textarea.ToggleOverwrite()
				textarea.TypeRune('d')
			},
			expectedContent: "abcd",
			expectedCursor:  4,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.DeleteToStartOfLine()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.DeleteToStartOfLine()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.DeleteToStartOfLine()
			},
			expectedContent: "ab",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('\n')
				textarea.DeleteToStartOfLine()
			},
			expectedContent: "ab",
			expectedCursor:  2,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('\n')
				textarea.TypeRune('c')
				textarea.TypeRune('d')
				textarea.DeleteToStartOfLine()
			},
			expectedContent: "ab\n",
			expectedCursor:  3,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.GoToStartOfLine()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.MoveCursorLeft()
				textarea.GoToStartOfLine()
			},
			expectedContent: "a",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('\n')
				textarea.TypeRune('c')
				textarea.TypeRune('d')
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.GoToStartOfLine()
			},
			expectedContent: "ab\ncd",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('\n')
				textarea.TypeRune('c')
				textarea.TypeRune('d')
				textarea.GoToStartOfLine()
			},
			expectedContent: "ab\ncd",
			expectedCursor:  3,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('\n')
				textarea.TypeRune('c')
				textarea.TypeRune('d')
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.GoToStartOfLine()
			},
			expectedContent: "ab\ncd",
			expectedCursor:  3,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.GoToEndOfLine()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('a')
				textarea.TypeRune('b')
				textarea.TypeRune('\n')
				textarea.TypeRune('c')
				textarea.TypeRune('d')
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.GoToEndOfLine()
			},
			expectedContent: "ab\ncd",
			expectedCursor:  5,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.SetCursor2D(10, 10)
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.SetCursor2D(-1, -1)
			},
			expectedContent: "",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("ab\ncd")
				textarea.SetCursor2D(0, 0)
			},
			expectedContent: "ab\ncd",
			expectedCursor:  0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("ab\ncd")
				textarea.SetCursor2D(2, 0)
			},
			expectedContent: "ab\ncd",
			expectedCursor:  2,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("ab\ncd\nef")
				textarea.SetCursor2D(2, 1)
			},
			expectedContent: "ab\ncd\nef",
			expectedCursor:  5,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("abcd\n\nijkl")
				textarea.MoveCursorUp()
			},
			expectedContent: "abcd\n\nijkl",
			expectedCursor:  5,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("abcdef\n老老老")
				textarea.MoveCursorLeft()
				textarea.MoveCursorUp()
			},
			expectedContent: "abcdef\n老老老",
			expectedCursor:  4,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("abcdef\n老老老")
				textarea.MoveCursorUp()
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.MoveCursorLeft()
				textarea.MoveCursorDown()
			},
			expectedContent: "abcdef\n老老老",
			expectedCursor:  9,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("abcd\nef")
				textarea.MoveCursorUp()
				textarea.GoToEndOfLine()
				textarea.MoveCursorDown()
			},
			expectedContent: "abcd\nef",
			expectedCursor:  7,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("abcd")
				textarea.MoveCursorUp()
			},
			expectedContent: "abcd",
			expectedCursor:  4,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString(`abcdefg`)
				textarea.Clear()
			},
			expectedContent: "",
			expectedCursor:  0,
		},
	}

	for _, test := range tests {
		textarea := &TextArea{}
		test.actions(textarea)
		assert.EqualValues(t, test.expectedContent, textarea.GetContent())
		assert.EqualValues(t, test.expectedCursor, textarea.cursor)
	}
}

func TestGetCursorXY(t *testing.T) {
	tests := []struct {
		actions   func(*TextArea)
		expectedX int
		expectedY int
	}{
		{
			actions: func(textarea *TextArea) {
				// do nothing
			},
			expectedX: 0,
			expectedY: 0,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("ab\ncd")
			},
			expectedX: 2,
			expectedY: 1,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeString("ab\n\n")
			},
			expectedX: 0,
			expectedY: 2,
		},
		{
			actions: func(textarea *TextArea) {
				textarea.TypeRune('漢')
				textarea.TypeRune('字')
			},
			expectedX: 4,
			expectedY: 0,
		},
	}

	for _, test := range tests {
		textarea := &TextArea{}
		test.actions(textarea)
		x, y := textarea.GetCursorXY()
		assert.EqualValues(t, test.expectedX, x)
		assert.EqualValues(t, test.expectedY, y)
	}
}