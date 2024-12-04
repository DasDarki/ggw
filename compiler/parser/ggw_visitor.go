// Code generated from ggw.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ggw
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ggwParser.
type ggwVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ggwParser#file.
	VisitFile(ctx *FileContext) interface{}

	// Visit a parse tree produced by ggwParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ggwParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ggwParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#comparisonExpression.
	VisitComparisonExpression(ctx *ComparisonExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#postfixExpression.
	VisitPostfixExpression(ctx *PostfixExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#postfixOperator.
	VisitPostfixOperator(ctx *PostfixOperatorContext) interface{}

	// Visit a parse tree produced by ggwParser#incrementalExpression.
	VisitIncrementalExpression(ctx *IncrementalExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#groupedExpression.
	VisitGroupedExpression(ctx *GroupedExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#funcCallExpression.
	VisitFuncCallExpression(ctx *FuncCallExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#funcCallChaining.
	VisitFuncCallChaining(ctx *FuncCallChainingContext) interface{}

	// Visit a parse tree produced by ggwParser#arrayExpression.
	VisitArrayExpression(ctx *ArrayExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#mapExpression.
	VisitMapExpression(ctx *MapExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#onStatement.
	VisitOnStatement(ctx *OnStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#moduleStatemnt.
	VisitModuleStatemnt(ctx *ModuleStatemntContext) interface{}

	// Visit a parse tree produced by ggwParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#argument.
	VisitArgument(ctx *ArgumentContext) interface{}

	// Visit a parse tree produced by ggwParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by ggwParser#argumentPassList.
	VisitArgumentPassList(ctx *ArgumentPassListContext) interface{}

	// Visit a parse tree produced by ggwParser#annotation.
	VisitAnnotation(ctx *AnnotationContext) interface{}

	// Visit a parse tree produced by ggwParser#annotationList.
	VisitAnnotationList(ctx *AnnotationListContext) interface{}

	// Visit a parse tree produced by ggwParser#annotationArgName.
	VisitAnnotationArgName(ctx *AnnotationArgNameContext) interface{}

	// Visit a parse tree produced by ggwParser#annotationArgument.
	VisitAnnotationArgument(ctx *AnnotationArgumentContext) interface{}

	// Visit a parse tree produced by ggwParser#annotationDeclaration.
	VisitAnnotationDeclaration(ctx *AnnotationDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#accessModifier.
	VisitAccessModifier(ctx *AccessModifierContext) interface{}

	// Visit a parse tree produced by ggwParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#deconstruct.
	VisitDeconstruct(ctx *DeconstructContext) interface{}

	// Visit a parse tree produced by ggwParser#variableName.
	VisitVariableName(ctx *VariableNameContext) interface{}

	// Visit a parse tree produced by ggwParser#variableKind.
	VisitVariableKind(ctx *VariableKindContext) interface{}

	// Visit a parse tree produced by ggwParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#variableMultiDeclaration.
	VisitVariableMultiDeclaration(ctx *VariableMultiDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#variableClassDeclaration.
	VisitVariableClassDeclaration(ctx *VariableClassDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#wireKind.
	VisitWireKind(ctx *WireKindContext) interface{}

	// Visit a parse tree produced by ggwParser#wireDeclaration.
	VisitWireDeclaration(ctx *WireDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#wireAuth.
	VisitWireAuth(ctx *WireAuthContext) interface{}

	// Visit a parse tree produced by ggwParser#functionDeclaration.
	VisitFunctionDeclaration(ctx *FunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#functionClassDeclaration.
	VisitFunctionClassDeclaration(ctx *FunctionClassDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#anonymousFunctionDeclaration.
	VisitAnonymousFunctionDeclaration(ctx *AnonymousFunctionDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#constructorHeadDeclaration.
	VisitConstructorHeadDeclaration(ctx *ConstructorHeadDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#constructorFuncDeclaration.
	VisitConstructorFuncDeclaration(ctx *ConstructorFuncDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#enumDeclaration.
	VisitEnumDeclaration(ctx *EnumDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#enumProperty.
	VisitEnumProperty(ctx *EnumPropertyContext) interface{}

	// Visit a parse tree produced by ggwParser#interfaceDeclaration.
	VisitInterfaceDeclaration(ctx *InterfaceDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#interfaceBodyDeclaration.
	VisitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#interfaceProperty.
	VisitInterfaceProperty(ctx *InterfacePropertyContext) interface{}

	// Visit a parse tree produced by ggwParser#interfaceFunction.
	VisitInterfaceFunction(ctx *InterfaceFunctionContext) interface{}

	// Visit a parse tree produced by ggwParser#classDeclaration.
	VisitClassDeclaration(ctx *ClassDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#classBodyDeclaration.
	VisitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#whenStatement.
	VisitWhenStatement(ctx *WhenStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#whenCase.
	VisitWhenCase(ctx *WhenCaseContext) interface{}

	// Visit a parse tree produced by ggwParser#whenCaseBlock.
	VisitWhenCaseBlock(ctx *WhenCaseBlockContext) interface{}

	// Visit a parse tree produced by ggwParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#forIncrementExpression.
	VisitForIncrementExpression(ctx *ForIncrementExpressionContext) interface{}

	// Visit a parse tree produced by ggwParser#forInStatement.
	VisitForInStatement(ctx *ForInStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#doWhileStatement.
	VisitDoWhileStatement(ctx *DoWhileStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#grantStatement.
	VisitGrantStatement(ctx *GrantStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#revokeStatement.
	VisitRevokeStatement(ctx *RevokeStatementContext) interface{}

	// Visit a parse tree produced by ggwParser#grantTree.
	VisitGrantTree(ctx *GrantTreeContext) interface{}

	// Visit a parse tree produced by ggwParser#grantTreeBody.
	VisitGrantTreeBody(ctx *GrantTreeBodyContext) interface{}

	// Visit a parse tree produced by ggwParser#grantTreeBranch.
	VisitGrantTreeBranch(ctx *GrantTreeBranchContext) interface{}

	// Visit a parse tree produced by ggwParser#grantTreeInheritance.
	VisitGrantTreeInheritance(ctx *GrantTreeInheritanceContext) interface{}

	// Visit a parse tree produced by ggwParser#sessionDeclaration.
	VisitSessionDeclaration(ctx *SessionDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#sessionBody.
	VisitSessionBody(ctx *SessionBodyContext) interface{}

	// Visit a parse tree produced by ggwParser#sessionProperty.
	VisitSessionProperty(ctx *SessionPropertyContext) interface{}

	// Visit a parse tree produced by ggwParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ggwParser#primitiveType.
	VisitPrimitiveType(ctx *PrimitiveTypeContext) interface{}

	// Visit a parse tree produced by ggwParser#arrayType.
	VisitArrayType(ctx *ArrayTypeContext) interface{}

	// Visit a parse tree produced by ggwParser#mapType.
	VisitMapType(ctx *MapTypeContext) interface{}

	// Visit a parse tree produced by ggwParser#tupleType.
	VisitTupleType(ctx *TupleTypeContext) interface{}

	// Visit a parse tree produced by ggwParser#genericTypeDeclaration.
	VisitGenericTypeDeclaration(ctx *GenericTypeDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#genericDeclaration.
	VisitGenericDeclaration(ctx *GenericDeclarationContext) interface{}

	// Visit a parse tree produced by ggwParser#genericUsage.
	VisitGenericUsage(ctx *GenericUsageContext) interface{}

	// Visit a parse tree produced by ggwParser#identifierUsage.
	VisitIdentifierUsage(ctx *IdentifierUsageContext) interface{}

	// Visit a parse tree produced by ggwParser#extends.
	VisitExtends(ctx *ExtendsContext) interface{}

	// Visit a parse tree produced by ggwParser#interfaceExtends.
	VisitInterfaceExtends(ctx *InterfaceExtendsContext) interface{}

	// Visit a parse tree produced by ggwParser#implements.
	VisitImplements(ctx *ImplementsContext) interface{}
}
