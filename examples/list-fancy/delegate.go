package main

import (
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/list"
)

func newItemDelegate(keys *delegateKeyMap, styles *styles) list.DefaultDelegate {
	_ = "STUB: not implemented"
	return *new(list.DefaultDelegate)
}

type delegateKeyMap struct {
	choose key.Binding
	remove key.Binding
}

func (d delegateKeyMap) ShortHelp() []key.Binding { _ = "STUB: not implemented"; return nil }

func (d delegateKeyMap) FullHelp() [][]key.Binding { _ = "STUB: not implemented"; return nil }

func newDelegateKeyMap() *delegateKeyMap { _ = "STUB: not implemented"; return nil }
