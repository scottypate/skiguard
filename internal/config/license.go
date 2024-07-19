package config

import (
	"context"
	"log"

	"github.com/denisbrodbeck/machineid"
	"github.com/keygen-sh/keygen-go/v3"
)

func ValidateLicenseKey(licenseKey string) {
	keygen.Account = "77dfdf77-4a4e-45ce-8a6c-f12dff2c2cd0"
	keygen.Product = "7ec7b618-a587-431f-90e6-bb4f169909c6"
	keygen.PublicKey = "730267b9dbb9091639f2a32a66f9f22e1797537d795d77052ba7aabffeb598a8"
	keygen.LicenseKey = licenseKey

	fingerprint, err := machineid.ProtectedID(keygen.Product)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Validate the license for the current fingerprint
	license, err := keygen.Validate(ctx, fingerprint)
	switch {
	case err == keygen.ErrLicenseNotActivated:
		// Activate the current fingerprint
		_, err := license.Activate(ctx, fingerprint)
		switch {
		case err == keygen.ErrMachineLimitExceeded:
			log.Fatalf("Machine limit has been exceeded!")
		case err != nil:
			log.Fatal("Machine activation failed!")
		}
	case err == keygen.ErrLicenseExpired:
		log.Fatal("Snowguard license is expired!")
	case err != nil:
		log.Fatalf("Snowguard license is invalid! %v", err)
	}
}
