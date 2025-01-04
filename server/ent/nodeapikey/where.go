// Code generated by ent, DO NOT EDIT.

package nodeapikey

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nullsploit01/iosync/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLTE(FieldID, id))
}

// APIKey applies equality check predicate on the "api_key" field. It's identical to APIKeyEQ.
func APIKey(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldAPIKey, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldDescription, v))
}

// IsRevoked applies equality check predicate on the "is_revoked" field. It's identical to IsRevokedEQ.
func IsRevoked(v bool) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldIsRevoked, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldUpdatedAt, v))
}

// APIKeyEQ applies the EQ predicate on the "api_key" field.
func APIKeyEQ(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldAPIKey, v))
}

// APIKeyNEQ applies the NEQ predicate on the "api_key" field.
func APIKeyNEQ(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNEQ(FieldAPIKey, v))
}

// APIKeyIn applies the In predicate on the "api_key" field.
func APIKeyIn(vs ...string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldIn(FieldAPIKey, vs...))
}

// APIKeyNotIn applies the NotIn predicate on the "api_key" field.
func APIKeyNotIn(vs ...string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNotIn(FieldAPIKey, vs...))
}

// APIKeyGT applies the GT predicate on the "api_key" field.
func APIKeyGT(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGT(FieldAPIKey, v))
}

// APIKeyGTE applies the GTE predicate on the "api_key" field.
func APIKeyGTE(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGTE(FieldAPIKey, v))
}

// APIKeyLT applies the LT predicate on the "api_key" field.
func APIKeyLT(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLT(FieldAPIKey, v))
}

// APIKeyLTE applies the LTE predicate on the "api_key" field.
func APIKeyLTE(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLTE(FieldAPIKey, v))
}

// APIKeyContains applies the Contains predicate on the "api_key" field.
func APIKeyContains(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldContains(FieldAPIKey, v))
}

// APIKeyHasPrefix applies the HasPrefix predicate on the "api_key" field.
func APIKeyHasPrefix(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldHasPrefix(FieldAPIKey, v))
}

// APIKeyHasSuffix applies the HasSuffix predicate on the "api_key" field.
func APIKeyHasSuffix(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldHasSuffix(FieldAPIKey, v))
}

// APIKeyEqualFold applies the EqualFold predicate on the "api_key" field.
func APIKeyEqualFold(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEqualFold(FieldAPIKey, v))
}

// APIKeyContainsFold applies the ContainsFold predicate on the "api_key" field.
func APIKeyContainsFold(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldContainsFold(FieldAPIKey, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldContainsFold(FieldDescription, v))
}

// IsRevokedEQ applies the EQ predicate on the "is_revoked" field.
func IsRevokedEQ(v bool) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldIsRevoked, v))
}

// IsRevokedNEQ applies the NEQ predicate on the "is_revoked" field.
func IsRevokedNEQ(v bool) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNEQ(FieldIsRevoked, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.NodeApiKey {
	return predicate.NodeApiKey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.NodeApiKey {
	return predicate.NodeApiKey(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasValues applies the HasEdge predicate on the "values" edge.
func HasValues() predicate.NodeApiKey {
	return predicate.NodeApiKey(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ValuesTable, ValuesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasValuesWith applies the HasEdge predicate on the "values" edge with a given conditions (other predicates).
func HasValuesWith(preds ...predicate.NodeValues) predicate.NodeApiKey {
	return predicate.NodeApiKey(func(s *sql.Selector) {
		step := newValuesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NodeApiKey) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NodeApiKey) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.NodeApiKey) predicate.NodeApiKey {
	return predicate.NodeApiKey(sql.NotPredicates(p))
}
