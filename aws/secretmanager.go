package aws

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/pkg/errors"
)

type Secrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func (a AWS) GetFromSM(key string) (Secrets, error) {
	var secretsVals Secrets
	output, err := a.sm.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: &key,
	})
	if err != nil {
		a.log.Println(err.Error())
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				a.log.Println(aerr.Error())
				a.log.Println("Secrets Manager could not decrypt the secret.")
			case secretsmanager.ErrCodeInternalServiceError:
				a.log.Println(aerr.Error())
				a.log.Println("Server side error.")
			case secretsmanager.ErrCodeInvalidParameterException:
				a.log.Println(aerr.Error())
				a.log.Println("Invalid parameter. Check inputs.")
			case secretsmanager.ErrCodeInvalidRequestException:
				a.log.Println(aerr.Error())
			case secretsmanager.ErrCodeResourceNotFoundException:
				a.log.Println(aerr.Error())
				a.log.Println("Is your secret name correct?")
			}
		}
		return secretsVals, err
	}
	dec := json.NewDecoder(strings.NewReader(*output.SecretString))
	if err := dec.Decode(&secretsVals); err != nil {
		return secretsVals, errors.Wrapf(err, "getting key %s from secret manager", key)
	}
	return secretsVals, nil
}
