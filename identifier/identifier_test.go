package identifier

import (
	"testing"

	"github.com/smokers10/go-infrastructure/config"
)

func TestIdentifier(t *testing.T) {
	c := config.Configuration{
		Application: config.Application{Port: ":8080", Secret: "ABSGDAYSTD67"},
	}

	identifier := Identifier(&c)

	t.Run("make identifier", func(t *testing.T) {
		id, err := identifier.MakeIdentifier()
		if err != nil {
			t.Fatalf("error make identifier : %v\n", err.Error())
		}

		t.Log(id)
	})

	t.Run("make otp", func(t *testing.T) {
		otp, err := identifier.GenerateOTP()
		if err != nil {
			t.Fatalf("error make otp : %v\n", err.Error())
		}

		t.Log(otp)
	})
}
