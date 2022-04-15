package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/renatoviolin/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

func TestInvalidToken(t *testing.T) {
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiYzBlOWM1MTItYzQxYy00YmRiLWJiOTQtNzAwZTY1NmZiNDIzIiwidXNlcm5hbWUiOiJydGFxZmQiLCJpc3N1ZWRfYXQiOiIyMDIyLTA0LTE1VDEyOjE4OjEwLjg1MzUwNS0wMzowMCIsImV4cGlyZWRfYXQiOiIyMDIyLTA0LTE1VDEyOjE5OjEwLjg1MzUwNS0wMzowMCJ9.fKVqNSBaR13r0Dk883Vv_7Uwt51s_vPwtPWP0-Fs_Y0")
	require.Error(t, err)
	require.Nil(t, payload)
	require.EqualError(t, err, ErrInvalidToken.Error())
}

func TestChangedToken(t *testing.T) {
	maker, err := NewJWTMaker("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	require.NoError(t, err)

	token, err := maker.CreateToken("renatoviolin", time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiZTRjMWY0YmUtNjYyNy00ODEyLWEwOWMtMjVjNGU1NWU4OTRkIiwidXNlcm5hbWUiOiJyZW5hdE92aW9saW4iLCJpc3N1ZWRfYXQiOiIyMDIyLTA0LTE1VDEyOjI4OjU2LjkyMDkzMS0wMzowMCIsImV4cGlyZWRfYXQiOiIyMDIyLTA0LTE1VDEyOjI5OjU2LjkyMDkzMS0wMzowMCJ9.ZhbjA57e6dE_tIUd_6A6wYSrsMzc0BJUKcArUtVgt-4")
	require.Error(t, err)
	require.Nil(t, payload)
	require.EqualError(t, err, ErrInvalidToken.Error())
}
