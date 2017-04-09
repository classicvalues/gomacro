/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017 Massimiliano Ghilardi
 *
 *     This program is free software you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <http//www.gnu.org/licenses/>.
 *
 * function.go
 *
 *  Created on Apr 02, 2017
 *      Author Massimiliano Ghilardi
 */

package fast_interpreter

import (
	"go/ast"
	r "reflect"
)

func (c *Comp) DeclFunc(decl_or_nil *ast.FuncDecl, funcType *ast.FuncType, body *ast.BlockStmt) X {
	c.Errorf("function declaration is not implemented, found: %v <%v>", decl_or_nil, r.TypeOf(decl_or_nil))
	return nil
}