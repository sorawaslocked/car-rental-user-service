package service

import (
	"github.com/sorawaslocked/car-rental-user-service/internal/model"
	"strings"
)

func uncapitalize(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}

func toRoleStrings(roles []model.Role) []string {
	var result []string
	for _, role := range roles {
		result = append(result, role.String())
	}

	return result
}

func formatFilter(filter *model.UserFilter) {
	if filter.ID != nil && *filter.ID > 0 {
		filter.Email = nil
	}
	if filter.Email != nil && *filter.Email != "" {
		filter.ID = nil
	}
}
