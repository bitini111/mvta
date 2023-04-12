package crypto

import (
	"github.com/bitini111/mvta/crypto/ecc"
	"github.com/bitini111/mvta/crypto/rsa"
	"github.com/bitini111/mvta/log"
)

type Encryptor interface {
	// Name 名称
	Name() string
	// Encrypt 加密
	Encrypt(data []byte) ([]byte, error)
}

var encryptors = make(map[string]Encryptor)

func init() {
	RegisterEncryptor(ecc.DefaultEncryptor)
	RegisterEncryptor(rsa.DefaultEncryptor)
}

// RegisterEncryptor 注册加密器
func RegisterEncryptor(encryptor Encryptor) {
	if encryptor == nil {
		log.Fatal("can't register a invalid encryptor")
	}

	name := encryptor.Name()

	if name == "" {
		log.Fatal("can't register a encryptor without name")
	}

	if _, ok := encryptors[name]; ok {
		log.Warnf("the old %s encryptor will be overwritten", name)
	}

	encryptors[name] = encryptor
}

// InvokeEncryptor 调用加密器
func InvokeEncryptor(name string) Encryptor {
	encryptor, ok := encryptors[name]
	if !ok {
		log.Fatalf("%s encryptor is not registered", name)
	}

	return encryptor
}
