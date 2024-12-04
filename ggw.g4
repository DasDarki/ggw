grammar ggw;

// Lexer rules
WS : [ \t\r\n]+ -> skip;
SEMICOLON : ';' -> skip;
SINGLE_LINE_COMMENT : '//' ~[\r\n]* -> skip;
MULTI_LINE_COMMENT : '/*' .*? '*/' -> skip;
DOC_COMMENT : '///' ~[\r\n]*;
INT : [0-9]+;
FLOAT : ([0-9]* '.') [0-9]+;
BOOL : 'true' | 'false';
STRING : '"' (~["\r\n])* '"' | '\'' (~['\r\n])* '\'' | '`' (~[`])* '`';
IDENTIFIER : [a-zA-Z_] [a-zA-Z_0-9]*;

// Parser rules
file : statement* EOF;
block : '{' statement* '}';

statement : onStatement | moduleStatemnt | importStatement | variableDeclaration | variableMultiDeclaration | functionDeclaration | constructorFuncDeclaration | enumDeclaration | interfaceDeclaration | classDeclaration | ifStatement | forStatement | forInStatement | whileStatement | doWhileStatement | funcCallExpression;

// Refactored expression rules
expression
    : logicalExpression
    ;

logicalExpression
    : comparisonExpression (('&&' | '||') comparisonExpression)*
    ;

comparisonExpression
    : additiveExpression (('<' | '<=' | '>' | '>=' | '==' | '!=') additiveExpression)*
    ;

additiveExpression
    : multiplicativeExpression (('+' | '-') multiplicativeExpression)*
    ;

multiplicativeExpression
    : postfixExpression (('*' | '/' | '%') postfixExpression)*
    ;

// Handling postfix expressions
postfixExpression
    : primaryExpression postfixOperator*
    ;

postfixOperator
    : '[' expression ']'          // Field access
    | '?' expression ':' expression // Ternary (unary if)
    ;

// Base expressions
primaryExpression
    : INT
    | FLOAT
    | BOOL
    | STRING
    | IDENTIFIER
    | groupedExpression
    | funcCallExpression
    ;

// Grouping expressions
groupedExpression
    : '(' expression ')'
    ;

// Function calls
funcCallExpression
    : identifierUsage '(' argumentPassList ')'
    ;

// Arrays and maps
arrayExpression : '[' expression (',' expression)* ']';
mapExpression : '[' expression ':' expression (',' expression ':' expression)* ']';

// Head statements
onStatement : 'on' ('backend'|'frontend'|'shared');
moduleStatemnt : 'module' IDENTIFIER onStatement?;
importStatement : 'import' IDENTIFIER ('as' IDENTIFIER)? ('from' STRING)?;

// Arguments
argument : IDENTIFIER typeDeclaration ('=' expression)?;
argumentList : argument (',' argument)* ('...' argument)?;
argumentPassList : expression (',' expression)*;

// Annotations
annotation : '@' IDENTIFIER ('(' (annotationArgument (',' annotationArgument)*)? ')')?;
annotationList : annotation*;
annotationArgName : IDENTIFIER '=';
annotationArgument : annotationArgName? (STRING | INT | FLOAT | BOOL | IDENTIFIER);
annotationDeclaration : 'annotation' IDENTIFIER '(' argumentList ')';

// Declarations
accessModifier : 'public' | 'private' | 'protected';
typeDeclaration : ':' type;

deconstruct : '(' IDENTIFIER (',' IDENTIFIER)* ')';

variableName : IDENTIFIER | deconstruct ;
variableKind : accessModifier? 'let' | 'const';
variableDeclaration : variableKind? variableName typeDeclaration? ('=' expression)?;
variableMultiDeclaration : variableKind? variableName (',' variableName)* typeDeclaration? ('=' expression (',' expression)*)?;

wireKind : 'WebSockets' | 'HTTP';
wireDeclaration : 'wire' wireKind? onStatement? ;

functionDeclaration : annotationList? 'func' IDENTIFIER genericDeclaration? '(' argumentList ')' typeDeclaration? wireDeclaration? block;
functionClassDeclaration : annotationList? accessModifier? 'static'? 'func' IDENTIFIER genericDeclaration? '(' argumentList ')' typeDeclaration? block;

constructorHeadDeclaration : '(' argumentList ')' ;
constructorFuncDeclaration : annotationList? accessModifier? 'constructor' constructorHeadDeclaration block;

enumDeclaration : annotationList accessModifier? 'enum' IDENTIFIER constructorHeadDeclaration '{' enumProperty (',' enumProperty)* functionClassDeclaration* '}';
enumProperty : annotationList IDENTIFIER ('(' argumentPassList ')')?;

interfaceDeclaration : accessModifier? 'interface' IDENTIFIER genericDeclaration? interfaceExtends? '{' interfaceBodyDeclaration* '}';
interfaceBodyDeclaration : interfaceProperty | interfaceFunction;
interfaceProperty : variableKind? IDENTIFIER typeDeclaration;
interfaceFunction : IDENTIFIER genericDeclaration? '(' argumentList ')' typeDeclaration;

classDeclaration : annotationList accessModifier? 'class' IDENTIFIER genericDeclaration? constructorHeadDeclaration? extends? implements? onStatement? '{' classBodyDeclaration* '}';
classBodyDeclaration : variableDeclaration | functionClassDeclaration | constructorFuncDeclaration | classDeclaration | interfaceDeclaration | enumDeclaration;

// Control flow
ifStatement : 'if' '('? expression ')'? block ('else if' '('? expression ')'? block)* ('else' block)?;

whenStatement : 'when' '('? expression ')'? '{' whenCase* ('else' '=>' whenCaseBlock) '}';
whenCase : expression '=>' whenCaseBlock;
whenCaseBlock : '{' statement* '}' | expression;

forStatement : 'for' '('? variableDeclaration? ';' expression? ';' forIncrementExpression? ')'? block;
forIncrementExpression : IDENTIFIER ('++' | '--') | ('++' | '--') IDENTIFIER | IDENTIFIER ('+=' | '-=') expression;

forInStatement : 'for' '('? variableName (',' variableName)* 'in' expression ')'? block;

whileStatement : 'while' '('? expression ')'? block;
doWhileStatement : 'do' block 'while' '('? expression ')'?;

// Type System
type : primitiveType | arrayType | mapType | tupleType | IDENTIFIER;
primitiveType : 'int' | 'float' | 'bool' | 'string' | 'none';
arrayType : '[]' type;
mapType: '[' type ']' type;
tupleType : '(' type (',' type)* ')';

genericTypeDeclaration : IDENTIFIER extends? implements? ;
genericDeclaration : '<' genericTypeDeclaration (',' genericTypeDeclaration)* '>';
genericUsage : '<' type '>' ;

identifierUsage : IDENTIFIER genericUsage? ('.' identifierUsage)*;
extends : 'extends' identifierUsage;
interfaceExtends : 'extends' identifierUsage (',' identifierUsage)*;
implements : 'implements' identifierUsage (',' identifierUsage)*;