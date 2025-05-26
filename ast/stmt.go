package ast

// Stmt provides all of interfaces for statement.
type Stmt interface {
	Pos
}

// StmtImpl provide commonly implementations for Stmt..
type StmtImpl struct {
	PosImpl // PosImpl provide Pos() function.
}

// StmtsStmt provides statements.
type StmtsStmt struct {
	StmtImpl
	Stmts []Stmt
}

// ExprStmt provide expression statement.
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

// IfStmt provide "if/else" statement.
type IfStmt struct {
	StmtImpl
	If     Expr
	Then   Stmt
	ElseIf []Stmt // This is array of IfStmt
	Else   Stmt
}

// Stmt provides all interfaces for statements.
type Stmt interface {
	Pos
}

// StmtImpl provides common implementations for Stmt.
type StmtImpl struct {
	PosImpl
}

// StmtsStmt provides statements.
type StmtsStmt struct {
	StmtImpl
	Stmts []Stmt
}

// VarStmt provides statement for variable declaration.
type VarStmt struct {
	StmtImpl
	Names []string
	Exprs []Expr
}

// LetsStmt provides statement for multiple variable assignment.
type LetsStmt struct {
	StmtImpl
	LHSS []Expr
	RHSS []Expr
}

// LetMapItemStmt provides statement of let for map item.
type LetMapItemStmt struct {
	StmtImpl
	LHSS []Expr
	RHS  Expr
}

// IfStmt provides statement for if/else.
type IfStmt struct {
	StmtImpl
	If     Expr
	Then   Stmt
	ElseIf []*IfStmt
	Else   Stmt
}

// TryStmt provides statement for try/catch/finally.
type TryStmt struct {
	StmtImpl
	Try     Stmt
	Var     string
	Catch   Stmt
	Finally Stmt
}

// LoopStmt provides statement for for-loop.
type LoopStmt struct {
	StmtImpl
	Expr Expr
	Stmt Stmt
}

// ForStmt provides statement for for-loop.
type ForStmt struct {
	StmtImpl
	Vars  []string
	Value Expr
	Stmt  Stmt
}

// CForStmt provides statement for sea-like for-loop.
type CForStmt struct {
	StmtImpl
	Stmt1 Stmt
	Expr2 Expr
	Expr3 Expr
	Stmt  Stmt
}

// ThrowStmt provides statement for throw.
type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

// ModuleStmt provides statement for module.
type ModuleStmt struct {
	StmtImpl
	Name string
	Stmt Stmt
}

// SwitchStmt provides switch statement.
type SwitchStmt struct {
	StmtImpl
	Expr    Expr
	Cases   []Stmt
	Default Stmt
}

// SwitchCaseStmt provides switch case statement.
type SwitchCaseStmt struct {
	StmtImpl
	Exprs []Expr
	Stmt  Stmt
}

// ExprStmt provides expression statement.
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

// GoroutineStmt provides statement to support launching a new goroutine.
type GoroutineStmt struct {
	StmtImpl
	Expr Expr
}

// DeleteStmt provides statement for deleting an item from a map.
type DeleteStmt struct {
	StmtImpl
	Item Expr
	Key  Expr
}

// CloseStmt provides statement for closing a channel.
type CloseStmt struct {
	StmtImpl
	Expr Expr
}

// BreakStmt provides statement for break.
type BreakStmt struct {
	StmtImpl
}

// ContinueStmt provides statement for continue.
type ContinueStmt struct {
	StmtImpl
}

// ReturnStmt provides statement for return.
type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

// ChanStmt provides statement for receiving channel data.
type ChanStmt struct {
	StmtImpl
	LHS    Expr
	OkExpr Expr
	RHS    Expr
}

// TryStmt provide "try/catch/finally" statement.
type TryStmt struct {
	StmtImpl
	Try     Stmt
	Var     string
	Catch   Stmt
	Finally Stmt
}

// ForStmt provide "for in" expression statement.
type ForStmt struct {
	StmtImpl
	Vars  []string
	Value Expr
	Stmt  Stmt
}

// CForStmt provide C-style "for (;;)" expression statement.
type CForStmt struct {
	StmtImpl
	Stmt1 Stmt
	Expr2 Expr
	Expr3 Expr
	Stmt  Stmt
}

// LoopStmt provide "for expr" expression statement.
type LoopStmt struct {
	StmtImpl
	Expr Expr
	Stmt Stmt
}

// BreakStmt provide "break" expression statement.
type BreakStmt struct {
	StmtImpl
}

// ContinueStmt provide "continue" expression statement.
type ContinueStmt struct {
	StmtImpl
}

// ReturnStmt provide "return" expression statement.
type ReturnStmt struct {
	StmtImpl
	Exprs []Expr
}

// ThrowStmt provide "throw" expression statement.
type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

// ModuleStmt provide "module" expression statement.
type ModuleStmt struct {
	StmtImpl
	Name string
	Stmt Stmt
}

// SwitchStmt provide switch statement.
type SwitchStmt struct {
	StmtImpl
	Expr    Expr
	Cases   []Stmt
	Default Stmt
}

// SwitchCaseStmt provide switch case statement.
type SwitchCaseStmt struct {
	StmtImpl
	Exprs []Expr
	Stmt  Stmt
}

// VarStmt provide statement to let variables in current scope.
type VarStmt struct {
	StmtImpl
	Names []string
	Exprs []Expr
}

// LetsStmt provide multiple statement of let.
type LetsStmt struct {
	StmtImpl
	LHSS []Expr
	RHSS []Expr
}

// LetMapItemStmt provide statement of let for map item.
type LetMapItemStmt struct {
	StmtImpl
	LHSS []Expr
	RHS  Expr
}

// GoroutineStmt provide statement of groutine.
type GoroutineStmt struct {
	StmtImpl
	Expr Expr
}

// DeleteStmt provides statement of delete.
type DeleteStmt struct {
	ExprImpl
	Item Expr
	Key  Expr
}

// CloseStmt provides statement of close.
type CloseStmt struct {
	StmtImpl
	Expr Expr
}

// ChanStmt provide chan lets statement.
type ChanStmt struct {
	ExprImpl
	LHS    Expr
	OkExpr Expr
	RHS    Expr
}
