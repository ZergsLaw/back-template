package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ZergsLaw/back-template/internal/dom"
)

func TestUserStatus_IsFreeze(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {true, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {false, app.UserStatusSupport},
		"admin":   {false, app.UserStatusAdmin},
		"jedi":    {false, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsFreeze())
		})
	}
}

func TestUserStatus_IsDefault(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {true, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {false, app.UserStatusSupport},
		"admin":   {false, app.UserStatusAdmin},
		"jedi":    {false, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsDefault())
		})
	}
}

func TestUserStatus_IsPremium(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {true, app.UserStatusPremium},
		"support": {false, app.UserStatusSupport},
		"admin":   {false, app.UserStatusAdmin},
		"jedi":    {false, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsPremium())
		})
	}
}

func TestUserStatus_IsSupport(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {true, app.UserStatusSupport},
		"admin":   {false, app.UserStatusAdmin},
		"jedi":    {false, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsSupport())
		})
	}
}

func TestUserStatus_IsAdmin(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {false, app.UserStatusSupport},
		"admin":   {true, app.UserStatusAdmin},
		"jedi":    {false, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsAdmin())
		})
	}
}

func TestUserStatus_IsJedi(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {false, app.UserStatusSupport},
		"admin":   {false, app.UserStatusAdmin},
		"jedi":    {true, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsJedi())
		})
	}
}

func TestUserStatus_IsSpecialist(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {true, app.UserStatusSupport},
		"admin":   {true, app.UserStatusAdmin},
		"jedi":    {true, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsSpecialist())
		})
	}
}

func TestUserStatus_IsManager(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		want   bool
		status app.UserStatus
	}{
		"freeze":  {false, app.UserStatusFreeze},
		"default": {false, app.UserStatusDefault},
		"premium": {false, app.UserStatusPremium},
		"support": {false, app.UserStatusSupport},
		"admin":   {true, app.UserStatusAdmin},
		"jedi":    {true, app.UserStatusJedi},
	}

	for name, tc := range testCases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			assert := require.New(t)
			assert.Equal(tc.want, tc.status.IsManager())
		})
	}
}
