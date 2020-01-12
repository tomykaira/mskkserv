// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mozc/commands/candidates.proto

/*
Package commands is a generated protocol buffer package.

It is generated from these files:
	mozc/commands/candidates.proto
	mozc/commands/commands.proto
	mozc/commands/engine_builder.proto
	mozc/commands/renderer_command.proto

It has these top-level messages:
	Annotation
	Information
	InformationList
	Footer
	CandidateWord
	CandidateList
	Candidates
	KeyEvent
	GenericStorageEntry
	SessionCommand
	Context
	Capability
	Request
	ApplicationInfo
	Input
	Result
	Preedit
	Status
	DeletionRange
	Output
	Command
	CommandList
	EngineReloadRequest
	EngineReloadResponse
	RendererCommand
*/
package commands

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Category describes the attribute of the words.
type Category int32

const (
	Category_CONVERSION      Category = 0
	Category_PREDICTION      Category = 1
	Category_SUGGESTION      Category = 2
	Category_TRANSLITERATION Category = 3
	Category_USAGE           Category = 4
)

var Category_name = map[int32]string{
	0: "CONVERSION",
	1: "PREDICTION",
	2: "SUGGESTION",
	3: "TRANSLITERATION",
	4: "USAGE",
}
var Category_value = map[string]int32{
	"CONVERSION":      0,
	"PREDICTION":      1,
	"SUGGESTION":      2,
	"TRANSLITERATION": 3,
	"USAGE":           4,
}

func (x Category) Enum() *Category {
	p := new(Category)
	*p = x
	return p
}
func (x Category) String() string {
	return proto.EnumName(Category_name, int32(x))
}
func (x *Category) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Category_value, data, "Category")
	if err != nil {
		return err
	}
	*x = Category(value)
	return nil
}
func (Category) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// DisplayType is a hint to UI renderers describing how the words are
// displayed.
type DisplayType int32

const (
	DisplayType_MAIN    DisplayType = 0
	DisplayType_CASCADE DisplayType = 1
)

var DisplayType_name = map[int32]string{
	0: "MAIN",
	1: "CASCADE",
}
var DisplayType_value = map[string]int32{
	"MAIN":    0,
	"CASCADE": 1,
}

func (x DisplayType) Enum() *DisplayType {
	p := new(DisplayType)
	*p = x
	return p
}
func (x DisplayType) String() string {
	return proto.EnumName(DisplayType_name, int32(x))
}
func (x *DisplayType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DisplayType_value, data, "DisplayType")
	if err != nil {
		return err
	}
	*x = DisplayType(value)
	return nil
}
func (DisplayType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The direction of candidates in the window.  This is just a
// suggestion from the server and client does not have to follow.
type Candidates_Direction int32

const (
	Candidates_VERTICAL   Candidates_Direction = 0
	Candidates_HORIZONTAL Candidates_Direction = 1
)

var Candidates_Direction_name = map[int32]string{
	0: "VERTICAL",
	1: "HORIZONTAL",
}
var Candidates_Direction_value = map[string]int32{
	"VERTICAL":   0,
	"HORIZONTAL": 1,
}

func (x Candidates_Direction) Enum() *Candidates_Direction {
	p := new(Candidates_Direction)
	*p = x
	return p
}
func (x Candidates_Direction) String() string {
	return proto.EnumName(Candidates_Direction_name, int32(x))
}
func (x *Candidates_Direction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Candidates_Direction_value, data, "Candidates_Direction")
	if err != nil {
		return err
	}
	*x = Candidates_Direction(value)
	return nil
}
func (Candidates_Direction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

// Annotation against a candidate.
type Annotation struct {
	// Annotation prepended to the value.
	Prefix *string `protobuf:"bytes,1,opt,name=prefix" json:"prefix,omitempty"`
	// Annotation appended to the value.
	Suffix *string `protobuf:"bytes,2,opt,name=suffix" json:"suffix,omitempty"`
	// Type of the candidate such as [HALF][KATAKANA], [GREEK],
	// [Black square], etc...
	Description *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// Shortcut key to select this candidate.
	Shortcut *string `protobuf:"bytes,4,opt,name=shortcut" json:"shortcut,omitempty"`
	// Set to true if this candidate can be deleted from history.
	Deletable        *bool  `protobuf:"varint,5,opt,name=deletable,def=0" json:"deletable,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Annotation) Reset()                    { *m = Annotation{} }
func (m *Annotation) String() string            { return proto.CompactTextString(m) }
func (*Annotation) ProtoMessage()               {}
func (*Annotation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

const Default_Annotation_Deletable bool = false

func (m *Annotation) GetPrefix() string {
	if m != nil && m.Prefix != nil {
		return *m.Prefix
	}
	return ""
}

func (m *Annotation) GetSuffix() string {
	if m != nil && m.Suffix != nil {
		return *m.Suffix
	}
	return ""
}

func (m *Annotation) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Annotation) GetShortcut() string {
	if m != nil && m.Shortcut != nil {
		return *m.Shortcut
	}
	return ""
}

func (m *Annotation) GetDeletable() bool {
	if m != nil && m.Deletable != nil {
		return *m.Deletable
	}
	return Default_Annotation_Deletable
}

// Additional information to a candidate word.  This message is
// used for describing a word usage for instance.
type Information struct {
	// Unique number specifying the information.
	Id *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// Title string of the information.  For usage, this value is
	// probably equal to Candidate::value or its canonicalized value.
	Title *string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// The content of the information.  For usage, this value actually
	// describes how to use the word.
	Description *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// The IDs of candidates which connect with the information.
	CandidateId      []int32 `protobuf:"varint,4,rep,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Information) Reset()                    { *m = Information{} }
func (m *Information) String() string            { return proto.CompactTextString(m) }
func (*Information) ProtoMessage()               {}
func (*Information) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Information) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Information) GetTitle() string {
	if m != nil && m.Title != nil {
		return *m.Title
	}
	return ""
}

func (m *Information) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Information) GetCandidateId() []int32 {
	if m != nil {
		return m.CandidateId
	}
	return nil
}

type InformationList struct {
	FocusedIndex *uint32        `protobuf:"varint,1,opt,name=focused_index,json=focusedIndex" json:"focused_index,omitempty"`
	Information  []*Information `protobuf:"bytes,2,rep,name=information" json:"information,omitempty"`
	// Category of the infolist.
	Category *Category `protobuf:"varint,3,opt,name=category,enum=commands.Category,def=0" json:"category,omitempty"`
	// Information to be used for rendering.
	DisplayType *DisplayType `protobuf:"varint,4,opt,name=display_type,json=displayType,enum=commands.DisplayType,def=1" json:"display_type,omitempty"`
	// How long rendere needs to wait before the infolist is displayed.
	// the default setting is 500 msec.
	Delay            *uint32 `protobuf:"varint,5,opt,name=delay,def=500" json:"delay,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *InformationList) Reset()                    { *m = InformationList{} }
func (m *InformationList) String() string            { return proto.CompactTextString(m) }
func (*InformationList) ProtoMessage()               {}
func (*InformationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

const Default_InformationList_Category Category = Category_CONVERSION
const Default_InformationList_DisplayType DisplayType = DisplayType_CASCADE
const Default_InformationList_Delay uint32 = 500

func (m *InformationList) GetFocusedIndex() uint32 {
	if m != nil && m.FocusedIndex != nil {
		return *m.FocusedIndex
	}
	return 0
}

func (m *InformationList) GetInformation() []*Information {
	if m != nil {
		return m.Information
	}
	return nil
}

func (m *InformationList) GetCategory() Category {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return Default_InformationList_Category
}

func (m *InformationList) GetDisplayType() DisplayType {
	if m != nil && m.DisplayType != nil {
		return *m.DisplayType
	}
	return Default_InformationList_DisplayType
}

func (m *InformationList) GetDelay() uint32 {
	if m != nil && m.Delay != nil {
		return *m.Delay
	}
	return Default_InformationList_Delay
}

// Message representing the footer part of the candidate window.
type Footer struct {
	// Message shown like a status bar.
	Label *string `protobuf:"bytes,1,opt,name=label" json:"label,omitempty"`
	// Whether index (e.g. 10/120) is visible or not.
	IndexVisible *bool `protobuf:"varint,2,opt,name=index_visible,json=indexVisible,def=0" json:"index_visible,omitempty"`
	// Whether the logo image is visible or not.
	LogoVisible *bool `protobuf:"varint,3,opt,name=logo_visible,json=logoVisible,def=0" json:"logo_visible,omitempty"`
	// Message modestly shown.  It is used for displaying the version on
	// dev-channel now.
	SubLabel         *string `protobuf:"bytes,4,opt,name=sub_label,json=subLabel" json:"sub_label,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Footer) Reset()                    { *m = Footer{} }
func (m *Footer) String() string            { return proto.CompactTextString(m) }
func (*Footer) ProtoMessage()               {}
func (*Footer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

const Default_Footer_IndexVisible bool = false
const Default_Footer_LogoVisible bool = false

func (m *Footer) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *Footer) GetIndexVisible() bool {
	if m != nil && m.IndexVisible != nil {
		return *m.IndexVisible
	}
	return Default_Footer_IndexVisible
}

func (m *Footer) GetLogoVisible() bool {
	if m != nil && m.LogoVisible != nil {
		return *m.LogoVisible
	}
	return Default_Footer_LogoVisible
}

func (m *Footer) GetSubLabel() string {
	if m != nil && m.SubLabel != nil {
		return *m.SubLabel
	}
	return ""
}

type CandidateWord struct {
	// Unique number specifing the candidate.  This may be a negative value.
	Id *int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The first index should be zero and index numbers should increase by one.
	Index *uint32 `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	// Reading of the value.  The value is only used when the key is
	// different from the input composition (e.g. suggestion/prediction).
	Key *string `protobuf:"bytes,3,opt,name=key" json:"key,omitempty"`
	// Converted value.  (e.g. Kanji value).
	Value            *string     `protobuf:"bytes,4,opt,name=value" json:"value,omitempty"`
	Annotation       *Annotation `protobuf:"bytes,5,opt,name=annotation" json:"annotation,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *CandidateWord) Reset()                    { *m = CandidateWord{} }
func (m *CandidateWord) String() string            { return proto.CompactTextString(m) }
func (*CandidateWord) ProtoMessage()               {}
func (*CandidateWord) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CandidateWord) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *CandidateWord) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *CandidateWord) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *CandidateWord) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *CandidateWord) GetAnnotation() *Annotation {
	if m != nil {
		return m.Annotation
	}
	return nil
}

type CandidateList struct {
	// This value represents the focused position of the next
	// |candidates|.  If the |candidates| is a part of the whole
	// candidate words (as a result of paging), this value indicates the
	// position from the beginning of that part.  (ex. where
	// |candidates| contatins 10th to 18th candidates, focused_index=0
	// means the 10th candidate, but not 1st candidate.
	//
	// The existense of |focused_index| does not represents whether this
	// candidate list is a 'suggestion' or not.  |category| represents
	// it.
	FocusedIndex *uint32          `protobuf:"varint,1,opt,name=focused_index,json=focusedIndex" json:"focused_index,omitempty"`
	Candidates   []*CandidateWord `protobuf:"bytes,2,rep,name=candidates" json:"candidates,omitempty"`
	// Category of the candidates.
	Category         *Category `protobuf:"varint,3,opt,name=category,enum=commands.Category,def=0" json:"category,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *CandidateList) Reset()                    { *m = CandidateList{} }
func (m *CandidateList) String() string            { return proto.CompactTextString(m) }
func (*CandidateList) ProtoMessage()               {}
func (*CandidateList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

const Default_CandidateList_Category Category = Category_CONVERSION

func (m *CandidateList) GetFocusedIndex() uint32 {
	if m != nil && m.FocusedIndex != nil {
		return *m.FocusedIndex
	}
	return 0
}

func (m *CandidateList) GetCandidates() []*CandidateWord {
	if m != nil {
		return m.Candidates
	}
	return nil
}

func (m *CandidateList) GetCategory() Category {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return Default_CandidateList_Category
}

// TODO(komatsu) rename it to CandidateWindow.
type Candidates struct {
	// TODO(komatsu): Use CandidateList.
	// When has_focused_index() is true, this message contains predicted and
	// normally converted candidates. Otherwise, when the field is not set,
	// this message contains a 'suggestion'.
	FocusedIndex *uint32 `protobuf:"varint,1,opt,name=focused_index,json=focusedIndex" json:"focused_index,omitempty"`
	// The size of the total candidates in this candidate list.  The
	// value does not include the size of subcandidate lists.  Note, the
	// next repeated-Candidate=3 may not contain all candidates.
	// all_candidates contains the values of subcandidate lists.
	Size      *uint32                 `protobuf:"varint,2,req,name=size" json:"size,omitempty"`
	Candidate []*Candidates_Candidate `protobuf:"group,3,rep,name=Candidate,json=candidate" json:"candidate,omitempty"`
	// The position on the composition in character counted by Util::CharsLen.
	// The number represents the left edge of the candidate window.  For example,
	// if the composition is "あいう" and the cursor is the position is between
	// "あ" and "い" (e.g. "あ|いう"), the number should be 1.
	// Note, Util::CharsLen does not take care of IVS or combining character
	// so much.  Thus CharsLen's behavoir on those characters might be changed.
	Position *uint32 `protobuf:"varint,6,req,name=position" json:"position,omitempty"`
	// Nested candidates aka cascading window.
	Subcandidates *Candidates `protobuf:"bytes,8,opt,name=subcandidates" json:"subcandidates,omitempty"`
	// Usages of candidates.
	Usages *InformationList `protobuf:"bytes,10,opt,name=usages" json:"usages,omitempty"`
	// TODO(komatsu): Use CandidateList.
	// Category of the candidates
	Category *Category `protobuf:"varint,11,opt,name=category,enum=commands.Category,def=0" json:"category,omitempty"`
	// Information to be used for rendering.
	DisplayType *DisplayType `protobuf:"varint,12,opt,name=display_type,json=displayType,enum=commands.DisplayType,def=0" json:"display_type,omitempty"`
	// Footer of the GUI window.
	Footer    *Footer               `protobuf:"bytes,13,opt,name=footer" json:"footer,omitempty"`
	Direction *Candidates_Direction `protobuf:"varint,14,opt,name=direction,enum=commands.Candidates_Direction,def=0" json:"direction,omitempty"`
	// The number of candidates per page.
	PageSize         *uint32 `protobuf:"varint,18,opt,name=page_size,json=pageSize,def=9" json:"page_size,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Candidates) Reset()                    { *m = Candidates{} }
func (m *Candidates) String() string            { return proto.CompactTextString(m) }
func (*Candidates) ProtoMessage()               {}
func (*Candidates) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

const Default_Candidates_Category Category = Category_CONVERSION
const Default_Candidates_DisplayType DisplayType = DisplayType_MAIN
const Default_Candidates_Direction Candidates_Direction = Candidates_VERTICAL
const Default_Candidates_PageSize uint32 = 9

func (m *Candidates) GetFocusedIndex() uint32 {
	if m != nil && m.FocusedIndex != nil {
		return *m.FocusedIndex
	}
	return 0
}

func (m *Candidates) GetSize() uint32 {
	if m != nil && m.Size != nil {
		return *m.Size
	}
	return 0
}

func (m *Candidates) GetCandidate() []*Candidates_Candidate {
	if m != nil {
		return m.Candidate
	}
	return nil
}

func (m *Candidates) GetPosition() uint32 {
	if m != nil && m.Position != nil {
		return *m.Position
	}
	return 0
}

func (m *Candidates) GetSubcandidates() *Candidates {
	if m != nil {
		return m.Subcandidates
	}
	return nil
}

func (m *Candidates) GetUsages() *InformationList {
	if m != nil {
		return m.Usages
	}
	return nil
}

func (m *Candidates) GetCategory() Category {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return Default_Candidates_Category
}

func (m *Candidates) GetDisplayType() DisplayType {
	if m != nil && m.DisplayType != nil {
		return *m.DisplayType
	}
	return Default_Candidates_DisplayType
}

func (m *Candidates) GetFooter() *Footer {
	if m != nil {
		return m.Footer
	}
	return nil
}

func (m *Candidates) GetDirection() Candidates_Direction {
	if m != nil && m.Direction != nil {
		return *m.Direction
	}
	return Default_Candidates_Direction
}

func (m *Candidates) GetPageSize() uint32 {
	if m != nil && m.PageSize != nil {
		return *m.PageSize
	}
	return Default_Candidates_PageSize
}

// TODO(komatsu): Use CandidateList.
type Candidates_Candidate struct {
	// The first index should be zero and index numbers should increase by one.
	Index            *uint32     `protobuf:"varint,4,req,name=index" json:"index,omitempty"`
	Value            *string     `protobuf:"bytes,5,req,name=value" json:"value,omitempty"`
	Id               *int32      `protobuf:"varint,9,opt,name=id" json:"id,omitempty"`
	Annotation       *Annotation `protobuf:"bytes,7,opt,name=annotation" json:"annotation,omitempty"`
	InformationId    *int32      `protobuf:"varint,10,opt,name=information_id,json=informationId" json:"information_id,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Candidates_Candidate) Reset()                    { *m = Candidates_Candidate{} }
func (m *Candidates_Candidate) String() string            { return proto.CompactTextString(m) }
func (*Candidates_Candidate) ProtoMessage()               {}
func (*Candidates_Candidate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6, 0} }

func (m *Candidates_Candidate) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *Candidates_Candidate) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

func (m *Candidates_Candidate) GetId() int32 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Candidates_Candidate) GetAnnotation() *Annotation {
	if m != nil {
		return m.Annotation
	}
	return nil
}

func (m *Candidates_Candidate) GetInformationId() int32 {
	if m != nil && m.InformationId != nil {
		return *m.InformationId
	}
	return 0
}

func init() {
	proto.RegisterType((*Annotation)(nil), "commands.Annotation")
	proto.RegisterType((*Information)(nil), "commands.Information")
	proto.RegisterType((*InformationList)(nil), "commands.InformationList")
	proto.RegisterType((*Footer)(nil), "commands.Footer")
	proto.RegisterType((*CandidateWord)(nil), "commands.CandidateWord")
	proto.RegisterType((*CandidateList)(nil), "commands.CandidateList")
	proto.RegisterType((*Candidates)(nil), "commands.Candidates")
	proto.RegisterType((*Candidates_Candidate)(nil), "commands.Candidates.Candidate")
	proto.RegisterEnum("commands.Category", Category_name, Category_value)
	proto.RegisterEnum("commands.DisplayType", DisplayType_name, DisplayType_value)
	proto.RegisterEnum("commands.Candidates_Direction", Candidates_Direction_name, Candidates_Direction_value)
}

func init() { proto.RegisterFile("mozc/commands/candidates.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 931 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xcf, 0x6f, 0xe3, 0x44,
	0x14, 0x5e, 0xe7, 0xd7, 0xda, 0xcf, 0x49, 0xea, 0x1d, 0x16, 0xc8, 0x2e, 0x52, 0x15, 0x02, 0x48,
	0xa1, 0x48, 0xe9, 0x52, 0x81, 0x10, 0x11, 0x42, 0xf2, 0xa6, 0xd9, 0xe2, 0xaa, 0xb4, 0xab, 0x49,
	0xb6, 0x48, 0xbd, 0x44, 0x93, 0xcc, 0x24, 0x1d, 0xea, 0x78, 0x2c, 0x8f, 0x5d, 0x6d, 0x7a, 0xe2,
	0xca, 0x11, 0xf1, 0x07, 0x70, 0xe1, 0xc2, 0x7f, 0x89, 0x66, 0xec, 0xd8, 0xce, 0xaa, 0x42, 0x5d,
	0x71, 0x9b, 0xf7, 0xbd, 0xf7, 0x66, 0x9e, 0xbf, 0x6f, 0xbe, 0x49, 0x60, 0x7f, 0x2d, 0xee, 0x16,
	0x87, 0x0b, 0xb1, 0x5e, 0x93, 0x80, 0xca, 0xc3, 0x05, 0x09, 0x28, 0xa7, 0x24, 0x66, 0x72, 0x10,
	0x46, 0x22, 0x16, 0xc8, 0xdc, 0xa6, 0x7a, 0x7f, 0x19, 0x00, 0x6e, 0x10, 0x88, 0x98, 0xc4, 0x5c,
	0x04, 0xe8, 0x23, 0x68, 0x84, 0x11, 0x5b, 0xf2, 0xb7, 0x1d, 0xa3, 0x6b, 0xf4, 0x2d, 0x9c, 0x45,
	0x0a, 0x97, 0xc9, 0x52, 0xe1, 0x95, 0x14, 0x4f, 0x23, 0xd4, 0x05, 0x9b, 0x32, 0xb9, 0x88, 0x78,
	0xa8, 0xda, 0x3b, 0x55, 0x9d, 0x2c, 0x43, 0xe8, 0x39, 0x98, 0xf2, 0x5a, 0x44, 0xf1, 0x22, 0x89,
	0x3b, 0x35, 0x9d, 0xce, 0x63, 0xf4, 0x19, 0x58, 0x94, 0xf9, 0x2c, 0x26, 0x73, 0x9f, 0x75, 0xea,
	0x5d, 0xa3, 0x6f, 0x0e, 0xeb, 0x4b, 0xe2, 0x4b, 0x86, 0x0b, 0xbc, 0xf7, 0x16, 0x6c, 0x2f, 0x58,
	0x8a, 0x68, 0x9d, 0x4e, 0xd8, 0x86, 0x0a, 0xa7, 0x7a, 0xba, 0x3a, 0xae, 0x70, 0x8a, 0x9e, 0x42,
	0x3d, 0xe6, 0xb1, 0xcf, 0xb2, 0xc1, 0xd2, 0xe0, 0x01, 0x73, 0x7d, 0x0a, 0xcd, 0x9c, 0x96, 0x19,
	0xa7, 0x9d, 0x5a, 0xb7, 0xda, 0xaf, 0x63, 0x3b, 0xc7, 0x3c, 0xda, 0xfb, 0xbd, 0x02, 0x7b, 0xa5,
	0xa3, 0xcf, 0xb8, 0x54, 0x23, 0xb7, 0x96, 0x62, 0x91, 0x48, 0x46, 0x67, 0x3c, 0xa0, 0x2c, 0xe5,
	0xa9, 0x85, 0x9b, 0x19, 0xe8, 0x29, 0x0c, 0x7d, 0x07, 0x36, 0x2f, 0xfa, 0x3a, 0x95, 0x6e, 0xb5,
	0x6f, 0x1f, 0x7d, 0x38, 0xd8, 0x92, 0x3e, 0x28, 0x6d, 0x8a, 0xcb, 0x95, 0x68, 0x08, 0xe6, 0x82,
	0xc4, 0x6c, 0x25, 0xa2, 0x8d, 0x9e, 0xb9, 0x7d, 0x84, 0x8a, 0xae, 0x51, 0x96, 0x19, 0xc2, 0xe8,
	0xe2, 0xfc, 0x72, 0x8c, 0x27, 0xde, 0xc5, 0x39, 0xce, 0xeb, 0x91, 0x0b, 0x4d, 0xca, 0x65, 0xe8,
	0x93, 0xcd, 0x2c, 0xde, 0x84, 0x4c, 0x93, 0xdd, 0x2e, 0x9f, 0x7a, 0x9c, 0x66, 0xa7, 0x9b, 0x90,
	0x0d, 0x1f, 0x8f, 0xdc, 0xc9, 0xc8, 0x3d, 0x1e, 0x63, 0x9b, 0x16, 0x28, 0x7a, 0x06, 0x75, 0xca,
	0x7c, 0xb2, 0xd1, 0x5a, 0xb4, 0x86, 0xd5, 0x6f, 0x5f, 0xbc, 0xc0, 0x29, 0xd2, 0xfb, 0xc3, 0x80,
	0xc6, 0x2b, 0x21, 0x62, 0x16, 0x29, 0xc6, 0x7d, 0x32, 0x67, 0x7e, 0x76, 0x45, 0xd2, 0x00, 0x1d,
	0x40, 0x4b, 0x13, 0x32, 0xbb, 0xe5, 0x92, 0xcf, 0x33, 0x3d, 0x72, 0x3d, 0x9b, 0x3a, 0x77, 0x99,
	0xa6, 0x50, 0x1f, 0x9a, 0xbe, 0x58, 0x89, 0xbc, 0xb4, 0x5a, 0x2e, 0xb5, 0x55, 0x6a, 0x5b, 0xf9,
	0x09, 0x58, 0x32, 0x99, 0xcf, 0xd2, 0xf3, 0xb6, 0xd7, 0x27, 0x99, 0x9f, 0xa9, 0xb8, 0xf7, 0xa7,
	0x01, 0xad, 0xd1, 0x56, 0xaf, 0x5f, 0x44, 0x44, 0xef, 0xbb, 0x1c, 0xa9, 0x4a, 0x15, 0xad, 0x52,
	0x1a, 0x20, 0x07, 0xaa, 0x37, 0x6c, 0x93, 0x5d, 0x0a, 0xb5, 0x54, 0x75, 0xb7, 0xc4, 0x4f, 0x58,
	0x76, 0x44, 0x1a, 0xa0, 0x6f, 0x00, 0x48, 0x6e, 0x0d, 0xcd, 0x89, 0x7d, 0xf4, 0xb4, 0xe0, 0xb3,
	0xb0, 0x0d, 0x2e, 0xd5, 0xf5, 0xfe, 0x29, 0x4f, 0xf5, 0x3e, 0x77, 0x06, 0x0a, 0x9b, 0x66, 0x57,
	0xe6, 0xe3, 0xb2, 0xf8, 0xa5, 0xef, 0xc4, 0xa5, 0xd2, 0xff, 0x73, 0x67, 0x7a, 0xbf, 0x35, 0x00,
	0x46, 0xc5, 0x56, 0x0f, 0x1a, 0x14, 0x41, 0x4d, 0xf2, 0x3b, 0xa5, 0x6f, 0xa5, 0xdf, 0xc2, 0x7a,
	0x8d, 0x7e, 0x00, 0x2b, 0x9f, 0xa8, 0x53, 0xed, 0x56, 0xfb, 0x70, 0xb4, 0x7f, 0xcf, 0xec, 0xa5,
	0x25, 0x2e, 0x1a, 0xd4, 0x13, 0x11, 0x0a, 0xc9, 0x35, 0xcb, 0x0d, 0xbd, 0x6b, 0x1e, 0xa3, 0x21,
	0xb4, 0x64, 0x32, 0x2f, 0x31, 0x63, 0xbe, 0x2b, 0x43, 0xb1, 0x3b, 0xde, 0x2d, 0x45, 0x5f, 0x43,
	0x23, 0x91, 0x64, 0xc5, 0x64, 0x07, 0x74, 0xd3, 0xb3, 0x7b, 0x1d, 0xa8, 0x24, 0xc2, 0x59, 0xe1,
	0x0e, 0x99, 0xf6, 0x7b, 0x1a, 0xf0, 0xc7, 0x77, 0x0c, 0xd8, 0xfc, 0x2f, 0x03, 0xd6, 0x7e, 0x76,
	0xbd, 0xf3, 0x5d, 0xf7, 0xf5, 0xa1, 0xb1, 0xd4, 0x0e, 0xeb, 0xb4, 0xf4, 0xb8, 0x4e, 0xd1, 0x99,
	0x3a, 0x0f, 0x67, 0x79, 0xf4, 0x0a, 0x2c, 0xca, 0x23, 0xb6, 0xd0, 0x8c, 0xb5, 0xf5, 0x31, 0xf7,
	0xd3, 0x7d, 0xbc, 0xad, 0x1a, 0x9a, 0x97, 0x63, 0x3c, 0xf5, 0x46, 0xee, 0x19, 0x2e, 0x5a, 0xd1,
	0x3e, 0x58, 0x21, 0x59, 0xb1, 0x99, 0xd6, 0x13, 0x69, 0xcf, 0x1b, 0xdf, 0x63, 0x53, 0x61, 0x13,
	0x7e, 0xc7, 0x9e, 0xff, 0x6d, 0x80, 0x95, 0xef, 0x56, 0x98, 0xa9, 0xa6, 0x35, 0xca, 0xcc, 0x94,
	0x5b, 0xa7, 0xde, 0xad, 0x14, 0xd6, 0x49, 0x8d, 0x68, 0xe5, 0x46, 0xdc, 0xb5, 0xd2, 0xe3, 0x87,
	0x59, 0x09, 0x7d, 0x01, 0xed, 0xd2, 0xeb, 0xa8, 0x5e, 0x69, 0xd0, 0x3b, 0xb6, 0x4a, 0xa8, 0x47,
	0x7b, 0x5f, 0x82, 0x95, 0x7f, 0x28, 0x6a, 0x42, 0xfe, 0xa9, 0xce, 0x23, 0xd4, 0x06, 0xf8, 0xe9,
	0x02, 0x7b, 0x57, 0x17, 0xe7, 0x53, 0xf7, 0xcc, 0x31, 0x4e, 0x6b, 0xe6, 0x9e, 0xe3, 0x9c, 0xd6,
	0x4c, 0xc7, 0x79, 0x72, 0x5a, 0x33, 0x9f, 0x38, 0xe8, 0xe0, 0x0a, 0xcc, 0xad, 0xae, 0xaa, 0xba,
	0x50, 0x36, 0xed, 0x7e, 0x8d, 0xc7, 0xc7, 0xde, 0x68, 0xaa, 0x62, 0x43, 0xc5, 0x93, 0x37, 0x27,
	0x27, 0xe3, 0x89, 0x8e, 0x2b, 0xe8, 0x03, 0xd8, 0x9b, 0x62, 0xf7, 0x7c, 0x72, 0xe6, 0x4d, 0xc7,
	0xd8, 0xd5, 0x60, 0x15, 0x59, 0x50, 0x7f, 0x33, 0x71, 0x4f, 0xc6, 0x4e, 0xed, 0xe0, 0x73, 0xb0,
	0x4b, 0x9a, 0x23, 0x13, 0xb4, 0xea, 0xce, 0x23, 0x64, 0xc3, 0xf6, 0x01, 0x76, 0x8c, 0x97, 0x01,
	0x0c, 0x44, 0xb4, 0x1a, 0xa8, 0x9f, 0xec, 0x01, 0x09, 0x68, 0x24, 0x38, 0x1d, 0xf0, 0x20, 0x4c,
	0xe2, 0x35, 0x8b, 0xaf, 0x05, 0x1d, 0xfc, 0x4a, 0x42, 0x12, 0x30, 0xc9, 0xd2, 0xdf, 0xee, 0x79,
	0xb2, 0x7c, 0xb9, 0xf7, 0x5a, 0xad, 0x0a, 0x9d, 0xaf, 0xbe, 0x5a, 0xf1, 0xf8, 0x3a, 0x99, 0x2b,
	0x42, 0x0f, 0x63, 0xb1, 0xde, 0xdc, 0x10, 0x1e, 0x91, 0xc3, 0xb5, 0xbc, 0xb9, 0x91, 0x2c, 0xba,
	0x3d, 0xdc, 0xf9, 0x37, 0xf0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0x52, 0xb0, 0x51, 0x1d,
	0x08, 0x00, 0x00,
}