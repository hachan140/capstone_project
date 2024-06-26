package middleware

import (
	"be-capstone-project/src/internal/core/web/constant"
	"be-capstone-project/src/internal/core/web/context"
	"github.com/google/uuid"
	"net/http"
)

// CorrelationId middleware responsible to inject correlationId to request attributes
// correlationId is usually sent in the request header by the client (see constant.HeaderCorrelationId),
// but sometimes it doesn't exist, we will generate it automatically by guid
func CorrelationId() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.GetOrCreateRequestAttributes(r).CorrelationId = getOrNewCorrelationId(r)
			next.ServeHTTP(w, r)
		})
	}
}

func getOrNewCorrelationId(r *http.Request) string {
	correlationId := r.Header.Get(constant.HeaderCorrelationId)
	if len(correlationId) > 0 {
		return correlationId
	}
	newCorrelationId, _ := uuid.NewUUID()
	return newCorrelationId.String()
}
