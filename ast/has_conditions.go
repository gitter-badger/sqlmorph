package ast

type HasConditions interface {
	AddCondition(*EqualsCondition)
}
