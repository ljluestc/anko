package ast

// Define missing types with required fields and methods.

type AnonCallExpr struct {
	Expr     Expr
	SubExprs []Expr
	VarArg   bool
	Go       bool
	position Position
}

func (e *AnonCallExpr) Position() Position {
	return e.position
}

func (e *AnonCallExpr) SetPosition(pos Position) {
	e.position = pos
}

type ItemExpr struct {
	Item     Expr
	Index    Expr
	position Position
}

func (e *ItemExpr) Position() Position {
	return e.position
}

func (e *ItemExpr) SetPosition(pos Position) {
	e.position = pos
}

type SliceExpr struct {
	Item     Expr
	Begin    Expr
	End      Expr
	Cap      Expr
	position Position
}

func (e *SliceExpr) Position() Position {
	return e.position
}

func (e *SliceExpr) SetPosition(pos Position) {
	e.position = pos
}

type TernaryOpExpr struct {
	Expr     Expr
	LHS      Expr
	RHS      Expr
	position Position
}

func (e *TernaryOpExpr) Position() Position {
	return e.position
}

func (e *TernaryOpExpr) SetPosition(pos Position) {
	e.position = pos
}

type NilCoalescingOpExpr struct {
	LHS      Expr
	RHS      Expr
	position Position
}

func (e *NilCoalescingOpExpr) Position() Position {
	return e.position
}

func (e *NilCoalescingOpExpr) SetPosition(pos Position) {
	e.position = pos
}

type FuncExpr struct {
	Name     string
	Params   []string
	Stmt     Stmt
	VarArg   bool
	position Position
}

func (e *FuncExpr) Position() Position {
	return e.position
}

func (e *FuncExpr) SetPosition(pos Position) {
	e.position = pos
}

type ArrayExpr struct {
	Exprs    []Expr
	TypeData *TypeStruct
	position Position
}

func (e *ArrayExpr) Position() Position {
	return e.position
}

func (e *ArrayExpr) SetPosition(pos Position) {
	e.position = pos
}

type ParenExpr struct {
	SubExpr  Expr
	position Position
}

func (e *ParenExpr) Position() Position {
	return e.position
}

func (e *ParenExpr) SetPosition(pos Position) {
	e.position = pos
}

type LenExpr struct {
	Expr     Expr
	position Position
}

func (e *LenExpr) Position() Position {
	return e.position
}

func (e *LenExpr) SetPosition(pos Position) {
	e.position = pos
}

type ImportExpr struct {
	Name     Expr
	position Position
}

func (e *ImportExpr) Position() Position {
	return e.position
}

func (e *ImportExpr) SetPosition(pos Position) {
	e.position = pos
}

type MakeExpr struct {
	TypeData *TypeStruct
	LenExpr  Expr
	CapExpr  Expr
	position Position
}

func (e *MakeExpr) Position() Position {
	return e.position
}

func (e *MakeExpr) SetPosition(pos Position) {
	e.position = pos
}

type MakeTypeExpr struct {
	Name     string
	Type     Expr
	position Position
}

func (e *MakeTypeExpr) Position() Position {
	return e.position
}

func (e *MakeTypeExpr) SetPosition(pos Position) {
	e.position = pos
}

type IncludeExpr struct {
	ItemExpr Expr
	ListExpr Expr
	position Position
}

func (e *IncludeExpr) Position() Position {
	return e.position
}

func (e *IncludeExpr) SetPosition(pos Position) {
	e.position = pos
}

// ChanExpr represents a channel operation expression
type ChanExpr struct {
	LHS      Expr
	RHS      Expr
	position Position
}

func (e *ChanExpr) Position() Position {
	return e.position
}

func (e *ChanExpr) SetPosition(pos Position) {
	e.position = pos
}

// UnaryExpr represents a unary operator expression
type UnaryExpr struct {
	Operator string
	Expr     Expr
	position Position
}

func (e *UnaryExpr) Position() Position {
	return e.position
}

func (e *UnaryExpr) SetPosition(pos Position) {
	e.position = pos
}

// AddrExpr represents an address-of expression
type AddrExpr struct {
	Expr     Expr
	position Position
}

func (e *AddrExpr) Position() Position {
	return e.position
}

func (e *AddrExpr) SetPosition(pos Position) {
	e.position = pos
}

// DerefExpr represents a dereference expression
type DerefExpr struct {
	Expr     Expr
	position Position
}

func (e *DerefExpr) Position() Position {
	return e.position
}

func (e *DerefExpr) SetPosition(pos Position) {
	e.position = pos
}

// OpExpr represents an operator expression
type OpExpr struct {
	Op       Operator
	position Position
}

func (e *OpExpr) Position() Position {
	return e.position
}

func (e *OpExpr) SetPosition(pos Position) {
	e.position = pos
}

// AddOperator represents an addition/subtraction operator
type AddOperator struct {
	LHS      Expr
	Operator string
	RHS      Expr
	position Position
}

func (o *AddOperator) Position() Position {
	return o.position
}

func (o *AddOperator) SetPosition(pos Position) {
	o.position = pos
}

func (o *AddOperator) String() string {
	return o.Operator
}

// MultiplyOperator represents a multiplication/division operator
type MultiplyOperator struct {
	LHS      Expr
	Operator string
	RHS      Expr
	position Position
}

func (o *MultiplyOperator) Position() Position {
	return o.position
}

func (o *MultiplyOperator) SetPosition(pos Position) {
	o.position = pos
}

func (o *MultiplyOperator) String() string {
	return o.Operator
}

// ComparisonOperator represents a comparison operator like == != < > <= >=
type ComparisonOperator struct {
	LHS      Expr
	Operator string
	RHS      Expr
	position Position
}

func (o *ComparisonOperator) Position() Position {
	return o.position
}

func (o *ComparisonOperator) SetPosition(pos Position) {
	o.position = pos
}

func (o *ComparisonOperator) String() string {
	return o.Operator
}

// BinaryOperator represents other binary operators like && || & | etc
type BinaryOperator struct {
	LHS      Expr
	Operator string
	RHS      Expr
	position Position
}

func (o *BinaryOperator) Position() Position {
	return o.position
}

func (o *BinaryOperator) SetPosition(pos Position) {
	o.position = pos
}

func (o *BinaryOperator) String() string {
	return o.Operator
}

// LetsExpr represents a let statement expression
type LetsExpr struct {
	LHSS     []Expr
	RHSS     []Expr
	position Position
}

func (e *LetsExpr) Position() Position {
	return e.position
}

func (e *LetsExpr) SetPosition(pos Position) {
	e.position = pos
}
