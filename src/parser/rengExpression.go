package parser

import (
	"RenG/src/ast"
	"RenG/src/token"
)

func (p *Parser) parseLabelExpression() ast.Expression {
	exp := &ast.LabelExpression{Token: p.curToken}

	p.nextToken()

	exp.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	exp.Body = p.parseBlockStatement()

	return exp
}

func (p *Parser) parseCallLabelExpression() ast.Expression {
	exp := &ast.CallLabelExpression{Token: p.curToken}

	p.nextToken()

	exp.Label = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return exp
}

func (p *Parser) parseJumpLabelExpression() ast.Expression {
	exp := &ast.JumpLabelExpression{Token: p.curToken}

	p.nextToken()

	exp.Label = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return exp
}

func (p *Parser) parseImageExpression() ast.Expression {
	exp := &ast.ImageExpression{Token: p.curToken}

	p.nextToken()

	exp.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	p.nextToken()

	exp.Path = p.parseStringLiteral()

	return exp
}

func (p *Parser) parseShowExpression() ast.Expression {
	exp := &ast.ShowExpression{Token: p.curToken}

	p.nextToken()

	exp.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if p.peekTokenIs(token.AT) {
		p.nextToken()
		p.nextToken()
		exp.Transform = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	} else {
		exp.Transform = &ast.Identifier{
			Token: token.Token{
				Type:    token.IDENT,
				Literal: "IDENT",
			},
			Value: "default",
		}
	}

	return exp
}

func (p *Parser) parseHideExpression() ast.Expression {
	exp := &ast.HideExpression{Token: p.curToken}

	p.nextToken()

	exp.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return exp
}

func (p *Parser) parseTranformExpression() ast.Expression {
	exp := &ast.TransformExpression{Token: p.curToken}

	p.nextToken()

	exp.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	exp.Body = p.parseBlockStatement()

	return exp
}

func (p *Parser) parseXposExpression() ast.Expression {
	exp := &ast.XPosExpression{Token: p.curToken}

	p.nextToken()

	exp.Value = p.parseExpression(PREFIX)

	return exp
}

func (p *Parser) parseYposExpression() ast.Expression {
	exp := &ast.YPosExpression{Token: p.curToken}

	p.nextToken()

	exp.Value = p.parseExpression(PREFIX)

	return exp
}
