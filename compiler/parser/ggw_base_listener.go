// Code generated from ggw.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ggw
import "github.com/antlr4-go/antlr/v4"

// BaseggwListener is a complete listener for a parse tree produced by ggwParser.
type BaseggwListener struct{}

var _ ggwListener = &BaseggwListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseggwListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseggwListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseggwListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseggwListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseggwListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseggwListener) ExitFile(ctx *FileContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseggwListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseggwListener) ExitBlock(ctx *BlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseggwListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseggwListener) ExitStatement(ctx *StatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseggwListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseggwListener) ExitExpression(ctx *ExpressionContext) {}

// EnterLogicalExpression is called when production logicalExpression is entered.
func (s *BaseggwListener) EnterLogicalExpression(ctx *LogicalExpressionContext) {}

// ExitLogicalExpression is called when production logicalExpression is exited.
func (s *BaseggwListener) ExitLogicalExpression(ctx *LogicalExpressionContext) {}

// EnterComparisonExpression is called when production comparisonExpression is entered.
func (s *BaseggwListener) EnterComparisonExpression(ctx *ComparisonExpressionContext) {}

// ExitComparisonExpression is called when production comparisonExpression is exited.
func (s *BaseggwListener) ExitComparisonExpression(ctx *ComparisonExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseggwListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseggwListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseggwListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseggwListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {}

// EnterPostfixExpression is called when production postfixExpression is entered.
func (s *BaseggwListener) EnterPostfixExpression(ctx *PostfixExpressionContext) {}

// ExitPostfixExpression is called when production postfixExpression is exited.
func (s *BaseggwListener) ExitPostfixExpression(ctx *PostfixExpressionContext) {}

// EnterPostfixOperator is called when production postfixOperator is entered.
func (s *BaseggwListener) EnterPostfixOperator(ctx *PostfixOperatorContext) {}

// ExitPostfixOperator is called when production postfixOperator is exited.
func (s *BaseggwListener) ExitPostfixOperator(ctx *PostfixOperatorContext) {}

// EnterIncrementalExpression is called when production incrementalExpression is entered.
func (s *BaseggwListener) EnterIncrementalExpression(ctx *IncrementalExpressionContext) {}

// ExitIncrementalExpression is called when production incrementalExpression is exited.
func (s *BaseggwListener) ExitIncrementalExpression(ctx *IncrementalExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseggwListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseggwListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterGroupedExpression is called when production groupedExpression is entered.
func (s *BaseggwListener) EnterGroupedExpression(ctx *GroupedExpressionContext) {}

// ExitGroupedExpression is called when production groupedExpression is exited.
func (s *BaseggwListener) ExitGroupedExpression(ctx *GroupedExpressionContext) {}

// EnterFuncCallExpression is called when production funcCallExpression is entered.
func (s *BaseggwListener) EnterFuncCallExpression(ctx *FuncCallExpressionContext) {}

// ExitFuncCallExpression is called when production funcCallExpression is exited.
func (s *BaseggwListener) ExitFuncCallExpression(ctx *FuncCallExpressionContext) {}

// EnterFuncCallChaining is called when production funcCallChaining is entered.
func (s *BaseggwListener) EnterFuncCallChaining(ctx *FuncCallChainingContext) {}

// ExitFuncCallChaining is called when production funcCallChaining is exited.
func (s *BaseggwListener) ExitFuncCallChaining(ctx *FuncCallChainingContext) {}

// EnterArrayExpression is called when production arrayExpression is entered.
func (s *BaseggwListener) EnterArrayExpression(ctx *ArrayExpressionContext) {}

// ExitArrayExpression is called when production arrayExpression is exited.
func (s *BaseggwListener) ExitArrayExpression(ctx *ArrayExpressionContext) {}

// EnterMapExpression is called when production mapExpression is entered.
func (s *BaseggwListener) EnterMapExpression(ctx *MapExpressionContext) {}

// ExitMapExpression is called when production mapExpression is exited.
func (s *BaseggwListener) ExitMapExpression(ctx *MapExpressionContext) {}

// EnterOnStatement is called when production onStatement is entered.
func (s *BaseggwListener) EnterOnStatement(ctx *OnStatementContext) {}

// ExitOnStatement is called when production onStatement is exited.
func (s *BaseggwListener) ExitOnStatement(ctx *OnStatementContext) {}

// EnterModuleStatemnt is called when production moduleStatemnt is entered.
func (s *BaseggwListener) EnterModuleStatemnt(ctx *ModuleStatemntContext) {}

// ExitModuleStatemnt is called when production moduleStatemnt is exited.
func (s *BaseggwListener) ExitModuleStatemnt(ctx *ModuleStatemntContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseggwListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseggwListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterArgument is called when production argument is entered.
func (s *BaseggwListener) EnterArgument(ctx *ArgumentContext) {}

// ExitArgument is called when production argument is exited.
func (s *BaseggwListener) ExitArgument(ctx *ArgumentContext) {}

// EnterArgumentList is called when production argumentList is entered.
func (s *BaseggwListener) EnterArgumentList(ctx *ArgumentListContext) {}

// ExitArgumentList is called when production argumentList is exited.
func (s *BaseggwListener) ExitArgumentList(ctx *ArgumentListContext) {}

// EnterArgumentPassList is called when production argumentPassList is entered.
func (s *BaseggwListener) EnterArgumentPassList(ctx *ArgumentPassListContext) {}

// ExitArgumentPassList is called when production argumentPassList is exited.
func (s *BaseggwListener) ExitArgumentPassList(ctx *ArgumentPassListContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseggwListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseggwListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterAnnotationList is called when production annotationList is entered.
func (s *BaseggwListener) EnterAnnotationList(ctx *AnnotationListContext) {}

// ExitAnnotationList is called when production annotationList is exited.
func (s *BaseggwListener) ExitAnnotationList(ctx *AnnotationListContext) {}

// EnterAnnotationArgName is called when production annotationArgName is entered.
func (s *BaseggwListener) EnterAnnotationArgName(ctx *AnnotationArgNameContext) {}

// ExitAnnotationArgName is called when production annotationArgName is exited.
func (s *BaseggwListener) ExitAnnotationArgName(ctx *AnnotationArgNameContext) {}

// EnterAnnotationArgument is called when production annotationArgument is entered.
func (s *BaseggwListener) EnterAnnotationArgument(ctx *AnnotationArgumentContext) {}

// ExitAnnotationArgument is called when production annotationArgument is exited.
func (s *BaseggwListener) ExitAnnotationArgument(ctx *AnnotationArgumentContext) {}

// EnterAnnotationDeclaration is called when production annotationDeclaration is entered.
func (s *BaseggwListener) EnterAnnotationDeclaration(ctx *AnnotationDeclarationContext) {}

// ExitAnnotationDeclaration is called when production annotationDeclaration is exited.
func (s *BaseggwListener) ExitAnnotationDeclaration(ctx *AnnotationDeclarationContext) {}

// EnterAccessModifier is called when production accessModifier is entered.
func (s *BaseggwListener) EnterAccessModifier(ctx *AccessModifierContext) {}

// ExitAccessModifier is called when production accessModifier is exited.
func (s *BaseggwListener) ExitAccessModifier(ctx *AccessModifierContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BaseggwListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BaseggwListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterDeconstruct is called when production deconstruct is entered.
func (s *BaseggwListener) EnterDeconstruct(ctx *DeconstructContext) {}

// ExitDeconstruct is called when production deconstruct is exited.
func (s *BaseggwListener) ExitDeconstruct(ctx *DeconstructContext) {}

// EnterVariableName is called when production variableName is entered.
func (s *BaseggwListener) EnterVariableName(ctx *VariableNameContext) {}

// ExitVariableName is called when production variableName is exited.
func (s *BaseggwListener) ExitVariableName(ctx *VariableNameContext) {}

// EnterVariableKind is called when production variableKind is entered.
func (s *BaseggwListener) EnterVariableKind(ctx *VariableKindContext) {}

// ExitVariableKind is called when production variableKind is exited.
func (s *BaseggwListener) ExitVariableKind(ctx *VariableKindContext) {}

// EnterVariableDeclaration is called when production variableDeclaration is entered.
func (s *BaseggwListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

// ExitVariableDeclaration is called when production variableDeclaration is exited.
func (s *BaseggwListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

// EnterVariableMultiDeclaration is called when production variableMultiDeclaration is entered.
func (s *BaseggwListener) EnterVariableMultiDeclaration(ctx *VariableMultiDeclarationContext) {}

// ExitVariableMultiDeclaration is called when production variableMultiDeclaration is exited.
func (s *BaseggwListener) ExitVariableMultiDeclaration(ctx *VariableMultiDeclarationContext) {}

// EnterVariableClassDeclaration is called when production variableClassDeclaration is entered.
func (s *BaseggwListener) EnterVariableClassDeclaration(ctx *VariableClassDeclarationContext) {}

// ExitVariableClassDeclaration is called when production variableClassDeclaration is exited.
func (s *BaseggwListener) ExitVariableClassDeclaration(ctx *VariableClassDeclarationContext) {}

// EnterWireKind is called when production wireKind is entered.
func (s *BaseggwListener) EnterWireKind(ctx *WireKindContext) {}

// ExitWireKind is called when production wireKind is exited.
func (s *BaseggwListener) ExitWireKind(ctx *WireKindContext) {}

// EnterWireDeclaration is called when production wireDeclaration is entered.
func (s *BaseggwListener) EnterWireDeclaration(ctx *WireDeclarationContext) {}

// ExitWireDeclaration is called when production wireDeclaration is exited.
func (s *BaseggwListener) ExitWireDeclaration(ctx *WireDeclarationContext) {}

// EnterWireAuth is called when production wireAuth is entered.
func (s *BaseggwListener) EnterWireAuth(ctx *WireAuthContext) {}

// ExitWireAuth is called when production wireAuth is exited.
func (s *BaseggwListener) ExitWireAuth(ctx *WireAuthContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BaseggwListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BaseggwListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterFunctionClassDeclaration is called when production functionClassDeclaration is entered.
func (s *BaseggwListener) EnterFunctionClassDeclaration(ctx *FunctionClassDeclarationContext) {}

// ExitFunctionClassDeclaration is called when production functionClassDeclaration is exited.
func (s *BaseggwListener) ExitFunctionClassDeclaration(ctx *FunctionClassDeclarationContext) {}

// EnterAnonymousFunctionDeclaration is called when production anonymousFunctionDeclaration is entered.
func (s *BaseggwListener) EnterAnonymousFunctionDeclaration(ctx *AnonymousFunctionDeclarationContext) {
}

// ExitAnonymousFunctionDeclaration is called when production anonymousFunctionDeclaration is exited.
func (s *BaseggwListener) ExitAnonymousFunctionDeclaration(ctx *AnonymousFunctionDeclarationContext) {
}

// EnterConstructorHeadDeclaration is called when production constructorHeadDeclaration is entered.
func (s *BaseggwListener) EnterConstructorHeadDeclaration(ctx *ConstructorHeadDeclarationContext) {}

// ExitConstructorHeadDeclaration is called when production constructorHeadDeclaration is exited.
func (s *BaseggwListener) ExitConstructorHeadDeclaration(ctx *ConstructorHeadDeclarationContext) {}

// EnterConstructorFuncDeclaration is called when production constructorFuncDeclaration is entered.
func (s *BaseggwListener) EnterConstructorFuncDeclaration(ctx *ConstructorFuncDeclarationContext) {}

// ExitConstructorFuncDeclaration is called when production constructorFuncDeclaration is exited.
func (s *BaseggwListener) ExitConstructorFuncDeclaration(ctx *ConstructorFuncDeclarationContext) {}

// EnterEnumDeclaration is called when production enumDeclaration is entered.
func (s *BaseggwListener) EnterEnumDeclaration(ctx *EnumDeclarationContext) {}

// ExitEnumDeclaration is called when production enumDeclaration is exited.
func (s *BaseggwListener) ExitEnumDeclaration(ctx *EnumDeclarationContext) {}

// EnterEnumProperty is called when production enumProperty is entered.
func (s *BaseggwListener) EnterEnumProperty(ctx *EnumPropertyContext) {}

// ExitEnumProperty is called when production enumProperty is exited.
func (s *BaseggwListener) ExitEnumProperty(ctx *EnumPropertyContext) {}

// EnterInterfaceDeclaration is called when production interfaceDeclaration is entered.
func (s *BaseggwListener) EnterInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// ExitInterfaceDeclaration is called when production interfaceDeclaration is exited.
func (s *BaseggwListener) ExitInterfaceDeclaration(ctx *InterfaceDeclarationContext) {}

// EnterInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is entered.
func (s *BaseggwListener) EnterInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// ExitInterfaceBodyDeclaration is called when production interfaceBodyDeclaration is exited.
func (s *BaseggwListener) ExitInterfaceBodyDeclaration(ctx *InterfaceBodyDeclarationContext) {}

// EnterInterfaceProperty is called when production interfaceProperty is entered.
func (s *BaseggwListener) EnterInterfaceProperty(ctx *InterfacePropertyContext) {}

// ExitInterfaceProperty is called when production interfaceProperty is exited.
func (s *BaseggwListener) ExitInterfaceProperty(ctx *InterfacePropertyContext) {}

// EnterInterfaceFunction is called when production interfaceFunction is entered.
func (s *BaseggwListener) EnterInterfaceFunction(ctx *InterfaceFunctionContext) {}

// ExitInterfaceFunction is called when production interfaceFunction is exited.
func (s *BaseggwListener) ExitInterfaceFunction(ctx *InterfaceFunctionContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BaseggwListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BaseggwListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterClassBodyDeclaration is called when production classBodyDeclaration is entered.
func (s *BaseggwListener) EnterClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// ExitClassBodyDeclaration is called when production classBodyDeclaration is exited.
func (s *BaseggwListener) ExitClassBodyDeclaration(ctx *ClassBodyDeclarationContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseggwListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseggwListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhenStatement is called when production whenStatement is entered.
func (s *BaseggwListener) EnterWhenStatement(ctx *WhenStatementContext) {}

// ExitWhenStatement is called when production whenStatement is exited.
func (s *BaseggwListener) ExitWhenStatement(ctx *WhenStatementContext) {}

// EnterWhenCase is called when production whenCase is entered.
func (s *BaseggwListener) EnterWhenCase(ctx *WhenCaseContext) {}

// ExitWhenCase is called when production whenCase is exited.
func (s *BaseggwListener) ExitWhenCase(ctx *WhenCaseContext) {}

// EnterWhenCaseBlock is called when production whenCaseBlock is entered.
func (s *BaseggwListener) EnterWhenCaseBlock(ctx *WhenCaseBlockContext) {}

// ExitWhenCaseBlock is called when production whenCaseBlock is exited.
func (s *BaseggwListener) ExitWhenCaseBlock(ctx *WhenCaseBlockContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseggwListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseggwListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForIncrementExpression is called when production forIncrementExpression is entered.
func (s *BaseggwListener) EnterForIncrementExpression(ctx *ForIncrementExpressionContext) {}

// ExitForIncrementExpression is called when production forIncrementExpression is exited.
func (s *BaseggwListener) ExitForIncrementExpression(ctx *ForIncrementExpressionContext) {}

// EnterForInStatement is called when production forInStatement is entered.
func (s *BaseggwListener) EnterForInStatement(ctx *ForInStatementContext) {}

// ExitForInStatement is called when production forInStatement is exited.
func (s *BaseggwListener) ExitForInStatement(ctx *ForInStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseggwListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseggwListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterDoWhileStatement is called when production doWhileStatement is entered.
func (s *BaseggwListener) EnterDoWhileStatement(ctx *DoWhileStatementContext) {}

// ExitDoWhileStatement is called when production doWhileStatement is exited.
func (s *BaseggwListener) ExitDoWhileStatement(ctx *DoWhileStatementContext) {}

// EnterGrantStatement is called when production grantStatement is entered.
func (s *BaseggwListener) EnterGrantStatement(ctx *GrantStatementContext) {}

// ExitGrantStatement is called when production grantStatement is exited.
func (s *BaseggwListener) ExitGrantStatement(ctx *GrantStatementContext) {}

// EnterRevokeStatement is called when production revokeStatement is entered.
func (s *BaseggwListener) EnterRevokeStatement(ctx *RevokeStatementContext) {}

// ExitRevokeStatement is called when production revokeStatement is exited.
func (s *BaseggwListener) ExitRevokeStatement(ctx *RevokeStatementContext) {}

// EnterGrantTree is called when production grantTree is entered.
func (s *BaseggwListener) EnterGrantTree(ctx *GrantTreeContext) {}

// ExitGrantTree is called when production grantTree is exited.
func (s *BaseggwListener) ExitGrantTree(ctx *GrantTreeContext) {}

// EnterGrantTreeBody is called when production grantTreeBody is entered.
func (s *BaseggwListener) EnterGrantTreeBody(ctx *GrantTreeBodyContext) {}

// ExitGrantTreeBody is called when production grantTreeBody is exited.
func (s *BaseggwListener) ExitGrantTreeBody(ctx *GrantTreeBodyContext) {}

// EnterGrantTreeBranch is called when production grantTreeBranch is entered.
func (s *BaseggwListener) EnterGrantTreeBranch(ctx *GrantTreeBranchContext) {}

// ExitGrantTreeBranch is called when production grantTreeBranch is exited.
func (s *BaseggwListener) ExitGrantTreeBranch(ctx *GrantTreeBranchContext) {}

// EnterGrantTreeInheritance is called when production grantTreeInheritance is entered.
func (s *BaseggwListener) EnterGrantTreeInheritance(ctx *GrantTreeInheritanceContext) {}

// ExitGrantTreeInheritance is called when production grantTreeInheritance is exited.
func (s *BaseggwListener) ExitGrantTreeInheritance(ctx *GrantTreeInheritanceContext) {}

// EnterSessionDeclaration is called when production sessionDeclaration is entered.
func (s *BaseggwListener) EnterSessionDeclaration(ctx *SessionDeclarationContext) {}

// ExitSessionDeclaration is called when production sessionDeclaration is exited.
func (s *BaseggwListener) ExitSessionDeclaration(ctx *SessionDeclarationContext) {}

// EnterSessionBody is called when production sessionBody is entered.
func (s *BaseggwListener) EnterSessionBody(ctx *SessionBodyContext) {}

// ExitSessionBody is called when production sessionBody is exited.
func (s *BaseggwListener) ExitSessionBody(ctx *SessionBodyContext) {}

// EnterSessionProperty is called when production sessionProperty is entered.
func (s *BaseggwListener) EnterSessionProperty(ctx *SessionPropertyContext) {}

// ExitSessionProperty is called when production sessionProperty is exited.
func (s *BaseggwListener) ExitSessionProperty(ctx *SessionPropertyContext) {}

// EnterType is called when production type is entered.
func (s *BaseggwListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseggwListener) ExitType(ctx *TypeContext) {}

// EnterPrimitiveType is called when production primitiveType is entered.
func (s *BaseggwListener) EnterPrimitiveType(ctx *PrimitiveTypeContext) {}

// ExitPrimitiveType is called when production primitiveType is exited.
func (s *BaseggwListener) ExitPrimitiveType(ctx *PrimitiveTypeContext) {}

// EnterArrayType is called when production arrayType is entered.
func (s *BaseggwListener) EnterArrayType(ctx *ArrayTypeContext) {}

// ExitArrayType is called when production arrayType is exited.
func (s *BaseggwListener) ExitArrayType(ctx *ArrayTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseggwListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseggwListener) ExitMapType(ctx *MapTypeContext) {}

// EnterTupleType is called when production tupleType is entered.
func (s *BaseggwListener) EnterTupleType(ctx *TupleTypeContext) {}

// ExitTupleType is called when production tupleType is exited.
func (s *BaseggwListener) ExitTupleType(ctx *TupleTypeContext) {}

// EnterGenericTypeDeclaration is called when production genericTypeDeclaration is entered.
func (s *BaseggwListener) EnterGenericTypeDeclaration(ctx *GenericTypeDeclarationContext) {}

// ExitGenericTypeDeclaration is called when production genericTypeDeclaration is exited.
func (s *BaseggwListener) ExitGenericTypeDeclaration(ctx *GenericTypeDeclarationContext) {}

// EnterGenericDeclaration is called when production genericDeclaration is entered.
func (s *BaseggwListener) EnterGenericDeclaration(ctx *GenericDeclarationContext) {}

// ExitGenericDeclaration is called when production genericDeclaration is exited.
func (s *BaseggwListener) ExitGenericDeclaration(ctx *GenericDeclarationContext) {}

// EnterGenericUsage is called when production genericUsage is entered.
func (s *BaseggwListener) EnterGenericUsage(ctx *GenericUsageContext) {}

// ExitGenericUsage is called when production genericUsage is exited.
func (s *BaseggwListener) ExitGenericUsage(ctx *GenericUsageContext) {}

// EnterIdentifierUsage is called when production identifierUsage is entered.
func (s *BaseggwListener) EnterIdentifierUsage(ctx *IdentifierUsageContext) {}

// ExitIdentifierUsage is called when production identifierUsage is exited.
func (s *BaseggwListener) ExitIdentifierUsage(ctx *IdentifierUsageContext) {}

// EnterExtends is called when production extends is entered.
func (s *BaseggwListener) EnterExtends(ctx *ExtendsContext) {}

// ExitExtends is called when production extends is exited.
func (s *BaseggwListener) ExitExtends(ctx *ExtendsContext) {}

// EnterInterfaceExtends is called when production interfaceExtends is entered.
func (s *BaseggwListener) EnterInterfaceExtends(ctx *InterfaceExtendsContext) {}

// ExitInterfaceExtends is called when production interfaceExtends is exited.
func (s *BaseggwListener) ExitInterfaceExtends(ctx *InterfaceExtendsContext) {}

// EnterImplements is called when production implements is entered.
func (s *BaseggwListener) EnterImplements(ctx *ImplementsContext) {}

// ExitImplements is called when production implements is exited.
func (s *BaseggwListener) ExitImplements(ctx *ImplementsContext) {}
