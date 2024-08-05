package config

import (
	"context"

	"github.com/denisbrodbeck/machineid"
	"github.com/keygen-sh/keygen-go/v3"
)

func ValidateLicenseKey(licenseKey string) error {
	keygen.Account = "77dfdf77-4a4e-45ce-8a6c-f12dff2c2cd0"
	keygen.Product = "7ec7b618-a587-431f-90e6-bb4f169909c6"
	keygen.PublicKey = "730267b9dbb9091639f2a32a66f9f22e1797537d795d77052ba7aabffeb598a8"
	keygen.LicenseKey = licenseKey

	fingerprint, err := machineid.ProtectedID(keygen.Product)
	if err != nil {
		return err
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
			return err
		case err != nil:
			return err
		}
	case err == keygen.ErrLicenseExpired:
		return err
	case err != nil:
		return err
	}

	return nil
}
