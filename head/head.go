package head

import (
	"github.com/smokers10/go-infrastructure/config"
	"github.com/smokers10/go-infrastructure/contract"
	"github.com/smokers10/go-infrastructure/database"
	"github.com/smokers10/go-infrastructure/encryption"
	"github.com/smokers10/go-infrastructure/identifier"
	"github.com/smokers10/go-infrastructure/jsonwebtoken"
	"github.com/smokers10/go-infrastructure/mailer"
)

type head struct {
	DB         contract.DatabaseContract
	Encryption contract.EncryptionContract
	Identfier  contract.IdentfierContract
	JWT        contract.JsonWebTokenContract
	Mailer     contract.MailerContract
}

func Head(path string) (*head, error) {
	config, err := config.Reader("config.yaml")
	if err != nil {
		return nil, err
	}

	result := head{
		DB:         database.Database(config),
		Encryption: encryption.Encryption(),
		Identfier:  identifier.Identifier(),
		JWT:        jsonwebtoken.JsonWebToken(config),
		Mailer:     mailer.Mailer(config),
	}

	return &result, nil
}
