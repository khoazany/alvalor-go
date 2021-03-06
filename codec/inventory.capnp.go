// Code generated by capnpc-go. DO NOT EDIT.

package codec

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Inventory struct{ capnp.Struct }

// Inventory_TypeID is the unique identifier for the type Inventory.
const Inventory_TypeID = 0x894854b19b0bdf88

func NewInventory(s *capnp.Segment) (Inventory, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Inventory{st}, err
}

func NewRootInventory(s *capnp.Segment) (Inventory, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return Inventory{st}, err
}

func ReadRootInventory(msg *capnp.Message) (Inventory, error) {
	root, err := msg.RootPtr()
	return Inventory{root.Struct()}, err
}

func (s Inventory) String() string {
	str, _ := text.Marshal(0x894854b19b0bdf88, s.Struct)
	return str
}

func (s Inventory) Hashes() (capnp.DataList, error) {
	p, err := s.Struct.Ptr(0)
	return capnp.DataList{List: p.List()}, err
}

func (s Inventory) HasHashes() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Inventory) SetHashes(v capnp.DataList) error {
	return s.Struct.SetPtr(0, v.List.ToPtr())
}

// NewHashes sets the hashes field to a newly
// allocated capnp.DataList, preferring placement in s's segment.
func (s Inventory) NewHashes(n int32) (capnp.DataList, error) {
	l, err := capnp.NewDataList(s.Struct.Segment(), n)
	if err != nil {
		return capnp.DataList{}, err
	}
	err = s.Struct.SetPtr(0, l.List.ToPtr())
	return l, err
}

// Inventory_List is a list of Inventory.
type Inventory_List struct{ capnp.List }

// NewInventory creates a new list of Inventory.
func NewInventory_List(s *capnp.Segment, sz int32) (Inventory_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return Inventory_List{l}, err
}

func (s Inventory_List) At(i int) Inventory { return Inventory{s.List.Struct(i)} }

func (s Inventory_List) Set(i int, v Inventory) error { return s.List.SetStruct(i, v.Struct) }

func (s Inventory_List) String() string {
	str, _ := text.MarshalList(0x894854b19b0bdf88, s.List)
	return str
}

// Inventory_Promise is a wrapper for a Inventory promised by a client call.
type Inventory_Promise struct{ *capnp.Pipeline }

func (p Inventory_Promise) Struct() (Inventory, error) {
	s, err := p.Pipeline.Struct()
	return Inventory{s}, err
}

const schema_9fd0f7eb12926b5d = "x\xda2pct`2d\xf5\xe7d`\x08\xcca" +
	"e\xfb\xdfq\x9f{\xf6\xc6\x10\x8fN\x06A\x01\xc6\xff" +
	"\xb1\xd9\x93\x84^\x7f\xbf0\x9f\x81\x95\x91\x9d\x81A\xb8" +
	"\x97\xe5\x92\xf0L\x16\x10k*\x8b=\xc3\x7f\x06\xf9\xff" +
	"\x99ye\xa9y%\xf9EL\x95z\xc9\x89\x05y\x05" +
	"V\x9eP\x01\xc6\xca\x00F\xc6@\x16f\x16\x06\x06\x16" +
	"F\x06\x06A^+\x06\x86@\x0ef\xc6@\x15&F" +
	"\xfb\x8c\xc4\xe2\x8c\xd4bF>\x06\xc6\x00fFF^" +
	"\x06&\x10\x13\x10\x00\x00\xff\xffi\xf0!j"

func init() {
	schemas.Register(schema_9fd0f7eb12926b5d,
		0x894854b19b0bdf88)
}
