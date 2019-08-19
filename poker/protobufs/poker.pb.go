// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobufs/poker.proto

package poker

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Card_Suit int32

const (
	Card_UNKNOWN_SUIT Card_Suit = 0
	Card_HEARTS       Card_Suit = 1
	Card_DIAMONDS     Card_Suit = 2
	Card_CLUBS        Card_Suit = 3
	Card_SPADES       Card_Suit = 4
)

var Card_Suit_name = map[int32]string{
	0: "UNKNOWN_SUIT",
	1: "HEARTS",
	2: "DIAMONDS",
	3: "CLUBS",
	4: "SPADES",
}

var Card_Suit_value = map[string]int32{
	"UNKNOWN_SUIT": 0,
	"HEARTS":       1,
	"DIAMONDS":     2,
	"CLUBS":        3,
	"SPADES":       4,
}

func (x Card_Suit) String() string {
	return proto.EnumName(Card_Suit_name, int32(x))
}

func (Card_Suit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{0, 0}
}

type Card_Rank int32

const (
	Card_UNKNOWN_RANK Card_Rank = 0
	Card_ACE          Card_Rank = 1
	Card_TWO          Card_Rank = 2
	Card_THREE        Card_Rank = 3
	Card_FOUR         Card_Rank = 4
	Card_FIVE         Card_Rank = 5
	Card_SIX          Card_Rank = 6
	Card_SEVEN        Card_Rank = 7
	Card_EIGHT        Card_Rank = 8
	Card_NINE         Card_Rank = 9
	Card_TEN          Card_Rank = 10
	Card_JACK         Card_Rank = 11
	Card_QUEEN        Card_Rank = 12
	Card_KING         Card_Rank = 13
)

var Card_Rank_name = map[int32]string{
	0:  "UNKNOWN_RANK",
	1:  "ACE",
	2:  "TWO",
	3:  "THREE",
	4:  "FOUR",
	5:  "FIVE",
	6:  "SIX",
	7:  "SEVEN",
	8:  "EIGHT",
	9:  "NINE",
	10: "TEN",
	11: "JACK",
	12: "QUEEN",
	13: "KING",
}

var Card_Rank_value = map[string]int32{
	"UNKNOWN_RANK": 0,
	"ACE":          1,
	"TWO":          2,
	"THREE":        3,
	"FOUR":         4,
	"FIVE":         5,
	"SIX":          6,
	"SEVEN":        7,
	"EIGHT":        8,
	"NINE":         9,
	"TEN":          10,
	"JACK":         11,
	"QUEEN":        12,
	"KING":         13,
}

func (x Card_Rank) String() string {
	return proto.EnumName(Card_Rank_name, int32(x))
}

func (Card_Rank) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{0, 1}
}

type Game_PlayerSlot int32

const (
	Game_slot_undefined Game_PlayerSlot = 0
	Game_slot_one       Game_PlayerSlot = 1
	Game_slot_two       Game_PlayerSlot = 2
	Game_slot_three     Game_PlayerSlot = 3
	Game_slot_four      Game_PlayerSlot = 4
	Game_slot_five      Game_PlayerSlot = 5
	Game_slot_six       Game_PlayerSlot = 6
	Game_slot_seven     Game_PlayerSlot = 7
	Game_slot_eight     Game_PlayerSlot = 8
)

var Game_PlayerSlot_name = map[int32]string{
	0: "slot_undefined",
	1: "slot_one",
	2: "slot_two",
	3: "slot_three",
	4: "slot_four",
	5: "slot_five",
	6: "slot_six",
	7: "slot_seven",
	8: "slot_eight",
}

var Game_PlayerSlot_value = map[string]int32{
	"slot_undefined": 0,
	"slot_one":       1,
	"slot_two":       2,
	"slot_three":     3,
	"slot_four":      4,
	"slot_five":      5,
	"slot_six":       6,
	"slot_seven":     7,
	"slot_eight":     8,
}

func (x Game_PlayerSlot) String() string {
	return proto.EnumName(Game_PlayerSlot_name, int32(x))
}

func (Game_PlayerSlot) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{6, 0}
}

type Card struct {
	Suit                 Card_Suit `protobuf:"varint,1,opt,name=suit,proto3,enum=poker.Card_Suit" json:"suit,omitempty"`
	Rank                 Card_Rank `protobuf:"varint,2,opt,name=rank,proto3,enum=poker.Card_Rank" json:"rank,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{0}
}

func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (m *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(m, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetSuit() Card_Suit {
	if m != nil {
		return m.Suit
	}
	return Card_UNKNOWN_SUIT
}

func (m *Card) GetRank() Card_Rank {
	if m != nil {
		return m.Rank
	}
	return Card_UNKNOWN_RANK
}

type Deck struct {
	Cards                []*Card  `protobuf:"bytes,1,rep,name=cards,proto3" json:"cards,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deck) Reset()         { *m = Deck{} }
func (m *Deck) String() string { return proto.CompactTextString(m) }
func (*Deck) ProtoMessage()    {}
func (*Deck) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{1}
}

func (m *Deck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deck.Unmarshal(m, b)
}
func (m *Deck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deck.Marshal(b, m, deterministic)
}
func (m *Deck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deck.Merge(m, src)
}
func (m *Deck) XXX_Size() int {
	return xxx_messageInfo_Deck.Size(m)
}
func (m *Deck) XXX_DiscardUnknown() {
	xxx_messageInfo_Deck.DiscardUnknown(m)
}

var xxx_messageInfo_Deck proto.InternalMessageInfo

func (m *Deck) GetCards() []*Card {
	if m != nil {
		return m.Cards
	}
	return nil
}

type Player struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Chips                int64    `protobuf:"varint,3,opt,name=chips,proto3" json:"chips,omitempty"`
	Hand                 *Hand    `protobuf:"bytes,4,opt,name=hand,proto3" json:"hand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{2}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetChips() int64 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func (m *Player) GetHand() *Hand {
	if m != nil {
		return m.Hand
	}
	return nil
}

type Players struct {
	Players              []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Players) Reset()         { *m = Players{} }
func (m *Players) String() string { return proto.CompactTextString(m) }
func (*Players) ProtoMessage()    {}
func (*Players) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{3}
}

func (m *Players) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Players.Unmarshal(m, b)
}
func (m *Players) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Players.Marshal(b, m, deterministic)
}
func (m *Players) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Players.Merge(m, src)
}
func (m *Players) XXX_Size() int {
	return xxx_messageInfo_Players.Size(m)
}
func (m *Players) XXX_DiscardUnknown() {
	xxx_messageInfo_Players.DiscardUnknown(m)
}

var xxx_messageInfo_Players proto.InternalMessageInfo

func (m *Players) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

type Hand struct {
	One                  *Card    `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two                  *Card    `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hand) Reset()         { *m = Hand{} }
func (m *Hand) String() string { return proto.CompactTextString(m) }
func (*Hand) ProtoMessage()    {}
func (*Hand) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{4}
}

func (m *Hand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hand.Unmarshal(m, b)
}
func (m *Hand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hand.Marshal(b, m, deterministic)
}
func (m *Hand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hand.Merge(m, src)
}
func (m *Hand) XXX_Size() int {
	return xxx_messageInfo_Hand.Size(m)
}
func (m *Hand) XXX_DiscardUnknown() {
	xxx_messageInfo_Hand.DiscardUnknown(m)
}

var xxx_messageInfo_Hand proto.InternalMessageInfo

func (m *Hand) GetOne() *Card {
	if m != nil {
		return m.One
	}
	return nil
}

func (m *Hand) GetTwo() *Card {
	if m != nil {
		return m.Two
	}
	return nil
}

type Flop struct {
	One                  *Card    `protobuf:"bytes,1,opt,name=one,proto3" json:"one,omitempty"`
	Two                  *Card    `protobuf:"bytes,2,opt,name=two,proto3" json:"two,omitempty"`
	Three                *Card    `protobuf:"bytes,3,opt,name=three,proto3" json:"three,omitempty"`
	Four                 *Card    `protobuf:"bytes,4,opt,name=four,proto3" json:"four,omitempty"`
	Five                 *Card    `protobuf:"bytes,5,opt,name=five,proto3" json:"five,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Flop) Reset()         { *m = Flop{} }
func (m *Flop) String() string { return proto.CompactTextString(m) }
func (*Flop) ProtoMessage()    {}
func (*Flop) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{5}
}

func (m *Flop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Flop.Unmarshal(m, b)
}
func (m *Flop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Flop.Marshal(b, m, deterministic)
}
func (m *Flop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Flop.Merge(m, src)
}
func (m *Flop) XXX_Size() int {
	return xxx_messageInfo_Flop.Size(m)
}
func (m *Flop) XXX_DiscardUnknown() {
	xxx_messageInfo_Flop.DiscardUnknown(m)
}

var xxx_messageInfo_Flop proto.InternalMessageInfo

func (m *Flop) GetOne() *Card {
	if m != nil {
		return m.One
	}
	return nil
}

func (m *Flop) GetTwo() *Card {
	if m != nil {
		return m.Two
	}
	return nil
}

func (m *Flop) GetThree() *Card {
	if m != nil {
		return m.Three
	}
	return nil
}

func (m *Flop) GetFour() *Card {
	if m != nil {
		return m.Four
	}
	return nil
}

func (m *Flop) GetFive() *Card {
	if m != nil {
		return m.Five
	}
	return nil
}

type Game struct {
	Players              *Players        `protobuf:"bytes,1,opt,name=players,proto3" json:"players,omitempty"`
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int64           `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	Small                Game_PlayerSlot `protobuf:"varint,10,opt,name=small,proto3,enum=poker.Game_PlayerSlot" json:"small,omitempty"`
	Big                  Game_PlayerSlot `protobuf:"varint,11,opt,name=big,proto3,enum=poker.Game_PlayerSlot" json:"big,omitempty"`
	Dealer               Game_PlayerSlot `protobuf:"varint,12,opt,name=Dealer,proto3,enum=poker.Game_PlayerSlot" json:"Dealer,omitempty"`
	Deck                 []*Card         `protobuf:"bytes,13,rep,name=deck,proto3" json:"deck,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_818c499f6358623d, []int{6}
}

func (m *Game) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Game.Unmarshal(m, b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Game.Marshal(b, m, deterministic)
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return xxx_messageInfo_Game.Size(m)
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetPlayers() *Players {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *Game) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Game) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Game) GetSmall() Game_PlayerSlot {
	if m != nil {
		return m.Small
	}
	return Game_slot_undefined
}

func (m *Game) GetBig() Game_PlayerSlot {
	if m != nil {
		return m.Big
	}
	return Game_slot_undefined
}

func (m *Game) GetDealer() Game_PlayerSlot {
	if m != nil {
		return m.Dealer
	}
	return Game_slot_undefined
}

func (m *Game) GetDeck() []*Card {
	if m != nil {
		return m.Deck
	}
	return nil
}

func init() {
	proto.RegisterEnum("poker.Card_Suit", Card_Suit_name, Card_Suit_value)
	proto.RegisterEnum("poker.Card_Rank", Card_Rank_name, Card_Rank_value)
	proto.RegisterEnum("poker.Game_PlayerSlot", Game_PlayerSlot_name, Game_PlayerSlot_value)
	proto.RegisterType((*Card)(nil), "poker.Card")
	proto.RegisterType((*Deck)(nil), "poker.Deck")
	proto.RegisterType((*Player)(nil), "poker.Player")
	proto.RegisterType((*Players)(nil), "poker.Players")
	proto.RegisterType((*Hand)(nil), "poker.Hand")
	proto.RegisterType((*Flop)(nil), "poker.Flop")
	proto.RegisterType((*Game)(nil), "poker.Game")
}

func init() { proto.RegisterFile("protobufs/poker.proto", fileDescriptor_818c499f6358623d) }

var fileDescriptor_818c499f6358623d = []byte{
	// 724 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0xe3, 0x78, 0x9d, 0x8f, 0xc9, 0x07, 0xab, 0x15, 0x20, 0xab, 0x12, 0xa2, 0x58, 0x20,
	0x82, 0x40, 0xa9, 0x08, 0x27, 0x48, 0x13, 0x37, 0x31, 0x01, 0xa7, 0xd8, 0x49, 0xcb, 0x5b, 0xe5,
	0xc6, 0xdb, 0xc6, 0x4a, 0x6a, 0x47, 0xb6, 0x93, 0xd2, 0x9b, 0x54, 0xdc, 0x80, 0x2b, 0x70, 0x33,
	0xde, 0xd0, 0xac, 0xdd, 0xd4, 0x49, 0xab, 0x16, 0x89, 0xb7, 0x9d, 0xfd, 0xff, 0x66, 0x3c, 0xfb,
	0x9f, 0x5d, 0x19, 0x9e, 0x2d, 0xc2, 0x20, 0x0e, 0x4e, 0x97, 0x67, 0xd1, 0xde, 0x22, 0x98, 0xf1,
	0xb0, 0x29, 0x62, 0xa6, 0x88, 0x40, 0xfb, 0x9d, 0x07, 0xd2, 0x71, 0x42, 0x97, 0xbd, 0x06, 0x12,
	0x2d, 0xbd, 0x58, 0x95, 0x76, 0xa5, 0x46, 0xbd, 0x45, 0x9b, 0x09, 0x8b, 0x52, 0xd3, 0x5e, 0x7a,
	0xb1, 0x25, 0x54, 0xa4, 0x42, 0xc7, 0x9f, 0xa9, 0xf9, 0xbb, 0x94, 0xe5, 0xf8, 0x33, 0x4b, 0xa8,
	0x9a, 0x01, 0x04, 0x73, 0x18, 0x85, 0xea, 0xd8, 0x1c, 0x98, 0xc3, 0x63, 0xf3, 0xc4, 0x1e, 0x1b,
	0x23, 0x9a, 0x63, 0x00, 0x85, 0xbe, 0xde, 0xb6, 0x46, 0x36, 0x95, 0x58, 0x15, 0x4a, 0x5d, 0xa3,
	0xfd, 0x75, 0x68, 0x76, 0x6d, 0x9a, 0x67, 0x65, 0x50, 0x3a, 0x5f, 0xc6, 0xfb, 0x36, 0x95, 0x11,
	0xb2, 0x0f, 0xdb, 0x5d, 0xdd, 0xa6, 0x44, 0xfb, 0x29, 0x01, 0xc1, 0xca, 0xd9, 0x5a, 0x56, 0xdb,
	0x1c, 0xd0, 0x1c, 0x2b, 0x82, 0xdc, 0xee, 0xe8, 0x54, 0xc2, 0xc5, 0xe8, 0x78, 0x98, 0xd4, 0x18,
	0xf5, 0x2d, 0x5d, 0xa7, 0x32, 0x2b, 0x01, 0x39, 0x18, 0x8e, 0x2d, 0x4a, 0xc4, 0xca, 0x38, 0xd2,
	0xa9, 0x82, 0x9c, 0x6d, 0x7c, 0xa7, 0x05, 0xe4, 0x6c, 0xfd, 0x48, 0x37, 0x69, 0x11, 0x97, 0xba,
	0xd1, 0xeb, 0x8f, 0x68, 0x09, 0x41, 0xd3, 0x30, 0x75, 0x5a, 0x16, 0x05, 0x75, 0x93, 0x02, 0x6e,
	0x7d, 0x6e, 0x77, 0x06, 0xb4, 0x82, 0xdc, 0xb7, 0xb1, 0xae, 0x9b, 0xb4, 0x8a, 0x9b, 0x03, 0xc3,
	0xec, 0xd1, 0x9a, 0xf6, 0x0e, 0x48, 0x97, 0x4f, 0x66, 0xec, 0x15, 0x28, 0x13, 0x27, 0x74, 0x23,
	0x55, 0xda, 0x95, 0x1b, 0x95, 0x56, 0x25, 0x63, 0x8b, 0x95, 0x28, 0xda, 0x04, 0x0a, 0x87, 0x73,
	0xe7, 0x8a, 0x87, 0xac, 0x0e, 0x79, 0xcf, 0x15, 0x36, 0xcb, 0x56, 0xde, 0x73, 0x19, 0x03, 0xe2,
	0x3b, 0x17, 0x5c, 0x58, 0x5a, 0xb6, 0xc4, 0x9a, 0x3d, 0x05, 0x65, 0x32, 0xf5, 0x16, 0x91, 0x2a,
	0x0b, 0x2c, 0x09, 0xd8, 0x4b, 0x20, 0x53, 0xc7, 0x77, 0x55, 0xb2, 0x2b, 0x65, 0xbe, 0xd2, 0x77,
	0x7c, 0xd7, 0x12, 0x82, 0xd6, 0x82, 0x62, 0xf2, 0x91, 0x88, 0xbd, 0x85, 0xe2, 0x22, 0x59, 0xa6,
	0x4d, 0xd5, 0x52, 0x3c, 0x01, 0xac, 0x1b, 0x55, 0xeb, 0x02, 0xc1, 0x0a, 0xec, 0x05, 0xc8, 0x81,
	0xcf, 0x45, 0x5f, 0x5b, 0x27, 0xc0, 0x7d, 0x94, 0xe3, 0xcb, 0x40, 0x34, 0xb9, 0x2d, 0xc7, 0x97,
	0x81, 0xf6, 0x4b, 0x02, 0x72, 0x30, 0x0f, 0x16, 0xff, 0x57, 0x06, 0x8d, 0x8c, 0xa7, 0x21, 0xe7,
	0xe2, 0xdc, 0xdb, 0x46, 0x0a, 0x05, 0x4d, 0x38, 0x0b, 0x96, 0xe1, 0x96, 0x09, 0x82, 0x10, 0x82,
	0x00, 0xbc, 0x15, 0x57, 0x95, 0xfb, 0x00, 0x6f, 0xc5, 0xb5, 0x6b, 0x19, 0x48, 0x0f, 0x5d, 0x6e,
	0x64, 0x3d, 0x42, 0xb8, 0xbe, 0xe1, 0x51, 0xb4, 0x36, 0xe9, 0xde, 0x19, 0x25, 0x73, 0x2c, 0xaf,
	0xe7, 0xf8, 0x01, 0x94, 0xe8, 0xc2, 0x99, 0xcf, 0x55, 0x10, 0x6f, 0xe3, 0x79, 0x5a, 0x0b, 0xbf,
	0x94, 0x16, 0xb4, 0xe7, 0x41, 0x6c, 0x25, 0x10, 0x6b, 0x80, 0x7c, 0xea, 0x9d, 0xab, 0x95, 0x07,
	0x59, 0x44, 0x58, 0x13, 0x0a, 0x5d, 0xee, 0xcc, 0x79, 0xa8, 0x56, 0x1f, 0x84, 0x53, 0x0a, 0xcf,
	0xef, 0xf2, 0xc9, 0x4c, 0xad, 0xdd, 0xbd, 0x8b, 0x42, 0xd0, 0xae, 0x25, 0x80, 0xdb, 0x3c, 0xc6,
	0xa0, 0x1e, 0xcd, 0x83, 0xf8, 0x64, 0xe9, 0xbb, 0xfc, 0xcc, 0xf3, 0xb9, 0x4b, 0x73, 0xf8, 0x34,
	0xc5, 0x5e, 0xe0, 0xf3, 0xe4, 0xa1, 0x8a, 0x28, 0xbe, 0x0c, 0x68, 0x9e, 0xd5, 0x01, 0x92, 0x08,
	0xc7, 0x41, 0x65, 0x56, 0x83, 0xb2, 0x88, 0xd1, 0x7c, 0x4a, 0x6e, 0x43, 0x6f, 0xc5, 0xa9, 0xb2,
	0xce, 0x8d, 0xbc, 0x1f, 0xb4, 0xb0, 0xce, 0x8d, 0xf8, 0x8a, 0xfb, 0xb4, 0xb8, 0x8e, 0xb9, 0x77,
	0x3e, 0x8d, 0x69, 0xa9, 0xf5, 0x27, 0x0f, 0xca, 0x21, 0xf6, 0xcb, 0x9a, 0x50, 0xed, 0x84, 0xdc,
	0x89, 0x79, 0xfa, 0x6a, 0x36, 0xaf, 0xef, 0xce, 0x66, 0xa8, 0xe5, 0xd8, 0x47, 0xa8, 0x65, 0xf9,
	0x88, 0x6d, 0xcd, 0x72, 0x67, 0x2b, 0xd6, 0x72, 0xec, 0x3d, 0x94, 0x7b, 0x3c, 0xfe, 0xe7, 0xfa,
	0x4f, 0xd6, 0xf0, 0xfe, 0x95, 0x89, 0x17, 0xe0, 0xb1, 0x94, 0x06, 0x40, 0xd2, 0x92, 0xb8, 0x6c,
	0x95, 0xcc, 0xd8, 0x76, 0xb2, 0x81, 0x96, 0x63, 0x6f, 0xa0, 0xd8, 0xe3, 0xf1, 0xa3, 0xd8, 0x1e,
	0xd4, 0x53, 0xec, 0xe6, 0x90, 0x1b, 0xf4, 0x7d, 0x27, 0xac, 0xa5, 0x09, 0x69, 0xcb, 0x0f, 0x54,
	0x3f, 0x2d, 0x88, 0xff, 0xc2, 0xa7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x97, 0xfd, 0xed,
	0x30, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PokerClient is the client API for Poker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PokerClient interface {
	CreatePlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Player, error)
	CreatePlayers(ctx context.Context, in *Players, opts ...grpc.CallOption) (*Players, error)
	GetPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Player, error)
	GetPlayerByName(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Player, error)
	CreateGame(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Game, error)
	GetGame(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Game, error)
	GetGamePlayers(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Players, error)
	GetGameByName(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Game, error)
}

type pokerClient struct {
	cc *grpc.ClientConn
}

func NewPokerClient(cc *grpc.ClientConn) PokerClient {
	return &pokerClient{cc}
}

func (c *pokerClient) CreatePlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/poker.Poker/CreatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) CreatePlayers(ctx context.Context, in *Players, opts ...grpc.CallOption) (*Players, error) {
	out := new(Players)
	err := c.cc.Invoke(ctx, "/poker.Poker/CreatePlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) GetPlayer(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/poker.Poker/GetPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) GetPlayerByName(ctx context.Context, in *Player, opts ...grpc.CallOption) (*Player, error) {
	out := new(Player)
	err := c.cc.Invoke(ctx, "/poker.Poker/GetPlayerByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) CreateGame(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Game, error) {
	out := new(Game)
	err := c.cc.Invoke(ctx, "/poker.Poker/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) GetGame(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Game, error) {
	out := new(Game)
	err := c.cc.Invoke(ctx, "/poker.Poker/GetGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) GetGamePlayers(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Players, error) {
	out := new(Players)
	err := c.cc.Invoke(ctx, "/poker.Poker/GetGamePlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pokerClient) GetGameByName(ctx context.Context, in *Game, opts ...grpc.CallOption) (*Game, error) {
	out := new(Game)
	err := c.cc.Invoke(ctx, "/poker.Poker/GetGameByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PokerServer is the server API for Poker service.
type PokerServer interface {
	CreatePlayer(context.Context, *Player) (*Player, error)
	CreatePlayers(context.Context, *Players) (*Players, error)
	GetPlayer(context.Context, *Player) (*Player, error)
	GetPlayerByName(context.Context, *Player) (*Player, error)
	CreateGame(context.Context, *Game) (*Game, error)
	GetGame(context.Context, *Game) (*Game, error)
	GetGamePlayers(context.Context, *Game) (*Players, error)
	GetGameByName(context.Context, *Game) (*Game, error)
}

// UnimplementedPokerServer can be embedded to have forward compatible implementations.
type UnimplementedPokerServer struct {
}

func (*UnimplementedPokerServer) CreatePlayer(ctx context.Context, req *Player) (*Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayer not implemented")
}
func (*UnimplementedPokerServer) CreatePlayers(ctx context.Context, req *Players) (*Players, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayers not implemented")
}
func (*UnimplementedPokerServer) GetPlayer(ctx context.Context, req *Player) (*Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayer not implemented")
}
func (*UnimplementedPokerServer) GetPlayerByName(ctx context.Context, req *Player) (*Player, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerByName not implemented")
}
func (*UnimplementedPokerServer) CreateGame(ctx context.Context, req *Game) (*Game, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (*UnimplementedPokerServer) GetGame(ctx context.Context, req *Game) (*Game, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGame not implemented")
}
func (*UnimplementedPokerServer) GetGamePlayers(ctx context.Context, req *Game) (*Players, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGamePlayers not implemented")
}
func (*UnimplementedPokerServer) GetGameByName(ctx context.Context, req *Game) (*Game, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGameByName not implemented")
}

func RegisterPokerServer(s *grpc.Server, srv PokerServer) {
	s.RegisterService(&_Poker_serviceDesc, srv)
}

func _Poker_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/CreatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).CreatePlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_CreatePlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Players)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).CreatePlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/CreatePlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).CreatePlayers(ctx, req.(*Players))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_GetPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).GetPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/GetPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).GetPlayer(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_GetPlayerByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Player)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).GetPlayerByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/GetPlayerByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).GetPlayerByName(ctx, req.(*Player))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Game)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).CreateGame(ctx, req.(*Game))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_GetGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Game)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).GetGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/GetGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).GetGame(ctx, req.(*Game))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_GetGamePlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Game)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).GetGamePlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/GetGamePlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).GetGamePlayers(ctx, req.(*Game))
	}
	return interceptor(ctx, in, info, handler)
}

func _Poker_GetGameByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Game)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PokerServer).GetGameByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poker.Poker/GetGameByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PokerServer).GetGameByName(ctx, req.(*Game))
	}
	return interceptor(ctx, in, info, handler)
}

var _Poker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poker.Poker",
	HandlerType: (*PokerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlayer",
			Handler:    _Poker_CreatePlayer_Handler,
		},
		{
			MethodName: "CreatePlayers",
			Handler:    _Poker_CreatePlayers_Handler,
		},
		{
			MethodName: "GetPlayer",
			Handler:    _Poker_GetPlayer_Handler,
		},
		{
			MethodName: "GetPlayerByName",
			Handler:    _Poker_GetPlayerByName_Handler,
		},
		{
			MethodName: "CreateGame",
			Handler:    _Poker_CreateGame_Handler,
		},
		{
			MethodName: "GetGame",
			Handler:    _Poker_GetGame_Handler,
		},
		{
			MethodName: "GetGamePlayers",
			Handler:    _Poker_GetGamePlayers_Handler,
		},
		{
			MethodName: "GetGameByName",
			Handler:    _Poker_GetGameByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/poker.proto",
}
