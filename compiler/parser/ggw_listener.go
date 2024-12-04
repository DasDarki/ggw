// Code generated from ggw.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // ggw
import "github.com/antlr4-go/antlr/v4"

// ggwListener is a complete listener for a parse tree produced by ggwParser.
type ggwListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterLogicalExpression is called when entering the logicalExpression production.
	EnterLogicalExpression(c *LogicalExpressionContext)

	// EnterComparisonExpression is called when entering the comparisonExpression production.
	EnterComparisonExpression(c *ComparisonExpressionContext)

	// EnterAdditiveExpression is called when entering the additiveExpression production.
	EnterAdditiveExpression(c *AdditiveExpressionContext)

	// EnterMultiplicativeExpression is called when entering the multiplicativeExpression production.
	EnterMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// EnterPostfixExpression is called when entering the postfixExpression production.
	EnterPostfixExpression(c *PostfixExpressionContext)

	// EnterPostfixOperator is called when entering the postfixOperator production.
	EnterPostfixOperator(c *PostfixOperatorContext)

	// EnterIncrementalExpression is called when entering the incrementalExpression production.
	EnterIncrementalExpression(c *IncrementalExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterGroupedExpression is called when entering the groupedExpression production.
	EnterGroupedExpression(c *GroupedExpressionContext)

	// EnterFuncCallExpression is called when entering the funcCallExpression production.
	EnterFuncCallExpression(c *FuncCallExpressionContext)

	// EnterFuncCallChaining is called when entering the funcCallChaining production.
	EnterFuncCallChaining(c *FuncCallChainingContext)

	// EnterArrayExpression is called when entering the arrayExpression production.
	EnterArrayExpression(c *ArrayExpressionContext)

	// EnterMapExpression is called when entering the mapExpression production.
	EnterMapExpression(c *MapExpressionContext)

	// EnterOnStatement is called when entering the onStatement production.
	EnterOnStatement(c *OnStatementContext)

	// EnterModuleStatemnt is called when entering the moduleStatemnt production.
	EnterModuleStatemnt(c *ModuleStatemntContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterArgument is called when entering the argument production.
	EnterArgument(c *ArgumentContext)

	// EnterArgumentList is called when entering the argumentList production.
	EnterArgumentList(c *ArgumentListContext)

	// EnterArgumentPassList is called when entering the argumentPassList production.
	EnterArgumentPassList(c *ArgumentPassListContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterAnnotationList is called when entering the annotationList production.
	EnterAnnotationList(c *AnnotationListContext)

	// EnterAnnotationArgName is called when entering the annotationArgName production.
	EnterAnnotationArgName(c *AnnotationArgNameContext)

	// EnterAnnotationArgument is called when entering the annotationArgument production.
	EnterAnnotationArgument(c *AnnotationArgumentContext)

	// EnterAnnotationDeclaration is called when entering the annotationDeclaration production.
	EnterAnnotationDeclaration(c *AnnotationDeclarationContext)

	// EnterAccessModifier is called when entering the accessModifier production.
	EnterAccessModifier(c *AccessModifierContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterDeconstruct is called when entering the deconstruct production.
	EnterDeconstruct(c *DeconstructContext)

	// EnterVariableName is called when entering the variableName production.
	EnterVariableName(c *VariableNameContext)

	// EnterVariableKind is called when entering the variableKind production.
	EnterVariableKind(c *VariableKindContext)

	// EnterVariableDeclaration is called when entering the variableDeclaration production.
	EnterVariableDeclaration(c *VariableDeclarationContext)

	// EnterVariableMultiDeclaration is called when entering the variableMultiDeclaration production.
	EnterVariableMultiDeclaration(c *VariableMultiDeclarationContext)

	// EnterVariableClassDeclaration is called when entering the variableClassDeclaration production.
	EnterVariableClassDeclaration(c *VariableClassDeclarationContext)

	// EnterWireKind is called when entering the wireKind production.
	EnterWireKind(c *WireKindContext)

	// EnterWireDeclaration is called when entering the wireDeclaration production.
	EnterWireDeclaration(c *WireDeclarationContext)

	// EnterWireAuth is called when entering the wireAuth production.
	EnterWireAuth(c *WireAuthContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterFunctionClassDeclaration is called when entering the functionClassDeclaration production.
	EnterFunctionClassDeclaration(c *FunctionClassDeclarationContext)

	// EnterAnonymousFunctionDeclaration is called when entering the anonymousFunctionDeclaration production.
	EnterAnonymousFunctionDeclaration(c *AnonymousFunctionDeclarationContext)

	// EnterConstructorHeadDeclaration is called when entering the constructorHeadDeclaration production.
	EnterConstructorHeadDeclaration(c *ConstructorHeadDeclarationContext)

	// EnterConstructorFuncDeclaration is called when entering the constructorFuncDeclaration production.
	EnterConstructorFuncDeclaration(c *ConstructorFuncDeclarationContext)

	// EnterEnumDeclaration is called when entering the enumDeclaration production.
	EnterEnumDeclaration(c *EnumDeclarationContext)

	// EnterEnumProperty is called when entering the enumProperty production.
	EnterEnumProperty(c *EnumPropertyContext)

	// EnterInterfaceDeclaration is called when entering the interfaceDeclaration production.
	EnterInterfaceDeclaration(c *InterfaceDeclarationContext)

	// EnterInterfaceBodyDeclaration is called when entering the interfaceBodyDeclaration production.
	EnterInterfaceBodyDeclaration(c *InterfaceBodyDeclarationContext)

	// EnterInterfaceProperty is called when entering the interfaceProperty production.
	EnterInterfaceProperty(c *InterfacePropertyContext)

	// EnterInterfaceFunction is called when entering the interfaceFunction production.
	EnterInterfaceFunction(c *InterfaceFunctionContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterClassBodyDeclaration is called when entering the classBodyDeclaration production.
	EnterClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhenStatement is called when entering the whenStatement production.
	EnterWhenStatement(c *WhenStatementContext)

	// EnterWhenCase is called when entering the whenCase production.
	EnterWhenCase(c *WhenCaseContext)

	// EnterWhenCaseBlock is called when entering the whenCaseBlock production.
	EnterWhenCaseBlock(c *WhenCaseBlockContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForIncrementExpression is called when entering the forIncrementExpression production.
	EnterForIncrementExpression(c *ForIncrementExpressionContext)

	// EnterForInStatement is called when entering the forInStatement production.
	EnterForInStatement(c *ForInStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterDoWhileStatement is called when entering the doWhileStatement production.
	EnterDoWhileStatement(c *DoWhileStatementContext)

	// EnterGrantStatement is called when entering the grantStatement production.
	EnterGrantStatement(c *GrantStatementContext)

	// EnterRevokeStatement is called when entering the revokeStatement production.
	EnterRevokeStatement(c *RevokeStatementContext)

	// EnterGrantTree is called when entering the grantTree production.
	EnterGrantTree(c *GrantTreeContext)

	// EnterGrantTreeBody is called when entering the grantTreeBody production.
	EnterGrantTreeBody(c *GrantTreeBodyContext)

	// EnterGrantTreeBranch is called when entering the grantTreeBranch production.
	EnterGrantTreeBranch(c *GrantTreeBranchContext)

	// EnterGrantTreeInheritance is called when entering the grantTreeInheritance production.
	EnterGrantTreeInheritance(c *GrantTreeInheritanceContext)

	// EnterSessionDeclaration is called when entering the sessionDeclaration production.
	EnterSessionDeclaration(c *SessionDeclarationContext)

	// EnterSessionBody is called when entering the sessionBody production.
	EnterSessionBody(c *SessionBodyContext)

	// EnterSessionProperty is called when entering the sessionProperty production.
	EnterSessionProperty(c *SessionPropertyContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterPrimitiveType is called when entering the primitiveType production.
	EnterPrimitiveType(c *PrimitiveTypeContext)

	// EnterArrayType is called when entering the arrayType production.
	EnterArrayType(c *ArrayTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterTupleType is called when entering the tupleType production.
	EnterTupleType(c *TupleTypeContext)

	// EnterGenericTypeDeclaration is called when entering the genericTypeDeclaration production.
	EnterGenericTypeDeclaration(c *GenericTypeDeclarationContext)

	// EnterGenericDeclaration is called when entering the genericDeclaration production.
	EnterGenericDeclaration(c *GenericDeclarationContext)

	// EnterGenericUsage is called when entering the genericUsage production.
	EnterGenericUsage(c *GenericUsageContext)

	// EnterIdentifierUsage is called when entering the identifierUsage production.
	EnterIdentifierUsage(c *IdentifierUsageContext)

	// EnterExtends is called when entering the extends production.
	EnterExtends(c *ExtendsContext)

	// EnterInterfaceExtends is called when entering the interfaceExtends production.
	EnterInterfaceExtends(c *InterfaceExtendsContext)

	// EnterImplements is called when entering the implements production.
	EnterImplements(c *ImplementsContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitLogicalExpression is called when exiting the logicalExpression production.
	ExitLogicalExpression(c *LogicalExpressionContext)

	// ExitComparisonExpression is called when exiting the comparisonExpression production.
	ExitComparisonExpression(c *ComparisonExpressionContext)

	// ExitAdditiveExpression is called when exiting the additiveExpression production.
	ExitAdditiveExpression(c *AdditiveExpressionContext)

	// ExitMultiplicativeExpression is called when exiting the multiplicativeExpression production.
	ExitMultiplicativeExpression(c *MultiplicativeExpressionContext)

	// ExitPostfixExpression is called when exiting the postfixExpression production.
	ExitPostfixExpression(c *PostfixExpressionContext)

	// ExitPostfixOperator is called when exiting the postfixOperator production.
	ExitPostfixOperator(c *PostfixOperatorContext)

	// ExitIncrementalExpression is called when exiting the incrementalExpression production.
	ExitIncrementalExpression(c *IncrementalExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitGroupedExpression is called when exiting the groupedExpression production.
	ExitGroupedExpression(c *GroupedExpressionContext)

	// ExitFuncCallExpression is called when exiting the funcCallExpression production.
	ExitFuncCallExpression(c *FuncCallExpressionContext)

	// ExitFuncCallChaining is called when exiting the funcCallChaining production.
	ExitFuncCallChaining(c *FuncCallChainingContext)

	// ExitArrayExpression is called when exiting the arrayExpression production.
	ExitArrayExpression(c *ArrayExpressionContext)

	// ExitMapExpression is called when exiting the mapExpression production.
	ExitMapExpression(c *MapExpressionContext)

	// ExitOnStatement is called when exiting the onStatement production.
	ExitOnStatement(c *OnStatementContext)

	// ExitModuleStatemnt is called when exiting the moduleStatemnt production.
	ExitModuleStatemnt(c *ModuleStatemntContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitArgument is called when exiting the argument production.
	ExitArgument(c *ArgumentContext)

	// ExitArgumentList is called when exiting the argumentList production.
	ExitArgumentList(c *ArgumentListContext)

	// ExitArgumentPassList is called when exiting the argumentPassList production.
	ExitArgumentPassList(c *ArgumentPassListContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitAnnotationList is called when exiting the annotationList production.
	ExitAnnotationList(c *AnnotationListContext)

	// ExitAnnotationArgName is called when exiting the annotationArgName production.
	ExitAnnotationArgName(c *AnnotationArgNameContext)

	// ExitAnnotationArgument is called when exiting the annotationArgument production.
	ExitAnnotationArgument(c *AnnotationArgumentContext)

	// ExitAnnotationDeclaration is called when exiting the annotationDeclaration production.
	ExitAnnotationDeclaration(c *AnnotationDeclarationContext)

	// ExitAccessModifier is called when exiting the accessModifier production.
	ExitAccessModifier(c *AccessModifierContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitDeconstruct is called when exiting the deconstruct production.
	ExitDeconstruct(c *DeconstructContext)

	// ExitVariableName is called when exiting the variableName production.
	ExitVariableName(c *VariableNameContext)

	// ExitVariableKind is called when exiting the variableKind production.
	ExitVariableKind(c *VariableKindContext)

	// ExitVariableDeclaration is called when exiting the variableDeclaration production.
	ExitVariableDeclaration(c *VariableDeclarationContext)

	// ExitVariableMultiDeclaration is called when exiting the variableMultiDeclaration production.
	ExitVariableMultiDeclaration(c *VariableMultiDeclarationContext)

	// ExitVariableClassDeclaration is called when exiting the variableClassDeclaration production.
	ExitVariableClassDeclaration(c *VariableClassDeclarationContext)

	// ExitWireKind is called when exiting the wireKind production.
	ExitWireKind(c *WireKindContext)

	// ExitWireDeclaration is called when exiting the wireDeclaration production.
	ExitWireDeclaration(c *WireDeclarationContext)

	// ExitWireAuth is called when exiting the wireAuth production.
	ExitWireAuth(c *WireAuthContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitFunctionClassDeclaration is called when exiting the functionClassDeclaration production.
	ExitFunctionClassDeclaration(c *FunctionClassDeclarationContext)

	// ExitAnonymousFunctionDeclaration is called when exiting the anonymousFunctionDeclaration production.
	ExitAnonymousFunctionDeclaration(c *AnonymousFunctionDeclarationContext)

	// ExitConstructorHeadDeclaration is called when exiting the constructorHeadDeclaration production.
	ExitConstructorHeadDeclaration(c *ConstructorHeadDeclarationContext)

	// ExitConstructorFuncDeclaration is called when exiting the constructorFuncDeclaration production.
	ExitConstructorFuncDeclaration(c *ConstructorFuncDeclarationContext)

	// ExitEnumDeclaration is called when exiting the enumDeclaration production.
	ExitEnumDeclaration(c *EnumDeclarationContext)

	// ExitEnumProperty is called when exiting the enumProperty production.
	ExitEnumProperty(c *EnumPropertyContext)

	// ExitInterfaceDeclaration is called when exiting the interfaceDeclaration production.
	ExitInterfaceDeclaration(c *InterfaceDeclarationContext)

	// ExitInterfaceBodyDeclaration is called when exiting the interfaceBodyDeclaration production.
	ExitInterfaceBodyDeclaration(c *InterfaceBodyDeclarationContext)

	// ExitInterfaceProperty is called when exiting the interfaceProperty production.
	ExitInterfaceProperty(c *InterfacePropertyContext)

	// ExitInterfaceFunction is called when exiting the interfaceFunction production.
	ExitInterfaceFunction(c *InterfaceFunctionContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitClassBodyDeclaration is called when exiting the classBodyDeclaration production.
	ExitClassBodyDeclaration(c *ClassBodyDeclarationContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhenStatement is called when exiting the whenStatement production.
	ExitWhenStatement(c *WhenStatementContext)

	// ExitWhenCase is called when exiting the whenCase production.
	ExitWhenCase(c *WhenCaseContext)

	// ExitWhenCaseBlock is called when exiting the whenCaseBlock production.
	ExitWhenCaseBlock(c *WhenCaseBlockContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForIncrementExpression is called when exiting the forIncrementExpression production.
	ExitForIncrementExpression(c *ForIncrementExpressionContext)

	// ExitForInStatement is called when exiting the forInStatement production.
	ExitForInStatement(c *ForInStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitDoWhileStatement is called when exiting the doWhileStatement production.
	ExitDoWhileStatement(c *DoWhileStatementContext)

	// ExitGrantStatement is called when exiting the grantStatement production.
	ExitGrantStatement(c *GrantStatementContext)

	// ExitRevokeStatement is called when exiting the revokeStatement production.
	ExitRevokeStatement(c *RevokeStatementContext)

	// ExitGrantTree is called when exiting the grantTree production.
	ExitGrantTree(c *GrantTreeContext)

	// ExitGrantTreeBody is called when exiting the grantTreeBody production.
	ExitGrantTreeBody(c *GrantTreeBodyContext)

	// ExitGrantTreeBranch is called when exiting the grantTreeBranch production.
	ExitGrantTreeBranch(c *GrantTreeBranchContext)

	// ExitGrantTreeInheritance is called when exiting the grantTreeInheritance production.
	ExitGrantTreeInheritance(c *GrantTreeInheritanceContext)

	// ExitSessionDeclaration is called when exiting the sessionDeclaration production.
	ExitSessionDeclaration(c *SessionDeclarationContext)

	// ExitSessionBody is called when exiting the sessionBody production.
	ExitSessionBody(c *SessionBodyContext)

	// ExitSessionProperty is called when exiting the sessionProperty production.
	ExitSessionProperty(c *SessionPropertyContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitPrimitiveType is called when exiting the primitiveType production.
	ExitPrimitiveType(c *PrimitiveTypeContext)

	// ExitArrayType is called when exiting the arrayType production.
	ExitArrayType(c *ArrayTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitTupleType is called when exiting the tupleType production.
	ExitTupleType(c *TupleTypeContext)

	// ExitGenericTypeDeclaration is called when exiting the genericTypeDeclaration production.
	ExitGenericTypeDeclaration(c *GenericTypeDeclarationContext)

	// ExitGenericDeclaration is called when exiting the genericDeclaration production.
	ExitGenericDeclaration(c *GenericDeclarationContext)

	// ExitGenericUsage is called when exiting the genericUsage production.
	ExitGenericUsage(c *GenericUsageContext)

	// ExitIdentifierUsage is called when exiting the identifierUsage production.
	ExitIdentifierUsage(c *IdentifierUsageContext)

	// ExitExtends is called when exiting the extends production.
	ExitExtends(c *ExtendsContext)

	// ExitInterfaceExtends is called when exiting the interfaceExtends production.
	ExitInterfaceExtends(c *InterfaceExtendsContext)

	// ExitImplements is called when exiting the implements production.
	ExitImplements(c *ImplementsContext)
}
