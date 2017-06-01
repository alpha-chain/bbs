package cxo

import (
	"github.com/skycoin/cxo/skyobject"
	"github.com/skycoin/skycoin/src/cipher"
)

type MsgMode int

const (
	RootFilled MsgMode = iota
	FeedAdded
	FeedDeleted
)

type Msg struct {
	pk cipher.PubKey
	m  MsgMode
}

func (m *Msg) PubKey() cipher.PubKey { return m.pk }
func (m *Msg) Mode() MsgMode         { return m.m }

func (c *Container) GetUpdatesChan() chan *Msg               { return c.msgs }
func (c *Container) rootFilledCallBack(root *skyobject.Root) { go c.sendMsg(root.Pub(), RootFilled) }
func (c *Container) feedAddedCallBack(pk cipher.PubKey)      { go c.sendMsg(pk, FeedAdded) }
func (c *Container) feedDeleted(pk cipher.PubKey)            { go c.sendMsg(pk, FeedDeleted) }
func (c *Container) sendMsg(pk cipher.PubKey, m MsgMode)     { c.msgs <- &Msg{pk: pk, m: m} }
