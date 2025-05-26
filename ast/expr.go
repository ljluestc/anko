package ast

import "reflect"

// Pos interface provides two functions to get/set the position for expression or statement.
type Pos interface {
	Position() Position
	SetPosition(Position)
}

// PosImpl provides common implementations for Pos.
type PosImpl struct {
	position Position
}

// Position returns the position of the expression or statement.
func (p *PosImpl) Position() Position {
	return p.position
}

// SetPosition sets the position of the expression or statement.
func (p *PosImpl) SetPosition(pos Position) {
	p.position = pos
}

// Operator provides interfaces for operators.
type Operator interface {
	Pos
	String() string
}

// OperatorImpl provides common implementations for Operator.
type OperatorImpl struct {
	PosImpl
	Op string
}

// String returns the operator as a string.
func (o *OperatorImpl) String() string {
	return o.Op
}

// Expr provides all interfaces for expressions.
type Expr interface {
	Pos
}

// ExprImpl provides common implementations for Expr.
type ExprImpl struct {
	PosImpl
}

// MemberExpr provides expression to refer to a member.
type MemberExpr struct {
	ExprImpl
	Expr Expr
	Name string
}

// IdentExpr provides identity expression.
type IdentExpr struct {
	ExprImpl
	Lit string
}

// MapExpr provides a map expression.
type MapExpr struct {
	ExprImpl
	Keys     []Expr
	Values   []Expr
	TypeData *TypeStruct
}

// CallExpr provides a calling expression.
type CallExpr struct {
	ExprImpl
	Func     reflect.Value
	Name     string
	SubExprs []Expr
	VarArg   bool
	Go       bool
}

// LiteralExpr provides literal expression.
type LiteralExpr struct {
	ExprImpl
	Literal reflect.Value
}
