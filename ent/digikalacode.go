// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/espitman/jbm-hr-backend/ent/digikalacode"
	"github.com/espitman/jbm-hr-backend/ent/user"
)

// DigikalaCode is the model entity for the DigikalaCode schema.
type DigikalaCode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Code holds the value of the "code" field.
	Code string `json:"code,omitempty"`
	// Used holds the value of the "used" field.
	Used bool `json:"used,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// AssignToUserID holds the value of the "assign_to_user_id" field.
	AssignToUserID int `json:"assign_to_user_id,omitempty"`
	// AssignAt holds the value of the "assign_at" field.
	AssignAt time.Time `json:"assign_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DigikalaCodeQuery when eager-loading is set.
	Edges        DigikalaCodeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// DigikalaCodeEdges holds the relations/edges for other nodes in the graph.
type DigikalaCodeEdges struct {
	// AssignedTo holds the value of the assigned_to edge.
	AssignedTo *User `json:"assigned_to,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AssignedToOrErr returns the AssignedTo value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DigikalaCodeEdges) AssignedToOrErr() (*User, error) {
	if e.AssignedTo != nil {
		return e.AssignedTo, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "assigned_to"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DigikalaCode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case digikalacode.FieldUsed:
			values[i] = new(sql.NullBool)
		case digikalacode.FieldID, digikalacode.FieldAssignToUserID:
			values[i] = new(sql.NullInt64)
		case digikalacode.FieldCode:
			values[i] = new(sql.NullString)
		case digikalacode.FieldCreatedAt, digikalacode.FieldAssignAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DigikalaCode fields.
func (dc *DigikalaCode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case digikalacode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dc.ID = int(value.Int64)
		case digikalacode.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				dc.Code = value.String
			}
		case digikalacode.FieldUsed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field used", values[i])
			} else if value.Valid {
				dc.Used = value.Bool
			}
		case digikalacode.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dc.CreatedAt = value.Time
			}
		case digikalacode.FieldAssignToUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field assign_to_user_id", values[i])
			} else if value.Valid {
				dc.AssignToUserID = int(value.Int64)
			}
		case digikalacode.FieldAssignAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field assign_at", values[i])
			} else if value.Valid {
				dc.AssignAt = value.Time
			}
		default:
			dc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DigikalaCode.
// This includes values selected through modifiers, order, etc.
func (dc *DigikalaCode) Value(name string) (ent.Value, error) {
	return dc.selectValues.Get(name)
}

// QueryAssignedTo queries the "assigned_to" edge of the DigikalaCode entity.
func (dc *DigikalaCode) QueryAssignedTo() *UserQuery {
	return NewDigikalaCodeClient(dc.config).QueryAssignedTo(dc)
}

// Update returns a builder for updating this DigikalaCode.
// Note that you need to call DigikalaCode.Unwrap() before calling this method if this DigikalaCode
// was returned from a transaction, and the transaction was committed or rolled back.
func (dc *DigikalaCode) Update() *DigikalaCodeUpdateOne {
	return NewDigikalaCodeClient(dc.config).UpdateOne(dc)
}

// Unwrap unwraps the DigikalaCode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dc *DigikalaCode) Unwrap() *DigikalaCode {
	_tx, ok := dc.config.driver.(*txDriver)
	if !ok {
		panic("ent: DigikalaCode is not a transactional entity")
	}
	dc.config.driver = _tx.drv
	return dc
}

// String implements the fmt.Stringer.
func (dc *DigikalaCode) String() string {
	var builder strings.Builder
	builder.WriteString("DigikalaCode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dc.ID))
	builder.WriteString("code=")
	builder.WriteString(dc.Code)
	builder.WriteString(", ")
	builder.WriteString("used=")
	builder.WriteString(fmt.Sprintf("%v", dc.Used))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(dc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("assign_to_user_id=")
	builder.WriteString(fmt.Sprintf("%v", dc.AssignToUserID))
	builder.WriteString(", ")
	builder.WriteString("assign_at=")
	builder.WriteString(dc.AssignAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// DigikalaCodes is a parsable slice of DigikalaCode.
type DigikalaCodes []*DigikalaCode
