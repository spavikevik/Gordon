package evaluator

import (
	"gordon/helpers"
	"gordon/lexer"
	"gordon/object"
	"gordon/parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-8", -8},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"5 * 2 ^ 2", 20},
		{"5 ^ 2 ^ 2", 625},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalRealExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"5.6", 5.6},
		{"3.14", 3.14},
		{"-4.9", -4.9},
		{"5.2 + 0.4", 5.6},
		{"100.0 ^ 0.5", 10},
		{"3 / 0.5", 6},
		{"4 - 2.45", 1.55},
		{"5 * 2.3 - 2.48", 9.02},
		{"-8.9 + 10", 1.1},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testRealObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"3.14 == 3.14", true},
		{"3.14 != 2.72", true},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!3.14", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!!3.14", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, expected=%d",
			result.Value, expected)
		return false
	}

	return true
}

func testRealObject(t *testing.T, obj object.Object, expected float64) bool {
	result, ok := obj.(*object.Real)
	if !ok {
		t.Errorf("object is not Real. got=%T (%+v)", obj, obj)
		return false
	}
	if !helpers.Realeq(expected, result.Value) {
		t.Errorf("object has wrong value. got=%s, expected=%s",
			helpers.Realf(result.Value), helpers.Realf(expected))
		return false
	}

	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object is not Boolean. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%t, wanted=%t",
			result.Value, expected)
		return false
	}
	return true
}
