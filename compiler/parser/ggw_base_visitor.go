// Code generated from ggw.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ggw
import "github.com/antlr4-go/antlr/v4"

type BaseggwVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseggwVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitPostfixExpression(ctx *PostfixExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitPostfixOperator(ctx *PostfixOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitIncrementalExpression(ctx *IncrementalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGroupedExpression(ctx *GroupedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitFuncCallExpression(ctx *FuncCallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitFuncCallChaining(ctx *FuncCallChainingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitArrayExpression(ctx *ArrayExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitMapExpression(ctx *MapExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitOnStatement(ctx *OnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitModuleStatemnt(ctx *ModuleStatemntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitArgument(ctx *ArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitArgumentPassList(ctx *ArgumentPassListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAnnotationList(ctx *AnnotationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAnnotationArgName(ctx *AnnotationArgNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAnnotationArgument(ctx *AnnotationArgumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAnnotationDeclaration(ctx *AnnotationDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAccessModifier(ctx *AccessModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitDeconstruct(ctx *DeconstructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitVariableName(ctx *VariableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitVariableKind(ctx *VariableKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitVariableMultiDeclaration(ctx *VariableMultiDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitVariableClassDeclaration(ctx *VariableClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWireKind(ctx *WireKindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWireDeclaration(ctx *WireDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWireAuth(ctx *WireAuthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitFunctionClassDeclaration(ctx *FunctionClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitAnonymousFunctionDeclaration(ctx *AnonymousFunctionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitConstructorHeadDeclaration(ctx *ConstructorHeadDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitConstructorFuncDeclaration(ctx *ConstructorFuncDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitEnumProperty(ctx *EnumPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitInterfaceProperty(ctx *InterfacePropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitInterfaceFunction(ctx *InterfaceFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitClassDeclaration(ctx *ClassDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWhenStatement(ctx *WhenStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWhenCase(ctx *WhenCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWhenCaseBlock(ctx *WhenCaseBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitForIncrementExpression(ctx *ForIncrementExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitForInStatement(ctx *ForInStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGrantStatement(ctx *GrantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitRevokeStatement(ctx *RevokeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGrantTree(ctx *GrantTreeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGrantTreeBody(ctx *GrantTreeBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGrantTreeBranch(ctx *GrantTreeBranchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGrantTreeInheritance(ctx *GrantTreeInheritanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitSessionDeclaration(ctx *SessionDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitSessionBody(ctx *SessionBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitSessionProperty(ctx *SessionPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitArrayType(ctx *ArrayTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitMapType(ctx *MapTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitTupleType(ctx *TupleTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGenericTypeDeclaration(ctx *GenericTypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGenericDeclaration(ctx *GenericDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitGenericUsage(ctx *GenericUsageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitIdentifierUsage(ctx *IdentifierUsageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitExtends(ctx *ExtendsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitInterfaceExtends(ctx *InterfaceExtendsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseggwVisitor) VisitImplements(ctx *ImplementsContext) interface{} {
	return v.VisitChildren(ctx)
}
