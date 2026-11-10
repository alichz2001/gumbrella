package dto

import (
	"github.com/alichz2001/gombrella/services/todo/api"
	"github.com/alichz2001/gombrella/services/todo/biz"
)

// convertTodoStatus resolves the shared enum name through the generated proto
// lookup table. Unknown names (including the empty Go zero value) map to the
// proto's zero-valued unspecified sentinel.
func ConvertTodoStatus(in biz.TodoStatus) api.TodoStatus {
	return api.TodoStatus(api.TodoStatus_value[in.String()])
}
