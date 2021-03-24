package sdk_core

// #include <github.com/VirgilSecurity/virgil-commkit-go/crypto/wrapper/sdk/core/vssc_core_sdk_public.h>
import "C"

type context interface {

	/* Get C context */
	Ctx() uintptr
}
